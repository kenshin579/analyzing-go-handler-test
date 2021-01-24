package handler

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/kenshin579/analyzing-go-handler-test/handler/app02"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestApiServer struct {
	router *mux.Router
}

func (api *TestApiServer) Run() string {

	api.router = mux.NewRouter()
	api.router.HandleFunc("/save/{serviceName}", api.Save).Methods("POST")
	api.router.HandleFunc("/save/{serviceName}/{fileName}", api.ReadFile).Methods("GET")

	logfile, _ := os.OpenFile("app02_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	server := httptest.NewServer(handlers.CombinedLoggingHandler(logfile, api.router))
	return server.URL
}

func (api TestApiServer) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceName := vars["serviceName"]
	if serviceName != "calc" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w, app02.ServiceOK)
}

func (api TestApiServer) ReadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceName := vars["serviceName"]
	fileName := vars["fileName"]
	if serviceName != "calc" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if fileName == "my.jpg" {
		//fmt.Fprintf(w, app02.ServiceOK)
		return
	}
}
