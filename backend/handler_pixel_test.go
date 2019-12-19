package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerPixel(t *testing.T) {
	setupData()
	router := setupRouter()

	t.Run("valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pixel", strings.NewReader(`{"r": 255, "g": 0, "b": 0, "x": 100, "y": 100}`))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `{"pixel":{"r":255,"g":0,"b":0,"x":100,"y":100},"status":"created"}`+"\n", w.Body.String())
	})

	t.Run("invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pixel", strings.NewReader(`{"r": 500}`))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, `{"error":"invalid post body"}`+"\n", w.Body.String())
	})
}
