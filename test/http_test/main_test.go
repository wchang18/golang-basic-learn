package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello world", string(w.Body.Bytes()))
}
