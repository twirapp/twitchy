package twitchy

import (
	"errors"

	"github.com/kvizyx/twitchy/internal/json"
)

func SetJSONMarshal(marshal json.MarshalFn) error {
	if marshal == nil {
		return errors.New("json marshal function is not provided")
	}

	json.Marshal = marshal
	return nil
}

func SetJSONUnmarshal(unmarshal json.UnmarshalFn) error {
	if unmarshal == nil {
		return errors.New("json unmarshal function is not provided")
	}

	json.Unmarshal = unmarshal
	return nil
}
