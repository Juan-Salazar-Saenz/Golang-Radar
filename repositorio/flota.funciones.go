package repositorio

import (
	Models "golang/models"
	"net/http"
)

// Persistencia
var allflota []*Models.Flota = []*Models.Flota{}

func CreateMessage(w http.ResponseWriter, r *http.Request) any {
	return allflota
}
