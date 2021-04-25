package utils

import "encoding/json"

func JsonStatus(message string) []byte {
	m, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	CheckErr(err)
	return m
}
