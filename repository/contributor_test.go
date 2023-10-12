package repository

import (
	"referral-system/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestContributorRepository_CheckContributor(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewContributorRepository(db)

	mock.ExpectQuery("SELECT EXISTS(SELECT * FROM contributors WHERE email = ?)").WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"EXISTS"}).AddRow(true))

	isReady, err := repo.CheckContributor("test@example.com")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, isReady, "Expected contributor to be ready, but got false")
}

func TestContributorRepository_InsertContributor(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewContributorRepository(db)

	mock.ExpectExec("INSERT INTO contributors(referral_code, email) VALUES(?, ?)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	contributor := entity.Contributor{
		ReferralCode: "ref123",
		Email:        "test@example.com",
	}

	err := repo.InsertContributor(contributor)
	if err != nil {
		t.Error(err)
	}
}
