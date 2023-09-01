package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Struct
var Products = []Product{
	{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Descripción del Producto 1", Category: "Categoría A"},
	{ID: 2, Name: "Producto 2", Price: 29.99, Description: "Descripción del Producto 2", Category: "Categoría B"},
	{ID: 3, Name: "Producto 3", Price: 39.99, Description: "Descripción del Producto 3", Category: "Categoría A"},
}

// Methods
func (p *Product) Save(product Product) {
	Products = append(Products, product)
	//fmt.Println(Products)
}

func GetAll() []Product {
	fmt.Println(Products)
	return Products
}

func GetByID(ID int) Product {
	for _, p := range Products {
		if p.ID == ID {
			fmt.Printf("Product founded %s\n", p)
			break
			return p
		}
	}
	return Product{}
}

func main() {
	newProduct := Product{
		ID:          4,
		Name:        "New",
		Price:       87.99,
		Description: "New Product",
		Category:    "New",
	}
	newProduct.Save(newProduct)
	foundProduct := GetByID(4)
	fmt.Println(foundProduct)
	//allProducts := GetAll()
	//fmt.Println(allProducts)
}
