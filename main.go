package main

import (
	"bitbucket.org/dream_yun/handlertest/handler"
	"net/http"
)

func main() {
	h := handler.Handler{}
	h.Init("http://127.0.0.1")
	http.ListenAndServe(":9998", nil)
}
