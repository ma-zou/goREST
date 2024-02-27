package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

var storageFile string = "requests.json"

func init() {
	_, err := os.Stat(storageFile)
	if os.IsNotExist(err) {
		file, err := os.Create(storageFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		data, _ := json.Marshal([]Request{})
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
	}
}

func SaveToJSON(data []Request) {
	file, err := os.Create(storageFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data saved successfully to", storageFile)
}

func ReadFromJSON() ([]Request, error) {
	// Open the file for reading
	file, err := os.Open(storageFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to store the data
	var users []Request

	// Decode JSON data from the file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
