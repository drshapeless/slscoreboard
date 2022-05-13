package data

import (
	"context"
	"database/sql"
	"time"
)

type Dee struct {
	ID         int64  `json:"id"`
	Winner     string `json:"winner"`
	Loser1     string `json:"loser1"`
	Loser1Card int    `json:"loser1_card"`
	Loser2     string `json:"loser2"`
	Loser2Card int    `json:"loser2_card"`
	Loser3     string `json:"loser3"`
	Loser3Card int    `json:"loser3_card"`
	Date       string `json:"date"`
}

type DeeModel struct {
	DB *sql.DB
}

func (m *DeeModel) Insert(dee *Dee) error {
	query := `
INSERT INTO dees (winner, loser1, loser1_card, loser2, loser2_card, loser3, loser3_card)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, date`

	args := []interface{}{
		dee.Winner,
		dee.Loser1,
		dee.Loser1Card,
		dee.Loser2,
		dee.Loser2Card,
		dee.Loser3,
		dee.Loser3Card,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&dee.ID, &dee.Date)
	if err != nil {
		return err
	}

	return nil
}

func (m *DeeModel) GetAll(page int) ([]*Dee, int, error) {
	query := `
SELECT COUNT(*) OVER(), id, winner, loser1, loser1_card, loser2, loser2_card, loser3, loser3_card, date
FROM dees
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
	dees := []*Dee{}

	for rows.Next() {
		var dee Dee

		err := rows.Scan(
			&totalRecords,
			&dee.ID,
			&dee.Winner,
			&dee.Loser1,
			&dee.Loser1Card,
			&dee.Loser2,
			&dee.Loser2Card,
			&dee.Loser3,
			&dee.Loser3Card,
		)
		if err != nil {
			return nil, 0, err
		}

		dees = append(dees, &dee)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	maxPage := totalRecords/PageSize + 1

	return dees, maxPage, nil
}
