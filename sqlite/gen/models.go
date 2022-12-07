// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gen

import (
	"database/sql"
)

type Meal struct {
	ID             int64
	CookSlackUid   string
	Year           int64
	Month          int64
	Day            int64
	Description    sql.NullString
	SlackMessageID sql.NullString
	CreatedAt      string
	UpdatedAt      string
}

type Member struct {
	ID          int64
	SlackUid    string
	FullName    string
	MealsEaten  int64
	MealsCooked int64
	Leader      int64
	CreatedAt   string
	UpdatedAt   string
}
