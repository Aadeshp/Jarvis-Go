package rest

import (
	"encoding/json"
)

func Serialize(val interface{}) ([]byte, error) {
	ret, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	return ret, err
}
