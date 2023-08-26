package handlers

import (
   "encoding/json"
   "fmt"
   "net/http"
   "errors"
   "avito_test/internals/services"
   "avito_test/internals/models"
   "github.com/gorilla/mux"
//    "strconv"
//    "github.com/sirupsen/logrus"
)

type Handler struct{
  service *services.Service
}

func NewHandler(Service *services.Service) *Handler{
	resultHandler := new(Handler)
	resultHandler.service = Service
	return resultHandler
}

func (h *Handler) FindUserSegment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 
	if vars["user_id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}


	segment, err := h.service.FindUserSegment(vars["user_id"])
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : segment,
	}

	WrapOK(w, m)
}
func (h *Handler) NewUserSegment(w http.ResponseWriter, r *http.Request){
	var newCar models.UserSegment

	err := json.NewDecoder(r.Body).Decode(&newCar)  
	if err != nil {
		WrapError(w, err) 
		return
	}

	err = h.service.NewUserSegment(newCar) 
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {  
		"result" : "OK",
		"data" : "",
	}

	WrapOK(w, m) //здесь возвращаем код 200 и успех
}

func (h *Handler) NewSegment(w http.ResponseWriter, r *http.Request){
	// vars := mux.Vars(r) 
	// if vars["order_uid"] == "" {
	// 	WrapError(w, errors.New("missing id"))
	// 	return
	// }


	// order, err := h.service.GetOrder(vars["order_uid"])
	// if err != nil {
	// 	WrapError(w, err)
	// 	return
	// }

	// var m = map[string]interface{} {
	// 	"result" : "OK",
	// 	"data" : order,
	// }

	// WrapOK(w, m)
}


func (h *Handler) DeleteSegment(w http.ResponseWriter, r *http.Request){
	// vars := mux.Vars(r) 
	// if vars["order_uid"] == "" {
	// 	WrapError(w, errors.New("missing id"))
	// 	return
	// }


	// order, err := h.service.GetOrder(vars["order_uid"])
	// if err != nil {
	// 	WrapError(w, err)
	// 	return
	// }

	// var m = map[string]interface{} {
	// 	"result" : "OK",
	// 	"data" : order,
	// }

	// WrapOK(w, m)
}

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string {
		"result" : "error",
		"data" : err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")  
	w.WriteHeader(httpStatus) 
	fmt.Fprintln(w, string(res))
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	WrapErrorWithStatus(w, errors.New("not found"), http.StatusNotFound)
}