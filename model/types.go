package model

import (
	"time"

	"gorm.io/datatypes"
)

// JobSpecialPriceExpiryModel details on products requiring special price expiry
type JobSpecialPriceExpiryModel struct {
	ID uint `gorm:"primaryKey" json:"-"`
	ProductMeta
	ProductUpdated bool
	CreatedAt      time.Time
	ExpiresAt      time.Time
	ExpiredAt      time.Time
}

// TableName overrides the table name for ProductModel
func (JobSpecialPriceExpiryModel) TableName() string {
	return "jobs_special_price_expiry"
}

type ProductMeta struct {
	StoreID   int
	StoreCode string
	EntityID  int
	TypeID    string

	// if a simple has parent(s), we need to ensure the special_price
	// is also updated on the configurables' record(s) for given store
	ParentEntityIDs *datatypes.JSON
}
