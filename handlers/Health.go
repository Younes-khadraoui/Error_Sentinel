package handlers

import (
	"fmt"
	"time"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func Health(w *internals.ResponseWriter, r *internals.Request) {
	w.WriteHeader(200)
	w.SetHeader("Content-Type", "text/html")

	uptime := time.Since(w.StartTime)
	body := fmt.Sprintf("<html><body>The server has been running for <span style='color:blue;text-align:center;'> %+v </span> seconds!</body></html>", uptime)
	w.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))

	w.Write([]byte(body))
}
