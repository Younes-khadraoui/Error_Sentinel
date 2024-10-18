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
	method Method
}

func ReadRequest(reqLen int, buf []byte) (Request, error) {
	usefulPart := buf[:reqLen]
	method := Method(buf[:3])
	if !isValidMethod(method) {
		return Request{}, errors.New("invalid HTTP Method")
	}
	//! split the first line by \r\n then by \n (32)
	parts := bytes.Split(usefulPart, []byte{32})
	for _, part := range parts {
		fmt.Println("part \n", string(part))
	}

	req := Request{
		method: method,
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
