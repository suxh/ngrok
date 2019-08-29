package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {
	path := "./htpasswd"
	auth, err := NewAuth(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(auth.Tokens)

	token := "username:password"
	if err := auth.Auth(token); err != nil {
		t.Fatal(err)
	}
}
