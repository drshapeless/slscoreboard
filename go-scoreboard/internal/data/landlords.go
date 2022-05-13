package data

import (
	"context"
	"database/sql"
	"time"
)

type Landlord struct {
	ID       int64  `json:"id"`
	Landlord string `json:"landlord"`
	Farmer1  string `json:"farmer1"`
	Farmer2  string `json:"farmer2"`
	Win      int    `json:"win"`
	Date     string `json:"date"`
}

type LandlordModel struct {
	DB *sql.DB
}

func (m *LandlordModel) Insert(landlord *Landlord) error {
	query := `
INSERT INTO landlords (landlord, farmer1, farmer2, win)
VALUES ($1, $2, $3, $4)
RETURNING id, date`

	args := []interface{}{landlord.Landlord, landlord.Farmer1, landlord.Farmer2, landlord.Win}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&landlord.ID, &landlord.Date)
	if err != nil {
		return err
	}

	return nil
}

func (m *LandlordModel) GetAll(page int) ([]*Landlord, int, error) {
	query := `
SELECT COUNT(*) OVER(), id, landlord, farmer1, farmer2, win, date
FROM landlords
ORDER BY id DESC
LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, PageSize, (page-1)*PageSize)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	totalRecords := 0
	landlords := []*Landlord{}

	for rows.Next() {
		var landlord Landlord

		err := rows.Scan(
			&totalRecords,
			&landlord.ID,
			&landlord.Landlord,
			&landlord.Farmer1,
			&landlord.Farmer2,
			&landlord.Win,
		)
		if err != nil {
			return nil, 0, err
		}

		landlords = append(landlords, &landlord)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	maxPage := totalRecords/PageSize + 1

	return landlords, maxPage, nil
}
