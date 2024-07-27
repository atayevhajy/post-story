package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Story struct {
	UserID    string    `bson:"user_id"`
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"timestamp"`
}

var mongoClient *mongo.Client

func init() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}
}

func postStory(userID, content string) error {
	collection := mongoClient.Database("stories_db").Collection("stories")

	story := Story{
		UserID:    userID,
		Content:   content,
		Timestamp: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), story)
	if err != nil {
		return fmt.Errorf("mongodb insert failed: %v", err)
	}

	key := "user:" + userID + ":latest_story"
	err = redisClient.Set(ctx, key, content, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("redis set failed: %v", err)
	}

	return nil
}

func getStoriesForUser(userID string) ([]string, error) {
	followerIDs, err := getFollowers(userID)
	if err != nil {
		return nil, err
	}

	var stories []string
	for _, id := range followerIDs {
		content, err := redisClient.Get(ctx, "user:"+id+":latest_story").Result()
		if err == nil {
			stories = append(stories, content)
		} else {
			collection := mongoClient.Database("stories_db").Collection("stories")
			cursor, err := collection.Find(context.TODO(), bson.M{"user_id": id})
			if err != nil {
				return nil, err
			}
			var results []Story
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}
			for _, story := range results {
				stories = append(stories, story.Content)
			}
		}
	}
	return stories, nil
}
