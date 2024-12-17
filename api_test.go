package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cep?cep=1234", nil)
	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, got %d", rr.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	req, _ := http.NewRequest("POST", "/cep?cep=41650000", nil)
	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405, got %d", rr.Code)
	}
}

func TestValidCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cep?cep=41650000", nil)
	rr := httptest.NewRecorder()
	handler(rr, req)
	fmt.Println(rr.Body.String())

	assert.Contains(t, rr.Body.String(), "temp_C")
}
