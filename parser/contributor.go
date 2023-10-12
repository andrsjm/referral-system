package parser

import (
	"fmt"
	"net/http"
	"referral-system/entity"
)

type contributorParser struct{}

func NewContributorParser() IContributorParser {
	return &contributorParser{}
}

func (p *contributorParser) ParseContributorEntity(r *http.Request) (entity.Contributor, error) {
	req := clientRequest{r}

	contributor := entity.Contributor{}

	contributor.ReferralCode = req.pathParam("code")
	contributor.Email = req.param("email")

	if contributor.ReferralCode == "" || contributor.Email == "" {
		return contributor, fmt.Errorf("error")
	}

	return contributor, nil
}
