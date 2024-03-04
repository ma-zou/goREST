package storage

import (
	"encoding/json"
	"fmt"
	"goREST/internal/observer"
	"goREST/internal/request"
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
		data, _ := json.Marshal([]request.RequestObject{})
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
	}
}

// func SaveToJSON(data []request.RequestObject) {
// 	file, err := os.Create(storageFile)
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error marshalling data:", err)
// 		return
// 	}

// 	_, err = file.Write(jsonData)
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}

// 	fmt.Println("Data saved successfully to", storageFile)
// }

type Storage struct {
	observers []observer.Observer
}

func (s *Storage) Register(observer observer.Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Storage) Remove(observer observer.Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Storage) Notify(event interface{}) {
	for _, observer := range s.observers {
		observer.Update(event)
	}
}

func (s *Storage) ReadFromJSON() ([]request.RequestObject, error) {
	file, err := os.Open(storageFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var requests []request.RequestObject

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&requests)

	if err != nil {
		return nil, err
	}

	return requests, nil
}
