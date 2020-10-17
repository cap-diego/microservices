package main

import (
	"os"
	"time"
	"context"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/go-openapi/runtime/middleware"
	"github.com/cap-diego/microservices/handlers"
	"os/signal"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()	
	getRouter.HandleFunc("/products", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)


	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create a server 
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