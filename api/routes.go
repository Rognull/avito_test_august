package api

import (
	"avito_test/internals/handlers"
	"github.com/gorilla/mux"
)

func CreateRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()  
	r.HandleFunc("/segment/find",h.FindUserSegment).Methods("POST")
	r.HandleFunc("/user_segment",h.NewUserSegment).Methods("POST")
	r.HandleFunc("/segment",h.NewSegment).Methods("POST")
	r.HandleFunc("/segment/delete",h.DeleteSegment).Methods("POST")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() 
	return r
}