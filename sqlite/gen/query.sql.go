// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package gen

import (
	"context"
	"database/sql"
)

const countMealsByDate = `-- name: CountMealsByDate :one
SELECT count(*) FROM meals WHERE year = ? AND month = ? AND day = ?
`

type CountMealsByDateParams struct {
	Year  int64
	Month int64
	Day   int64
}

func (q *Queries) CountMealsByDate(ctx context.Context, arg CountMealsByDateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countMealsByDate, arg.Year, arg.Month, arg.Day)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createMeal = `-- name: CreateMeal :one
INSERT INTO meals (
    cook_slack_uid, year, month, day
) VALUES (
    ?, ?, ?, ?
)
RETURNING id, cook_slack_uid, year, month, day, description, slack_message_id, created_at, updated_at
`

type CreateMealParams struct {
	CookSlackUid string
	Year         int64
	Month        int64
	Day          int64
}

func (q *Queries) CreateMeal(ctx context.Context, arg CreateMealParams) (Meal, error) {
	row := q.db.QueryRowContext(ctx, createMeal,
		arg.CookSlackUid,
		arg.Year,
		arg.Month,
		arg.Day,
	)
	var i Meal
	err := row.Scan(
		&i.ID,
		&i.CookSlackUid,
		&i.Year,
		&i.Month,
		&i.Day,
		&i.Description,
		&i.SlackMessageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createMember = `-- name: CreateMember :one
INSERT INTO members (
    slack_uid, full_name, leader
) VALUES (
    ?, ?, ?
)
RETURNING id, slack_uid, full_name, meals_eaten, meals_cooked, leader, created_at, updated_at
`

type CreateMemberParams struct {
	SlackUid string
	FullName string
	Leader   int64
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, createMember, arg.SlackUid, arg.FullName, arg.Leader)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.SlackUid,
		&i.FullName,
		&i.MealsEaten,
		&i.MealsCooked,
		&i.Leader,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteMeal = `-- name: DeleteMeal :exec
DELETE FROM meals
WHERE id = ?
`

func (q *Queries) DeleteMeal(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMeal, id)
	return err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM members
WHERE id = ?
`

func (q *Queries) DeleteMember(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMember, id)
	return err
}

const findMealByDate = `-- name: FindMealByDate :one
SELECT id, cook_slack_uid, year, month, day, description, slack_message_id, created_at, updated_at FROM meals
WHERE year = ? AND month = ? AND day = ? LIMIT 1
`

type FindMealByDateParams struct {
	Year  int64
	Month int64
	Day   int64
}

func (q *Queries) FindMealByDate(ctx context.Context, arg FindMealByDateParams) (Meal, error) {
	row := q.db.QueryRowContext(ctx, findMealByDate, arg.Year, arg.Month, arg.Day)
	var i Meal
	err := row.Scan(
		&i.ID,
		&i.CookSlackUid,
		&i.Year,
		&i.Month,
		&i.Day,
		&i.Description,
		&i.SlackMessageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findMealByID = `-- name: FindMealByID :one
SELECT id, cook_slack_uid, year, month, day, description, slack_message_id, created_at, updated_at FROM meals
WHERE id = ? LIMIT 1
`

func (q *Queries) FindMealByID(ctx context.Context, id int64) (Meal, error) {
	row := q.db.QueryRowContext(ctx, findMealByID, id)
	var i Meal
	err := row.Scan(
		&i.ID,
		&i.CookSlackUid,
		&i.Year,
		&i.Month,
		&i.Day,
		&i.Description,
		&i.SlackMessageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findMealBySlackMessageID = `-- name: FindMealBySlackMessageID :one
SELECT id, cook_slack_uid, year, month, day, description, slack_message_id, created_at, updated_at FROM meals
WHERE slack_message_id = ? LIMIT 1
`

func (q *Queries) FindMealBySlackMessageID(ctx context.Context, slackMessageID sql.NullString) (Meal, error) {
	row := q.db.QueryRowContext(ctx, findMealBySlackMessageID, slackMessageID)
	var i Meal
	err := row.Scan(
		&i.ID,
		&i.CookSlackUid,
		&i.Year,
		&i.Month,
		&i.Day,
		&i.Description,
		&i.SlackMessageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findMemberByID = `-- name: FindMemberByID :one
SELECT id, slack_uid, full_name, meals_eaten, meals_cooked, leader, created_at, updated_at FROM members
WHERE id = ? LIMIT 1
`

func (q *Queries) FindMemberByID(ctx context.Context, id int64) (Member, error) {
	row := q.db.QueryRowContext(ctx, findMemberByID, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.SlackUid,
		&i.FullName,
		&i.MealsEaten,
		&i.MealsCooked,
		&i.Leader,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findMemberBySlackUID = `-- name: FindMemberBySlackUID :one
SELECT id, slack_uid, full_name, meals_eaten, meals_cooked, leader, created_at, updated_at FROM members
WHERE slack_uid = ? LIMIT 1
`

func (q *Queries) FindMemberBySlackUID(ctx context.Context, slackUid string) (Member, error) {
	row := q.db.QueryRowContext(ctx, findMemberBySlackUID, slackUid)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.SlackUid,
		&i.FullName,
		&i.MealsEaten,
		&i.MealsCooked,
		&i.Leader,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT id, slack_uid, full_name, meals_eaten, meals_cooked, leader, created_at, updated_at FROM members
ORDER BY meals_cooked ASC, meals_eaten DESC
`

func (q *Queries) ListMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Member
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.SlackUid,
			&i.FullName,
			&i.MealsEaten,
			&i.MealsCooked,
			&i.Leader,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMealDescription = `-- name: UpdateMealDescription :exec
UPDATE meals
set description = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMealDescriptionParams struct {
	Description sql.NullString
	ID          int64
}

func (q *Queries) UpdateMealDescription(ctx context.Context, arg UpdateMealDescriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateMealDescription, arg.Description, arg.ID)
	return err
}

const updateMealSlackMessageID = `-- name: UpdateMealSlackMessageID :exec
UPDATE meals
set slack_message_id = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMealSlackMessageIDParams struct {
	SlackMessageID sql.NullString
	ID             int64
}

func (q *Queries) UpdateMealSlackMessageID(ctx context.Context, arg UpdateMealSlackMessageIDParams) error {
	_, err := q.db.ExecContext(ctx, updateMealSlackMessageID, arg.SlackMessageID, arg.ID)
	return err
}

const updateMealSlackUID = `-- name: UpdateMealSlackUID :exec
UPDATE meals
set cook_slack_uid = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMealSlackUIDParams struct {
	CookSlackUid string
	ID           int64
}

func (q *Queries) UpdateMealSlackUID(ctx context.Context, arg UpdateMealSlackUIDParams) error {
	_, err := q.db.ExecContext(ctx, updateMealSlackUID, arg.CookSlackUid, arg.ID)
	return err
}

const updateMemberLeaderStatus = `-- name: UpdateMemberLeaderStatus :exec
UPDATE members
set leader = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMemberLeaderStatusParams struct {
	Leader int64
	ID     int64
}

func (q *Queries) UpdateMemberLeaderStatus(ctx context.Context, arg UpdateMemberLeaderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateMemberLeaderStatus, arg.Leader, arg.ID)
	return err
}

const updateMemberMealsCooked = `-- name: UpdateMemberMealsCooked :exec
UPDATE members
set meals_cooked = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMemberMealsCookedParams struct {
	MealsCooked int64
	ID          int64
}

func (q *Queries) UpdateMemberMealsCooked(ctx context.Context, arg UpdateMemberMealsCookedParams) error {
	_, err := q.db.ExecContext(ctx, updateMemberMealsCooked, arg.MealsCooked, arg.ID)
	return err
}

const updateMemberMealsEaten = `-- name: UpdateMemberMealsEaten :exec
UPDATE members
set meals_eaten = ?, updated_at = datetime('now')
WHERE id = ?
`

type UpdateMemberMealsEatenParams struct {
	MealsEaten int64
	ID         int64
}

func (q *Queries) UpdateMemberMealsEaten(ctx context.Context, arg UpdateMemberMealsEatenParams) error {
	_, err := q.db.ExecContext(ctx, updateMemberMealsEaten, arg.MealsEaten, arg.ID)
	return err
}