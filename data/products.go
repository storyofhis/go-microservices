package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID 			int64 		`json:"id"`
	Name 		string		`json:"name"`
	Description	string 		`json:"description"`
	Price 		float64		`json:"price"`
	SKU			string		`json:sku`
	CreatedOn	string		`json:"-"`
	UpdatedOn	string		`json:"-"`
	DeletedOn	string		`json:"-"`
}

// Products is a collection of Product
type Products[] *Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID: 2,
		Name: "Espresso",
		Description: "short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
