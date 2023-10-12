package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func SetEnvMysqlConnection(t *testing.T) {
	t.Setenv("DB_CONNECTION", "mysql")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_NAME", "db_komunitas_mea")
	t.Setenv("DB_USERNAME", "root")
	t.Setenv("DB_PASSWORD", "")
}

func OpenSqlAndMock() (*sqlx.DB, sqlmock.Sqlmock) {
	options := sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)
	mockDB, mock, _ := sqlmock.New(options)

	sqlxDb := sqlx.NewDb(mockDB, "sqlmock")
	return sqlxDb, mock
}
