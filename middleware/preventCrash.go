package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func PreventCrash(next internals.Callback) internals.Callback {
	return func(w *internals.ResponseWriter, r *internals.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic occurred: %v\n", err)
				fmt.Println("Stack trace:")
				fmt.Println(string(debug.Stack()))

				w.WriteHeader(http.StatusInternalServerError)

				isDevMode := true
				if isDevMode {
					w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n\nStack Trace:\n%s", err, string(debug.Stack()))))
				} else {
					w.Write([]byte("Something went wrong."))
				}
			}
		}()

		next(w, r)
	}
}
