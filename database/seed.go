package database

import (
	"io/ioutil"
	"test/database/model"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
)

func SeedData(db *gorm.DB) {
	if db == nil {
		panic("db argument must not be nil")
	}

	products := make([]model.Product, 0)
	dat, err := ioutil.ReadFile("database/products.yml")
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(dat, &products)
	if err != nil {
		panic(err.Error())
	}

	db.Transaction(func(tx *gorm.DB) error {
		err := db.Exec("DELETE FROM product_cards").Error
		if err != nil {
			return err
		}

		err = db.Exec("DELETE FROM products").Error
		if err != nil {
			return err
		}

		for _, product := range products {
			err = db.Create(&product).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
