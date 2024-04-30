package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	// "net/http"

	"github.com/joho/godotenv"
)

var T_URL string

type TRequest struct {
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	T_URL = os.Getenv("T_PARTY")
}

func PostTransactions(content io.Reader, path string) {
	fmt.Println("what is t_url ", T_URL)
	url, _ := url.JoinPath(T_URL, path)
	fmt.Println("query url ", url)
	http.Post(url, "application/json", content)
}
