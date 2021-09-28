package handler

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shota-tech/ddd-demo/domain/model"
	"github.com/shota-tech/ddd-demo/interfaces/dto"
	"github.com/shota-tech/ddd-demo/usecase"
)

type UserHandler interface {
	GetUserByID(http.ResponseWriter, *http.Request)
	GetUserList(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
	EditUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
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
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userResponse := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	res, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *userHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetUserList()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponse := dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	res, err := json.Marshal(userResponses)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userRequest dto.UserRequest
	err := json.Unmarshal(body, &userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := model.NewUser(userRequest.Name, userRequest.Email)

	id, err := h.userUsecase.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	url := path.Join(r.Host, r.URL.Path, strconv.Itoa(id))
	w.Header().Set("Location", url)
	w.WriteHeader(201)
}

func (h *userHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userRequest dto.UserRequest
	err = json.Unmarshal(body, &userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := model.NewUser(userRequest.Name, userRequest.Email)

	err = h.userUsecase.EditUser(id, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.userUsecase.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}
