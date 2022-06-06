package repositorio

import (
	"encoding/json"
	Models "golang/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Persistencia
var allflota []*Models.Flota = []*Models.Flota{}

const delimitador string = "."

var nombresateliteanterior = ""
var tamañoAnterior int = 0

var xG int = 0
var yG int = 0

func CreateMessage(w http.ResponseWriter, r *http.Request) any {
	newFlota := &Models.Flota{}

	log.Println("1")
	err := json.NewDecoder(r.Body).Decode(newFlota)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	log.Println(newFlota)
	log.Println("2")
	validacionejes := validacionesEjes(newFlota)
	if validacionejes != "ok" {
		w.WriteHeader(http.StatusBadRequest)
		return validacionejes
	}

	log.Println(validacionejes)
	log.Println("3")
	/*Codificamos el mensaje*/
	mensaje := [100]*Models.Item{}
	for i, satelite := range newFlota.Secret {
		for j, item := range satelite.Message {
			if i > 0 {
				/*Validamos si tiene algun valor y cambiamos el espacio por el valor*/
				if item.Item != "" {
					if mensaje[j].Item != item.Item {
						mensaje[j].Item = item.Item
					}
				}
			}
			/*Cargamos el primer array tal cual es*/
			if i == 0 {
				mensaje[j].Item = item.Item
			}
		}
	}
	log.Println("4")

	/*unificamos el mensaje recibido*/
	var mensajeUltrasecreto = ""
	log.Println(mensajeUltrasecreto)
	var contador = len(mensajeUltrasecreto)
	if contador > 0 {
		for _, mensajes := range mensaje {
			log.Println("4.1")
			if string(mensajes.Item) != "" {
				log.Println("4.2")
				mensajeUltrasecreto += string(mensajes.Item) + " "
			}
		}
	}

	log.Println("5")

	/*guardamos el mensaje*/
	nuevoNave := &Models.Nave{}
	nuevoNave.ID = len(allNaves) + 1
	log.Println("5.1")
	nuevoNave.Message = string(mensajeUltrasecreto)
	log.Println("5.2")
	nuevoNave.X = xG
	log.Println("5.3")
	nuevoNave.Y = yG
	log.Println("5.4")
	allNaves = append(allNaves, nuevoNave)

	log.Println("6")
	/*Guardamos el mensaje de la flota que se envio despues de las distintas validaciones*/
	newFlota.ID = len(allflota) + 1
	allflota = append(allflota, newFlota)

	log.Println("7")
	/*Mostramos el mensaje tal cual fue decodificado*/
	return nuevoNave
}

func validacionesEjes(flota *Models.Flota) string {
	/*Aqui hacemos la validaciones del mensaje*/
	for i, satelite := range flota.Secret {

		/*obtenemos variables del satelite a validar*/
		var tamanoMessage int = len(satelite.Message)
		var nombresatelite = satelite.Name

		/*Por cada arreglo , validamos los valores en x / y con el fin de poder graficar en angular*/
		Ejes := strings.Split(satelite.Distance, delimitador)

		x, err := strconv.Atoi(Ejes[0])
		/*Validamos si es un caracter distinto a numero eje x */
		if err != nil {
			return "Invalid X of satelite " + satelite.Name + " " + strconv.Itoa(x)
		}

		y, err := strconv.Atoi(Ejes[1])
		/*Validamos si es un caracter distinto a numero eje y */
		if err != nil {
			return "Invalid Y of satelite " + satelite.Name + " " + strconv.Itoa(y)
		}

		/*Variables de localizacion*/
		var xP int = 0
		var yP int = 0
		var xN int = 0
		var yN int = 0
		var h int = 0
		var k int = 0

		/*Validamos que este dentro de los parametro -50 a 50 eje x */
		error := validafCoordenadas(x, "X", satelite.Name)
		if error != "ok" {
			return error
		}

		/*Validamos que este dentro de los parametro -50 a 50 eje y */
		error2 := validafCoordenadas(y, "Y", satelite.Name)
		if error2 != "ok" {
			return error2
		}

		/*validamos que las coordenadas enviadas a los distintos satelites esten correctas*/

		var existenave = ""
		/*Consultamos la existencia del satelite y consultamos sus coordenadas*/
		h, k, existenave = GetOneSateliteName(satelite.Name)
		if existenave == "ok" {
			return "Satelite no existe " + satelite.Name
		}

		/*obtenemos x prima y y prima*/
		xP = x
		yP = y

		/*calculamos la posicion en el plano de la nave*/
		xN = validacionCoordenadas(xP, h)
		yN = validacionCoordenadas(yP, k)

		/*Validamos que el resultado de los calculos sean los mismos y asi verificar que los datos enviados esten correctos*/
		if i > 0 {
			if xG != xN {
				return "coordenadas errores eje x " + nombresatelite + " y " + nombresateliteanterior
			}
			if yG != yN {
				return "coordenadas errores eje Y " + nombresatelite + " y " + nombresateliteanterior
			}
		}

		/*Localizacion el punto origen del mensaje */
		xG = xN
		yG = yN

		/*validamos que todos los mensajes tengan el mismo tamaño*/
		if i > 0 {
			if tamanoMessage != tamañoAnterior {
				return "Mensaje diferente entre " + nombresatelite + " y " + nombresateliteanterior
			}
		}

		/*Validamos que todos los mensajes tengan el mismo tamaño*/
		nombresateliteanterior = satelite.Name
		tamañoAnterior = len(satelite.Message)

	}

	return "ok"
}

//Validamos que las coordenadas esten dentro de los parametros iniciales
func validafCoordenadas(eje int, texto string, name string) string {
	if eje > 50 || eje < -50 {
		return "the position of the satellite in the " + texto + " axis must be in (-50 to 50), Satelite " + name
	}
	return "ok"
}

func validacionCoordenadas(xP int, h int) int {
	/*De acuerdo a la definicon de Traslación de ejes en el plano cartesiano tenemos
	Origen(h,k) Punto(x,y)
	xP = x-h    yP = y-k
	Pero nosotros tenemos xP y yP asi que despejamos
	x = xP + h     y = yP + k
	*/
	var x int = xP + h
	return x
}

//Metodo para consultar todos los satelites existentes
func GetMessageAll(w http.ResponseWriter, r *http.Request) any {
	w.WriteHeader(http.StatusOK)
	return allflota
}
