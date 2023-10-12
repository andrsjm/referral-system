package parser

import (
	"net/http"
	"net/url"
	"referral-system/entity"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUserParser_ParseContributorEntity(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		contributor := entity.Contributor{
			ReferralCode: "12345",
			Email:        "wadidaw@gmail.com",
		}

		r := &http.Request{}
		r.URL = &url.URL{
			RawQuery: "email=wadidaw@gmail.com",
		}

		r = mux.SetURLVars(r, map[string]string{"code": "12345"})

		parser := NewContributorParser()
		result, err := parser.ParseContributorEntity(r)

		assert.Nil(t, err)
		assert.Equal(t, contributor, result)
	})

	t.Run("error", func(t *testing.T) {
		r := &http.Request{}
		r.URL = &url.URL{
			RawQuery: "email=wadidaw@gmail.com",
		}

		parser := NewContributorParser()
		_, err := parser.ParseContributorEntity(r)

		assert.NotNil(t, err)
	})

}
