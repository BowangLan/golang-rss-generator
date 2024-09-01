package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKeyFromHeaders(r *http.Request) (string, error) {
	apiKey := r.Header.Get("Authorization")
	if apiKey == "" {
		return "", fmt.Errorf("API key is required")
	}

	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", fmt.Errorf("Invalid API key format")
	}

	if vals[0] != "Bearer" {
		return "", fmt.Errorf("Invalid API key format")
	}

	return vals[1], nil
}
