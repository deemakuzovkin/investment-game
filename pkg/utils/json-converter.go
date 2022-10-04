package utils

import "encoding/json"

// ToJsonBytes to JSON bytes
func ToJsonBytes(input interface{}) []byte {
	dataBytes, err := json.Marshal(input)
	if err != nil {
		return nil
	}
	return dataBytes
}
