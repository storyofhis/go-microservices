package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/storyofhis/microservices-go/data"
)


type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	data, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		// log.Println("error : ", err)
		// os.Exit(1)
	}

	w.Write(data)
}