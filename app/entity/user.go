package entity

import (
	"mygram/pkg/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// MyGramUser represents the model for a user
type MyGramUser struct {
	//uint 32bit dan tidak boleh minus
	MyGramID
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password must be 6 characters or more"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(9|60)~Your age should be above 8 years old"`
	MyGramDate
}

func (u *MyGramUser) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}
