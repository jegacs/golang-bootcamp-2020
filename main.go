package main

import (
	"fmt"

	v1 "github.com/jegacs/golang-bootcamp-2020/handlers/v1"
)

func main() {
	fmt.Println("Starting server... ")
	v1.Run(":8000")
}
