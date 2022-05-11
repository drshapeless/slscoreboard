package data

import "database/sql"

type Models struct {
	Snookers  SnookerModel
	Dees      DeeModel
	Landlords LandlordModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Snookers:  SnookerModel{DB: db},
		Dees:      DeeModel{DB: db},
		Landlords: LandlordModel{DB: db},
	}
}
