package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"referral-system/entity"
	"referral-system/flow"
	"referral-system/parser"
	"referral-system/repository"
	"referral-system/util"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestContributorHandler_Referral_E2E(t *testing.T) {
	router := mux.NewRouter()
	db := repository.NewConnectMysqlDb()
	parser := parser.NewContributorParser()
	presenter := util.NewJsonPresenter()
	repoContributor := repository.NewContributorRepository(db)
	repoReferral := repository.NewReferralRepository(db)
	repoUser := repository.NewUserRepository(db)
	flow := flow.NewContributorFlow(repoContributor, repoReferral, repoUser)
	handler := NewContributorHandler(parser, presenter, flow)

	router.HandleFunc("/referral/{code}", handler.Referral).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()

	requestData := entity.Contributor{
		ReferralCode: "test_code",
		Email:        "test@example.com",
	}

	requestBody, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", server.URL+"/referral/test_code", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	// Send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	responseData := util.BaseResponse{
		Status:  200,
		Message: "Success.",
	}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	assert.NoError(t, err)

	assert.Equal(t, "OK", responseData.Message)
}
