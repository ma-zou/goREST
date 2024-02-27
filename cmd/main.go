package main

import (
	"flag"
	"fmt"
	"goREST/internal/storage"
	"goREST/internal/ui"
)

func main() {
	n := flag.Bool("n", false, "new request")
	d := flag.Bool("d", false, "delete request")
	e := flag.Bool("e", false, "edit request")

	flag.Parse()

	store, err := storage.ReadFromJSON()
	if err != nil {
		fmt.Println(err)
	}

	if *n || len(store) == 0 {
		var newRequest = ui.RenderRequestForm("", "", "", nil, false, false)
		store = append(store, newRequest)
		storage.SaveToJSON(store)
		fmt.Println(newRequest.Url)
	}

	if *e {
		fmt.Println("edit")
		// var newRequest ui.NewRequest = ui.RenderRequestForm()
		// fmt.Println(newRequest.Url)
	}

	if *d {
		fmt.Println("delete")
	}

	ui.RenderTable(store)
}
