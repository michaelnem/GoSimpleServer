package n_http

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	// "PATCH", "HEAD", "OPTIONS"
)
