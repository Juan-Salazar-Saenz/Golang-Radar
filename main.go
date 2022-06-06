package main

import (
	"context"
	Server "golang/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	server := Server.New(":3003")

	/*http*/
	//err := server.ListenAndServe()
	/*https*/
	go func() {
		err := server.ListenAndServeTLS("cert/my_cert.crt", "cert/my_cert.key")
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Server Started")

	<-serverDoneChan

	server.Shutdown(ctx)
	log.Println("Server Stopped")
}
