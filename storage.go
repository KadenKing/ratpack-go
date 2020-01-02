package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pointChange struct {
	User         string `bson:"user"`
	Points       int64  `bson:"points"`
	Reason       string `bson:"reason"`
	UserChanging string `bson:"changer"`
}

type storage interface {
	pointIncrementer
}

type pointIncrementer interface {
	IncrementPoints(userID string, pointValue int64) error
}

type mongodb struct {
	points *mongo.Collection
	log    *mongo.Collection
}

func newMongodb(env environment) *mongodb {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Get("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("ratpack")
	points := db.Collection("points")
	log := db.Collection("log")
	return &mongodb{points, log}
}

func (p *mongodb) IncrementPoints(userID string, pointValue int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var options options.UpdateOptions
	_, err := p.points.UpdateOne(ctx, bson.M{"user": userID}, bson.M{"$inc": bson.M{"points": pointValue}}, options.SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}

func logPointChange(log *mongo.Collection, pc pointChange) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := log.InsertOne(ctx, pc)
	if err != nil {
		return err
	}

	return nil
}
