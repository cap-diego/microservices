package main

import (
	"os"
	"time"
	"context"
	"net/http"
	"log"
	"github.com/cap-diego/microservices/handlers"
	"os/signal"
	"fmt"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHome(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	server := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() { //Non-blocking
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) //Broadcast message whenever os.Int/os.Kill
	signal.Notify(sigChan, os.Kill)
	
	sig := <- sigChan //Block

	l.Println("Terminate: ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	server.Shutdown(tc)
}