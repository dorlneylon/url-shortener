package shorten

import (
	"testing"
)

func TestShorten(t *testing.T) {
	shortUrl := Shorten("test")
	t.Log(shortUrl)
}
