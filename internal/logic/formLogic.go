package logic

import (
	"gin-demo/common"
	"gin-demo/model"

	"github.com/jinzhu/gorm"
)

type IFormLogic interface {
	Insert(name string) (*model.Form, error)
	Update(form model.Form, name string) (*model.Form, error)
	SelectById(id int64) (*model.Form, error)
	DeleteById(id int64) error
}

type FormLogic struct {
	DB *gorm.DB
}

func (c FormLogic) DeleteById(id int64) error {
	if err := c.DB.Delete(model.Form{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c FormLogic) SelectById(id int64) (*model.Form, error) {
	var form model.Form
	if err := c.DB.First(&form, id).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func (c FormLogic) Update(form model.Form, name string) (*model.Form, error) {
	if err := c.DB.Model(&form).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func (c FormLogic) Insert(name string) (*model.Form, error) {
	form := model.Form{
		Name: name,
	}

	if err := c.DB.Create(&form).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func NewFormLogic() IFormLogic {
	return FormLogic{DB: common.GetDB()}
}
