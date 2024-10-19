package internals

import (
	"bytes"
	"errors"
)

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PUT    Method = "PUT"
)

type Method string

type Request struct {
	Method      Method
	Endpoint    string
	HttpVersion string
	Headers     map[string]string
	Body        string
}

func ReadRequest(reqLen int, buf []byte) (Request, error) {
	usefulPart := buf[:reqLen]

	var parts [][][]byte

	lines := bytes.Split(usefulPart, []byte{13, 10})
	for _, line := range lines {
		parts = append(parts, bytes.Split(line, []byte{32}))

	}

	method := Method(parts[0][0])
	if !isValidMethod(method) {
		return Request{}, errors.New("invalid HTTP Method")
	}
	endpoint := string(parts[0][1])

	req := Request{
		Method:   method,
		Endpoint: endpoint,
	}
	return req, nil
}

func isValidMethod(m Method) bool {
	switch m {
	case GET, POST, DELETE, PUT:
		return true
	default:
		return false
	}
}
