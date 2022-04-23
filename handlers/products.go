package handlers

import (
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
	// lp := data.GetProducts()
	// data, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	// 	// log.Println("error : ", err)
	// 	// os.Exit(1)
	// }
	// w.Write(data)
	if r.Method == http.MethodGet {
		p.getProducts(w,r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}