package data

import (
	"project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &cartData{
		db: db,
	}
}

func (repo *cartData) SelectByToken(token int) ([]cart.Core, error) {

	var dataCart []Cart
	tx := repo.db.Model(&Cart{}).Where("user_id = ?", token).Find(&dataCart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataProduct []Product
	txProduct := repo.db.Find(&dataProduct)
	if txProduct.Error != nil {
		return nil, txProduct.Error
	}

	return toCoreList(dataCart, dataProduct), nil

}
