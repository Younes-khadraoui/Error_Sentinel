package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
)

type webServer struct{}

//? Methodes to set status codes , header and send the response
type ResponseWriter struct{}
//? HTTP method , path , headers , body
type Request struct{}

type callback func(w ResponseWriter, r *Request)

func (w webServer) GET(endpoint string, handler callback) {
	//! map the "/" to handler functions , store these mappings in (ex: map) so u can call the right handler based on the requested path
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Panic("Please Provide a PORT number")
	}
	port := args[1]
	var re = regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(port) {
		log.Panic("PORT must contain only numbers.")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Panic("Error converting Port to INT ", err)
	}
	if portInt > 65353 {
		log.Panic("Plese Provide a PORT number < 65353")
	}
	app := new(webServer)
	app.GET("/", Home)
	err = app.Start(port)
	if err != nil {
		log.Panic("Error Starting The Server :( ")
	}
	fmt.Println("Server Running on Port", port)
}

func Home(w ResponseWriter, r *Request) {

}

func (w webServer) Start(PORT string) error {
	ln, err := net.Listen("tcp4", ":"+PORT)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		//! Parse the data ( split to headers , method , path ... and convert to a Request object)
		fmt.Print("the read data", n)
		_, err = c.Write(buf[0:n])
		if err != nil {
			c.Close()
			break
		}
		//! construct the HTTP response (HTTP/1.1 200 OK) , headers (content-types: text/html)
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
