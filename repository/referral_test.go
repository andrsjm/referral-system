package repository

import (
	"referral-system/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestReferralRepository_InsertReferral(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewReferralRepository(db)

	referral := entity.Referral{
		ReferralCode: "wadidaw",
		ExpiredDate:  time.Now(),
		UserID:       1,
	}

	mock.ExpectExec("INSERT INTO referrals(referral_code, user_id, expired_date) VALUES(?, ?, ?)").
		WithArgs(referral.ReferralCode, referral.UserID, referral.ExpiredDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.InsertReferral(referral)
	if err != nil {
		t.Error(err)
	}
}

func TestReferralRepository_GetReferral(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewReferralRepository(db)

	referralCode := "referral123"

	expectedReferral := entity.Referral{
		ID:           1,
		ReferralCode: referralCode,
		ExpiredDate:  time.Now(),
		UserID:       1,
	}

	mock.ExpectQuery("SELECT * FROM referrals WHERE referral_code=?").
		WithArgs(referralCode).
		WillReturnRows(sqlmock.NewRows([]string{"id", "referral_code", "expired_date", "user_id"}).
			AddRow(expectedReferral.ID, expectedReferral.ReferralCode, expectedReferral.ExpiredDate, expectedReferral.UserID))

	resultReferral, err := repo.GetReferral(referralCode)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedReferral, resultReferral, "Expected referral to match, but got a different referral")
}
