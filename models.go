package main

import (
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type FeedFollowList struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Url  string    `json:"url"`
}

type Post struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedtoFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedstoFeeds(feed []database.Feed) []Feed {
	jsonFeed := make([]Feed, len(feed))
	for idx, val := range feed {
		jsonFeed[idx] = databaseFeedtoFeed(val)
	}
	return jsonFeed
}

func databaseFeedFollowtoFeedFollow(feed_follow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feed_follow.ID,
		CreatedAt: feed_follow.CreatedAt,
		UpdatedAt: feed_follow.UpdatedAt,
		UserID:    feed_follow.UserID,
		FeedID:    feed_follow.FeedID,
	}
}

func databaseFeedFollowstoFeedFollows(feed_follows []database.GetMyFollowsRow) []FeedFollowList {
	jsonFeedList := make([]FeedFollowList, len(feed_follows))
	for idx, val := range feed_follows {
		jsonFeedList[idx] = FeedFollowList{
			ID:   val.ID,
			Name: val.Name,
			Url:  val.Url,
		}
	}
	return jsonFeedList
}

func databasePoststoPosts(posts []database.GetPostListForUserRow) []Post {
	jsonPostList := make([]Post, len(posts))
	for idx, val := range posts {
		jsonPostList[idx] = Post{
			Title:       val.Title,
			Url:         val.Url,
			Description: val.Description.String,
		}
	}
	return jsonPostList
}
