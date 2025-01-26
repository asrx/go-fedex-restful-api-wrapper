package go_fedex_restful_api_wrapper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Fedex_authorization(grantType, clientId, clientSecret string) ([]byte, error) {
	params := fmt.Sprintf("grant_type=%s&client_id=%s&client_secret=%s", grantType, clientId, clientSecret)
	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, API_AuthTokenURL(), payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
