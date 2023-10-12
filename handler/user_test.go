package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"referral-system/entity"
	flowMock "referral-system/flow/mock"
	parserMock "referral-system/parser/mock"
	presenterMock "referral-system/util/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Register(t *testing.T) {
	data := entity.User{
		Email:    "mail@mail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	referralLink := "localhost:8000/referral/code"

	testCases := []struct {
		name             string
		mockFlowArgs     []interface{}
		mockFlowReturn   []interface{}
		mockParserReturn []interface{}
		expectCode       int
		expectMessage    string
	}{
		{
			name: "error parsing data",
			mockParserReturn: []interface{}{
				entity.User{},
				fmt.Errorf("parse form error"),
			},
			expectCode:    400,
			expectMessage: "Error Parsing",
		},
		{
			name: "error",
			mockParserReturn: []interface{}{
				data, nil,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{
				referralLink, fmt.Errorf("error"),
			},
			expectCode:    400,
			expectMessage: "Error Insert",
		},
		{
			name: "success",
			mockParserReturn: []interface{}{
				data, nil,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{referralLink, nil},
			expectCode:     200,
			expectMessage:  "Success.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFlow := new(flowMock.MockUserFlow)
			mockFlow.On("Register", tc.mockFlowArgs...).Return(tc.mockFlowReturn...)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/category", nil)

			mockParser := new(parserMock.MockUserParser)
			mockParser.On("ParseUserEntity", r).Return(tc.mockParserReturn...)

			mockPresenter := new(presenterMock.MockJsonPresenter)
			if tc.expectCode == 200 {
				mockPresenter.On("SendSuccess", w, tc.expectMessage)
			} else {
				mockPresenter.On("SendError", w, tc.expectMessage)
			}

			categoryHandler := NewUserHandler(mockParser, mockPresenter, mockFlow)
			categoryHandler.Register(w, r)

			res := getBodyResponse(w)

			assert.Equal(t, res.Status, tc.expectCode)
			assert.Equal(t, res.Message, tc.expectMessage)
		})
	}
}

func TestUserHandler_Login(t *testing.T) {
	dataLogin := entity.User{
		Email:    "mail@mail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
	}

	data := entity.User{
		Email:    "mail@mail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	testCases := []struct {
		name             string
		mockFlowArgs     []interface{}
		mockFlowReturn   []interface{}
		mockParserReturn []interface{}
		expectCode       int
		expectMessage    string
	}{
		{
			name: "error parsing data",
			mockParserReturn: []interface{}{
				entity.User{},
				fmt.Errorf("parse form error"),
			},
			expectCode:    400,
			expectMessage: "Error Parsing",
		},
		{
			name: "error",
			mockParserReturn: []interface{}{
				dataLogin, nil,
			},
			mockFlowArgs: []interface{}{
				dataLogin,
			},
			mockFlowReturn: []interface{}{
				data, fmt.Errorf("error"),
			},
			expectCode:    400,
			expectMessage: "Error Login",
		},
		{
			name: "success",
			mockParserReturn: []interface{}{
				dataLogin, nil,
			},
			mockFlowArgs: []interface{}{
				dataLogin,
			},
			mockFlowReturn: []interface{}{data, nil},
			expectCode:     200,
			expectMessage:  "Success.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFlow := new(flowMock.MockUserFlow)
			mockFlow.On("Login", tc.mockFlowArgs...).Return(tc.mockFlowReturn...)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/category", nil)

			mockParser := new(parserMock.MockUserParser)
			mockParser.On("ParseUserEntity", r).Return(tc.mockParserReturn...)

			mockPresenter := new(presenterMock.MockJsonPresenter)
			if tc.expectCode == 200 {
				mockPresenter.On("SendSuccess", w, tc.expectMessage)
			} else {
				mockPresenter.On("SendError", w, tc.expectMessage)
			}

			categoryHandler := NewUserHandler(mockParser, mockPresenter, mockFlow)
			categoryHandler.Login(w, r)

			res := getBodyResponse(w)

			assert.Equal(t, res.Status, tc.expectCode)
			assert.Equal(t, res.Message, tc.expectMessage)
		})
	}
}

func TestUserHandler_Update(t *testing.T) {
	data := entity.User{
		Email:    "mail@mail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	testCases := []struct {
		name             string
		mockFlowArgs     []interface{}
		mockFlowReturn   []interface{}
		mockParserReturn []interface{}
		expectCode       int
		expectMessage    string
	}{
		{
			name: "error parsing data",
			mockParserReturn: []interface{}{
				entity.User{},
				fmt.Errorf("parse form error"),
			},
			expectCode:    400,
			expectMessage: "Error Parsing",
		},
		{
			name: "error",
			mockParserReturn: []interface{}{
				data, nil,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{
				fmt.Errorf("error"),
			},
			expectCode:    400,
			expectMessage: "Error Update",
		},
		{
			name: "success",
			mockParserReturn: []interface{}{
				data, nil,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{nil},
			expectCode:     200,
			expectMessage:  "Success.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFlow := new(flowMock.MockUserFlow)
			mockFlow.On("Update", tc.mockFlowArgs...).Return(tc.mockFlowReturn...)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/category", nil)

			mockParser := new(parserMock.MockUserParser)
			mockParser.On("ParseUserEntity", r).Return(tc.mockParserReturn...)

			mockPresenter := new(presenterMock.MockJsonPresenter)
			if tc.expectCode == 200 {
				mockPresenter.On("SendSuccess", w, tc.expectMessage)
			} else {
				mockPresenter.On("SendError", w, tc.expectMessage)
			}

			categoryHandler := NewUserHandler(mockParser, mockPresenter, mockFlow)
			categoryHandler.Update(w, r)

			res := getBodyResponse(w)

			assert.Equal(t, res.Status, tc.expectCode)
			assert.Equal(t, res.Message, tc.expectMessage)
		})
	}
}

func TestUserHandler_GenerateNewReferral(t *testing.T) {
	data := entity.User{
		Email:    "mail@mail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	referralLink := "localhost:8000/referral/code"

	testCases := []struct {
		name             string
		mockFlowArgs     []interface{}
		mockFlowReturn   []interface{}
		mockParserReturn []interface{}
		expectCode       int
		expectMessage    string
	}{
		{
			name: "error",
			mockParserReturn: []interface{}{
				data,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{
				referralLink, fmt.Errorf("error"),
			},
			expectCode:    400,
			expectMessage: "Error Generate",
		},
		{
			name: "success",
			mockParserReturn: []interface{}{
				data,
			},
			mockFlowArgs: []interface{}{
				data,
			},
			mockFlowReturn: []interface{}{referralLink, nil},
			expectCode:     200,
			expectMessage:  "Success.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFlow := new(flowMock.MockUserFlow)
			mockFlow.On("GenerateNewCode", tc.mockFlowArgs...).Return(tc.mockFlowReturn...)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/category", nil)

			mockParser := new(parserMock.MockUserParser)
			mockParser.On("ParseUserFromCoockies", r).Return(tc.mockParserReturn...)

			mockPresenter := new(presenterMock.MockJsonPresenter)
			if tc.expectCode == 200 {
				mockPresenter.On("SendSuccess", w, tc.expectMessage)
			} else {
				mockPresenter.On("SendError", w, tc.expectMessage)
			}

			categoryHandler := NewUserHandler(mockParser, mockPresenter, mockFlow)
			categoryHandler.GenerateNewReferral(w, r)

			res := getBodyResponse(w)

			assert.Equal(t, res.Status, tc.expectCode)
			assert.Equal(t, res.Message, tc.expectMessage)
		})
	}
}
