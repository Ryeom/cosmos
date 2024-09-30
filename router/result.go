package router

import "net/http"

type HttpResult struct {
	ResultCode interface{} `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func GetDefaultResult() HttpResult {
	return HttpResult{
		ResultCode: http.StatusBadRequest,
		ResultMsg:  "Fail",
	}
}

func (r *HttpResult) OK() HttpResult {
	r.ResultCode = http.StatusOK
	r.ResultMsg = "OK"
	return *r
}
