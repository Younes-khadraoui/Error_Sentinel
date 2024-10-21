package internals

import (
	"fmt"
	"time"
)

type ResponseWriter struct {
	StatusLine string
	Headers    map[string]string
	Body       string
	StartTime  time.Time
}

func (r *ResponseWriter) Write(body []byte) {
	r.Body = string(body)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.StatusLine = StatusLineFromCode(statusCode)
}

func (r *ResponseWriter) SetHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[key] = value
}

func CreateResponse(w WebServer, r Request) []byte {
	responseWriter := &ResponseWriter{
		Headers:   make(map[string]string),
		StartTime: w.StartTime,
	}

	handler, ok := w.Router[r.Method][r.Endpoint]
	if !ok {
		return []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}

	handler(responseWriter, &r)

	response := responseWriter.StatusLine + "\r\n"
	for key, value := range responseWriter.Headers {
		response += key + ": " + value + "\r\n"
	}
	response += "\r\n" + responseWriter.Body

	return []byte(response)
}

func StatusLineFromCode(code int) string {
	switch code {
	case 200:
		return "HTTP/1.1 200 OK"
	case 404:
		return "HTTP/1.1 404 Not Found"
	case 500:
		return "HTTP/1.1 500 Internal Server Error"
	default:
		return fmt.Sprintf("HTTP/1.1 %d", code)
	}
}
