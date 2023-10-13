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
	r.HandleFunc("/user", http.HandlerFunc(userHandler.Register)).Methods("POST")
	r.HandleFunc("/user", util.Authenticate(userHandler.Update, 1)).Methods("PUT")
	r.HandleFunc("/login", http.HandlerFunc(userHandler.Login)).Methods("POST")
	r.HandleFunc("/logout", util.Authenticate(userHandler.Logout, 1)).Methods("POST")
	r.HandleFunc("/generate/referral", util.Authenticate(userHandler.GenerateNewReferral, 1)).Methods("POST")

	//Contributor Routes
	r.HandleFunc("/referral/{code}", http.HandlerFunc(contributorHandler.Referral)).Methods("POST")

	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	return r
}
