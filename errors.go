package andromeda

// respErr400 содержит структуру для парсинга ошибки 400 от сервера
type respErr400 struct {
	Message      string `json:"Message"`
	SpResultCode int    `json:"SpResultCode"`
}
