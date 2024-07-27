package main

import (
	"fmt"
	"log"
)

func main() {
	initDB()
	initRedis()

	err := postStory("user123", "This is a new story")
	if err != nil {
		log.Fatalf("error posting story: %v", err)
	}

	fmt.Println("Story posted successfully")

	stories, err := getStoriesForUser("user123")
	if err != nil {
		log.Fatalf("Error getting stories: %v", err)
	}

	for _, story := range stories {
		fmt.Println("Story:", story)
	}
}
