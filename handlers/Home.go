package handlers

import (
	"fmt"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func Home(w *internals.ResponseWriter, r *internals.Request) {
	w.WriteHeader(200)
	w.SetHeader("Content-Type", "text/html")

	body := "<html><body>Hello, world!</body></html>"
	w.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))

	w.Write([]byte(body))
}
