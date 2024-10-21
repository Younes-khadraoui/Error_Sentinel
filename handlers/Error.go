package handlers

import (
	"fmt"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func Error(w *internals.ResponseWriter, r *internals.Request) {
	w.WriteHeader(500)
	w.SetHeader("Content-Type", "text/html")

	body := "<html><body><h1>500 Internal Server Error</h1><p>Something went wrong.</p></body></html>"
	w.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))

	w.Write([]byte(body))
}
