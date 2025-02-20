package go_fedex_restful_api_wrapper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Authorization(clientId, clientSecret string) ([]byte, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	// params := fmt.Sprintf("grant_type=%s&client_id=%s&client_secret=%s", grantType, clientId, clientSecret)
	// payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, ApiAuthTokenURL(), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return body, nil
}
