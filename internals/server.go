package internals

import (
	"fmt"
	"io"
	"log"
	"net"
)

type WebServer struct{}

// ? Methodes to set status codes , header and send the response
type ResponseWriter struct{}

type callback func(w ResponseWriter, r *Request)

func (w WebServer) GET(endpoint string, handler callback) {
	//! map the "/" to handler functions , store these mappings in (ex: map) so u can call the right handler based on the requested path
}

func (w WebServer) Start(PORT string) error {
	ln, err := net.Listen("tcp4", ":"+PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go HandleConnection(conn)
	}
}

func HandleConnection(c net.Conn) {
	buf := make([]byte, 4096)

	for {
		reqLen, err := c.Read(buf)
		if err != nil || reqLen == 0 {
			c.Close()
			break
		}
		req, err := ReadRequest(reqLen, buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The Request is %+v\n", req)
		io.Copy(c, c)
		//! construct the HTTP response (HTTP/1.1 200 OK) , headers (content-types: text/html)
		
		fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
