package handler

import (
	"github.com/gorilla/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server  *httptest.Server
	testUrl string
)

func Test_Init(t *testing.T) {
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	assert.Nil(t, err, "")

	FileServer := TestApiServer{}
	fileServerAddr := FileServer.Run() //app02 외부 서버를 띄움 - 이 로직은 mock으로 처리하는게 좋아보임
	h := Handler{}
	h.Init(fileServerAddr)

	//todo : http server를 시작하고 나서 어디서 teardown을 하나? - 해야 하지 않나?
	server = httptest.NewServer(handlers.CombinedLoggingHandler(logfile, http.DefaultServeMux))
	testUrl = server.URL
}

func Test_Ping(t *testing.T) {
	res, err := DoGet(testUrl + "/ping")
	assert.Nil(t, err, "")
	assert.Equal(t, 200, res.Code, "PING API")
	assert.Equal(t, "pong", res.Content, "PONG Message")
}

func Test_APINotFound(t *testing.T) {
	res, err := DoGet(testUrl + "/myfunc")
	assert.Nil(t, err, "")
	assert.Equal(t, 404, res.Code, "Unknown API")
}

func Test_Div(t *testing.T) {
	res, err := DoGet(testUrl + "/div/4/2")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "2", res.Content)

	res, err = DoGet(testUrl + "/div/4/a")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusInternalServerError, res.Code, "Invalide argument")

	res, err = DoGet(testUrl + "/div/0/4")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusNotAcceptable, res.Code, "Invalide argument")

	res, err = DoGet(testUrl + "/div/4/0")
	assert.Nil(t, err, "")
	assert.Equal(t, http.StatusNotAcceptable, res.Code, "Invalide argument")
}

func DoGet(url string) (*Response, error) {
	response, err := http.Get(url)
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
