package Controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"webapp/Model"
	"webapp/Utlis/httpResp"

	"github.com/gorilla/mux"
)

func CreateDue(w http.ResponseWriter, r *http.Request) {
	var due Model.Due

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&due); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	defer r.Body.Close()

	saveErr := due.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	// No error, respond with success message
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "due added"})
}

func GetAllDue(w http.ResponseWriter, r *http.Request) {
	dues, err := Model.GetAllDue()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, dues)
}

func UpdateDue(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	did, err := convertIDToInt64(id)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var due Model.Due

	// Decode the JSON body of the request into the due struct.
	err = json.NewDecoder(r.Body).Decode(&due)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	// Set the ID of the due to the parsed ID.
	due.ID = did

	// Call the UpdateDue method to update the due in the database.
	err = due.UpdateDue()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// If the due is not found, respond with a "Not Found" error message.
			httpResp.RespondWithError(w, http.StatusNotFound, "Due not found")
		default:
			// If there is an internal server error, respond with the error message.
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		// Respond with a success message.
		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Due updated successfully"})
	}
}

func DeleteDue(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	dueID, idErr := convertIDToInt64(id)

	if idErr != nil {
		// If there is an error parsing the due ID, respond with an error message.
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	due := Model.Due{ID: dueID}

	// Call the DeleteDue method of the due model to delete the due from the database.
	if err := due.DeleteDue(); err != nil {
		// If there is an error while deleting the due, respond with an error message.
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Respond with a JSON message indicating the successful deletion of the due.
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Deleted"})
}
