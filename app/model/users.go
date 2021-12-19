package model

type Users struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	BaseModel
}

func (Users) TableName() string {
	return "Users"
}
