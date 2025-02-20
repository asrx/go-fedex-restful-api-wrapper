package go_fedex_restful_api_wrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/golang-module/carbon"
)

// ===== 跟踪-请求参数结构体 start ======
type TrackReq struct {
	TrackingInfo         []TrackingInfo `json:"trackingInfo"`
	IncludeDetailedScans bool           `json:"includeDetailedScans"`
}

type TrackingInfo struct {
	ShipDateBegin      string             `json:"shipDateBegin,omitempty"`
	ShipDateEnd        string             `json:"shipDateEnd,omitempty"`
	TrackingNumberInfo TrackingNumberInfo `json:"trackingNumberInfo"`
}

type TrackingNumberInfo struct {
	TrackingNumber string `json:"trackingNumber"`
}

func NewTrackingInfoList(numbers []string) []TrackingInfo {
	var tracks = []TrackingInfo{}

	for _, v := range numbers {
		if v != "" {
			tracks = append(tracks, NewTrackingInfo(v))
		}
	}

	return tracks
}

func NewTrackingInfo(number string) TrackingInfo {
	return TrackingInfo{
		ShipDateBegin: carbon.Now().AddDays(-180).ToDateString(),
		ShipDateEnd:   carbon.Now().ToDateString(),
		TrackingNumberInfo: TrackingNumberInfo{
			TrackingNumber: number,
		},
	}
}

// ===== 跟踪-请求参数结构体 end ======
// ===== 跟踪-响应参数结构体 start ======

type TrackResp struct {
	TransactionId    string `json:"transactionId"`
	Output           Output `json:"output"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewTrackResp() *TrackResp {
	return &TrackResp{}
}

type Output struct {
	CompleteTrackResults []CompleteTrackResults `json:"completeTrackResults"`
}

type CompleteTrackResults struct {
	TrackingNumber string        `json:"trackingNumber"`
	TrackResults   []TrackResult `json:"trackResults"`
}

type TrackResult struct {
	// TrackingNumberInfo struct {
	// 	TrackingNumber         string `json:"trackingNumber"`
	// 	TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
	// 	CarrierCode            string `json:"carrierCode"`
	// } `json:"trackingNumberInfo"`
	// AdditionalTrackingInfo struct {
	// 	Nickname           string `json:"nickname"`
	// 	PackageIdentifiers []struct {
	// 		Type                   string   `json:"type"`
	// 		Values                 []string `json:"values"`
	// 		TrackingNumberUniqueId string   `json:"trackingNumberUniqueId"`
	// 		CarrierCode            string   `json:"carrierCode"`
	// 	} `json:"packageIdentifiers"`
	// 	HasAssociatedShipments bool `json:"hasAssociatedShipments"`
	// } `json:"additionalTrackingInfo"`
	// ShipperInformation struct {
	// 	Contact struct {
	// 	} `json:"contact"`
	// 	Address struct {
	// 		City                string `json:"city"`
	// 		StateOrProvinceCode string `json:"stateOrProvinceCode"`
	// 		CountryCode         string `json:"countryCode"`
	// 		Residential         bool   `json:"residential"`
	// 		CountryName         string `json:"countryName"`
	// 	} `json:"address"`
	// } `json:"shipperInformation"`
	RecipientInformation struct {
		Contact struct {
		} `json:"contact"`
		Address struct {
			City                string `json:"city"`
			StateOrProvinceCode string `json:"stateOrProvinceCode"`
			CountryCode         string `json:"countryCode"`
			Residential         bool   `json:"residential"`
			CountryName         string `json:"countryName"`
		} `json:"address"`
	} `json:"recipientInformation"`
	// LatestStatusDetail struct {
	// 	Code           string `json:"code"`
	// 	DerivedCode    string `json:"derivedCode"`
	// 	StatusByLocale string `json:"statusByLocale"`
	// 	Description    string `json:"description"`
	// 	ScanLocation   struct {
	// 		City                string `json:"city"`
	// 		StateOrProvinceCode string `json:"stateOrProvinceCode"`
	// 		CountryCode         string `json:"countryCode"`
	// 		Residential         bool   `json:"residential"`
	// 		CountryName         string `json:"countryName"`
	// 	} `json:"scanLocation"`
	// } `json:"latestStatusDetail"`
	// DateAndTimes []struct {
	// 	Type     string    `json:"type"`
	// 	DateTime time.Time `json:"dateTime"`
	// } `json:"dateAndTimes"`
	// AvailableImages []struct {
	// 	Type string `json:"type"`
	// } `json:"availableImages"`
	// ShipmentDetails struct {
	// 	PossessionStatus bool `json:"possessionStatus"`
	// 	Weight           []struct {
	// 		Value string `json:"value"`
	// 		Unit  string `json:"unit"`
	// 	} `json:"weight"`
	// } `json:"shipmentDetails"`
	// AvailableNotifications []interface{} `json:"availableNotifications"`
	DeliveryDetails struct {
		// ActualDeliveryAddress struct {
		// 	City                string `json:"city"`
		// 	StateOrProvinceCode string `json:"stateOrProvinceCode"`
		// 	CountryCode         string `json:"countryCode"`
		// 	Residential         bool   `json:"residential"`
		// 	CountryName         string `json:"countryName"`
		// } `json:"actualDeliveryAddress"`
		// DeliveryAttempts string `json:"deliveryAttempts"`
		// 签收人
		ReceivedByName string `json:"receivedByName"`
		// DeliveryOptionEligibilityDetails []struct {
		// 	Option      string `json:"option"`
		// 	Eligibility string `json:"eligibility"`
		// } `json:"deliveryOptionEligibilityDetails"`
	} `json:"deliveryDetails"`
	// OriginLocation struct {
	// 	LocationContactAndAddress struct {
	// 		Address struct {
	// 			City                string `json:"city"`
	// 			StateOrProvinceCode string `json:"stateOrProvinceCode"`
	// 			CountryCode         string `json:"countryCode"`
	// 			Residential         bool   `json:"residential"`
	// 			CountryName         string `json:"countryName"`
	// 		} `json:"address"`
	// 	} `json:"locationContactAndAddress"`
	// } `json:"originLocation"`
	// LastUpdatedDestinationAddress struct {
	// 	City                string `json:"city"`
	// 	StateOrProvinceCode string `json:"stateOrProvinceCode"`
	// 	CountryCode         string `json:"countryCode"`
	// 	Residential         bool   `json:"residential"`
	// 	CountryName         string `json:"countryName"`
	// } `json:"lastUpdatedDestinationAddress"`
	// ServiceDetail struct {
	// 	Type             string `json:"type"`
	// 	Description      string `json:"description"`
	// 	ShortDescription string `json:"shortDescription"`
	// } `json:"serviceDetail"`
	// StandardTransitTimeWindow struct {
	// 	Window struct {
	// 		Ends time.Time `json:"ends"`
	// 	} `json:"window"`
	// } `json:"standardTransitTimeWindow"`
	// EstimatedDeliveryTimeWindow struct {
	// 	Window struct {
	// 	} `json:"window"`
	// } `json:"estimatedDeliveryTimeWindow"`
	// GoodsClassificationCode string `json:"goodsClassificationCode"`
	// ReturnDetail            struct {
	// } `json:"returnDetail"`
	PackageDetails PackageDetails `json:"packageDetails"`
	ScanEvents     []ScanEvent    `json:"scanEvents"`
}

// 包裹信息
type PackageDetails struct {
	PackagingDescription struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"packagingDescription"`
	PhysicalPackagingType string `json:"physicalPackagingType"`
	SequenceNumber        string `json:"sequenceNumber"`
	Count                 string `json:"count"`
	WeightAndDimensions   struct {
		Weight []struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"weight"`
		Dimensions []struct {
			Length int    `json:"length"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Units  string `json:"units"`
		} `json:"dimensions"`
	} `json:"weightAndDimensions"`
	PackageContent []interface{} `json:"packageContent"`
}

