package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/j-santos2/cinema-vault/internal/assert"
	"github.com/j-santos2/cinema-vault/internal/validator"
)

func TestReadIDParam(t *testing.T) {
	router := httprouter.New()
	a := application{}

	var id int64
	var err error
	router.HandlerFunc(
		http.MethodGet,
		"/resource/:id",
		func(w http.ResponseWriter, r *http.Request) {
			id, err = a.readIDParam(r)
		},
	)

	r, err := http.NewRequest(http.MethodGet, "/resource/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, r)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, 2)
}

func TestReadString(t *testing.T) {
	t.Run("test read query string key as string", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()
		q.Add("a_key", "val")

		a := application{}
		s := a.readString(q, "a_key", "def")

		assert.Equal(t, s, "val")
	})
	t.Run("test read absent query string key as string", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()

		a := application{}
		s := a.readString(q, "a_key", "def")

		assert.Equal(t, s, "def")
	})
}

func TestReadCSV(t *testing.T) {
	t.Run("test read query string key as CSV", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()
		q.Add("a_key", "val1,val2,val3")

		a := application{}
		s := a.readCSV(q, "a_key", []string{"def1", "def2"})

		assert.Equal(t, len(s), 3)
		assert.Equal(t, s[0], "val1")
		assert.Equal(t, s[1], "val2")
		assert.Equal(t, s[2], "val3")
	})
	t.Run("test read absent query string key as string", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()

		a := application{}
		s := a.readCSV(q, "a_key", []string{"def1", "def2"})

		assert.Equal(t, len(s), 2)
		assert.Equal(t, s[0], "def1")
		assert.Equal(t, s[1], "def2")
	})
}

func TestReadInt(t *testing.T) {
	t.Run("test read query string key as int", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()
		q.Add("a_key", "11")

		a := application{}
		i := a.readInt(q, "a_key", 0, validator.New())

		assert.Equal(t, i, 11)
	})
	t.Run("test read absent query string key as int", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()

		a := application{}
		i := a.readInt(q, "a_key", 0, validator.New())

		assert.Equal(t, i, 0)
	})
	t.Run("test read non int query string key as int", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		q := req.URL.Query()
		q.Add("a_key", "NaN")

		a := application{}
		v := validator.New()
		i := a.readInt(q, "a_key", 0, v)

		assert.Equal(t, i, 0)
		assert.Equal(t, v.Errors["a_key"], "must be an integer value")
	})
}
