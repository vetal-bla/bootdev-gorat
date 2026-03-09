package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/vetal-bla/bootdev-gorat/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBeetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Can't parse duration %s\n error: %v", cmd.Args[0], err)
	}

	log.Printf("collecting feeds every %s", timeBeetweenReqs)

	ticker := time.NewTicker(timeBeetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: %s <title> <url>", cmd.Name)
	}

	title := cmd.Args[0]
	urlFeed := cmd.Args[1]
	currentUser := user

	feed := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       urlFeed,
		UserID:    currentUser.ID,
	}

	newFeed, err := s.db.AddFeed(context.Background(), feed)
	if err != nil {
		fmt.Errorf("Cant add feed. Sql error: %v", err)
	}
	if newFeed.Url == "" {
		fmt.Printf("Url %s alread exist in database", urlFeed)
	} else {
		followArguments := database.FollowFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    currentUser.ID,
			FeedID:    newFeed.ID,
		}
		_, err := s.db.FollowFeed(context.Background(), followArguments)
		if err != nil {
			return fmt.Errorf("Can't follow feed: %v", err)
		}

		fmt.Printf("ID: %s\n", newFeed.ID)
		fmt.Printf("Name (title): %s\n", newFeed.Name)
		fmt.Printf("URL: %s\n", newFeed.Url)
		fmt.Printf("User id: %s\n", newFeed.UserID)
	}

	return nil
}

func handdlerGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Cant get feeds from database: %v", err)
	}

	for _, item := range feeds {
		fmt.Printf("User: %s\n", item.Username)
		fmt.Printf("Feed name: %s\n", item.Name)
		fmt.Printf("Feed url: %s\n", item.Url)
		fmt.Println("---")
	}
	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("can't get error")
	}

	if feed.Name == "" {
		return fmt.Errorf("Feed was empty. Cant follow it")
	}

	currentUser := user
	followArguments := database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	insertedFeed, err := s.db.FollowFeed(context.Background(), followArguments)
	if err != nil {
		return fmt.Errorf("Can't follow feed: %v", err)
	}

	fmt.Printf("%s you are now following %s\n", insertedFeed[0].UserName, insertedFeed[0].FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Can't get feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("UserName: %s | FeedName: %s\n", feed.UserName, feed.FeedName)
		fmt.Println("---")
	}

	return nil

}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Cant get feed")
	}
	if feed.Url == "" {
		return fmt.Errorf("No feed with this url")
	}

	deleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), deleteParams)
	if err != nil {
		return fmt.Errorf("Cant unfollow: %v", err)
	}

	return nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Can't get feed for fetch: %v", err)
		return
	}
	log.Println("Found a feed to fetch!")

	scrapeFeed(s.db, feed)

}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Can't update feed: %v\n", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Can't fetch feed url: %s\n error: %v\n", feed.Url, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	fmt.Printf("Feed %s collected. Post founds - %d", feed.Name, len(feedData.Channel.Item))
}
