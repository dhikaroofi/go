package model

import (
	"github.com/dhikaroofi/go/app/helper"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedBy int64     `json:"updated_by"`
}

func (model *BaseModel) InitCreated(userID int64) {
	model.CreatedAt = helper.GenerateTime()
	model.CreatedBy = userID
}

func (model *BaseModel) InitUpdated(userID int64) {
	model.UpdatedAt = helper.GenerateTime()
	model.UpdatedBy = userID
}
