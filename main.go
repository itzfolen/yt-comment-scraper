package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <videoID>", os.Args[0])
	}

	videoID := os.Args[1]
	apiKey := "INPUT_YOU_YOUTUBE_DATA_API_V3_KEY"
	outputFileName := "comments.csv"

	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Could not create file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Nick", "Time", "Comment", "Likes"})

	service, err := youtube.NewService(nil, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	var nextPageToken string

	for {
		call := service.CommentThreads.List([]string{"snippet"}).VideoId(videoID).TextFormat("plainText").PageToken(nextPageToken)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error fetching comments: %v", err)
		}

		for _, item := range response.Items {
			snippet := item.Snippet.TopLevelComment.Snippet
			nick := snippet.AuthorDisplayName
			time := snippet.PublishedAt
			comment := snippet.TextDisplay
			likes := fmt.Sprintf("%d", snippet.LikeCount)

			fmt.Printf("Nick: %s\nTime: %s\nComment: %s\nLikes: %s\n\n", nick, time, comment, likes)

			writer.Write([]string{nick, time, comment, likes})
		}

		if response.NextPageToken == "" {
			break
		}

		nextPageToken = response.NextPageToken
	}

	fmt.Printf("Comments have been saved to %s\n", outputFileName)
}
