package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/nikolasnorth/microservices/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		// Grab id from request
		path := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)

		subs := reg.FindAllStringSubmatch(path, -1)
		if len(subs) != 1 || len(subs[0]) != 2 {
			http.Error(w, "Request must contain exactly one ID", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(subs[0][1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p.updateProduct(id, w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	w.Header().Set("Content-Type", "application/json")
	list := data.GetProducts()
	err := list.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = data.UpdateProduct(prod, id)
	if err != nil {
		if err == data.ErrProductNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
