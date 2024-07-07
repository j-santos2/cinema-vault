package data

import (
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
)

func TestMarshalJSON(t *testing.T) {
	t.Run("MarshalJSON valid input", func(t *testing.T) {
		r := Runtime(2)
		marshal, err := r.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, string(marshal), `"2 mins"`)
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("UnmarshalJSON valid input", func(t *testing.T) {
		jsonValue := `"2 mins"`
		r := Runtime(0)
		err := r.UnmarshalJSON([]byte(jsonValue))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, r, 2)
	})

	t.Run("MarshalJSON unquoted value", func(t *testing.T) {
		jsonValue := `2 mins`
		r := Runtime(0)
		err := r.UnmarshalJSON([]byte(jsonValue))
		assert.Equal(t, err.Error(), "invalid runtime format")
	})

	t.Run("MarshalJSON invalid units", func(t *testing.T) {
		jsonValue := `"2 minutes"`
		r := Runtime(0)
		err := r.UnmarshalJSON([]byte(jsonValue))
		assert.Equal(t, err.Error(), "invalid runtime format")
	})

	t.Run("MarshalJSON integer greater than int32", func(t *testing.T) {
		jsonValue := `"8589934592 mins"`
		r := Runtime(0)
		err := r.UnmarshalJSON([]byte(jsonValue))
		assert.Equal(t, err.Error(), "invalid runtime format")
	})
}
