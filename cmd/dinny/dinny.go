package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants from env file
var botSigningKey string
var mongoDBConnectionUri string
var defaultServerTimeout time.Duration
var currentChannelID string

// DB connection
var userCollection *mongo.Collection
var eatingCollection *mongo.Collection
var ctx = context.TODO()

// api
var api *slack.Client

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

type Eating struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	MessageID string             `bson:"message_id"`
	ChannelID string             `bson:"channel_id"`
}

// creates the environment variables within the .env file to global constants
func initializeEnvironmentVariables() {
	// loads environment variables in .env to ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botSigningKey = os.Getenv("botSigningKey")
	mongoDBConnectionUri = os.Getenv("mongoDBConnectionUri")
	currentChannelID = os.Getenv("currentChannelID")
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

// this is the first function to be run right when the package is 'called' for the first time. This
// runs before the main function.
func init() {
	initializeEnvironmentVariables()
	initializeMongoDBConnection()
	api = slack.New(botSigningKey, slack.OptionDebug(true))
}

func main() {
	err := getApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getApp() *cli.App {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "dinny",
		Authors: []*cli.Author{
			{
				Name:  "Dominik Ritzenhoff",
				Email: "ddritzenhoff@gmail.com",
			},
		},
		Usage: "Get dinner_rotation_bot info",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get a value",
				Subcommands: []*cli.Command{
					{
						Name:  "user_info",
						Usage: "gets information about the user",
						Action: func(c *cli.Context) error {
							if c.NArg() > 0 {
								cliPrintUserInfo(c.Args().Get(0))
							}
							return nil
						},
					},
					{
						Name:  "members",
						Usage: "Gives basic stats about all the members in the dinner rotation channel",
						Action: func(c *cli.Context) error {
							cliPrintUsersInTopicDinnerRotation()
							return nil
						},
					},
				},
			},
			{
				Name:  "server",
				Usage: "starts the dinner rotation backend",
				Subcommands: []*cli.Command{
					{
						Name:  "start_semester",
						Usage: "starts the dinner rotation semester",
						Action: func(c *cli.Context) error {
							cliStartDinnerRotation()
							return nil
						},
					},
					{
						Name:  "is_eating_today",
						Usage: "asks who is eating today",
						Action: func(c *cli.Context) error {
							cliIsEatingToday()
							return nil
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))
	return app
}

// this creates the block used to ask user if he/she wants to join dinner rotation this semester.
func 
startDinnerRotationBlock(userID string) slack.MsgOption {

	// Shared Assets for example
	divSection := slack.NewDividerBlock()

	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", "hey <@"+userID+">, would you like to join dinner rotation this semester?", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// div

	// Join and Lurk Buttons
	approveBtnTxt := slack.NewTextBlockObject("plain_text", "Join", false, false)
	approveBtn := slack.NewButtonBlockElement("interaction_join_dinny", "click_me_123", approveBtnTxt).WithStyle(slack.StylePrimary)

	denyBtnTxt := slack.NewTextBlockObject("plain_text", "Lurk", false, false)
	denyBtn := slack.NewButtonBlockElement("interaction_lurk_dinny", "click_me_123", denyBtnTxt).WithStyle(slack.StyleDefault)

	actionBlock := slack.NewActionBlock("", approveBtn, denyBtn)

	return slack.MsgOptionBlocks(
		headerSection,
		divSection,
		actionBlock,
	)
}

func isEatingTodayBlock() slack.MsgOption {
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", "hey <!channel>, please react to this message (:thumbsup:) if you are eating tomorrow", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	return slack.MsgOptionBlocks(
		headerSection,
	)
}

func cliIsEatingToday() {
	// passing bson.D{{}} matches all documents in the collection
	respChannel, respTimestamp, err := api.PostMessage(currentChannelID, isEatingTodayBlock())
	if err != nil {
		log.Fatal(err)
	}

	var eating Eating
	if err = eatingCollection.FindOne(ctx, bson.M{}).Decode(&eating); err != nil {
		fmt.Printf("Nothing found\n")
		newMessage := &Eating{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			MessageID: respTimestamp,
			ChannelID: respChannel,
		}
		_, err = eatingCollection.InsertOne(ctx, newMessage)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("FOUND SOMETHING\n")
		_, err := eatingCollection.UpdateOne(
			ctx,
			bson.M{"_id": eating.ID},
			bson.D{
				{"$set", bson.D{{"message_id", respTimestamp}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cliStartDinnerRotation() {
	userIDs, _, err := api.GetUsersInConversation(&slack.GetUsersInConversationParameters{ChannelID: currentChannelID, Cursor: "", Limit: 100})
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, userID := range userIDs {
		_, _, err := api.PostMessage(userID, startDinnerRotationBlock(userID))
		// this is to prevent rate limiting. see https://api.slack.com/docs/rate-limits
		time.Sleep(2 * time.Second)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	}
}

func cliPrintUsersInTopicDinnerRotation() {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	users, err := filterUsers(filter)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		cliPrintUserInfo(user.SlackUserId)
	}
}

func filterUsers(filter interface{}) ([]*User, error) {
	var users []*User

	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var u User
		err := cur.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &u)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil
}

func cliPrintUserInfo(userID string) {
	// UQZN1SE8N
	user, err := api.GetUserInfo(userID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("ID: %s, Fullname: %s, DisplayName: %s, isBot: %t\n", user.ID, user.Profile.RealName, user.Profile.DisplayName, user.IsBot)
}
