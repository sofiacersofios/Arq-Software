package main

import (
	"errors" // para errores
	"fmt"    // para mensajes
)

func main() {
	div, err := division(7, 0)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Dvision: ", div)
}
func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("No se puede dividir")
	}

	return a / b, nil
}
