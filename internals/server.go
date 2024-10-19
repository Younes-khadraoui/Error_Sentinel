package internals

import (
	"fmt"
	"log"
	"net"
)

type WebServer struct {
	Router Routes
}

type callback func(w ResponseWriter, r *Request)

type Routes map[Method]map[string]callback

//? Constructor for WebServer
func NewWebServer() *WebServer {
	return &WebServer{
		Router: make(Routes),
	}
}

func (w *WebServer) GET(endpoint string, handler callback) {
	if w.Router == nil {
		w.Router = make(Routes)
	}
	if w.Router["GET"] == nil {
		w.Router["GET"] = make(map[string]callback)
	}
	w.Router["GET"][endpoint] = handler
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

		res := CreateResponse(req)
		c.Write(res)
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
