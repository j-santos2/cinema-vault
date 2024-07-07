package data

import (
	"strings"
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
	"github.com/j-santos2/cinema-vault/internal/validator"
)

func TestValidateUser(t *testing.T) {
	t.Run("Valid user", func(t *testing.T) {
		pass := "A_Password"
		user := User{
			Name:  "Joe",
			Email: "joe@example.com",
		}
		user.Password.Set(pass)
		v := validator.New()
		ValidateUser(v, &user)
		assert.Equal(t, v.Valid(), true)
	})

	t.Run("Invalid user without name", func(t *testing.T) {
		pass := "A_Password"
		user := User{
			Name:  "",
			Email: "joe@example.com",
		}
		user.Password.Set(pass)
		v := validator.New()
		ValidateUser(v, &user)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Invalid user with a name larger than 500 bytes", func(t *testing.T) {
		pass := "A_Password"
		user := User{
			Name:  strings.Repeat("a", 501),
			Email: "joe@example.com",
		}
		user.Password.Set(pass)
		v := validator.New()
		ValidateUser(v, &user)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Invalid email format", func(t *testing.T) {
		email := "joeexample.com"
		v := validator.New()
		ValidateEmail(v, email)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Invalid password shorter than 8 bytes", func(t *testing.T) {
		pass := "short"
		v := validator.New()
		ValidatePasswordPlaintext(v, pass)
		assert.Equal(t, !v.Valid(), true)
	})

	t.Run("Invalid password longer than 72 bytes", func(t *testing.T) {
		pass := strings.Repeat("a", 73)
		v := validator.New()
		ValidatePasswordPlaintext(v, pass)
		assert.Equal(t, !v.Valid(), true)
	})
}
