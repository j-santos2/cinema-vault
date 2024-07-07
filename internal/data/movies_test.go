package data

import (
	"strings"
	"testing"
	"time"

	"github.com/j-santos2/cinema-vault/internal/assert"
	"github.com/j-santos2/cinema-vault/internal/validator"
)

func TestValidateMovie(t *testing.T) {
	t.Run("Valid movie", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1981,
			Runtime: 115,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, v.Valid(), true)
	})

	t.Run("Movie with empty title", func(t *testing.T) {
		movie := Movie{
			Title:   "",
			Year:    1990,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie with title bigger than 500 bytes", func(t *testing.T) {
		movie := Movie{
			Title:   strings.Repeat("A", 501),
			Year:    1990,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie year before 1888", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1663,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie year after current year", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    int32(time.Now().Year()) + 1,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie year before 1888", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1663,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie year is not 0", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    0,
			Runtime: 89,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie runtime is greater than 0", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1981,
			Runtime: 0,
			Genres:  []string{"adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie genres has at least 1 genre", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1981,
			Runtime: 115,
			Genres:  []string{},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie genres has at most 5 genres ", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1981,
			Runtime: 115,
			Genres: []string{
				"action",
				"adventure",
				"historic",
				"supernatural",
				"classic",
				"mistery",
			},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Movie genres does not contain duplicated values ", func(t *testing.T) {
		movie := Movie{
			Title:   "Raiders of the Lost Ark",
			Year:    1981,
			Runtime: 115,
			Genres:  []string{"action", "adventure", "action"},
		}
		v := validator.New()
		ValidateMovie(v, &movie)
		assert.Equal(t, !v.Valid(), true)
	})
}
