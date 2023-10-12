package flow

import (
	"fmt"
	"referral-system/entity"
	repoMock "referral-system/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContributorFlow_Referral(t *testing.T) {
	user := entity.User{
		ID:       1,
		Email:    "mail@gmail.com",
		Password: "827ccb0eea8a706c4c34a16891f84e7b",
		UserName: "username",
		Name:     "name",
	}

	contributor := entity.Contributor{
		ReferralCode: "12345",
		Email:        "mail@mail.com",
	}

	referral := entity.Referral{
		ReferralCode: "12345",
		ExpiredDate:  time.Now().Add(7 * 24 * time.Hour),
		UserID:       1,
	}

	testCases := []struct {
		name                 string
		param                entity.Contributor
		mockRepoCheckArgs    []interface{}
		mockRepoCheckReturn  []interface{}
		mockRepoGetArgs      []interface{}
		mockRepoGetReturn    []interface{}
		mockRepoInsertArgs   []interface{}
		mockRepoInsertReturn []interface{}
		mockRepoUpdateArgs   []interface{}
		mockRepoUpdateReturn []interface{}
		expectData           interface{}
		expectError          interface{}
	}{
		{
			name:                 "general error",
			param:                contributor,
			mockRepoCheckArgs:    []interface{}{contributor.Email},
			mockRepoCheckReturn:  []interface{}{true, fmt.Errorf("error")},
			mockRepoGetArgs:      []interface{}{referral},
			mockRepoGetReturn:    []interface{}{fmt.Errorf("error")},
			mockRepoInsertArgs:   []interface{}{user},
			mockRepoInsertReturn: []interface{}{referral, fmt.Errorf("error")},
			mockRepoUpdateArgs:   []interface{}{referral},
			mockRepoUpdateReturn: []interface{}{fmt.Errorf("error")},
			expectError:          fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepoUser := new(repoMock.MockUserRepo)
			mockRepoUser.On("UpdateReferralHit", tc.mockRepoUpdateArgs...).Return(tc.mockRepoUpdateReturn...)

			mockRepoReferral := new(repoMock.MockReferralRepo)
			mockRepoReferral.On("GetReferral", tc.mockRepoGetArgs...).Return(tc.mockRepoGetReturn...)

			mockRepoContributor := new(repoMock.MockContributorRepo)
			mockRepoContributor.On("CheckContributor", tc.mockRepoCheckArgs...).Return(tc.mockRepoCheckReturn...)
			mockRepoContributor.On("InsertContributor", tc.mockRepoInsertArgs...).Return(tc.mockRepoInsertReturn...)

			flow := NewContributorFlow(mockRepoContributor, mockRepoReferral, mockRepoUser)
			_, _, actualErr := flow.Referral(tc.param)

			assert.Equal(t, tc.expectError, actualErr)
		})
	}
}
