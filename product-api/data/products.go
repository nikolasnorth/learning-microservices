package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

type Products []*Product

var productList = []*Product{
	{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee without milk",
		Price:     1.99,
		SKU:       "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
