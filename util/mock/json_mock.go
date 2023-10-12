package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"referral-system/util"

	"github.com/stretchr/testify/mock"
)

type MockJsonPresenter struct {
	mock.Mock
}

func (m *MockJsonPresenter) SendError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := util.BaseResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}

	errEncode := json.NewEncoder(w).Encode(response)
	if errEncode != nil {
		fmt.Println(errEncode)
	}
}

func (m *MockJsonPresenter) SendSuccess(w http.ResponseWriter, data ...interface{}) {
	var response = util.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success.",
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	errEncode := json.NewEncoder(w).Encode(response)
	if errEncode != nil {
		fmt.Println(errEncode)
	}
}

func (m *MockJsonPresenter) SendSuccessWithReferral(w http.ResponseWriter, referralLink string, data ...interface{}) {
	var response = util.BaseResponse{
		Status:       http.StatusOK,
		Message:      "Success.",
		ReferralLink: referralLink,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	errEncode := json.NewEncoder(w).Encode(response)
	if errEncode != nil {
		fmt.Println(errEncode)
	}
}
