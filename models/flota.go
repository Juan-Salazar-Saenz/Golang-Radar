package models

/*definimos el primer nivel de los mensajes*/
type Item struct {
	Item string `json:"item"`
}

/*definimos el segundo nuvel de los mensajes que son por satelites*/
type Satelites struct {
	Name     string `json:"name"`
	Distance string `json:"distance"`
	Message  []Item `json:"message"`
}

/*definimos el tercer nuvel de los mensajes que son las flota que contiene todos los mensajes de los satelites*/
type Flota struct {
	ID     int         `json:"id"`
	Secret []Satelites `json:"satelites"`
}
