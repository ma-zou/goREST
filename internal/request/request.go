package request

import (
	"encoding/base64"
	"goREST/internal/observer"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	observers []observer.Observer
}

type Auth struct {
	Type     string `json:"type"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type RequestObject struct {
	URL     string            `json:"url"`
	Method  string            `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Params  map[string]string `json:"params,omitempty"`
	Body    string            `json:"body,omitempty"`
	Auth    Auth              `json:"auth"`
	LastMod time.Time         `json:"lastMod"`
}

func (r *Request) Register(observer observer.Observer) {
	r.observers = append(r.observers, observer)
}

func (r *Request) Remove(observer observer.Observer) {
	for i, obs := range r.observers {
		if obs == observer {
			r.observers = append(r.observers[:i], r.observers[i+1:]...)
		}
	}
}

func (r *Request) Notify(event interface{}) {
	for _, observer := range r.observers {
		observer.Update(event)
	}
}

func (r *Request) SendRequest(requestObj RequestObject) (*http.Response, error) {

	var payload = requestObj.Body

	req, err := http.NewRequest(requestObj.Method, requestObj.URL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	if len(requestObj.Headers) != 0 {
		addHeaders(req, requestObj.Headers)
	}

	if len(requestObj.Params) != 0 {
		q := req.URL.Query()
		addParams(&q, requestObj.Params)
		req.URL.RawQuery = q.Encode()
	}

	handleAuth(req, requestObj.Auth)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func addHeaders(r *http.Request, headers map[string]string) {
	for key, value := range headers {
		r.Header.Add(key, value)
	}
}

func handleAuth(r *http.Request, auth Auth) {
	switch auth.Type {
	case "basic":
		addBasicAuth(r, auth)
	case "bearer":
		addBearerAuth(r, auth)
	}
}

func addBasicAuth(r *http.Request, auth Auth) {
	encoded := base64.StdEncoding.EncodeToString([]byte(auth.Username + ":" + auth.Password))
	r.Header.Add("Authorization", "Basic "+encoded)
}

func addBearerAuth(r *http.Request, auth Auth) {
	r.Header.Add("Authorization", "Bearer "+auth.Token)
}

func addParams(q *url.Values, params map[string]string) {
	for key, value := range params {
		q.Add(key, value)
	}
}
