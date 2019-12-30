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
	user         string
	points       int64
	reason       string
	userChanging string
}

type storage interface {
	pointIncrementer
}

type pointIncrementer interface {
	IncrementPoints(pc pointChange) error
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

func (p *mongodb) IncrementPoints(pc pointChange) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var options options.UpdateOptions
	_, err := p.points.UpdateOne(ctx, bson.M{"user": pc.user}, bson.M{"$inc": bson.M{"points": pc.points}}, options.SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}
