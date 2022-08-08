package db

import (
	"gorm.io/gorm"
	"hotel"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	db.AutoMigrate(&hotel.Customer{})
	return &CustomerRepo{
		db: db,
	}
}

func (c CustomerRepo) CreateCustomer(model CustomerModel) error {
	err := c.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepo) DeleteCustomer(iden int64) error {
	err := c.db.Where("identity_num=?", iden).Delete(&CustomerModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepo) UpdateCustomer(model CustomerModel) error {
	err := c.db.Model(&model).Omit("identity_num").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepo) GetListCustomer() ([]CustomerModel, error) {
	var models []CustomerModel
	err := c.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (c CustomerRepo) GetCustomer(iden int64) (*CustomerModel, error) {
	var model CustomerModel
	err := c.db.Where("identity_num= ?", iden).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
