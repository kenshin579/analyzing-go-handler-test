package main

import (
	"github.com/kenshin579/analyzing-go-handler-test/handler"
	"net/http"
)

func main() {
	h := handler.Handler{}
	h.Init("http://127.0.0.1")
	http.ListenAndServe(":9998", nil)
}
