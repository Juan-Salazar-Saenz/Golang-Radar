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

	fmt.Fprintf(w, "Hello Word %s", "visitor")
}

func getCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.GetSecretAll(w, r))
}

func addCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Funciones.AddSecretNew(w, r))
}
