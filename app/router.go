package app

import (
	"net/http"
	"referral-system/flow"
	"referral-system/handler"
	"referral-system/parser"
	"referral-system/repository"
	"referral-system/util"

	_ "referral-system/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

var jsonPresenter = util.NewJsonPresenter()
var db = repository.NewConnectMysqlDb()

var userParser = parser.NewUserParser()
var userRepo = repository.NewUserRepository(db)
var userFlow = flow.NewUserFlow(userRepo, referralRepo, referralFlow)
var userHandler = handler.NewUserHandler(userParser, jsonPresenter, userFlow)

var referralRepo = repository.NewReferralRepository(db)
var referralFlow = flow.NewReferralFlow(referralRepo)

var contributorParser = parser.NewContributorParser()
var contributorRepo = repository.NewContributorRepository(db)
var contributorFlow = flow.NewContributorFlow(contributorRepo, referralRepo, userRepo)
var contributorHandler = handler.NewContributorHandler(contributorParser, jsonPresenter, contributorFlow)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	//User Routes
	// r.Handle("/user", UserTypeMiddleware([]int{}, http.HandlerFunc(userHandler.Register))).Methods("POST")
	r.HandleFunc("/user", http.HandlerFunc(userHandler.Register)).Methods("POST")
	// r.Handle("/user", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(userHandler.Update))).Methods("PUT")
	r.HandleFunc("/user", util.Authenticate(userHandler.Update, 1)).Methods("PUT")
	// r.Handle("/login", UserTypeMiddleware([]int{}, http.HandlerFunc(userHandler.Login))).Methods("POST")
	r.HandleFunc("/login", http.HandlerFunc(userHandler.Login)).Methods("POST")
	// r.Handle("/logout", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(userHandler.Logout))).Methods("POST")
	r.HandleFunc("/logout", util.Authenticate(userHandler.Logout, 1)).Methods("POST")
	// r.Handle("/generate/referral", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(userHandler.GenerateNewReferral))).Methods("POST")
	r.HandleFunc("/generate/referral", util.Authenticate(userHandler.GenerateNewReferral, 1)).Methods("POST")

	//Contributor Routes
	// r.Handle("/referral/{code}", UserTypeMiddleware([]int{}, http.HandlerFunc(contributorHandler.Referral))).Methods("POST")
	r.HandleFunc("/referral/{code}", http.HandlerFunc(contributorHandler.Referral)).Methods("POST")

	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	return r
}
