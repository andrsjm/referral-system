package parser

import (
	"net/http"
	"referral-system/entity"
	"referral-system/util"
)

type userParser struct{}

func NewUserParser() IUserParser {
	return &userParser{}
}

func (p *userParser) ParseUserEntity(r *http.Request) (entity.User, error) {
	user := entity.User{}

	err := jsonParser(r, &user)
	if err != nil {
		return user, err
	}

	user.Password = HashPassword(user.Password)

	return user, nil
}

func (p *userParser) ParseBlogID(r *http.Request) int {
	req := clientRequest{r}
	return req.getEntityID()
}

func (p *userParser) ParseUserFromCoockies(r *http.Request) entity.User {
	var id int
	var username string

	id = util.Getid(r)
	username = util.GetUsername(r)

	user := entity.User{
		ID:       id,
		UserName: username,
	}

	return user
}
