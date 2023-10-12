package parser

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"referral-system/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserParser_ParseUserEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expect := entity.User{
			Email:    "mail@mail.com",
			Password: "827ccb0eea8a706c4c34a16891f84e7b",
			UserName: "username",
			Name:     "name",
		}

		jsonData := []byte(`{
			"email": "mail@mail.com",
			"password": "12345",
			"username": "username",
			"name": "name"
		}`)

		r := &http.Request{}
		r.Method = http.MethodPost
		r.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

		parser := NewUserParser()
		result, err := parser.ParseUserEntity(r)

		assert.Nil(t, err)
		assert.JSONEq(t, EntityStringify(expect), EntityStringify(result))
	})

	t.Run("error", func(t *testing.T) {
		expect := entity.User{}

		jsonData := []byte(``)

		r := &http.Request{}
		r.Method = http.MethodPost
		r.Body = ioutil.NopCloser(bytes.NewReader(jsonData))

		parser := NewUserParser()
		result, err := parser.ParseUserEntity(r)

		assert.NotNil(t, err)
		assert.JSONEq(t, EntityStringify(expect), EntityStringify(result))
	})
}

func TestUserParser_ParseBlogID(t *testing.T) {
	expect := 1

	r := &http.Request{}
	r.URL = &url.URL{
		RawQuery: "id=1",
	}

	parser := NewUserParser()
	result := parser.ParseBlogID(r)

	assert.Equal(t, expect, result)
}
