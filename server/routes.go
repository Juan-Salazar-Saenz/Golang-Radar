package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/localitationsecret", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountry(w, r)
		case http.MethodPost:
			addCountry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})
}
