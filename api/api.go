package api

import (
	"api-server/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is default welcome page")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is test page")
}
func players(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(database.Data)
}
func player(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var player database.Player

	json.NewDecoder(r.Body).Decode(&player)

	fmt.Println(player)
}

// Router
var router = mux.NewRouter()

func routerCalling() {

	//Calling handler functions by their path
	router.HandleFunc("/", welcome)
	router.HandleFunc("/test", test)
	router.HandleFunc("/api/players", players).Methods("GET")
	router.HandleFunc("/api/players/{id}", player).Methods("GET")
}
func StartServer() {
	log.Printf("-------------server started at port %d -------\n", 8080)
	routerCalling()

	//Server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
