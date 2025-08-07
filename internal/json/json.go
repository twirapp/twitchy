package json

import (
	"encoding/json"
)

type RawMessage = json.RawMessage

type (
	Marshaller   func(v any) ([]byte, error)
	UnMarshaller func(data []byte, v any) error
)

var (
	Marshal   Marshaller
	Unmarshal UnMarshaller
)

func init() {
	Marshal = json.Marshal
	Unmarshal = json.Unmarshal
}
