package Controller

import (
	"encoding/json"
	"net/http"
	"webapp/Model"
	"webapp/Utlis/httpResp"

	"github.com/gorilla/mux"
)

func CreatNODue(w http.ResponseWriter, r *http.Request) {
	var due Model.NoDue

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
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "nodue added"})
}

func GetAllNoDue(w http.ResponseWriter, r *http.Request) {
	dues, err := Model.GetAllNoDue()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, dues)
}

func DeleteNoDue(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	dueID, idErr := convertIDToInt64(id)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	due := Model.NoDue{ID: dueID}

	if err := due.DeleteNoDue(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Deleted"})
}
