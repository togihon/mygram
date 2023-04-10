package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type MyGramSocialMedia struct {
	MyGramID
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Social media name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	MyGramUserID   uint
	MyGramDate
}

func (sm *MyGramSocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}
	return nil
}

func (sm *MyGramSocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(sm)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	return nil
}
