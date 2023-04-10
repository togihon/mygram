package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// MyGramPhoto represents the photo
type MyGramPhoto struct {
	MyGramID
	Title        string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption      string `gorm:"not null" json:"caption" form:"caption"`
	Photo_URL    string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	MyGramUserID uint
	MyGramDate
}

func (ph *MyGramPhoto) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ph)

	if errCreate != nil {
		err = errCreate
		return
	}
	return nil
}

func (ph *MyGramPhoto) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(ph)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	return nil
}
