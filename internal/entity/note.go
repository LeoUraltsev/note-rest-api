package entity

import "time"

type Note struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	ModifiedAt  time.Time `db:"modified_at"`
}
