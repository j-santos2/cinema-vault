package data

import (
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
	"github.com/j-santos2/cinema-vault/internal/validator"
)

func TestValidateFilters(t *testing.T) {
	t.Run("Valid filters", func(t *testing.T) {
		f := Filters{
			Page:         1,
			PageSize:     10,
			Sort:         "safeItem",
			SortSafeList: []string{"safeItem"},
		}
		v := validator.New()
		ValidateFilters(v, f)
		assert.Equal(t, v.Valid(), true)
	})

	t.Run("Invalid filters sort resource not in safe list", func(t *testing.T) {
		f := Filters{
			Page:         1,
			PageSize:     10,
			Sort:         "unsafe",
			SortSafeList: []string{"safeItem"},
		}
		v := validator.New()
		ValidateFilters(v, f)
		assert.Equal(t, !v.Valid(), true)
	})
}
