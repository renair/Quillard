package coder

import (
	"encoding/json"
	"io"
	"strings"
)

func DecodeJson(io io.Reader, v interface{}) error {
	decoder := json.NewDecoder(io)
	return decoder.Decode(v)
}

func EncodeJson(v interface{}) string {
	var builder strings.Builder
	encoder := json.NewEncoder(&builder)
	encoder.Encode(v)
	return builder.String()
}
