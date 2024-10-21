package internals

import (
	"fmt"
	"log"
	"net"
	"time"
)

type WebServer struct {
	Router    Routes
	StartTime time.Time
}

type Callback func(w *ResponseWriter, r *Request)

type Routes map[Method]map[string]Callback

func (w *WebServer) GET(endpoint string, handler Callback) {
	if w.Router == nil {
		w.Router = make(Routes)
	}
	if w.Router["GET"] == nil {
		w.Router["GET"] = make(map[string]Callback)
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
		go HandleConnection(w, conn)
	}
}

// ? Constructor for WebServer
func NewWebServer() *WebServer {
	return &WebServer{
		Router:    make(Routes),
		StartTime: time.Now(),
	}
}

func HandleConnection(w WebServer, c net.Conn) {
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

		res := CreateResponse(w, req)
		c.Write(res)
		break
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
	c.Close()
}
