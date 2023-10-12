package handler

import (
	"encoding/json"
	"net/http/httptest"
	"referral-system/util"
)

func getBodyResponse(w *httptest.ResponseRecorder) (response util.BaseResponse) {
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		return response
	}

	return response
}
