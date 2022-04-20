package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"os/signal"
	"github.com/storyofhis/microservices-go/handlers"
)

func main () {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)	// log
	h := handlers.NewMainHandler(l) 						// handler

	// server mux
	sm := http.NewServeMux()
	sm.Handle("/", h)
	
	server := &http.Server{
		Addr: ":8080",						// as Address
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func () {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// signal chanel
	sigChan := make(chan os.Signal)
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