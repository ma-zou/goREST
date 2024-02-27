package request

import (
	"fmt"
	"goREST/internal/storage"
	"net/http"
)

func SendRequest(request storage.Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	for key, value := range request.Headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	return resp, err
}
