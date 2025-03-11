package andromeda

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// request содержит данные для HTTP-запроса
type request struct {
	URL    string
	body   []byte
	apiKey string
}

func buildURL(host, endpoint string, params map[string]string) string {
	baseURL, _ := url.Parse(host + endpoint)
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}
	baseURL.RawQuery = q.Encode()
	return baseURL.String()
}

// generateRequest создает запрос для метода GetSites
func (i GetSitesInput) generateRequest() request {

	param := make(map[string]string)
	param["id"] = i.Id
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetSites, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCustomers
func (i GetCustomersInput) generateRequest() request {

	param := make(map[string]string)
	param["siteId"] = i.SiteId
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetCustomers, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCustomer
func (i GetCustomerInput) generateRequest() request {

	param := make(map[string]string)
	param["id"] = i.Id
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetCustomers, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PostCheckPanic
func (i PostCheckPanicInput) generateRequest() request {

	param := make(map[string]string)
	param["siteId"] = i.SiteId
	param["stopOnEvent"] = "True"
	if i.CheckInterval != 0 {
		param["checkInterval"] = strconv.Itoa(i.CheckInterval)
	}
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointCheckPanic, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCheckPanic
func (i GetCheckPanicInput) generateRequest() request {

	param := make(map[string]string)
	param["checkPanicId"] = i.CheckPanicId
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointCheckPanic, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetUsersMyAlarm
func (i GetUsersMyAlarmInput) generateRequest() request {

	param := make(map[string]string)
	param["siteId"] = i.SiteId
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointMyAlarm, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PutChangeUserMyAlarm
func (i PutChangeUserMyAlarmInput) generateRequest() request {

	param := make(map[string]string)
	param["custId"] = i.CustId
	param["role"] = i.Role
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointMyAlarm, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PutChangeKTSUserMyAlarm
func (i PutChangeKTSUserMyAlarmInput) generateRequest() request {

	param := make(map[string]string)
	param["custId"] = i.CustId
	param["isPanic"] = strconv.FormatBool(i.IsPanic)
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointMyAlarm, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetUserObjectMyAlarm
func (i GetUserObjectMyAlarmInput) generateRequest() request {

	jsonData, _ := json.Marshal(i)

	param := make(map[string]string)
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetUserObjectMyAlarm, param),
		body:   jsonData,
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetParts
func (i GetPartsInput) generateRequest() request {

	param := make(map[string]string)
	param["siteId"] = i.SiteId
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetParts, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetZones
func (i GetZonesInput) generateRequest() request {

	param := make(map[string]string)
	param["siteId"] = i.SiteId
	if i.UserName != "" {
		param["userName"] = i.UserName
	}

	return request{
		URL:    buildURL(i.Host, endpointGetZones, param),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}
