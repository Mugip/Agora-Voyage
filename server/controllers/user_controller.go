go
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/your-package/models"
	"github.com/your-package/utils"
)

// UserController represents the controller for user-related operations
type UserController struct {
	userModel  models.UserModel
}

// NewUserController initializes a new instance of UserController
func NewUserController(userModel models.UserModel) *UserController {
	return &UserController{
		userModel: userModel,
	}
}

// GetAllUsers retrieves all users
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Implement code to retrieve all users from the model
	users, err := uc.userModel.GetUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	id := // parse the ID from the request params, e.g., mux.Vars(r)["id"]

	// Implement code to retrieve the user by ID from the model
	user, err := uc.userModel.GetUserByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	// Decode the JSON request body into the newUser object
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Implement code to create the user in the model
	err = uc.userModel.CreateUser(&newUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, newUser)
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	id := // parse the ID from the request params, e.g., mux.Vars(r)["id"]

	var updatedUser models.User

	// Decode the JSON request body into the updatedUser object
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Implement code to update the user in the model
	err = uc.userModel.UpdateUser(id, &updatedUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedUser)
}

// DeleteUser deletes an existing user
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	id := // parse the ID from the request params, e.g., mux.Vars(r)["id"]

	// Implement code to delete the user from the model
	err := uc.userModel.DeleteUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
