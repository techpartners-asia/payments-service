package entity

import databaseBase "github.com/techpartners-asia/payments-service/infrastructure/database/base"

type PaymentType string

const (
	PaymentTypeTokiay     PaymentType = "tokipay"
	PaymentTypeQpay       PaymentType = "qpay"
	PaymentTypeBalc       PaymentType = "balc"
	PaymentTypeBpay       PaymentType = "bpay"
	PaymentTypeEcommerce  PaymentType = "ecommerce"
	PaymentTypeHipay      PaymentType = "hipay"
	PaymentTypeMongolchat PaymentType = "mongolchat"
	PaymentTypeMonpay     PaymentType = "monpay"
	PaymentTypePass       PaymentType = "pass"
	PaymentTypePocket     PaymentType = "pocket"
	PaymentTypeSimple     PaymentType = "simple"
	PaymentTypeSocialpay  PaymentType = "socialpay"
	PaymentTypeStorepay   PaymentType = "storepay"
	PaymentTypeTino       PaymentType = "tino"
	PaymentTypeTokipay    PaymentType = "tokipay"
	PaymentTypeUpoint     PaymentType = "upoint"
)

type (
	MerchantEntity struct {
		databaseBase.BaseEntity
		Name string `gorm:"name" json:"name"`
	}

	MerchantEbarimtEntity struct {
		databaseBase.BaseEntity
		MerchantID   uint            `gorm:"merchant_id" json:"merchant_id"`
		Merchant     *MerchantEntity `gorm:"foreignKey:MerchantID" json:"merchant"`
		Url          string          `gorm:"url" json:"url"`
		Tin          string          `gorm:"tin" json:"tin"`
		PosNo        string          `gorm:"pos_no" json:"pos_no"`
		BranchNo     string          `gorm:"branch_no" json:"branch_no"`
		DistrictCode string          `gorm:"district_code" json:"district_code"`
	}
)
