package internals

type ResponseWriter struct{}

func CreateResponse(w WebServer, r Request) []byte {
	handler, ok := w.Router[r.Method][r.Endpoint]
	if !ok {
		return []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}

	return handler(w)
}
