package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ddritzenhoff/dinny/pkg/api"
	"github.com/ddritzenhoff/dinny/pkg/app"
	"github.com/ddritzenhoff/dinny/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up db connections, routers etc
func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize environment variables
	botSigningKey := os.Getenv("botSigningKey")
	mongoDBConnectionUri := os.Getenv("mongoDBConnectionUri")
	signingSecret := os.Getenv("signingSecret")
	val, err := strconv.Atoi(os.Getenv("defaultServerTimeout"))
	if err != nil {
		log.Fatal(err)
	}
	defaultServerTimeout := time.Duration(val)

	// setup database connection
	db, err := setupDatabase(mongoDBConnectionUri, defaultServerTimeout)

	if err != nil {
		return err
	}

	slackApi := slack.New(botSigningKey, slack.OptionDebug(true))

	// create storage dependency
	storage := repository.NewStorage(db)

	// create router dependecy
	router := gin.Default()

	// create a cook service
	cookService := api.NewCookService(storage)

	// create a participant service
	participantService := api.NewParticipantService(storage, slackApi)

	server := app.NewServer(router, cookService, participantService, signingSecret)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connectionString string, defaultServerTimeout time.Duration) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connectionString).SetServerSelectionTimeout(defaultServerTimeout * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// ping the DB to know that it is connected

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
