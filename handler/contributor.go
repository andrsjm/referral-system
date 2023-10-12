package handler

import (
	"net/http"
	"referral-system/flow"
	"referral-system/parser"
	"referral-system/util"
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
