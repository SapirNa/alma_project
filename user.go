package main

import (
	"errors"
)

var Users = []string{}

func CreateUser(name string) (err error){
	if name == "" {
		return errors.New("empty name")
	}
	for _, value := range Users{
		if value == name{
			return errors.New("user is alrady exist")
		}
	}
	Users = append(Users, name)
	addUserToFollowers(name)
	addUserPosts(name)
	return nil
}

func GetUser(name string) (User string, err error){
	for _, value := range Users {
		if value == name{
			return name, nil
		}
	}
	return "", errors.New("user not found")
}

func UpdateUser(oldName string, newName string) (err error){
	for idx, value := range Users {
		if value == oldName{
			Users[idx] = newName
			return nil
		}
	}
	return errors.New("user not found, name not update")

}

func DeleteUser(name string) (err error){
	for idx, value := range Users {
		if value == name{
			newUsers := make([]string, 0)
			newUsers = append(newUsers, Users[:idx]...)
			Users = append(newUsers, Users[idx+1:]...)
			return nil
		}
	}
	return errors.New("user not found, name not delete")
}


