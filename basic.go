package go_fedex_restful_api_wrapper

const (
	api_basic_url         = "https://apis.fedex.com"
	api_basic_url_sandbox = "https://apis-sandbox.fedex.com"

	api_auth_token        = "/oauth/token"
	api_tracking          = "/track/v1/trackingnumbers"
	api_addressValidation = "/address/v1/addresses/resolve"
	api_rating            = "/"
	api_shipping          = "/"
	api_void              = "/"
)

func getBasicUrl() string {
	// return api_basic_url_sandbox
	return api_basic_url
}

func ApiAuthTokenURL() string {
	return getBasicUrl() + api_auth_token
}
func ApiTrackingURL() string {
	return getBasicUrl() + api_tracking
}
func ApiAddressValidationURL() string {
	return getBasicUrl() + api_addressValidation
}
func ApiRatingURL() string {
	return getBasicUrl() + api_rating
}
func ApiShippingURL() string {
	return getBasicUrl() + api_shipping
}
func ApiVoidURL() string {
	return getBasicUrl() + api_void
}

// Token 结构体
type FedExTokenSt struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
