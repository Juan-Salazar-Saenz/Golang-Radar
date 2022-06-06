package models

// Types para los naves y sus mensajes

type Nave struct {
	ID      int    `json:"ID"`
	Message string `json:"Message"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}
