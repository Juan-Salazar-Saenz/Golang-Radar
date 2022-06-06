package main

import (
	Server "golang/server"
)

func main() {

	server := Server.New(":3003")

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
