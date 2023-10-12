package repository

import (
	"referral-system/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Login(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewUserRepository(db)

	user := entity.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	mock.ExpectQuery("SELECT * FROM users WHERE email=? and password=?").
		WithArgs(user.Email, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).
			AddRow(user.Email, user.Password))

	resultUser, err := repo.Login(user)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, user, resultUser, "Expected user to match, but got a different user")
}

func TestUserRepository_Register(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewUserRepository(db)

	user := entity.User{
		Email:    "test@example.com",
		Password: "password123",
		UserName: "testuser",
		Name:     "Test User",
	}

	mock.ExpectExec("INSERT INTO users(email, password, username, name, referral_hit, user_type) VALUES(?, ?, ?, ?, 0, 1)").
		WithArgs(user.Email, user.Password, user.UserName, user.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := repo.Register(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewUserRepository(db)

	user := entity.User{
		Email:    "test@example.com",
		Password: "123",
		Name:     "Test User",
	}

	mock.ExpectExec("UPDATE users SET email=?, password=?, name=? WHERE email=?").
		WithArgs(user.Email, user.Password, user.Name, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_UpdateReferralHit(t *testing.T) {
	SetEnvMysqlConnection(t)
	db, mock := OpenSqlAndMock()
	defer db.Close()

	repo := NewUserRepository(db)

	ID := 5

	mock.ExpectExec("UPDATE users SET referral_hit = referral_hit + 1 WHERE id = ?").
		WithArgs(ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.UpdateReferralHit(ID)
	if err != nil {
		t.Error(err)
	}
}
