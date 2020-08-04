package tim_utils_http

import (
	"io/ioutil"
	"net/http"
)

func SendGetMsg(iMicroServiceName string,
	iPort string,
	iServicePath string,
	iFormData map[string]string) (eResultAsByteArray []byte, eRespStatus string, eRespStatusCode int, eException ExceptionStruct) {

	eException = ExceptionStruct{}

	serviceadr := "http://" + iMicroServiceName + ":" + iPort + iServicePath
	println("serviceadr:" + serviceadr)
	urlQueryStr := ""
	for k, v := range iFormData {
		if len(urlQueryStr) == 0 {
			urlQueryStr = "?" + k + "=" + v
		} else {
			urlQueryStr = urlQueryStr + "&" + k + "=" + v
		}
	}
	serviceadr += urlQueryStr

	resp, err := http.Get(serviceadr)
	eRespStatus = resp.Status
	eRespStatusCode = resp.StatusCode
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "Technical error: http.Get failed." + err.Error()
	}

	defer resp.Body.Close()
	eResultAsByteArray, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "Technical error: read response body failed:" + err.Error()
	}

	return
}
