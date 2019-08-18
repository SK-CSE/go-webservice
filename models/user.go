package models

import (
	"errors"
	"fmt"
)

// User  : Fields that start with lower case characters are package internal and not exposed, If you want to reference the field from another package it needs to start with an upper case character
type User struct {
	ID    int
	FName string
	LName string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers : GetUsers
func GetUsers() []*User {
	return users
}

// AddUser : AddUser
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user ID will be auto incremented")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID : GetUserByID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
	return User{}, fmt.Errorf("user with id %v not present", id)
}

// UpdateUser : UpdateUser
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
		}
	}
	return User{}, fmt.Errorf("user with id %v not present", u.ID)
}

// RemoveUserByID : RemoveUserByID
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id %v not present", id)
}
