package sharedDTO

type (
	SharedPaymentConfigDTO struct {
		Qpay      QpayAdapterDTO      `json:"qpay"`
		Tokiay    TokipayAdapterDTO   `json:"tokiay"`
		Monpay    MonpayAdapterDTO    `json:"monpay"`
		Golomt    GolomtAdapterDTO    `json:"golomt"`
		Socialpay SocialPayAdapterDTO `json:"socialpay"`
		Storepay  StorePayAdapterDTO  `json:"storepay"`
		Pocket    PocketAdapterDTO    `json:"pocket"`
		Simple    SimpleAdapterDTO    `json:"simple"`
		Balc      BalcAdapterDTO      `json:"balc"`
	}

	QpayAdapterDTO struct {
		Username    string
		Password    string
		Endpoint    string
		Callback    string
		InvoiceCode string
		MerchantID  string
	}

	TokipayAdapterDTO struct {
		Endpoint      string
		APIKey        string
		IMAPIKey      string
		Authorization string
		MerchantID    string
		SuccessURL    string
		FailureURL    string
		AppSchemaIOS  string
	}

	StorePayAdapterDTO struct {
		AppUserName string
		AppPassword string
		Username    string
		Password    string
		AuthUrl     string
		BaseUrl     string
		StoreId     string
		CallbackUrl string
	}

	SocialPayAdapterDTO struct {
		Terminal string
		Secret   string
		Endpoint string
	}
	SimpleAdapterDTO struct {
		UserName    string
		Password    string
		BaseUrl     string
		CallbackUrl string
	}
	PocketAdapterDTO struct {
		Merchant      string
		ClientID      string
		ClientSecret  string
		Environment   string
		TerminalIDRaw int64
	}
	MonpayAdapterDTO struct {
		Endpoint  string
		Username  string
		AccountID string
		Callback  string
	}
	GolomtAdapterDTO struct {
		BaseURL     string
		Secret      string
		BearerToken string
		ReturnType  string
		CallbackURL string
	}
	BalcAdapterDTO struct {
		Endpoint string
		Token    string
	}
)
