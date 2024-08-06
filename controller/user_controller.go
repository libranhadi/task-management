package controller

import (
	"encoding/json"
	"net/http"
	"task-management/service"
	"task-management/utils"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var userRequest map[string]string

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	email, emailOk := userRequest["email"]
	password, passwordOk := userRequest["password"]

	if !emailOk || !passwordOk {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	err := c.service.Register(email, password, userRequest["name"])
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	email, emailOk := user["email"]
	password, passwordOk := user["password"]

	if !emailOk || !passwordOk {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	authenticatedUser, err := c.service.Login(email, password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Account not found")
		return
	}

	token := utils.SignToken(authenticatedUser.Email)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Login successfully",
		"token":   token,
	})
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := utils.ExtractBearerToken(r)
	authenticatedUser, err := c.service.Profile(tokenString)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	responseUser := authenticatedUser.ToResponse()
	json.NewEncoder(w).Encode(map[string]any{
		"user": responseUser,
	})
}
