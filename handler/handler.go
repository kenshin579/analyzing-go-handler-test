package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	router     *mux.Router
	fileServer string
}

type Response struct {
	Content string
	Code    int
}

func (h Handler) Init(fileServer string) {
	h.router = mux.NewRouter()
	h.fileServer = fileServer
	h.router.HandleFunc("/ping", h.Ping).Methods("GET")
	h.router.HandleFunc("/div/{a}/{b}", h.Div).Methods("GET")
	http.Handle("/", h.router)
}
func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (h Handler) Div(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]

	ai, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if bi == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if ai == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	DoPost(h.fileServer+"/save/calc", "a/b")
	fmt.Fprintf(w, "%d", ai/bi)
}

func DoPost(url string, data string) (*Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(data))
	request.ContentLength = int64(len(data))
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}
