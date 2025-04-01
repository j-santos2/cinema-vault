package data

import (
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
)

func TestIncludes(t *testing.T) {
	t.Run("Permission is included", func(t *testing.T) {
		p := Permissions{"P1", "P2"}
		included := p.Includes("P1")
		assert.Equal(t, included, true)
	})

	t.Run("Permission is not included", func(t *testing.T) {
		p := Permissions{"P1", "P2"}
		included := p.Includes("does-not-exist")
		assert.Equal(t, included, false)
	})
}
