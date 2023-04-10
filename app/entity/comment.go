package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// MyGramComment represents Comment
type MyGramComment struct {
	MyGramID
	MyGramUserID  uint   `json:"user_id" example:"2"`
	MyGramPhotoID uint   `json:"photo_id" example:"3"`
	Message       string `gorm:"not null" json:"message" form:"message" valid:"required~Comment is required"`
	MyGramDate
}

func (c *MyGramComment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return nil
}

func (c *MyGramComment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	return nil
}
