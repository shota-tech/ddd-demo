package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shota-tech/layered-architecture-demo/src/domain/model"
	"github.com/shota-tech/layered-architecture-demo/src/usecase"
)

type UserHandler interface {
	GetUserByID(http.ResponseWriter, *http.Request)
	GetUserList(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_ := vars["id"]
	id, err := strconv.Atoi(id_)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := h.userUsecase.GetUserByID(id)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func (h *userHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetUserList()
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func (h *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id, err := h.userUsecase.AddUser(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res := fmt.Sprintf(`{"id": %s}`, strconv.Itoa(id))
	w.Write([]byte(res))
}
