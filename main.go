package main

import (
	Server "golang/server"
)

func main() {

	server := Server.New(":3003")

	/*http*/
	//err := server.ListenAndServe()
	/*https*/
	err := server.ListenAndServeTLS("cert/my_cert.crt", "cert/my_cert.key")

	if err != nil {
		panic(err)
	}
}
