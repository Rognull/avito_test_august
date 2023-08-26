package api

import (
	"avito_test/internals/handlers"
	"github.com/gorilla/mux"
)

func CreateRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()  
	r.HandleFunc("/user_segment/find/{user_id:[0-9]+}",h.FindUserSegment).Methods("GET")
	r.HandleFunc("/user_segment",h.NewUserSegment).Methods("POST")
	r.HandleFunc("/segment",h.NewSegment).Methods("POST")
	r.HandleFunc("/segment",h.DeleteSegment).Methods("DELETE")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() 
	return r
}