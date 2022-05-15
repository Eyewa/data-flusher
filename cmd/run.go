package cmd

import (
	"github.com/eyewa/eyewa-go-lib/db"
	"github.com/eyewa/eyewa-go-lib/log"
	"github.com/eyewa/migrator/algolia"
	"github.com/eyewa/migrator/config"
	"github.com/eyewa/migrator/model"
)

func Migrate() {
	err := config.Init()
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("Starting Migration.")

	// connect to DB
	pClient := db.NewPostgresClientFromConfig(config.Config.DB)
	_, err = pClient.OpenConnection()
	if err != nil {
		log.Error(err.Error())
		return
	}
	model.InitDBConnection(pClient.Gorm)

	// Truncate products, variants, jobs tables
	err = model.TruncateTables()
	if err != nil {
		return
	}
	log.Info("All tables successfully truncated")

	// Clear root indices
	err = algolia.ClearIndices()
	if err != nil {
		return
	}
	log.Info("All indices successfully truncated")
}
