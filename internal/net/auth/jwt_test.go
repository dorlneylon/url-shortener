package auth

import (
	"testing"
)

func TestMakeJWT(t *testing.T) {
	token, err := MakeJWT("1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestDecodeJWT(t *testing.T) {
	token, err := MakeJWT("1")
	if err != nil {
		t.Fatal(err)
	}
	userId, err := DecodeJWT(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userId)
}
