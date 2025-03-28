package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(responseRecorder.Body.String())
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(cafes, 2)
}

func TestMainHandlerUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=paris&count=2", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, responseRecorder.Code)
	assert.Equal("wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, responseRecorder.Code)
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(cafes, 4)
	assert.Equal(strings.Join(cafeList["moscow"], ","), responseRecorder.Body.String())
}
