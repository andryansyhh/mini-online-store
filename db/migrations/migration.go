package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "00001",
			Migrate: CreateUserTable,
		},
		{
			ID:      "00002",
			Migrate: CreateUserTokenTable,
		},
		{
			ID:      "00003",
			Migrate: CreateProduct,
		},
		{
			ID:      "00004",
			Migrate: CreateProductCategory,
		},
		{
			ID:      "00005",
			Migrate: CreateShoppingCart,
		},
		{
			ID:      "00006",
			Migrate: CreateTrx,
		},
	})

	if err := m.Migrate(); err != nil {
		panic(err)
	}
}
