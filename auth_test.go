package go_fedex_restful_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_fedexAuthorization(t *testing.T) {

	var grantType, clientId, clientSecret = "client_credentials", "l775dddff0953f4d4aa29aa9b30e68d24e", "dd4b25024bf3467a93606b1c57bf8d3e"

	authorization, err := Fedex_authorization(grantType, clientId, clientSecret)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println("Token:")
	fmt.Println(string(authorization))
}

func TestTokenSt(t *testing.T) {
	respBody := `{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJDWFMtVFAiXSwiUGF5bG9hZCI6eyJjbGllbnRJZGVudGl0eSI6eyJjbGllbnRLZXkiOiJsNzc1ZGRkZmYwOTUzZjRkNGFhMjlhYTliMzBlNjhkMjRlIn0sImF1dGhlbnRpY2F0aW9uUmVhbG0iOiJDTUFDIiwiYWRkaXRpb25hbElkZW50aXR5Ijp7InRpbWVTdGFtcCI6IjI2LUphbi0yMDI1IDA0OjE3OjAxIEVTVCIsImdyYW50X3R5cGUiOiJjbGllbnRfY3JlZGVudGlhbHMiLCJhcGltb2RlIjoiU2FuZGJveCIsImN4c0lzcyI6Imh0dHBzOi8vY3hzYXV0aHNlcnZlci1zdGFnaW5nLmFwcC5wYWFzLmZlZGV4LmNvbS90b2tlbi9vYXV0aDIifSwicGVyc29uYVR5cGUiOiJEaXJlY3RJbnRlZ3JhdG9yX0IyQiJ9LCJleHAiOjE3Mzc4ODY2MjEsImp0aSI6IjEzOTRjNDM3LTQ5MmQtNGU4ZC04YzU0LTg5MDJkMjIxMjg0ZCJ9.UsXbhMiLpJmKo8_ttC4wtpGuObUzLoFBpVKnJpyuaESigsdOSW99XWSN3K_hjq1L_SIexitx50ymimNKIs1-0UZBnJX3qz4sSokoSCoNMAl7tKJwWYeZUystpTVd8aF34EcvyJYfwWQnW04nT03jf_FWY25Ss0ACWjDIKK4tQ2lnfMEsU8pSAncGemo87Gtioa-C7A-c7i4tFCxtzutpb1p8NGMA_Fc6gF7aU7QuE9PBKQBIw0xN75BuPWX9cEvdEKIYYo0M3q_zwouJ8OE3MdwhAKMKCy7cDlpJUUVZ0gb-6-JIH6iYGGJL-eulk1_4NnogABZ1mFN30P_jL0pmiae9YVkibCHhAnxUyAiH9rga1qqjZcWDTiM2a_Hz7Nz9GX95LEdIPOksoXa0lSKzSsJ-Rrr8GNK1NrODJlFPS_SVJM0Je-nOTTmfb8GlxcaJhwHKpKh9LSw9jY5xyC01uxndvzAmg8GnWKZINxyHX3jA1qWHu2tfr3LLc1n3FOZETGchjtyGj-MQtfDHlF-7dKmenQcMhaoeASCsPPV6Zi3RbQh-glLn_VGZIZFdgaUCA6DFsto1ZKO8wDDrYyLIF8JraY7ydaN3_oQtBP66FD0NJDs_KxNGwuXebzYxzN6KJ-uc3obiY0qsywniV7b2JzbgxKl65E9gstQoFe7V-Q0","token_type":"bearer","expires_in":3599,"scope":"CXS-TP"}`
	fedExTokenSt := &FedExTokenSt{
		AccessToken: "",
		TokenType:   "",
		ExpiresIn:   0,
		Scope:       "",
	}
	err := json.Unmarshal([]byte(respBody), fedExTokenSt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", fedExTokenSt)
}
