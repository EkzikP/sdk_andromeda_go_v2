package andromeda

import "github.com/pkg/errors"

func (c *Config) validateBase() error {
	if c.ApiKey == "" {
		return errors.New("неверно задан API ключ")
	}
	if c.Host == "" {
		return errors.New("неверно задан адрес сервера")
	}
	return nil
}

// validate проверяет заполнение обязательных полей метода GetSites
func (i GetSitesInput) validate() error {
	if i.Id == "" {
		return errors.New("неверно задан номер объекта")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetCustomers
func (i GetCustomersInput) validate() error {
	if i.SiteId == "" {
		return errors.New("неверно задан номер объекта")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetCustomer
func (i GetCustomerInput) validate() error {
	if i.Id == "" {
		return errors.New("неверно задан идентификатор ответственного лица")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода PostCheckPanic
func (i PostCheckPanicInput) validate() error {
	if i.SiteId == "" {
		return errors.New("неверно задан номер объекта")
	}

	if i.CheckInterval != 0 && (i.CheckInterval < 30 || i.CheckInterval > 180) {
		return errors.New("неверно задано время ожидания проверки, должно быть от 30 до 180 секунд.")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetCheckPanic
func (i GetCheckPanicInput) validate() error {
	if i.CheckPanicId == "" {
		return errors.New("неверно задан идентификатор проверки")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetUsersMyAlarm
func (i GetUsersMyAlarmInput) validate() error {
	if i.SiteId == "" {
		return errors.New("неверно задан идентификатор объекта")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetUserObjectMyAlarm
func (i GetUserObjectMyAlarmInput) validate() error {
	if i.Phone == "" {
		return errors.New("неверно задан номер телефона")
	}

	if len(i.Phone) != 12 && i.Phone[:2] != "+7" {
		return errors.New("неверно задан номер телефона")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода PutChangeUserMyAlarm
func (i PutChangeUserMyAlarmInput) validate() error {
	if i.CustId == "" {
		return errors.New("неверно задан идентификатор пользователя")
	}

	if i.Role != "admin" && i.Role != "user" && i.Role != "unlink" {
		return errors.New("неверно задана роль пользователя")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода PutChangeKTSUserMyAlarm
func (i PutChangeKTSUserMyAlarmInput) validate() error {
	if i.CustId == "" {
		return errors.New("неверно задан идентификатор пользователя")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetParts
func (i GetPartsInput) validate() error {
	if i.SiteId == "" {
		return errors.New("неверно задан идентификатор пользователя")
	}

	return i.Config.validateBase()

}

// validate проверяет заполнение обязательных полей метода GetZones
func (i GetZonesInput) validate() error {
	if i.SiteId == "" {
		return errors.New("неверно задан идентификатор пользователя")
	}

	return i.Config.validateBase()

}
