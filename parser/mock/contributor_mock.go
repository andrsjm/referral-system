package mock

import (
	"net/http"
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockContributorParser struct {
	mock.Mock
}

func (m *MockContributorParser) ParseContributorEntity(r *http.Request) (entity.Contributor, error) {
	args := m.Called(r)
	return args.Get(0).(entity.Contributor), args.Error(1)
}
