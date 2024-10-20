package handlers

import (
	"fmt"
	"time"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func Health(w internals.WebServer) []byte {
	uptime := time.Since(w.StartTime)
	body := fmt.Sprintf("<html><body>The server has been running for <span style='color:blue;text-align:center;'> %+v </span> seconds!</body></html>", uptime)
	bodyLength := len(body)

	statusLine := "HTTP/1.1 200 OK\r\n"
	headers := fmt.Sprintf("Content-Type: text/html\r\nContent-Length: %d\r\n\r\n", bodyLength)

	return []byte(statusLine + headers + body)
}
