package repository

import (
	"context"

	"github.com/ddritzenhoff/dinny/pkg/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	// cook
	CreateCook(request api.NewCookRequest) error
	GetCook(eventMessageID string) (api.NewCookRequest, error)

	// participant
	CreateParticipant(request api.NewParticipantRequest) error
	GetParticipant(slackUserID string) (api.NewParticipantRequest, error)
	UpdateMealsEaten(slackUserID string, newMealsEaten int) error
}

type storage struct {
	db *mongo.Client
}

func NewStorage(db *mongo.Client) Storage {
	return &storage{db: db}
}

func (s *storage) CreateParticipant(request api.NewParticipantRequest) error {
	return nil
}

func (s *storage) CreateCook(request api.NewCookRequest) error {
	return nil
}

func (s *storage) GetCook(eventMessageID string) (api.NewCookRequest, error) {
	cookCollection := s.db.Database("dinnyDB").Collection("cook")
	var cook api.NewCookRequest
	filter := bson.D{{Key: "MessageID", Value: eventMessageID}}
	err := cookCollection.FindOne(context.TODO(), filter).Decode(&cook)
	if err != nil {
		// this means that there isn't a document in the eating collection
		return cook, err
	}

	return cook, err
}

func (s *storage) GetParticipant(slackUserID string) (api.NewParticipantRequest, error) {
	participantCollection := s.db.Database("dinnyDB").Collection("participant")
	var participant api.NewParticipantRequest
	filter := bson.D{{Key: "SlackUserID", Value: slackUserID}}
	err := participantCollection.FindOne(context.TODO(), filter).Decode(&participant)
	if err != nil {
		// this means that there isn't a document in the eating collection
		return participant, err
	}

	return participant, err
}

func (s *storage) UpdateMealsEaten(slackUserID string, newMealsEaten int) error {
	// TODO (ddritzenhoff) finish this.
	participantCollection := s.db.Database("dinnyDB").Collection("participant")
	filter := bson.D{primitive.E{Key: "SlackUserID", Value: slackUserID}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "meals_eaten", Value: newMealsEaten},
		}},
	}
	_, err := participantCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		// TODO (ddritzenhoff) probably makes sense to log here and to add a bit of context
		// to the rror.
		return err
	}

	return nil
}
