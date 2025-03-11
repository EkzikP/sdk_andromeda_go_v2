package andromeda

// Структуры входных параметров
type (
	// GetSitesInput содержит входные параметры для метода GetSites
	GetSitesInput struct {
		Id       string // Номер или идентификатор объекта
		UserName string // Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// GetCustomersInput содержит входные параметры для метода GetCustomers
	GetCustomersInput struct {
		SiteId   string //Идентификатор объекта, список ответственных лиц которого нужно получить. Соответствует полю Id карточки объекта
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}
	// GetCustomerInput содержит входные параметры для метода GetCustomer
	GetCustomerInput struct {
		Id       string //Идентификатор ответственного лица, информацию которого нужно получить
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// PostCheckPanicInput содержит входные параметры для метода PostCheckPanic
	PostCheckPanicInput struct {
		SiteId        string //Идентификатор объекта, по которому нужно проверить КТС
		CheckInterval int    //Интервал в секундах, в течении которого будет продолжаться процедура проверки КТС. (необязательное поле)
		StopOnEvent   bool   //Признак остановки проверки КТС. (необязательное поле)
		UserName      string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config        `json:"-"`
	}

	// GetCheckPanicInput содержит входные параметры для метода GetCheckPanic
	GetCheckPanicInput struct {
		CheckPanicId string //Идентификатор процедуры проверки, для которой нужно получить результат
		UserName     string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config       `json:"-"`
	}

	// GetUsersMyAlarmInput содержит входные параметры для метода GetUsersMyAlarm
	GetUsersMyAlarmInput struct {
		SiteId   string //Идентификатор объекта, список пользователей MyAlarm которого нужно получить. Соответствует полю Id карточки объекта
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// GetUserObjectMyAlarmInput содержит входные параметры для метода GetUserObjectMyAlarm
	GetUserObjectMyAlarmInput struct {
		Phone    string `json:"Phone"` //Телефон пользователя MyAlarm, для которого нужно получить список объектов
		UserName string `json:"-"`     //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// PutChangeUserMyAlarmInput содержит входные параметры для метода PutChangeUserMyAlarm
	PutChangeUserMyAlarmInput struct {
		CustId   string //Идентификатор пользователя
		Role     string //Роль пользователя, допустимые значения: “unlink”, “user”, “admin”
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// PutChangeKTSUserMyAlarmInput содержит входные параметры для метода PutChangeKTSUserMyAlarm
	PutChangeKTSUserMyAlarmInput struct {
		CustId   string //Идентификатор пользователя
		IsPanic  bool   //true разрешить использование КТС, false - запретить
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// GetPartsInput содержит входные параметры для метода GetParts
	GetPartsInput struct {
		SiteId   string //Идентификатор пользователя
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}

	// GetZonesInput содержит входные параметры для метода GetZones
	GetZonesInput struct {
		SiteId   string //Идентификатор пользователя
		UserName string //Имя пользователя, от которого делается запрос (необязательное поле)
		Config   `json:"-"`
	}
)

// Структуры выходных данных
type (
	// GetSitesResponse содержит ответ от сервера метода GetSites
	GetSitesResponse struct {
		RowNumber                  int     `json:"RowNumber"`
		Id                         string  `json:"Id"`
		AccountNumber              int     `json:"AccountNumber"`
		CloudObjectID              int     `json:"CloudObjectID"`
		Name                       string  `json:"Name"`
		ObjectPassword             string  `json:"ObjectPassword"`
		Address                    string  `json:"Address"`
		Phone1                     string  `json:"Phone1"`
		Phone2                     string  `json:"Phone2"`
		TypeName                   string  `json:"TypeName"`
		IsFire                     bool    `json:"IsFire"`
		IsArm                      bool    `json:"IsArm"`
		IsPanic                    bool    `json:"IsPanic"`
		DeviceTypeName             string  `json:"DeviceTypeName"`
		EventTemplateName          string  `json:"EventTemplateName"`
		ContractNumber             string  `json:"ContractNumber"`
		ContractPrice              float64 `json:"ContractPrice"`
		MoneyBalance               float64 `json:"MoneyBalance"`
		PaymentDate                string  `json:"PaymentDate"`
		DebtInformLevel            int     `json:"DebtInformLevel"`
		Disabled                   bool    `json:"Disabled"`
		DisableReason              int     `json:"DisableReason"`
		DisableDate                string  `json:"DisableDate"`
		AutoEnable                 bool    `json:"AutoEnable"`
		AutoEnableDate             string  `json:"AutoEnableDate"`
		CustomersComment           string  `json:"CustomersComment"`
		CommentForOperator         string  `json:"CommentForOperator"`
		CommentForGuard            string  `json:"CommentForGuard"`
		MapFileName                string  `json:"MapFileName"`
		WebLink                    string  `json:"WebLink"`
		ControlTime                int     `json:"ControlTime"`
		CTIgnoreSystemEvent        bool    `json:"CTIgnoreSystemEvent"`
		IsContractPriceForceUpdate bool    `json:"IsContractPriceForceUpdate"`
		IsMoneyBalanceForceUpdate  bool    `json:"IsMoneyBalanceForceUpdate"`
		IsPaymentDateForceUpdate   bool    `json:"IsPaymentDateForceUpdate"`
		IsStateArm                 bool    `json:"IsStateArm"`
		IsStateAlarm               bool    `json:"IsStateAlarm"`
		IsStatePartArm             bool    `json:"IsStatePartArm"`
		StateArmDisArmDateTime     string  `json:"StateArmDisArmDateTime"`
	}

	// GetCustomerResponse содержит ответ от сервера методов GetCustomers и GetCustomer
	GetCustomerResponse struct {
		Id                 string `json:"Id"`
		OrderNumber        int    `json:"OrderNumber"`
		UserNumber         int    `json:"UserNumber"`
		ObjCustName        string `json:"ObjCustName"`
		ObjCustTitle       string `json:"ObjCustTitle"`
		ObjCustPhone1      string `json:"ObjCustPhone1"`
		ObjCustPhone2      string `json:"ObjCustPhone2"`
		ObjCustPhone3      string `json:"ObjCustPhone3"`
		ObjCustPhone4      string `json:"ObjCustPhone4"`
		ObjCustPhone5      string `json:"ObjCustPhone5"`
		ObjCustAddress     string `json:"ObjCustAddress"`
		IsVisibleInCabinet bool   `json:"IsVisibleInCabinet"`
		ReclosingRequest   bool   `json:"ReclosingRequest"`
		ReclosingFailure   bool   `json:"ReclosingFailure"`
		PINCode            string `json:"PINCode"`
	}

	// PostCheckPanicResponse содержит ответ от сервера метода PostCheckPanic
	PostCheckPanicResponse struct {
		Status       int    `json:"Status"`
		Description  string `json:"Description"`
		CheckPanicId string `json:"CheckPanicId"`
	}

	// GetCheckPanicResponse содержит ответ от сервера метода GetCheckPanic
	GetCheckPanicResponse struct {
		Status      int    `json:"Status"`
		Description string `json:"Description"`
	}

	// UserMyAlarmResponse содержит ответ от сервера метода GetUsersMyAlarm
	UserMyAlarmResponse struct {
		CustomerID   string `json:"CustomerID"`
		MobilePhone  string `json:"MobilePhone"`
		MyAlarmPhone string `json:"MyAlarmPhone"`
		Role         string `json:"Role"`
		IsPanic      bool   `json:"IsPanic"`
	}

	// GetUserObjectMyAlarmResponse содержит ответ от сервера метода GetUserObjectMyAlarm
	GetUserObjectMyAlarmResponse struct {
		ObjectGUID string `json:"ObjectGUID"`
		CustomerID string `json:"CustomerID"`
		Role       string `json:"Role"`
		IsPanic    bool   `json:"IsPanic"`
	}

	// GetPartsResponse содержит ответ от сервера метода GetParts
	GetPartsResponse struct {
		Id                     string `json:"Id"`
		PartNumber             int    `json:"PartNumber"`
		ObjectNumber           int    `json:"ObjectNumber"`
		PartDesc               string `json:"PartDesc"`
		PartEquip              string `json:"PartEquip"`
		IsStateArm             bool   `json:"IsStateArm"`
		IsStateAlarm           bool   `json:"IsStateAlarm"`
		StateArmDisArmDateTime string `json:"StateArmDisArmDateTime"`
	}

	// GetZonesResponse содержит ответ от сервера метода GetZones
	GetZonesResponse struct {
		Id         string `json:"Id"`
		ZoneNumber int    `json:"ZoneNumber"`
		ZoneDesc   string `json:"ZoneDesc"`
		ZoneEquip  string `json:"ZoneEquip"`
	}

	// PutChangeUserMyAlarmResponse содержит ответ от сервера метода PutChangeUserMyAlarm
	PutChangeUserMyAlarmResponse struct {
		Message string `json:"Message"`
	}
)
