package main

import (
	"context"
	"fmt"
	"nombre_del_modulo/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)
	srv := server.New(":8080")

	go srv.ListenAndServe()
	fmt.Println("hola usuario el server se prende")
	<-serverDoneChan
	srv.Shutdown(ctx)
	fmt.Println("adios usuario el server se apaga")
}
