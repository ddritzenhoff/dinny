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
				},
			},
				},
			},
	sort.Sort(cli.CommandsByName(app.Commands))
	return app
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
