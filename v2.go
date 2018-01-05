package orzo

import (
	"crypto/hmac"
	"crypto/sha512"

	"encoding/base64"
)

const (
	version = "v2"
)

// Hash
func NewHMACKey() *[32]byte {
	key := &[32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return key
}

func Auth(key *[32]byte, data []byte) string {
	header := version + ".auth."

	h := hmac.New(sha512.New512_256(), key[:])
	h.Write(data)

	return HEADER + base64.URLEncoding.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func AuthWithFooter(key *[32]byte, data []byte, string footer) string {
	return fmt.Sprintf("%s.%s", auth(key, data), footer)
}

func AuthVerify(authMsg string, key *[32]byte) {

}
