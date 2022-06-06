package repositorio

import (
	"encoding/json"
	"fmt"
	Models "golang/models"
	"net/http"
)

// Persistencia
var allNaves []*Models.Nave = []*Models.Nave{}

//Metodo para consultar todos las posiciones y mensajes ocultos
func GetSecretAll(w http.ResponseWriter, r *http.Request) any {
	return allNaves
}

func AddSecretNew(w http.ResponseWriter, r *http.Request) any {
	newNave := &Models.Nave{}

	err := json.NewDecoder(r.Body).Decode(newNave)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return err
	}

	allNaves = append(allNaves, newNave)
	return allNaves
}
