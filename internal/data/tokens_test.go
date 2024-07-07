package data

import (
	"strings"
	"testing"
	"time"

	"github.com/j-santos2/cinema-vault/internal/assert"
	"github.com/j-santos2/cinema-vault/internal/validator"
)

func TestValidateToken(t *testing.T) {
	t.Run("Valid token plain text", func(t *testing.T) {
		token := strings.Repeat("A", 26)
		v := validator.New()
		ValidateTokenPlaintext(v, token)
		assert.Equal(t, v.Valid(), true)
	})

	t.Run("Invalid token plain text is not 26 chars long", func(t *testing.T) {
		token := "ABCABC"
		v := validator.New()
		ValidateTokenPlaintext(v, token)
		assert.Equal(t, !v.Valid(), true)
	})
}

func TestGenerateToken(t *testing.T) {
	t.Run("Generated token plain text is valid", func(t *testing.T) {
		token, _ := generateToken(1, 1*time.Hour, "scope")
		v := validator.New()
		ValidateTokenPlaintext(v, token.Plaintext)
		assert.Equal(t, v.Valid(), true)
	})

	t.Run("Generated token expiry time is valid", func(t *testing.T) {
		now := time.Now().Unix()
		token, _ := generateToken(1, 1*time.Hour, "scope")
		time_diff := token.Expiry.Unix() - now
		assert.Equal(t, time_diff, int64(1*time.Hour.Seconds()))
	})
}
