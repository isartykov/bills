package handler

import (
	"github.com/gorilla/mux"
	"github.com/isartykov/user/pkg/service"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/isartykov/user/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	router.HandleFunc("/auth/sign-in", h.signIn).Methods("POST")

	s := router.PathPrefix("/api").Subrouter()

	s.HandleFunc("/bills/calculate", h.userIdentity(h.calculate))
	s.HandleFunc("/bills/random", h.userIdentity(h.random)).Methods("GET")

	s.HandleFunc("/bills", h.userIdentity(h.createBill)).Methods("POST")
	s.HandleFunc("/bills", h.userIdentity(h.getAllBills)).Methods("GET")
	s.HandleFunc("/bills/{id}", h.userIdentity(h.getBill)).Methods("GET")
	s.HandleFunc("/bills/{id}", h.userIdentity(h.updateBill)).Methods("PUT")
	s.HandleFunc("/bills/{id}", h.userIdentity(h.deleteBill)).Methods("DELETE")

	return router
}
