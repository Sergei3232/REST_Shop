package models

import (
	u "github.com/Sergei3232/REST_Shop/utils"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (product *Product) Validate() (map[string]interface{}, bool) {
	//товар должен быть уникален
	temp := &Product{}

	err := GetDB().Table("products").Where("name = ?", product.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(500, "Connection error. Please retry"), false
	}

	if temp.Name != "" {
		return u.Message(400, "A product with this name already exists"), false
	}

	return u.Message(200, "The product has been created"), true
}

func (product *Product) Create() map[string]interface{} {
	if resp, ok := product.Validate(); !ok {
		return resp
	}

	GetDB().Create(product)

	if product.ID <= 0 {
		return u.Message(500, "Failed to create a product, connection error.")
	}

	response := u.Message(201, "Product has been created")
	response["product"] = product
	return response
}

func GetProduct(product string) *Product {
	acc := &Product{}
	GetDB().Table("products").Where("name = ?", product).First(acc)
	if acc.Name == "" { //Товар не найден!
		return nil
	}

	return acc
}
