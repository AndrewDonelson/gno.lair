// Package models provides functionality for generic models. This file is for managing users.
package models

import "fmt"

// UserInterface defines the contract for the User class.
type UserInterface interface {
	GetID() int
	SetID(int)
	GetFullName() string
	SetFullName(string)
	GetUserName() string
	SetUserName(string)
	GetEmail() string
	SetEmail(string)
	Display() string
}

// User represents a user with basic properties and methods.
type User struct {
	ID       int
	FullName string
	Username string
	Email    string
}

// GetID returns the ID of a user.
func (u *User) GetID() int {
	return u.ID
}

// SetID sets the ID of a user.
func (u *User) SetID(id int) {
	u.ID = id
}

// GetFullName returns the name of a user.
func (u *User) GetFullName() string {
	return u.FullName
}

// SetFullName sets the name of a user.
func (u *User) SetFullName(name string) {
	u.FullName = name
}

// GetUserName returns the name of a user.
func (u *User) GetUserName() string {
	return u.Username
}

// SeUserName sets the name of a user.
func (u *User) SetUserName(name string) {
	u.Username = name
}

// GetEmail returns the email address of a user.
func (u *User) GetEmail() string {
	return u.Email
}

// SetEmail sets the email address of a user.
func (u *User) SetEmail(email string) {
	u.Email = email
}

// Display returns a formatted string containing the user's details.
func (u *User) Display() string {
	return fmt.Sprintf("ID: %d, FullName: %s, Email: %s", u.GetID(), u.GetFullName(), u.GetEmail())
}
