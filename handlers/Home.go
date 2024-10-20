package handlers

import "github.com/Younes-khadraoui/Error_Sentinel/internals"

func Home(w internals.WebServer) []byte {
	statusLine := "HTTP/1.1 200 OK\r\n"
	headers := "Content-Type: text/html\r\nContent-Length: 45\r\n\r\n"
	body := "<html><body>Hello, world!</body></html>"

	return []byte(statusLine + headers + body)
}
