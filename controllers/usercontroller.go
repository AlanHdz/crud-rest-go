package controllers

import (
	"alanhedz/golang-crud/database"
	"alanhedz/golang-crud/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.Instance.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]

	var user models.User
	resultFirst := database.Instance.First(&user, userId)

	if resultFirst.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultSave := database.Instance.Save(&user)

	if resultSave.Error != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	var user models.User
	resultFrist := database.Instance.First(&user, userId)

	if resultFrist.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	database.Instance.Delete(&user, userId)
	json.NewEncoder(w).Encode("Product Deleted succesfully")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	var user models.User

	resultFirst := database.Instance.First(&user, userId)

	if resultFirst.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
