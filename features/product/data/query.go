package data

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &productData{
		db: db,
	}
}

func (repo *productData) InsertData(data product.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)

	return int(tx.RowsAffected), tx.Error
}
