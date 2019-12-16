package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetIndexRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestPostPixelRoute(t *testing.T) {
// 	router := setupRouter()
//
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/pixel", strings.NewReader(`{"color": "ff0000", "pos": [10, 20]}`))
// 	req.Header.Add("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
//
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	t.Log("===", w.Body.String())
// 	// assert.Equal(t, "pong", w.Body.String())
// }
