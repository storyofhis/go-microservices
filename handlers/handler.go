package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MainHandler struct {
	l *log.Logger
}

func NewMainHandler(l *log.Logger) *MainHandler {
	return &MainHandler{l}
}

func (m *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.l.Println("Hello World!");
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "error nya mas", http.StatusBadRequest)
		log.Println(err)
		return
	}
	// log.Printf("Data %s\n", d)	-> curl -v -d "message" 127.0.0.1:8080. (using Request)
	fmt.Fprintf(w, "Hello %s\n", d)	// -> curl -d "message" 127.0.0.1:8080. will show on another terminal  (using Response)		
}