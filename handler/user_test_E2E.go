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

func TestUserHandler_Register_E2E(t *testing.T) {
	router := mux.NewRouter()
	db := repository.NewConnectMysqlDb()
	parser := parser.NewUserParser()
	presenter := util.NewJsonPresenter()
	repoReferral := repository.NewReferralRepository(db)
	repoUser := repository.NewUserRepository(db)
	flowReferral := flow.NewReferralFlow(repoReferral)
	flowUser := flow.NewUserFlow(repoUser, repoReferral, flowReferral)
	handler := NewUserHandler(parser, presenter, flowUser)

	router.HandleFunc("/user", handler.Register).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()

	requestData := entity.User{
		Email:    "mail@mail.com",
		Password: "12345",
		UserName: "username",
		Name:     "name",
	}

	requestBody, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", server.URL+"/user", bytes.NewBuffer(requestBody))
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

func TestUserHandler_Login_E2E(t *testing.T) {
	router := mux.NewRouter()
	db := repository.NewConnectMysqlDb()
	parser := parser.NewUserParser()
	presenter := util.NewJsonPresenter()
	repoReferral := repository.NewReferralRepository(db)
	repoUser := repository.NewUserRepository(db)
	flowReferral := flow.NewReferralFlow(repoReferral)
	flowUser := flow.NewUserFlow(repoUser, repoReferral, flowReferral)
	handler := NewUserHandler(parser, presenter, flowUser)

	router.HandleFunc("/user", handler.Register).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()

	requestData := entity.User{
		Email:    "mail@mail.com",
		Password: "12345",
	}

	requestBody, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer(requestBody))
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

func TestUserHandler_Update_E2E(t *testing.T) {
	router := mux.NewRouter()
	db := repository.NewConnectMysqlDb()
	parser := parser.NewUserParser()
	presenter := util.NewJsonPresenter()
	repoReferral := repository.NewReferralRepository(db)
	repoUser := repository.NewUserRepository(db)
	flowReferral := flow.NewReferralFlow(repoReferral)
	flowUser := flow.NewUserFlow(repoUser, repoReferral, flowReferral)
	handler := NewUserHandler(parser, presenter, flowUser)

	router.HandleFunc("/user", handler.Register).Methods("PUT")

	server := httptest.NewServer(router)
	defer server.Close()

	requestData := entity.User{
		Email:    "mail@mail.com",
		Password: "12345",
		UserName: "username",
		Name:     "Name",
	}

	requestBody, _ := json.Marshal(requestData)
	req, err := http.NewRequest("PUT", server.URL+"/user", bytes.NewBuffer(requestBody))
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

func TestUserHandler_GenerateNewReferral_E2E(t *testing.T) {
	router := mux.NewRouter()
	db := repository.NewConnectMysqlDb()
	parser := parser.NewUserParser()
	presenter := util.NewJsonPresenter()
	repoReferral := repository.NewReferralRepository(db)
	repoUser := repository.NewUserRepository(db)
	flowReferral := flow.NewReferralFlow(repoReferral)
	flowUser := flow.NewUserFlow(repoUser, repoReferral, flowReferral)
	handler := NewUserHandler(parser, presenter, flowUser)

	router.HandleFunc("/generate/referral", handler.Register).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()

	requestData := entity.User{
		Email:    "mail@mail.com",
		UserName: "username",
	}

	requestBody, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", server.URL+"/generate/referral", bytes.NewBuffer(requestBody))
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
