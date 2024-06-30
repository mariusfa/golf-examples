package auth

import "testing"

func TestCreateAndParseToken(t *testing.T) {
	secret := "secret"
	token, err := CreateToken("user", secret)
	if err != nil {
		t.Errorf("Error creating token: %s", err)
	}

	parsedUser, err := ParseToken(token, secret)
	if err != nil {
		t.Errorf("Error parsing token: %s", err)
	}
	if parsedUser != "user" {
		t.Errorf("Expected user to be 'user', got '%s'", parsedUser)
	}
}
