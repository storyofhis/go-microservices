package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/storyofhis/microservices-go/handlers"
)

func main () {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)	// log

	h := handlers.NewMainHandler(l) 						// create handler
	productHandler := handlers.NewProducts(l)

	// create a new server mux and register the handler 
	sm := http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/products", productHandler)
	
	server := &http.Server{
		Addr: ":8080",						// as Address. configure the bind address
		Handler: sm,						// set the default handler
		ErrorLog : l,						// set the logger for the server 
		IdleTimeout: 120 * time.Second,		// max time to connecting using TCP Keep-Alive
		ReadTimeout: 1 * time.Second,		// max time to read request from the client	
		WriteTimeout: 1 * time.Second,
	}

	// start the server 
	go func () {
		l.Println("Starting server on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server : %s\n", err )
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	// signal chanel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received Terminate, graceful shutdown", sig)
	
	// timeout context
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Fatal("error : ", err)
	}

	server.Shutdown(tc)
}