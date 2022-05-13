package data

import (
	"context"
	"database/sql"
	"time"
)

type Snooker struct {
	ID     int64  `json:"id"`
	Winner string `json:"winner"`
	Loser  string `json:"loser"`
	Diff   int    `json:"diff"`
	Date   string `json:"date"`
}

type SnookerModel struct {
	DB *sql.DB
}

func (m *SnookerModel) Insert(snooker *Snooker) error {
	query := `
INSERT INTO snookers (winner, loser, diff)
VALUES ($1, $2, $3)
RETURNING id, date`

	args := []interface{}{snooker.Winner, snooker.Loser, snooker.Diff}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&snooker.ID, &snooker.Date)
	if err != nil {
		return err
	}

	return nil
}

func (m *SnookerModel) GetAll(page int) ([]*Snooker, int, error) {
	query := `
SELECT COUNT(*) OVER(), id, winner, loser, diff, date
FROM snookers
ORDER BY id DESC
LIMIT $1 OFFSET $2
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, PageSize, (page-1)*PageSize)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	totalRecords := 0
	snookers := []*Snooker{}

	for rows.Next() {
		var snooker Snooker

		err := rows.Scan(
			&totalRecords,
			&snooker.ID,
			&snooker.Winner,
			&snooker.Loser,
			&snooker.Diff,
			&snooker.Date,
		)
		if err != nil {
			return nil, 0, err
		}

		snookers = append(snookers, &snooker)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	maxPage := totalRecords/PageSize + 1

	return snookers, maxPage, nil
}
