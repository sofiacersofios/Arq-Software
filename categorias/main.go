package main

import (
	"fmt"
	"os"

	"github.com/sofiacersofios/Arq-Software/categorias/cat"
)

func main() {
	categories, err := cat.GetCategories("MLA")
	if err != nil {
		fmt.Println("Error getting categories: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Categorias obtenidas de la API de Mercado Libre:")
	for _, category := range categories {
		fmt.Println(category.String())
	}
}
