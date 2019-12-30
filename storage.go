package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storage interface {
	pointIncrementer
}

type pointIncrementer interface {
	IncrementPoints(user string, points int64) error
}

type mongodb struct {
	points *mongo.Collection
}

func newMongodb(env environment) *mongodb {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Get("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	points := client.Database("ratpack").Collection("points")
	return &mongodb{points}
}

func (p *mongodb) IncrementPoints(user string, points int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var options options.UpdateOptions
	p.points.UpdateOne(ctx, bson.M{"user": user}, bson.M{"$inc": bson.M{"points": points}}, options.SetUpsert(true))
	return nil
}
