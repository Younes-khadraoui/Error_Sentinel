package handlers

import "github.com/Younes-khadraoui/Error_Sentinel/internals"

func Panic(w *internals.ResponseWriter, r *internals.Request) {
	panic("Erm something went wrong")
}
