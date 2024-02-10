package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.NotEmpty(t, status)
	assert.Equal(t, status, http.StatusOK)
	assert.NotEqual(t, body, "moscow")
	assert.Equal(t, totalCount, len(list))

}
