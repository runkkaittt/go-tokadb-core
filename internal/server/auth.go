package server

import (
	"crypto/sha256"
	"fmt"
)

func createAuthToken(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

func (s *DBServer) SetAuthToken(str string) {
	s.AuthToken = createAuthToken(str)
}
