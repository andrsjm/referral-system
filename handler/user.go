package handler

import (
	"net/http"
	"referral-system/flow"
	"referral-system/parser"
	"referral-system/util"
)

type userHandler struct {
	parser    parser.IUserParser
	presenter util.IPresenterJSON
	flow      flow.IUserFlow
}

func NewUserHandler(parser parser.IUserParser, presenter util.IPresenterJSON, flow flow.IUserFlow) *userHandler {
	return &userHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	referralLink, err := h.flow.Register(user)
	if err != nil {
		h.presenter.SendError(w, "Error Insert")
		return
	}

	h.presenter.SendSuccessWithReferral(w, referralLink)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	user, err = h.flow.Login(user)
	if err != nil {
		h.presenter.SendError(w, "Error Login")
		return
	}

	util.GenerateToken(w, user.ID, user.UserName, 1)

	h.presenter.SendSuccess(w)
}

func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	util.ResetUsersToken(w)

	h.presenter.SendSuccess(w)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Update(user)
	if err != nil {
		h.presenter.SendError(w, "Error Update")
		return
	}

	h.presenter.SendSuccess(w)
}

func (h *userHandler) GenerateNewReferral(w http.ResponseWriter, r *http.Request) {
	user := h.parser.ParseUserFromCoockies(r)

	referralLink, err := h.flow.GenerateNewCode(user)
	if err != nil {
		h.presenter.SendError(w, "Error Generate")
		return
	}

	h.presenter.SendSuccessWithReferral(w, referralLink)
}
