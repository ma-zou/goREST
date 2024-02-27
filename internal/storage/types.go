package storage

type Request struct {
	Method        string
	Url           string
	Auth          bool
	CustomHeaders bool
	Headers       map[string]string
	AuthToken     string
}
