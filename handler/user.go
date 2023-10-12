package handler

import (
	"net/http"
	"referral-system/flow"
	"referral-system/parser"
	"referral-system/util"

	_ "referral-system/docs"
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

// Register handles user registration.
// @Summary Register a new user
// @Description This endpoint allows users to register with an email, password, and name.
// @ID register
// @Accept json
// @Produce json
// @Param user body entity.User true "User registration information"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [post]
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

// Login handles user login.
// @Summary User login
// @Description This endpoint allows users to log in with an email and password.
// @ID login
// @Accept json
// @Produce json
// @Param user body entity.User true "User login information"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /login [post]
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

// Logout handles user logout.
// @Summary User logout
// @Description This endpoint allows users to log out.
// @ID logout
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /logout [post]
func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	util.ResetUsersToken(w)

	h.presenter.SendSuccess(w)
}

// Update handles user profile updates.
// @Summary Update user profile
// @Description This endpoint allows users to update their profile information.
// @ID update
// @Accept json
// @Produce json
// @Param user body entity.User true "User update information"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [put]
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

// GenerateNewReferral generates a new referral code.
// @Summary Generate a new referral code
// @Description This endpoint generates a new referral code for the user.
// @ID generateReferral
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /generate/referral [post]
func (h *userHandler) GenerateNewReferral(w http.ResponseWriter, r *http.Request) {
	user := h.parser.ParseUserFromCoockies(r)

	referralLink, err := h.flow.GenerateNewCode(user)
	if err != nil {
		h.presenter.SendError(w, "Error Generate")
		return
	}

	h.presenter.SendSuccessWithReferral(w, referralLink)
}
