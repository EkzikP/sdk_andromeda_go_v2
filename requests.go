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

// generateRequest создает запрос для метода GetSites
func (i GetSitesInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetSites)
	param := url.Values{}
	param.Add("id", i.Id)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCustomers
func (i GetCustomersInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetCustomers)
	param := url.Values{}
	param.Add("siteId", i.SiteId)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCustomer
func (i GetCustomerInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetCustomers)
	param := url.Values{}
	param.Add("id", i.Id)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PostCheckPanic
func (i PostCheckPanicInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointCheckPanic)
	param := url.Values{}
	param.Add("siteId", i.SiteId)
	param.Add("stopOnEvent", "True")
	if i.CheckInterval != 0 {
		param.Add("checkInterval", strconv.Itoa(i.CheckInterval))
	}
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetCheckPanic
func (i GetCheckPanicInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointCheckPanic)
	param := url.Values{}
	param.Add("checkPanicId", i.CheckPanicId)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetUsersMyAlarm
func (i GetUsersMyAlarmInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointMyAlarm)
	param := url.Values{}
	param.Add("siteId", i.SiteId)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PutChangeUserMyAlarm
func (i PutChangeUserMyAlarmInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointMyAlarm)
	param := url.Values{}
	param.Add("custId", i.CustId)
	param.Add("role", i.Role)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода PutChangeKTSUserMyAlarm
func (i PutChangeKTSUserMyAlarmInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointMyAlarm)
	param := url.Values{}
	param.Add("custId", i.CustId)
	param.Add("isPanic", strconv.FormatBool(i.IsPanic))
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetUserObjectMyAlarm
func (i GetUserObjectMyAlarmInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetUserObjectMyAlarm)
	param := url.Values{}
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()
	jsonData, _ := json.Marshal(i)

	return request{
		URL:    baseURL.String(),
		body:   jsonData,
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetParts
func (i GetPartsInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetParts)
	param := url.Values{}
	param.Add("siteId", i.SiteId)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}

// generateRequest создает запрос для метода GetZones
func (i GetZonesInput) generateRequest() request {
	baseURL, _ := url.Parse(i.Host + endpointGetZones)
	param := url.Values{}
	param.Add("siteId", i.SiteId)
	if i.UserName != "" {
		param.Add("userName", i.UserName)
	}
	baseURL.RawQuery = param.Encode()

	return request{
		URL:    baseURL.String(),
		body:   []byte{},
		apiKey: i.ApiKey,
	}
}
