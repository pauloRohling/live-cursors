package json

import (
	"bytes"
	"encoding/json"
)

func Serialize[T any](message T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(message)
	return buffer.Bytes(), err
}

func Deserialize[T any](data []byte) (*T, error) {
	var message T
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)
	err := decoder.Decode(&message)
	return &message, err
}
