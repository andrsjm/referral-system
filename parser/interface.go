package parser

import (
	"net/http"
	"referral-system/entity"
)

type IUserParser interface {
	ParseUserEntity(r *http.Request) (entity.User, error)
	ParseBlogID(r *http.Request) int
	ParseUserFromCoockies(r *http.Request) entity.User
}

type IContributorParser interface {
	ParseContributorEntity(r *http.Request) (entity.Contributor, error)
}
