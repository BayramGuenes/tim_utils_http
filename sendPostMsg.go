package tim_utils_http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendPostMsg(iMicroServiceName string,
	iPort string,
	iServicePath string,
	iData []byte) (eResultAsByteArray []byte, eRespStatus string, eRespStatusCode int, eException ExceptionStruct) {

	eException = ExceptionStruct{}

	serviceadr := "http://" + iMicroServiceName + ":" + iPort + iServicePath
	//println("serviceadr:" + serviceadr)
	println("serviceadr:" + serviceadr)

	eResultAsByteArray = []byte{}
	resp, err := http.Post(serviceadr, "application/json", bytes.NewBuffer(iData)) //<--
	eRespStatus = resp.Status
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "Technical error: http.Post failed." + err.Error()
	}
	//println(lServiceUrl)

	defer resp.Body.Close()

	eResultAsByteArray, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "Technical error: http.Post failed." + err.Error()
	}

	return
}
