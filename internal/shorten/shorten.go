package shorten

import "math/rand"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var alphabetLen = len(alphabet)

func Shorten(link string) string {
	var (
		shortUrl string
	)

	for i := uint32(0); i < min(8, uint32(len(link))); i++ {
		shortUrl += string(alphabet[rand.Intn(alphabetLen)])
	}

	return shortUrl
}
