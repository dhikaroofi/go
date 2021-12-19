package model

type Outlets struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	MerchantID int64  `json:"merchant_id"`
	OutletName string `json:"outlet_name"`
	BaseModel
}

func (Outlets) TableName() string {
	return "Outlets"
}
