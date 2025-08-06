package json

import (
	"encoding/json"
)

type RawMessage = json.RawMessage

type (
	MarshalFn   func(v any) ([]byte, error)
	UnmarshalFn func(data []byte, v any) error
)

var (
	Marshal   MarshalFn
	Unmarshal UnmarshalFn
)

func init() {
	Marshal = json.Marshal
	Unmarshal = json.Unmarshal
}
