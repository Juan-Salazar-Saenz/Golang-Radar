package models

// Types para los satelites

type Satelite struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}
