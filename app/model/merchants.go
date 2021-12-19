package model

type Merchants struct {
	ID           int64  `json:"id" gorm:"primaryKey"`
	MerchantName string `json:"merchant_name"`
	UserID       int64  `json:"user_id"`
	BaseModel
}

func (Merchants) TableName() string {
	return "Merchants"
}
