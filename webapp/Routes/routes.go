package Routes

import (
	"log"
	"net/http"
	"webapp/Controller"

	// "webapp/Controller"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", Controller.SignUp).Methods("POST")
	router.HandleFunc("/login", Controller.Login).Methods("POST")
	router.HandleFunc("/users", Controller.GetALlUsers).Methods("GET")
	router.HandleFunc("/users/{id}", Controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", Controller.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", Controller.DeleteUser).Methods("DELETE")

	//due routes
	router.HandleFunc("/due", Controller.CreateDue).Methods("POST")
	router.HandleFunc("/due", Controller.GetAllDue).Methods("GET")
	router.HandleFunc("/due/{id}", Controller.UpdateDue).Methods("PUT")
	router.HandleFunc("/due/{id}", Controller.DeleteDue).Methods("DELETE")
	//Nodue
	router.HandleFunc("/nodue", Controller.CreatNODue).Methods("POST")
	router.HandleFunc("/nodue", Controller.GetAllNoDue).Methods("GET")
	router.HandleFunc("/nodue/{id}", Controller.DeleteNoDue).Methods("DELETE")

	fhandler := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
