package handlers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Younes-khadraoui/Error_Sentinel/internals"
)

func Retry(w *internals.ResponseWriter, r *internals.Request) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	attempts := 0
	maxAttempts := 3
	success := false

	for attempts < maxAttempts {
		attempts++

		if rng.Intn(2) == 0 {
			fmt.Printf("Attempt %d: Simulated failure\n", attempts)
			continue
		}

		success = true
		break
	}

	if success {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Request succeeded on attempt %d", attempts)))
	} else {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Failed after %d attempts.", maxAttempts)))
	}
}
