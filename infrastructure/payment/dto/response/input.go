package paymentServiceResponseDTO

type (
	CheckInvoiceResult struct {
		IsPaid bool   `json:"is_paid"`
		Msg    string `json:"msg"`
	}

	InvoiceResult struct {
		BankInvoiceID string     `json:"bank_invoice_id"`
		BankQRCode    string     `json:"bank_qr_code"`
		Deeplinks     []Deeplink `json:"deeplinks"`
		IsPaid        bool       `json:"is_paid"`
		Raw           any        `json:"raw"`
	}

	Deeplink struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Logo        string `json:"logo"`
	}
)
