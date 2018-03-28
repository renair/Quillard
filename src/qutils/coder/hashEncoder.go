package coder

import (
	"crypto/sha1"
	"fmt"
)

func EncodeSha1(pass string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(pass)))
}
