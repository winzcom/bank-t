package http

import (
	"fmt"
	"io"
	// "net/http"
	"os"
)

var T_URL = os.Getenv("T_PARTY")

type TRequest struct {
}

func PostTransactions(content io.Reader) {
	fmt.Println("what is t_url ", T_URL)
	// http.Post(T_URL, "application/json", content)
}
