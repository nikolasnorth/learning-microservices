package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var ErrProductNotFound = errors.New("product not found")

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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

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

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(p *Product, id int) error {
	_, i, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	return nil
}

// Simulate find Product by id database query
func FindProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, 0, ErrProductNotFound
}

// Simulate function to generate unique ID
func getNextID() int {
	last := productList[len(productList)-1]
	return last.ID + 1
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