type ScanEvent struct {
	Date                 string `json:"date"`
	EventType            string `json:"eventType"`
	EventDescription     string `json:"eventDescription"`
	ExceptionCode        string `json:"exceptionCode"`
	ExceptionDescription string `json:"exceptionDescription"`
	ScanLocation         struct {
		StreetLines         []string `json:"streetLines"`
		City                string   `json:"city,omitempty"`
		StateOrProvinceCode string   `json:"stateOrProvinceCode,omitempty"`
		PostalCode          string   `json:"postalCode,omitempty"`
		CountryCode         string   `json:"countryCode,omitempty"`
		Residential         bool     `json:"residential"`
		CountryName         string   `json:"countryName,omitempty"`
	} `json:"scanLocation"`
	LocationType      string `json:"locationType"`
	DerivedStatusCode string `json:"derivedStatusCode"`
	DerivedStatus     string `json:"derivedStatus"`
	LocationId        string `json:"locationId,omitempty"`
	DelayDetail       struct {
		Type string `json:"type"`
	} `json:"delayDetail,omitempty"`
}

// ===== 跟踪-响应参数结构体 end ======
func Tracking(trackingNumbers []string, token string) ([]*TrackResp, error) {
	count := len(trackingNumbers)
	if count == 0 {
		return nil, errors.New("trackingNumbers is empty!")
	}

	const groupLimit = 30
	var wg sync.WaitGroup

	pipLineCount := count / 30
	i := count % 30
	if i > 0 {
		pipLineCount += 1
	}
	// 开启缓冲信道
	pipline := make(chan *TrackResp, pipLineCount)

	for i := 0; i < count; i += groupLimit {
		end := i + groupLimit
		if end > count {
			end = count
		}
		groupTrackings := trackingNumbers[i:end]
		wg.Add(1)
		go func(trackings []string) {
			defer wg.Done()
			getTracking(groupTrackings, pipline, token)
		}(groupTrackings)
	}

	// 关闭信道
	go func() {
		wg.Wait()
		close(pipline)
	}()

	results := make([]*TrackResp, 0)

	// results = append(results, <-pipline)
	for result := range pipline {
		if result != nil {
			results = append(results, result)
		}
	}
	return results, nil
}

func getTracking(trackingNumbers []string, pipline chan *TrackResp, token string) {
	params := TrackReq{
		TrackingInfo:         NewTrackingInfoList(trackingNumbers),
		IncludeDetailedScans: true,
	}
	reqBytes, err := json.Marshal(params)
	if err != nil {
		pipline <- nil // 或者可以返回一个包含错误信息的TrackResp
		return
	}

	payload := bytes.NewReader(reqBytes)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, ApiTrackingURL(), payload)
	if err != nil {
		pipline <- nil
		return
	}

	req.Header.Add("X-locale", "en_US")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		pipline <- nil
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	trackResp := NewTrackResp()
	_ = json.Unmarshal(body, trackResp)
	if trackResp.Error != "" || trackResp.Output.CompleteTrackResults == nil || len(trackResp.Output.CompleteTrackResults) == 0 {
		pipline <- nil
		return
	}

	pipline <- trackResp
}
