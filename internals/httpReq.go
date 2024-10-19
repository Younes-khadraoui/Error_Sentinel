package internals

import (
	"bytes"
	"errors"
	"fmt"
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
	Path        string
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
	for _,part := range parts {
		fmt.Println("part: ",string(part[0]))
	}
	//? 1st line 1st word
	method := Method(parts[0][0])
	if !isValidMethod(method) {
		return Request{}, errors.New("invalid HTTP Method")
	}
	path := string(parts[0][1])

	req := Request{
		Method: method,
		Path: path,
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
