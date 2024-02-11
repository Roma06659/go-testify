package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, totalCount, len(list))

}

func TestMainHandlerWhenStatusOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.Equal(t, status, http.StatusOK)
	assert.NotEmpty(t, body)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	var city string
	req := httptest.NewRequest("GET", "/cafe?count=10&city=perm", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code
	city = req.URL.Query().Get("city")

	assert.Equal(t, status, http.StatusBadRequest)
	assert.NotEqual(t, "moscow", city, "wrong city value")

}
