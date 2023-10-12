package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"referral-system/entity"
	flowMock "referral-system/flow/mock"
	parserMock "referral-system/parser/mock"
	presenterMock "referral-system/util/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContributorHandler_Referral(t *testing.T) {
	data := entity.Contributor{
		ReferralCode: "12345",
		Email:        "mail@mail.com",
	}

	testCases := []struct {
		name             string
		mockParserReturn []interface{}
		mockFlowArgs     []interface{}
		mockFlowReturn   []interface{}
		payload          url.Values
		expectCode       int
		expectMessage    string
	}{
		{
			name: "error parsing data",
			mockParserReturn: []interface{}{
				entity.Contributor{},
				fmt.Errorf("parse form error"),
			},
			expectCode:    400,
			expectMessage: "Error Parsing",
		},
		{
			name:             "general error",
			mockFlowArgs:     []interface{}{data},
			mockFlowReturn:   []interface{}{false, false, fmt.Errorf("error")},
			mockParserReturn: []interface{}{data, nil},
			expectCode:       400,
			expectMessage:    "Error Hit Referral",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFlow := new(flowMock.MockContributorFlow)
			mockFlow.On("Referral", tc.mockFlowArgs...).Return(tc.mockFlowReturn...)

			var payload io.Reader

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/api/admin", payload)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			mockParser := new(parserMock.MockContributorParser)
			mockParser.On("ParseContributorEntity", r).Return(tc.mockParserReturn...)

			mockPresenter := new(presenterMock.MockJsonPresenter)
			if tc.expectCode == 200 {
				mockPresenter.On("SendSuccess", w, tc.expectMessage)
			} else {
				mockPresenter.On("SendError", w, tc.expectMessage)
			}

			adminFlow := NewContributorHandler(mockParser, mockPresenter, mockFlow)
			adminFlow.Referral(w, r)

			res := getBodyResponse(w)

			assert.Equal(t, res.Status, tc.expectCode)
			assert.Equal(t, res.Message, tc.expectMessage)
		})
	}
}
