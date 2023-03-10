package api

import (
	JWTauth "api-server/auth"
	"api-server/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	validToken, err := JWTauth.GenerateJWT()
	if err != nil {
		fmt.Fprint(w, "Error Generating token", err.Error())
	}

	fmt.Fprint(w, "This is default welcome page\n\n\nToken\n", validToken)

}
func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is test page")
}
func players(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(database.Data)
}
func player(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queries := mux.Vars(r)

	for _, player := range database.Data {
		if player.ID == queries["id"] {

			json.NewEncoder(w).Encode(player)
			w.WriteHeader(http.StatusOK)

			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
func addPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPlayer database.Player
	json.NewDecoder(r.Body).Decode(&newPlayer)

	database.Data = append(database.Data, newPlayer)

	fmt.Fprint(w, "New player added\n", newPlayer)
	w.WriteHeader(http.StatusCreated)
}
func updatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedUser database.Player
	json.NewDecoder(r.Body).Decode(&updatedUser)

	for i, player := range database.Data {
		if player.ID == params["id"] {

			updatedUser.ID = player.ID
			database.Data[i] = updatedUser

			w.WriteHeader(http.StatusAccepted)
			fmt.Fprint(w, "Updated\n", database.Data[i])

			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
func deletePlayer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	queries := mux.Vars(r)

	for i, player := range database.Data {
		if player.ID == queries["id"] {

			database.Data = append(database.Data[:i], database.Data[i+1:]...)

			fmt.Fprint(w, "Deleted\n", player)

			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Router
var router = mux.NewRouter()

func routerCalling() {

	//Calling handler functions by their path
	router.HandleFunc("/", welcome)
	router.HandleFunc("/test", JWTauth.IsAuthorized(testPage))
	router.HandleFunc("/api/players", JWTauth.IsAuthorized(players)).Methods("GET")
	router.HandleFunc("/api/players/{id}", JWTauth.IsAuthorized(player)).Methods("GET")
	router.HandleFunc("/api/players/add", JWTauth.IsAuthorized(addPlayer)).Methods("POST")
	router.HandleFunc("/api/players/update/{id}", JWTauth.IsAuthorized(updatePlayer)).Methods("PUT")
	router.HandleFunc("/api/players/delete/{id}", JWTauth.IsAuthorized(deletePlayer)).Methods("DELETE")

}
func StartServer(Port int) {

	log.Println("-------------------starting sever at port :", Port)
	log.Println("Home Page: ", "http://localhost:"+strconv.Itoa(Port))
	log.Println("You will get a new token at home page")

	routerCalling()

	//Server
	server := http.Server{
		Addr:    ":" + strconv.Itoa(Port),
		Handler: router,
	}

	server.ListenAndServe()
}
