package main

import (
	"fmt"
	"net/http"
	"projectc/internal/models"
)

var allResources models.Resources

func main() {
	allResources = resources
	http.Handle("/", allResources)
	er := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
