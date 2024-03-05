package main

import (
	"fmt"
	"sort"
)

var Followers = make(map[string][]string)

func addUserToFollowers(name string){
	Followers[name] = append(Followers[name])
}

func FollowUser(followerName string, followingName string) (err error){
	if _, ok := Followers[followerName]; ok{
		if _, ok := Followers[followingName]; ok{
			Followers[followerName] = append(Followers[followerName], followingName)
			return nil
		}
		return fmt.Errorf("user %s not found in system", followingName)
	}
	return fmt.Errorf("user %s not found in system", followerName)

}

func UnFollowUser(followerName string, followingName string) (err error){
	for idx, value := range Followers[followerName]{
		if value == followingName{
			newUsers := make([]string, 0)
			newUsers = append(newUsers, Users[:idx]...)
			Users = append(newUsers, Users[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user %s not following %s", followerName, followingName)
}

func GetMutualFollowers(userName1 string, userName2 string) (commonUsers []string, err error){
	if _, ok := Followers[userName1]; !ok{
		return commonUsers, fmt.Errorf("user %s not found in system", userName1)
	}
	if _, ok := Followers[userName2]; !ok{
		return commonUsers, fmt.Errorf("user %s not found in system", userName2)
	}

	for _, user1_follower := range Followers[userName1]{
		for _, user2_follower := range Followers[userName2]{
			if user1_follower == user2_follower{
				commonUsers = append(commonUsers, user1_follower)
				break
			}
		}
	}
	return commonUsers, nil
}

func GetTopInfluencers(n int) (influencers []string, err error){
	var totalFollowers = make(map[string]int)
	for idx, value := range Followers{
		totalFollowers[idx] = len(value)
	}
	totalFollowers_values := make([]int, 0, len(totalFollowers))
	for _, followers_count := range totalFollowers{
		totalFollowers_values = append(totalFollowers_values, followers_count)

	}
	sort.Ints(totalFollowers_values)

	totalFollowers_keys := make([]string, 0, len(totalFollowers))
	for follower_name, _ := range totalFollowers{
		totalFollowers_keys = append(totalFollowers_keys, follower_name)
	}
	return totalFollowers_keys[:n], nil
}