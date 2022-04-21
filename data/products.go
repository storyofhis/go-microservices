package data

import (
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
// type Products[] *Product

func GetProducts() []*Product {
	return productList
}
var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}