package main

import (
	"errors"
	"time"
	"sort"
	"strings"
)

type Tweets struct{
	contant string
	time string
}

var Posts = make(map[string][]Tweets)

func addUserPosts(name string){
	Posts[name] = append(Posts[name])
}

func PostTweet(userName string, tweetContent string) (err error){
	var newTweet Tweets
	now := time.Now()

	if _, ok := Posts[userName]; ok{
		newTweet.contant = tweetContent
		newTweet.time = now.Format("2006-01-02 15:04:05")
		Posts[userName] = append(Posts[userName], newTweet)
		return nil
	}
	return errors.New("user not found in system")

}

func GetUserFeed(userName string) (tweets []Tweets, err error){
	var userFollowers = Followers[userName]
	var allFollowersTweets = make([]Tweets, 0)

	for _, follower := range userFollowers{
		if _, ok := Posts[follower]; ok {
			for _, value := range Posts[follower]{
				allFollowersTweets = append(allFollowersTweets, value)
			}
		}
	}

	sort.Slice(allFollowersTweets, func(i, j int) bool{
		switch strings.Compare(allFollowersTweets[j].time, allFollowersTweets[i].time){
		case -1:
			return true
		case 1:
			return false
		}
		return allFollowersTweets[j].contant > allFollowersTweets[i].contant
	})

	return allFollowersTweets, nil
}
