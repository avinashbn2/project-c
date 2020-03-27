package main

import (
	"cproject/internal/models"
	"fmt"
	"net/http"
)

var allResources models.Resources

func main() {
	allResources = resources
	http.Handle("/", allResources)
	err := http.ListenAndServe(":3001", nil)
	fmt.Println(err)
}
