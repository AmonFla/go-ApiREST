package main

import (
	"fmt"
	"go-ApiREST/endpoint"
	"net/http"

	"github.com/gorilla/mux"
)

// Función principal que incializa el servidor
func main() {

	//Habilito los handler de cada ruta utilizando Gorilla/mux
	router := mux.NewRouter()

	// Configuro cada una de las rutas y metodos que aceptan
	router.HandleFunc("/personas", endpoint.GetPersonas).Methods("GET")
	router.HandleFunc("/personas/{ID:[0-9]+}", endpoint.GetPersonaById).Methods("GET")
	router.HandleFunc("/personas", endpoint.SavePersona).Methods("POST")
	router.HandleFunc("/personas/{ID:[0-9]+}", endpoint.EditPersona).Methods("PUT")
	router.HandleFunc("/personas/{ID:[0-9]+}", endpoint.DelPersona).Methods("DELETE")

	//Inicializo el server
	fmt.Printf("Starting server 127.0.0.1:4000 ")
	if err := http.ListenAndServe(":4000", router); err != nil {
		panic(err)
	}
}
