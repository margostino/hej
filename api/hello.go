package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!", os.Getenv("TEST_SUPER_KEY"))
}
