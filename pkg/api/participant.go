package api

import (
	"time"

	"github.com/slack-go/slack"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantService interface {
	New(slackUserID string) error
	Get(slackUserID string) (NewParticipantRequest, error)
	LikedEating(slackUserID string) error
	UnlikedEating(slackUserID string) error
}

type ParticipantRepository interface {
	CreateParticipant(NewParticipantRequest) error
	GetParticipant(slackUserID string) (NewParticipantRequest, error)
	UpdateMealsEaten(slackUserID string, newMealsEaten int) error
}

type participantService struct {
	storage  ParticipantRepository
	slackApi *slack.Client
}

func NewParticipantService(participantRepo ParticipantRepository, slackApi *slack.Client) ParticipantService {
	return &participantService{
		storage:  participantRepo,
		slackApi: slackApi,
	}
}

func (p *participantService) New(slackUserID string) error {
	slackUser, err := p.slackApi.GetUserInfo(slackUserID)
	if err != nil {
		return err
	}
	return p.storage.CreateParticipant(NewParticipantRequest{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		SlackUserID: slackUserID,
		RealName:    slackUser.RealName,
		MealsEaten:  0,
		MealsCooked: 0,
		Points:      0,
	})
}

// TODO (ddritzenhoff) take another look at this b/c this binds me to the mongoDB
func (p *participantService) Get(slackUserID string) (NewParticipantRequest, error) {
	return p.storage.GetParticipant(slackUserID)
}

// TODO (ddritzenhoff) is there a better way of doing this?
func (p *participantService) LikedEating(slackUserID string) error {
	participant, err := p.storage.GetParticipant(slackUserID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = p.New(participant.SlackUserID)
			if err != nil {
				return err
			}
			return p.LikedEating(slackUserID)
		} else {
			return err
		}
	}
	return p.storage.UpdateMealsEaten(slackUserID, participant.MealsEaten + 1)
}

func (p *participantService) UnlikedEating(slackUserID string) error {
	participant, err := p.storage.GetParticipant(slackUserID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = p.New(participant.SlackUserID)
			if err != nil {
				return err
			}
			return p.LikedEating(slackUserID)
		} else {
			return err
		}
	}
	return p.storage.UpdateMealsEaten(slackUserID, participant.MealsEaten - 1)
}
