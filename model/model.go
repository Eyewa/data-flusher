package model

import (
	"errors"
	"fmt"

	"github.com/eyewa/eyewa-go-lib/base"
	"github.com/eyewa/eyewa-go-lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var conn *gorm.DB

func InitDBConnection(connection *gorm.DB) {
	conn = connection
	migrations()
}

func migrations() {
	// https://gorm.io/docs/migration.html
	log.Debug("AutoMigrating Products table.")
	err := conn.AutoMigrate(&base.ProductModel{})
	if err != nil {
		log.Error("Products table migration failure.", zap.Error(err))
	}

	log.Debug("AutoMigrating JobSpecialPriceExpiry table.")
	err = conn.AutoMigrate(&JobSpecialPriceExpiryModel{})
	if err != nil {
		log.Error("JobSpecialPriceExpiryModel table migration failure.", zap.Error(err))
	}

	log.Debug("AutoMigrating Variants table.")
	err = conn.AutoMigrate(&base.VariantModel{})
	if err != nil {
		log.Error("Variants table migration failure.", zap.Error(err))
	}
}

func TruncateTables() error {
	tx := conn.Where("1 = 1").Delete(&JobSpecialPriceExpiryModel{})
	if tx.Error != nil {
		log.Error(`Got error while truncating "jobs" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}
	log.Info("Jobs tables successfully truncated")

	var count int64
	tx = conn.Model(&JobSpecialPriceExpiryModel{}).Where("1=1").Count(&count)
	if tx.Error != nil {
		log.Error(`Got error while getting counts of "jobs" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}

	countMsg := fmt.Sprintf("Jobs tables count is %d", count)
	if count != 0 {
		log.Error(countMsg)
		return errors.New(countMsg)
	} else {
		log.Info(countMsg)
	}

	tx = conn.Where("1 = 1").Delete(&base.ProductModel{})
	if tx.Error != nil {
		log.Error(`Got error while truncating "products" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}
	log.Info("Products tables successfully truncated")

	tx = conn.Model(&base.ProductModel{}).Where("1=1").Count(&count)
	if tx.Error != nil {
		log.Error(`Got error while getting counts of "products" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}

	countMsg = fmt.Sprintf("Products tables count is %d", count)
	if count != 0 {
		log.Error(countMsg)
		return errors.New(countMsg)
	} else {
		log.Info(countMsg)
	}

	tx = conn.Where("1 = 1").Delete(&base.VariantModel{})
	if tx.Error != nil {
		log.Error(`Got error while truncating "variants" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}
	log.Info("Variants tables successfully truncated")

	tx = conn.Model(&base.VariantModel{}).Where("1=1").Count(&count)
	if tx.Error != nil {
		log.Error(`Got error while getting counts of "variants" table`)
		log.Error(tx.Error.Error())
		return tx.Error
	}

	countMsg = fmt.Sprintf("Variants tables count is %d", count)
	if count != 0 {
		log.Error(countMsg)
		return errors.New(countMsg)
	} else {
		log.Info(countMsg)
	}

	return nil
}
