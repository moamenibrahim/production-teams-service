package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"production_service/models"
	"production_service/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var response models.UserResponse
	fmt.Println("Getting users...")
	
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		utils.CheckErr(w, r, err)
		return
	}

	var users []models.User
	for rows.Next() {
		var userID string
		var userName string
		var userRole string

		err = rows.Scan(&userID, &userName, &userRole)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		users = append(users, models.User{ID: userID, Name: userName, Role: userRole})
	}

	response = models.UserResponse{Type: "success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userid")
	userName := r.FormValue("username")
	userRole := r.FormValue("userrole")

	var response = models.UserResponse{}
	if userID == "" || userName == "" {
		response = models.UserResponse{Type: "error", Message: "You are missing userID or userName parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Inserting new user with ID: " + userID + " and name: " + userName + " and role: " + userRole)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO users(id, name, role) VALUES($1, $2, $3) returning id;", userID, userName, userRole).Scan(&lastInsertID)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		data := []models.User{}
		data = append(data, models.User{ID: userID, Name: userName, Role: userRole})
		response = models.UserResponse{Type: "success", Data: data, Message: "The user has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userid"]
	var response = models.UserResponse{}
	if userID == "" {
		response = models.UserResponse{Type: "error", Message: "You are missing userID parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Deleting user from DB")

		result, err := db.Exec("DELETE FROM users where id=$1", userID)
		rows, _ := result.RowsAffected()
		if rows == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound),
				http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		response = models.UserResponse{Type: "success", Message: "The users has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userid"]
	var response = models.UserResponse{}
	if userID == "" {
		response = models.UserResponse{Type: "error", Message: "You are missing userID parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Getting user from DB")

		var ID string
		var userName string
		var userRole string
		err := db.QueryRow("SELECT id, name, role FROM users WHERE name=$1", userID).Scan(&ID, &userName, &userRole)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		users := []models.User{}
		users = append(users, models.User{ID: ID, Name: userName, Role: userRole})
		response = models.UserResponse{Type: "success", Data: users}
	}
	json.NewEncoder(w).Encode(response)
}
