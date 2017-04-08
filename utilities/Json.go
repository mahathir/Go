package utilities

import (
	"encoding/json"
	"io"
)

func JsonDecode(body io.Reader, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}

func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
