package repositorio

import (
	"encoding/json"
	Models "golang/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Persistencia
var allSatelites []*Models.Satelite = []*Models.Satelite{}

//Metodo para crear un nuevo satelite
func CreateSatelite(w http.ResponseWriter, r *http.Request) any {
	newSatelite := &Models.Satelite{}

	err := json.NewDecoder(r.Body).Decode(newSatelite)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	error := ValidacionSatelite(newSatelite)
	if error != "ok" {
		w.WriteHeader(http.StatusBadRequest)
		return error
	}

	newSatelite.ID = len(allSatelites)
	allSatelites = append(allSatelites, newSatelite)
	w.WriteHeader(http.StatusCreated)

	return "New Satelite " + newSatelite.Name
}

func ValidacionSatelite(satelite *Models.Satelite) string {
	error := validaCoordenadas(satelite.X, "X")
	if error != "ok" {
		return error
	}

	error2 := validaCoordenadas(satelite.Y, "Y")
	if error2 != "ok" {
		return error2
	}

	error3 := GetOneSateliteName(satelite.Name)
	if error3 != "ok" {
		return error3
	}

	return "ok"
}

//Validamos que las coordenadas esten dentro de los parametros iniciales
func validaCoordenadas(eje int, texto string) string {
	if eje > 50 || eje < -50 {
		return "the position of the satellite in the " + texto + " axis must be in (-50 to 50)"
	}
	return "ok"
}

//Validamos que el nombre sea unico
func GetOneSateliteName(name string) string {

	for _, satelite := range allSatelites {
		if satelite.Name == name {
			return "Satelite exist"
		}
	}

	return "ok"
}

//Metodo para consultar todos los satelites existentes
func GetSateliteAll(w http.ResponseWriter, r *http.Request) any {
	w.WriteHeader(http.StatusOK)
	return allSatelites
}

//Metodo para consultar un solo satelite
func GetOneSatelite(w http.ResponseWriter, r *http.Request) any {
	vars := mux.Vars(r)
	sateliteID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "Invalid ID of Satelite"
	}

	for _, satelite := range allSatelites {
		if satelite.ID == sateliteID {
			w.WriteHeader(http.StatusOK)
			return satelite
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return "Satelite does not exist"
}

//Metodo para eliminar un solo satelite
func DeleteSatelite(w http.ResponseWriter, r *http.Request) any {
	vars := mux.Vars(r)
	sateliteID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "Invalid ID of Satelite"
	}

	for i, satelite := range allSatelites {
		if satelite.ID == sateliteID {
			allSatelites = append(allSatelites[:i], allSatelites[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return "The satelite with ID " + strconv.Itoa(sateliteID) + " has been delete successfully"
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return "Satelite does not exist"
}

//Metodo para modificar un solo satelite
func UpdateSatelite(w http.ResponseWriter, r *http.Request) any {
	vars := mux.Vars(r)
	sateliteID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "Invalid ID of Satelite"
	}

	updateSatelite := &Models.Satelite{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "Please Enter Valid Satelite"
	}
	json.Unmarshal(reqBody, &updateSatelite)

	error := ValidacionSatelite(updateSatelite)
	if error != "ok" {
		w.WriteHeader(http.StatusBadRequest)
		return error
	}

	for i, satelite := range allSatelites {
		if satelite.ID == sateliteID {
			allSatelites = append(allSatelites[:i], allSatelites[i+1:]...)
			allSatelites = append(allSatelites, updateSatelite)
			w.WriteHeader(http.StatusOK)
			return "The satelite with ID " + strconv.Itoa(sateliteID) + " has been delete successfully"
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return "Satelite does not exist"
}
