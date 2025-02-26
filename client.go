package andromeda

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

// Client представляет клиент для работы с API
type Client struct {
	client *http.Client
}

// NewClient создает новый экземпляр клиента
func NewClient() *Client {
	return &Client{
		client: &http.Client{Timeout: defaultTimeout},
	}
}

// GetSites выполняет запрос к методу GetSites
func (c *Client) GetSites(ctx context.Context, input GetSitesInput) (GetSitesResponse, error) {
	if err := input.validate(); err != nil {
		return GetSitesResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return GetSitesResponse{}, err
	}

	var resp GetSitesResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return GetSitesResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// Customers выполняет запрос к методу GetCustomers
func (c *Client) Customers(ctx context.Context, input GetCustomersInput) ([]GetCustomerResponse, error) {
	if err := input.validate(); err != nil {
		return []GetCustomerResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return []GetCustomerResponse{}, err
	}

	resp := []GetCustomerResponse{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []GetCustomerResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// GetCustomer выполняет запрос к методу GetCustomer
func (c *Client) GetCustomer(ctx context.Context, input GetCustomerInput) (GetCustomerResponse, error) {
	if err := input.validate(); err != nil {
		return GetCustomerResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return GetCustomerResponse{}, err
	}

	resp := GetCustomerResponse{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return GetCustomerResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// PostCheckPanic выполняет запрос к методу PostCheckPanic
func (c *Client) PostCheckPanic(ctx context.Context, input PostCheckPanicInput) (PostCheckPanicResponse, error) {
	if err := input.validate(); err != nil {
		return PostCheckPanicResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodPost, req)
	if err != nil {
		return PostCheckPanicResponse{}, err
	}

	var resp PostCheckPanicResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return PostCheckPanicResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// GetCheckPanic выполняет запрос к методу GetCheckPanic
func (c *Client) GetCheckPanic(ctx context.Context, input GetCheckPanicInput) (GetCheckPanicResponse, error) {
	if err := input.validate(); err != nil {
		return GetCheckPanicResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return GetCheckPanicResponse{}, err
	}

	var resp GetCheckPanicResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return GetCheckPanicResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// GetUsersMyAlarm выполняет запрос к методу GetUsersMyAlarm
func (c *Client) GetUsersMyAlarm(ctx context.Context, input GetUsersMyAlarmInput) ([]UserMyAlarmResponse, error) {
	if err := input.validate(); err != nil {
		return []UserMyAlarmResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return []UserMyAlarmResponse{}, err
	}

	var resp []UserMyAlarmResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []UserMyAlarmResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// PutChangeUserMyAlarm выполняет запрос к методу PutChangeUserMyAlarm
func (c *Client) PutChangeUserMyAlarm(ctx context.Context, input PutChangeUserMyAlarmInput) (PutChangeUserMyAlarmResponse, error) {
	if err := input.validate(); err != nil {
		return PutChangeUserMyAlarmResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodPut, req)
	if err != nil {
		return PutChangeUserMyAlarmResponse{}, err
	}

	var resp PutChangeUserMyAlarmResponse

	if len(body) != 0 {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return PutChangeUserMyAlarmResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
		}
	}

	return resp, nil
}

// GetUserObjectMyAlarm выполняет запрос к методу GetUserObjectMyAlarm
func (c *Client) GetUserObjectMyAlarm(ctx context.Context, input GetUserObjectMyAlarmInput) ([]GetUserObjectMyAlarmResponse, error) {
	if err := input.validate(); err != nil {
		return []GetUserObjectMyAlarmResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return []GetUserObjectMyAlarmResponse{}, err
	}

	var resp []GetUserObjectMyAlarmResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []GetUserObjectMyAlarmResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// PutChangeKTSUserMyAlarm выполняет запрос к методу PutChangeKTSUserMyAlarm
func (c *Client) PutChangeKTSUserMyAlarm(ctx context.Context, input PutChangeKTSUserMyAlarmInput) error {
	if err := input.validate(); err != nil {
		return err
	}

	req := input.generateRequest()
	_, err := c.doHTTP(ctx, http.MethodPut, req)
	if err != nil {
		return err
	}

	return nil
}

// GetParts выполняет запрос к методу GetParts
func (c *Client) GetParts(ctx context.Context, input GetPartsInput) ([]GetPartsResponse, error) {
	if err := input.validate(); err != nil {
		return []GetPartsResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return []GetPartsResponse{}, err
	}

	var resp []GetPartsResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []GetPartsResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// GetZones выполняет запрос к методу GetZones
func (c *Client) GetZones(ctx context.Context, input GetZonesInput) ([]GetZonesResponse, error) {
	if err := input.validate(); err != nil {
		return []GetZonesResponse{}, err
	}

	req := input.generateRequest()
	body, err := c.doHTTP(ctx, http.MethodGet, req)
	if err != nil {
		return []GetZonesResponse{}, err
	}

	var resp []GetZonesResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return []GetZonesResponse{}, errors.WithMessage(err, "Не удалось парсить ответ")
	}

	return resp, nil
}

// doHTTP выполняет HTTP-запрос
func (c *Client) doHTTP(ctx context.Context, method string, r request) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, r.URL, bytes.NewBuffer(r.body))
	if err != nil {
		return []byte{}, errors.WithMessage(err, "Не удалось создать запрос")
	}

	req.Header.Set("apiKey", r.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, errors.WithMessage(err, "Не удалось выполнить запрос")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, resp.Body); err != nil {
			return []byte{}, errors.WithMessage(err, "Не удалось выполнить запрос")
		}
		err400 := respErr400{}
		err = json.Unmarshal(buf.Bytes(), &err400)
		if err != nil {
			return []byte{}, err
		}
		return []byte{}, errors.New(err400.Message)
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, errors.New("Не удалось выполнить запрос")
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return []byte{}, errors.WithMessage(err, "Не удалось парсинг ответа")
	}

	return buf.Bytes(), nil
}
