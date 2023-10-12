package flow

import (
	"fmt"
	"referral-system/entity"
	flowMock "referral-system/flow/mock"
	repoMock "referral-system/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserFlow_GenerateNewCode(t *testing.T) {
	user := entity.User{
		ID:       1,
		Email:    "mail@gmail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	referral := entity.Referral{
		ReferralCode: "12345",
		ExpiredDate:  time.Now(),
		UserID:       1,
	}

	testCases := []struct {
		name                   string
		param                  entity.User
		mockFlowGenerateArgs   []interface{}
		mockFlowGenerateReturn []interface{}
		mockRepoInsertArgs     []interface{}
		mockRepoInsertReturn   []interface{}
		expectData             interface{}
		expectError            interface{}
	}{
		{
			name:                   "general error",
			param:                  user,
			mockFlowGenerateArgs:   []interface{}{user},
			mockFlowGenerateReturn: []interface{}{referral, fmt.Errorf("error")},
			mockRepoInsertArgs:     []interface{}{referral},
			mockRepoInsertReturn:   []interface{}{fmt.Errorf("error")},
			expectError:            fmt.Errorf("error"),
		},
		{
			name:                   "success",
			param:                  user,
			mockFlowGenerateArgs:   []interface{}{user},
			mockFlowGenerateReturn: []interface{}{referral, nil},
			mockRepoInsertArgs:     []interface{}{referral},
			mockRepoInsertReturn:   []interface{}{nil},
			expectError:            nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepoUser := new(repoMock.MockUserRepo)

			mockRepoReferral := new(repoMock.MockReferralRepo)
			mockRepoReferral.On("InsertReferral", tc.mockRepoInsertArgs...).Return(tc.mockRepoInsertReturn...)

			mockFlowReferral := new(flowMock.MockReferralFlow)
			mockFlowReferral.On("GenerateReferralCode", tc.mockFlowGenerateArgs...).Return(tc.mockFlowGenerateReturn...)

			flow := NewUserFlow(mockRepoUser, mockRepoReferral, mockFlowReferral)
			_, actualErr := flow.GenerateNewCode(tc.param)

			assert.Equal(t, tc.expectError, actualErr)
		})
	}
}

func TestUserFlow_Update(t *testing.T) {
	data := entity.User{
		ID:       1,
		Email:    "mail@gmail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	testCases := []struct {
		name           string
		param          entity.User
		mockRepoArgs   []interface{}
		mockRepoReturn []interface{}
		expectData     interface{}
		expectError    interface{}
	}{
		{
			name:           "general error",
			param:          data,
			mockRepoArgs:   []interface{}{data},
			mockRepoReturn: []interface{}{fmt.Errorf("error")},
			expectError:    fmt.Errorf("error"),
		},
		{
			name:           "success",
			param:          data,
			mockRepoArgs:   []interface{}{data},
			mockRepoReturn: []interface{}{nil},
			expectError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepoUser := new(repoMock.MockUserRepo)
			mockRepoUser.On("Update", tc.mockRepoArgs...).Return(tc.mockRepoReturn...)

			mockRepoReferral := new(repoMock.MockReferralRepo)
			mockFlowReferral := new(flowMock.MockReferralFlow)

			flow := NewUserFlow(mockRepoUser, mockRepoReferral, mockFlowReferral)
			actualErr := flow.Update(tc.param)

			assert.Equal(t, tc.expectError, actualErr)
		})
	}
}
