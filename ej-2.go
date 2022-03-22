// curl link -v
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Categorias []Categoria

type Categoria struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func main() {
	cats, err := GetCategorias("MLA")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Las categorias son ....")
}

func GetCategorias(siteID string) (Categorias, error) {
	response := http.Get()
	bytes := ioutil.ReadAll(response.Bytes)
	json.Unmarshal(bytes, &Categoria)
}
