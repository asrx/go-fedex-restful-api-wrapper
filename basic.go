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
	return api_basic_url_sandbox
	// return api_basic_url
}

func API_AuthTokenURL() string {
	return getBasicUrl() + api_auth_token
}
func API_TrackingURL() string {
	return getBasicUrl() + api_tracking
}
func API_AddressValidationURL() string {
	return getBasicUrl() + api_addressValidation
}
func API_RatingURL() string {
	return getBasicUrl() + api_rating
}
func API_ShippingURL() string {
	return getBasicUrl() + api_shipping
}
func API_VoidURL() string {
	return getBasicUrl() + api_void
}

type FedExTokenSt struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
