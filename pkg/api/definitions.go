package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DinnerRotationCookingDays int

const (
	Monday DinnerRotationCookingDays = iota
	Tuesday
	Wednesday
	Thursday
)

type NewParticipantRequest struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	SlackUserID string             `bson:"slack_user_id"`
	RealName    string             `bson:"name"`
	MealsEaten  int                `bson:"meals_eaten"`
	MealsCooked int                `bson:"meals_cooked"`
	Points      int                `bson:"points"`
}

type NewCookRequest struct {
	ID                  primitive.ObjectID        `bson:"_id"`
	CreatedAt           time.Time                 `bson:"created_at"`
	UpdatedAt           time.Time                 `bson:"updated_at"`
	ParticipantObjectID primitive.ObjectID        `bson:"participant_object_id"`
	SlackUserID         string                    `bson:"slack_user_id"`
	MessageID           string                    `bson:"message_id"`
	CookingDay          DinnerRotationCookingDays `bson:"cooking_day"`
	MealDescription     string                    `bson:"meal_description"`
}
