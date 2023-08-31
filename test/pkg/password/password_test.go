package password

import (
	"github.com/adykaaa/online-notes/lib/random"
	"github.com/stretchr/testify/require"
	"go-online-notes/pkg/password"
	"testing"
)

func TestHashPasswordPositive(t *testing.T) {
	t.Run("password hashing works correctly",
		func(t *testing.T) {
			hpw, err := password.Hash(random.NewString(5))
			require.NoError(t, err)
			require.NotEmpty(t, hpw)
		})
}

func TestHashPasswordThrowException(t *testing.T) {
	t.Run("password hashing throws exception with string shorter than 5 characters",
		func(t *testing.T) {
			hpw, err := password.Hash(random.NewString(4))
			require.Error(t, err)
			require.Empty(t, hpw)
		})
}

func TestHashPasswordWithEmptyString(t *testing.T) {
	t.Run("password hashing throws exception with empty string",
		func(t *testing.T) {
			hpw, err := password.Hash("")
			require.Error(t, err)
			require.Empty(t, hpw)
		})
}

func TestValidatePositive(t *testing.T) {
	const passwordSample = "abs123!"
	t.Run("password validation works correctly",
		func(t *testing.T) {
			hpw, _ := password.Hash(passwordSample)
			err := password.Validate(hpw, passwordSample)
			require.NoError(t, err)
		})
}

func TestValidateNegative(t *testing.T) {
	const passwordCorrect = "abs123!"
	const passwordIncorrect = "abc123!"
	t.Run("password validation throws exception with incorrect password",
		func(t *testing.T) {
			hpw, _ := password.Hash(passwordCorrect)
			err := password.Validate(hpw, passwordIncorrect)
			require.Error(t, err)
		})
}
