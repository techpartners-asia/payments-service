package entity

import databaseBase "git.techpartners.asia/gateway-services/payment-service/infrastructure/database/base"

type (
	EbarimtEntity struct {
		databaseBase.BaseEntity
		TotalAmount  float64                `gorm:"total_amount" json:"total_amount"`
		TotalVat     float64                `gorm:"total_vat" json:"total_vat"`
		TotalCityTax float64                `gorm:"total_city_tax" json:"total_city_tax"`
		BranchNo     string                 `gorm:"branch_no" json:"branch_no"`
		DistrictCode string                 `gorm:"district_code" json:"district_code"`
		MerchantTin  string                 `gorm:"merchant_tin" json:"merchant_tin"`
		PosNo        string                 `gorm:"pos_no" json:"pos_no"`
		CustomerTin  string                 `gorm:"customer_tin" json:"customer_tin"`
		ConsumerNo   string                 `gorm:"consumer_no" json:"consumer_no"`
		Type         string                 `gorm:"type:varchar(20);column:receipt_type" json:"type"`
		BillID       string                 `gorm:"bill_id" json:"bill_id"`
		InvoiceID    string                 `gorm:"invoice_id" json:"invoice_id"`
		PosID        float64                `gorm:"pos_id" json:"pos_id"`
		Message      string                 `gorm:"message" json:"message"`
		QrData       string                 `gorm:"qr_data" json:"qr_data"`
		Lottery      string                 `gorm:"lottery" json:"lottery"`
		Date         string                 `gorm:"date" json:"date"`
		IsRefund     bool                   `gorm:"is_refund" json:"is_refund"`
		Receipts     []EbarimtReceiptEntity `gorm:"foreignKey:EbarimtID" json:"receipts"`
	}

	EbarimtReceiptEntity struct {
		databaseBase.BaseEntity
		EbarimtID     int64                      `gorm:"ebarimt_id" json:"ebarimt_id"`
		BillID        string                     `gorm:"bill_id" json:"bill_id"`
		TotalAmount   float64                    `gorm:"total_amount" json:"total_amount"`
		TotalVat      float64                    `gorm:"total_vat" json:"total_vat"`
		TotalCityTax  float64                    `gorm:"total_city_tax" json:"total_city_tax"`
		TaxType       string                     `gorm:"tax_type" json:"tax_type"`
		MerchantTin   string                     `gorm:"merchant_tin" json:"merchant_tin"`
		BankAccountNo string                     `gorm:"bank_account_no" json:"bank_account_no"`
		Items         []EbarimtReceiptItemEntity `gorm:"foreignKey:ReceiptID" json:"items"`
		IsRefund      bool                       `gorm:"is_refund" json:"is_refund"`
	}

	EbarimtReceiptItemEntity struct {
		databaseBase.BaseEntity
		Name               string                `gorm:"name" json:"name"`
		BarCode            string                `gorm:"bar_code" json:"bar_code"`
		BarCodeType        string                `gorm:"bar_code_type" json:"bar_code_type"`
		ClassificationCode string                `gorm:"classification_code" json:"classification_code"`
		MeasureUnit        string                `gorm:"measure_unit" json:"measure_unit"`
		TaxProductCode     string                `gorm:"tax_product_code" json:"tax_product_code"`
		Qty                float64               `gorm:"qty" json:"qty"`
		UnitPrice          float64               `gorm:"unit_price" json:"unit_price"`
		TotalAmount        float64               `gorm:"total_amount" json:"total_amount"`
		TotalVat           float64               `gorm:"total_vat" json:"total_vat"`
		TotalCityTax       float64               `gorm:"total_city_tax" json:"total_city_tax"`
		TotalBonus         float64               `gorm:"total_bonus" json:"total_bonus"`
		ReceiptID          int64                 `gorm:"receipt_id" json:"receipt_id"`
		Receipt            *EbarimtReceiptEntity `gorm:"foreignKey:ReceiptID" json:"receipt"`
	}
)
