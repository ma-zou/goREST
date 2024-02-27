package ui

import (
	"errors"
	"goREST/internal/storage"
	"log"
	"net/url"

	"github.com/charmbracelet/huh"
)

type NewRequest struct {
	Method        string
	Url           string
	Auth          bool
	CustomHeaders bool
	Headers       map[string]string
	AuthToken     string
}

func RenderRequestForm(method, urlString, authToken string, headers map[string]string, auth, customHeaders bool) storage.Request {

	if urlString == "" {
		urlString = "https://"
	}

	if headers == nil {
		headers = make(map[string]string)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a method").
				Options(
					huh.NewOption("GET", "GET"),
					huh.NewOption("POST", "POST"),
					huh.NewOption("PUT", "PUT"),
					huh.NewOption("DELETE", "DELETE"),
				).
				Value(&method),
			huh.NewInput().
				Title("URL").
				Value(&urlString).
				Validate(func(s string) error {
					_, err := url.ParseRequestURI(s)
					if err != nil {
						return errors.New(s + " is an invalid URL")
					}

					return nil
				}),
			huh.NewConfirm().
				Title("Authentification?").
				Affirmative("Yes").
				Negative("No").
				Value(&auth),
			huh.NewConfirm().
				Title("Custom Headers?").
				Affirmative("Yes").
				Negative("No").
				Value(&customHeaders),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if auth {
		authToken = RenderAuthForm()
	}

	if customHeaders {
		for customHeaders {
			var headersKey string
			var headerValue string
			customHeaders, headersKey, headerValue = RenderHeadersForm()

			if customHeaders {
				headers[headersKey] = headerValue
			}
		}
	}

	return storage.Request{
		Method:        method,
		Url:           urlString,
		Auth:          auth,
		CustomHeaders: customHeaders,
		Headers:       headers,
		AuthToken:     authToken,
	}
}
