package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Server) Pong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string {
			"status": "success",
			"data": "pong",
		}

		c.JSON(http.StatusOK, response)
	}
}

// TODO (ddritzenhoff) what is a gin.HandlerFunc??

func (s *Server) HandleEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sv, err := slack.NewSecretsVerifier(c.Request.Header, s.signingSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := sv.Write(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := sv.Ensure(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, r.Challenge)
	}

	// TODO (ddritzenhoff) fix this
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		switch innerEvent := eventsAPIEvent.InnerEvent.Data.(type) {
		case *slackevents.ReactionAddedEvent:
			s.handleReactionAddedEvent(innerEvent)
		case *slackevents.ReactionRemovedEvent:
			s.handleReactionRemovedEvent(innerEvent)
		}
	}
	}
}


func (s *Server) handleReactionRemovedEvent(reactionEvent *slackevents.ReactionRemovedEvent) {
	eventMessageID := reactionEvent.Item.Timestamp
	fmt.Printf("handleReactionRemovedEvent with event message ID %s\n", eventMessageID)
	cook, err := s.cookService.Get(eventMessageID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Couldn't find the cook document. Either it doesn't exist, or a separate message had a reaction")
			return
		}
		log.Fatal(err)
	}

	if cook.MessageID != eventMessageID || reactionEvent.Reaction != "+1" {
		return
	}

	// at this point, you know that the user liked the correct message
	// at this point, you've found the correct user from within the User collection.
	s.participantService.LikedEating(cook.SlackUserID)
}

func (s *Server) handleReactionAddedEvent(reactionEvent *slackevents.ReactionAddedEvent) {
	eventMessageID := reactionEvent.Item.Timestamp
	fmt.Printf("handleReactionAddedEvent with event message ID %s\n", eventMessageID)
	cook, err := s.cookService.Get(eventMessageID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Couldn't find the cook document. Either it doesn't exist, or a separate message had a reaction")
			return
		}
		log.Fatal(err)
	}

	if cook.MessageID != eventMessageID || reactionEvent.Reaction != "+1" {
		return
	}

	// at this point, you know that the user liked the correct message
	// at this point, you've found the correct user from within the User collection.
	s.participantService.LikedEating(cook.SlackUserID)
}