package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
)

func TestHealcheckHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	a := application{}

	a.healthcheckHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	env := envelope{}
	err = json.Unmarshal(body, &env)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, env["status"], "available")
}
