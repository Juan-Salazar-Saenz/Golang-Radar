package server

import (
	"encoding/json"
	"fmt"
	Funciones "golang/repositorio"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo no permitido")
		return
	}

	fmt.Fprintf(w, "Hello world, welcome to the resistance!")
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.CreateMessage(w, r))
}

func createSatelite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.CreateSatelite(w, r))
}

func getSateliteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.GetSateliteAll(w, r))
}

func getOneSatelite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.GetOneSatelite(w, r))
}

func deleteSatelite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.DeleteSatelite(w, r))
}

func updateSatelite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.UpdateSatelite(w, r))
}

/*FAse 3 */
func getSecretAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.GetSecretAll(w, r))
}

func addSecretNew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.AddSecretNew(w, r))
}
