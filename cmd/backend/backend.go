// TODO (ddritzenhoff): Figure out a way to make the user-id another primary key in addition to the default key within mongo
// TODO (ddritzenhoff): Figure out a way to add a ping function to see if the backend is running.
// TODO (ddritzenhoff): I'm probably doing a lot of copying of structs. Figure out a way to decrease that.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants from env file
var signingSecret string
var botSigningKey string
var mongoDBConnectionUri string
var defaultServerTimeout time.Duration

// DB connection
var userCollection *mongo.Collection
var eatingCollection *mongo.Collection
var ctx = context.TODO()

// api
var api *slack.Client

// define structs used
type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	SlackUserId string             `bson:"slack_user_id"`
	Name        string             `bson:"name"`
	MealsEaten  int                `bson:"meals_eaten"`
	MealsCooked int                `bson:"meals_cooked"`
	Points      int                `bson:"points"`
}

func (user *User) likedEatingToday() {
	user.updateMealsEaten(user.MealsEaten + 1)
}

func (user *User) unlikedEatingToday() {
	user.updateMealsEaten(user.MealsEaten - 1)
}

func (user *User) updateMealsEaten(updatedMealsEaten int) {
	// at this point, you've found the correct user from within the User collection.
	_, err := userCollection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.D{
			{"$set", bson.D{{"meals_eaten", updatedMealsEaten}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

type Eating struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	MessageID string             `bson:"message_id"`
	ChannelID string             `bson:"channel_id"`
}

func (eating *Eating) isCorrectMessageID(messageID string) bool {
	return eating.MessageID == messageID
}

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	TeamID   string `json:"team_id"`
}

type ActionInfo struct {
	ActionID string `json:"action_id"`
	BlockID  string `json:"block_id"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	ActionTS string `json:"action_ts"`
}

type ChannelInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContainerInfo struct {
	Type        string `json:"type"`
	MessageTS   string `json:"message_ts"`
	ChannelID   string `json:"channel_id"`
	IsEphemeral bool   `json:"is_ephemeral"`
}

type Interaction struct {
	User        UserInfo      `json:"user"`
	Channel     ChannelInfo   `json:"channel"`
	Container   ContainerInfo `json:"container"`
	Actions     []ActionInfo  `json:"actions"`
	ResponseUrl string        `json:"response_url"`
}

// this is the first function to be run right when the package is 'called' for the first time. This
// runs before the main function.
func init() {
	initializeEnvironmentVariables()
	initializeMongoDBConnection()
	api = slack.New(botSigningKey, slack.OptionDebug(true))
}

// creates the environment variables within the .env file to global constants
func initializeEnvironmentVariables() {
	// loads environment variables in .env to ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	signingSecret = os.Getenv("signingSecret")
	botSigningKey = os.Getenv("botSigningKey")
	mongoDBConnectionUri = os.Getenv("mongoDBConnectionUri")
	val, err := strconv.Atoi(os.Getenv("defaultServerTimeout"))
	if err != nil {
		log.Fatal(err)
	}
	defaultServerTimeout = time.Duration(val)
}

// Gets the mongodb client
func getMongoDBClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoDBConnectionUri).SetServerSelectionTimeout(defaultServerTimeout * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// gets the handles for the mongo collections and creates them if they didn't exist
func initializeMongoCollections(client *mongo.Client) {
	userCollection = client.Database("dinnyDB").Collection("users")
	eatingCollection = client.Database("dinnyDB").Collection("eating")
}

// verifies that the deployment is successfully connected and the Client was correctly configured.
func testMongoDBConnection(client *mongo.Client) {
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// create connection to MongoDB
func initializeMongoDBConnection() {
	client := getMongoDBClient()
	testMongoDBConnection(client)
	initializeMongoCollections(client)
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/event", eventsEndpoint)
		v1.POST("/interaction", interactionEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func interactionEndpoint(c *gin.Context) {
	userData := Interaction{}
	err := json.Unmarshal([]byte(c.PostForm("payload")), &userData)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%+v\n", userData)

	if len(userData.Actions) <= 0 {
		fmt.Printf("No actions recorded")
		return
	}

	// at this point, you know that the number of actions are greater than 0
	if userData.Actions[0].ActionID == "interaction_join_dinny" {
		err := addUserToDinnerRotation(userData)
		if err != nil {
			log.Fatal(err)
		}
		acknowledgeDecision(&userData)
	} else if userData.Actions[0].ActionID == "interaction_lurk_dinny" {
		acknowledgeDecision(&userData)
	}

	c.String(http.StatusOK, "")
}

func eventsEndpoint(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sv, err := slack.NewSecretsVerifier(c.Request.Header, signingSecret)
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

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		// here, you are extracting the underlying type of the interface type.
		switch innerEvent.Data.(type) {
		case *slackevents.ReactionAddedEvent:
			// TODO (ddritzenhoff) take another look at this. I'm not sure if this is the optimal solution
			reactionAddedEvent, _ := innerEvent.Data.(*slackevents.ReactionAddedEvent)
			handleReactionAddedEvent(reactionAddedEvent)
		case *slackevents.ReactionRemovedEvent:
			reactionRemovedEvent, _ := innerEvent.Data.(*slackevents.ReactionRemovedEvent)
			handleReactionRemovedEvent(reactionRemovedEvent)
		}
	}
}

func addUserToDinnerRotation(userData Interaction) error {
	newUser := &User{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		SlackUserId: userData.User.ID,
		Name:        userData.User.Name,
		MealsEaten:  0,
		MealsCooked: 0,
		Points:      0,
	}

	_, err := userCollection.InsertOne(ctx, newUser)
	return err
}

func getEatingDocument(eventMessageID string) Eating {
	var eating Eating
	// Todo (ddritzenhoff) can this be done differently?
	// the rational here was that there will only every be one document in the eating collection
	// so you don't have to filter to find it.
	err := eatingCollection.FindOne(ctx, bson.M{}).Decode(&eating)
	if err != nil {
		// this means that there isn't a document in the eating collection
		log.Fatal(err)
	}

	return eating
}

func getUser(slackUserID string) *User {
	// at this point, you know that the user liked the correct message
	var user User
	err := userCollection.FindOne(ctx, bson.M{"slack_user_id": slackUserID}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

// Remove 1 from the meals eaten if the correct 'will you eat today' message like is un-liked
func handleReactionRemovedEvent(reactionEvent *slackevents.ReactionRemovedEvent) {
	eventMessageID := reactionEvent.Item.Timestamp
	eating := getEatingDocument(eventMessageID)
	if !eating.isCorrectMessageID(eventMessageID) || isDesiredReaction(reactionEvent.Reaction, "+1") {
		return
	}

	// at this point, you know that the user liked the correct message
	// at this point, you've found the correct user from within the User collection.
	getUser(reactionEvent.User).unlikedEatingToday()
}

// Add 1 to the meals eaten if the correct 'will you eat today' message is liked
func handleReactionAddedEvent(reactionEvent *slackevents.ReactionAddedEvent) {
	eventMessageID := reactionEvent.Item.Timestamp
	eating := getEatingDocument(eventMessageID)
	if !eating.isCorrectMessageID(eventMessageID) || isDesiredReaction(reactionEvent.Reaction, "+1") {
		return
	}

	// at this point, you know that the user liked the correct message
	// at this point, you've found the correct user from within the User collection.
	getUser(reactionEvent.User).likedEatingToday()
}

func isDesiredReaction(eventReaction string, desiredReaction string) bool {
	return eventReaction == desiredReaction
}

func acknowledgeDecision(userData *Interaction) {
	_, _, err := api.PostMessage(userData.User.ID, slack.MsgOptionReplaceOriginal(userData.ResponseUrl), slack.MsgOptionText("Recorded, Thanks!", false))
	if err != nil {
		log.Fatal(err)
	}
}
