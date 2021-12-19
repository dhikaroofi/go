package model

type Transactions struct {
	ID         int64   `json:"id" gorm:"primaryKey"`
	MerchantID int64   `json:"merchant_id"`
	OutletID   int64   `json:"outlet_id"`
	BillTotal  float64 `json:"bill_total"`
	BaseModel
}

func (Transactions) TableName() string {
	return "Transactions"
}
