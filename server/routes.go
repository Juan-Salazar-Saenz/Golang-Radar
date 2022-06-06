package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func initRoutes() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	/*Fase 1 en contruccion*/
	router.HandleFunc("/topsecret", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createMessage(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido 0")
			return
		}
	})

	/*Fase 2*/
	router.HandleFunc("/hightopsecret/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getOneSatelite(w, r)
		case http.MethodDelete:
			deleteSatelite(w, r)
		case http.MethodPut:
			updateSatelite(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido 1")
			return
		}
	})

	router.HandleFunc("/hightopsecret", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createSatelite(w, r)
		case http.MethodGet:
			getSateliteAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido 1.1")
			return
		}
	})

	/*Fase 3*/
	router.HandleFunc("/localitationsecret", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getSecretAll(w, r)
		case http.MethodPost:
			addSecretNew(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido 2")
			return
		}
	})

	return router
}
