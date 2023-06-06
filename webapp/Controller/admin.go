package Controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/Model"
	"webapp/Utlis/httpResp"

	"github.com/gorilla/mux"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var admin Model.Admin

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "user added"})

}

func Login(w http.ResponseWriter, r *http.Request) {
	var use Model.Admin
	err := json.NewDecoder(r.Body).Decode(&use)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := use.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}
	response := struct {
		Message string      `json:"message"`
		Data    Model.Admin `json:"data"`
	}{
		Message: "Success",
		Data:    use,
	}

	httpResp.RespondWithJSON(w, http.StatusOK, response)
}

func GetALlUsers(w http.ResponseWriter, r *http.Request) {
	users, getErr := Model.GetAllUsers()

	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, users)
}

func convertIDToInt64(id string) (int64, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return intID, nil
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	uid, err := convertIDToInt64(id)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user Model.Admin

	// Decode the JSON body of the request into the user struct.
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	// Set the ID of the user to the parsed ID.
	user.ID = uid

	// Call the UpdateUser method to update the user in the database.
	err = user.UpdateUser(uid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// If the user is not found, respond with a "Not Found" error message.
			httpResp.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			// If there is an internal server error, respond with the error message.
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		// Respond with a success message.
		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, idErr := convertIDToInt64(id)

	if idErr != nil {
		// If there is an error parsing the user ID, respond with an error message.
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	user := Model.Admin{ID: userId}

	// Call the ReadUser method of the user model to retrieve the user from the database.
	getErr := user.ReadUser()

	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			// If the user is not found, respond with a "Not Found" error message.
			httpResp.RespondWithError(w, http.StatusNotFound, "User Not Found")
		default:
			// If there is an internal server error, respond with the error message.
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
			return
		}
	}

	// Respond with the retrieved user in JSON format.
	httpResp.RespondWithJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID, err := convertIDToInt64(id)

	if err != nil {
		// If there is an error parsing the user ID, respond with an error message.
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := Model.Admin{ID: userID}

	// Call the DeleteUser method of the user model to delete the user from the database.
	if err := user.DeleteUser(); err != nil {
		// If there is an error while deleting the user, respond with an error message.
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Respond with a JSON message indicating the successful deletion of the user.
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Deleted"})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	uid, err := convertIDToInt64(id)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user Model.Admin

	// Decode the JSON body of the request into the user struct.
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	// Set the ID of the user to the parsed ID.
	user.ID = uid

	// Call the UpdatePassword method to update the user's password in the database.
	err = user.UpdatePassword(uid)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			// If the user is not found, respond with a "Not Found" error message.
			httpResp.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			// If there is an internal server error, respond with the error message.
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Respond with a success message.
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
}
