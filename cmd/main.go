package main

import (
	"fmt"
	"goREST/internal/storage"
	"goREST/internal/ui"
)

func main() {
	storage := storage.Storage{}
	ui := ui.UI{}

	storage.Register(&ui)

	requests, err := storage.ReadFromJSON()
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		panic("Issues loading from JSON")
	}
	storage.Notify(requests)

}
