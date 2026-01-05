package adapterTino

type InvoiceStatus string

const (
	INVOICE_STATUS_CREATED InvoiceStatus = "created"
	INVOICE_STATUS_PENDING InvoiceStatus = "pending"
	INVOICE_STATUS_PAID    InvoiceStatus = "paid"
	INVOICE_STATUS_EXPIRED InvoiceStatus = "expired"
)

type (
	InitInput struct {
		Url         string
		User        string
		Password    string
		CallbackUrl string
	}

	AuthResponse struct {
		Status  bool             `json:"status"`
		Data    AuthDataResponse `json:"data"`
		Message string           `json:"message"`
	}

	AuthDataResponse struct {
		User            map[string]interface{} `json:"user"`
		Token           string                 `json:"token"`
		ExpireAt        int64                  `json:"expire_at"`
		RefreshToken    string                 `json:"refresh_token"`
		RefreshExpireAt int64                  `json:"refresh_expire_at"`
	}

	CreateInvoiceInput struct {
		CallbackURL   string  `json:"callback_url"`
		Description   string  `json:"description"`
		Amount        float64 `json:"amount"`
		UserID        string  `json:"user_id"`
		TransactionID string  `json:"transaction_id"`
		WalletID      string  `json:"wallet_id"`
	}

	CreateInvoiceResponse struct {
		Status  bool                      `json:"status"`
		Message string                    `json:"message"`
		Data    CreateInvoiceDataResponse `json:"data"`
	}

	CreateInvoiceDataResponse struct {
		InvoiceID  string        `json:"invoice_id"`
		MerchantID string        `json:"merchant_id"`
		WalletID   string        `json:"wallet_id"`
		Amount     float64       `json:"amount"`
		Status     InvoiceStatus `json:"status"`
	}

	CheckInvoiceResponse struct {
		Status  bool                     `json:"status"`
		Message string                   `json:"message"`
		Data    CheckInvoiceDataResponse `json:"data"`
	}

	CheckInvoiceDataResponse struct {
		InvoiceID string        `json:"invoice_id"`
		Status    InvoiceStatus `json:"status"`
		Amount    float64       `json:"amount"`
		PaidAt    string        `json:"paid_at"`
	}

	CancelInvoiceResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	RefundInvoiceInput struct {
		InvoiceID string `json:"invoice_id"`
		Reason    string `json:"reason"`
	}

	RefundInvoiceResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
)
