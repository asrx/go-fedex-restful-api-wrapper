package go_fedex_restful_api_wrapper

// 响应参数
type ValidateAddressResp struct {
	TransactionId string      `json:"transactionId"`
	Output        output      `json:"output"`
	Errors        interface{} `json:"errors"`
}

type output struct {
	ResolvedAddresses []resolvedAddress `json:"resolvedAddresses"`
}

type resolvedAddress struct {
	StreetLinesToken                  []string         `json:"streetLinesToken"`
	City                              string           `json:"city"`
	StateOrProvinceCode               string           `json:"stateOrProvinceCode"`
	PostalCode                        string           `json:"postalCode"`
	ParsedPostalCode                  parsedPostalCode `json:"parsedPostalCode"`
	CountryCode                       string           `json:"countryCode"`
	Classification                    string           `json:"classification"`
	RuralRouteHighwayContract         bool             `json:"ruralRouteHighwayContract"`
	GeneralDelivery                   bool             `json:"generalDelivery"`
	CustomerMessages                  []interface{}    `json:"customerMessages"`
	StandardizedStatusNameMatchSource string           `json:"standardizedStatusNameMatchSource"`
	ResolutionMethodName              string           `json:"resolutionMethodName"`
	Attributes                        attributes       `json:"attributes"`
}

type parsedPostalCode struct {
	Base          string `json:"base"`
	AddOn         string `json:"addOn"`
	DeliveryPoint string `json:"deliveryPoint"`
}

type attributes struct {
	POBox                     string `json:"POBox"`
	Intersection              string `json:"Intersection"`
	SuiteRequiredButMissing   string `json:"SuiteRequiredButMissing"`
	InvalidSuiteNumber        string `json:"InvalidSuiteNumber"`
	ResolutionInput           string `json:"ResolutionInput"`
	DPV                       string `json:"DPV"`
	ResolutionMethod          string `json:"ResolutionMethod"`
	DataVintage               string `json:"DataVintage"`
	MatchSource               string `json:"MatchSource"`
	CountrySupported          string `json:"CountrySupported"`
	ValidlyFormed             string `json:"ValidlyFormed"`
	Matched                   string `json:"Matched"`
	Resolved                  string `json:"Resolved"`
	Inserted                  string `json:"Inserted"`
	InterpolatedStreetAddress string `json:"InterpolatedStreetAddress"`
	MultiUnitBase             string `json:"MultiUnitBase"`
	ZIP11Match                string `json:"ZIP11Match"`
	ZIP4Match                 string `json:"ZIP4Match"`
	StreetRange               string `json:"StreetRange"`
	UniqueZIP                 string `json:"UniqueZIP"`
	RRConversion              string `json:"RRConversion"`
	ValidMultiUnit            string `json:"ValidMultiUnit"`
	AddressType               string `json:"AddressType"`
	MultipleMatches           string `json:"MultipleMatches"`
}

func NewValidateAddressResp() *ValidateAddressResp {
	return &ValidateAddressResp{}
}

func FedEx_AddressValidation(streetLines []string, city, stateOrProvinceCode, postalCode, countryCode string) {

}
