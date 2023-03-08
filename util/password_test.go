package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hasedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasedPassword1)

	err = CheckPassword(password, hasedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hasedPassword1)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hasedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasedPassword2)
	require.NotEqual(t, hasedPassword1, hasedPassword2)
}
