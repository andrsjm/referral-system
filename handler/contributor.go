package handler

import (
	"net/http"
	"referral-system/flow"
	"referral-system/parser"
	"referral-system/util"

	_ "referral-system/docs"
)

type contributorHandler struct {
	parser    parser.IContributorParser
	presenter util.IPresenterJSON
	flow      flow.IContributorFlow
}

func NewContributorHandler(parser parser.IContributorParser, presenter util.IPresenterJSON, flow flow.IContributorFlow) *contributorHandler {
	return &contributorHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

// Referral handles the referral endpoint.
// @Summary Handle referral requests
// @Description This endpoint handles referral requests. It requires a path parameter 'code' and a query parameter 'email'.
// @ID referral
// @Accept json
// @Produce json
// @Param code path string true "Referral code"
// @Param email query string true "User email"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /referral/{code} [post]
func (h *contributorHandler) Referral(w http.ResponseWriter, r *http.Request) {
	contributor, err := h.parser.ParseContributorEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	isReady, expired, err := h.flow.Referral(contributor)
	if err != nil {
		h.presenter.SendError(w, "Error Hit Referral")
		return
	}

	if isReady {
		h.presenter.SendError(w, "User has used the referral")
		return
	}

	if expired {
		h.presenter.SendError(w, "The referral has expired")
		return
	}

	h.presenter.SendSuccess(w)
}
