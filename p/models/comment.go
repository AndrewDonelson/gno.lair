// Package models provides functionality for generic models. This file is for managing comments
package models

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// CommentInterface defines the contract for the Comment class.
type CommentInterface interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	GetPostID() uuid.UUID
	SetPostID(uuid.UUID)
	GetUserID() uuid.UUID
	SetUserID(uuid.UUID)
	GetContent() string
	SetContent(string)
	Display() string
}

// Comment represents a comment on a post with basic properties and methods.
type Comment struct {
	ID      uuid.UUID
	PostID  uuid.UUID
	UserID  uuid.UUID
	Content string
}

// GetID returns the ID of a comment.
func (c *Comment) GetID() uuid.UUID {
	return c.ID
}

// SetID sets the ID of a comment.
func (c *Comment) SetID(id uuid.UUID) {
	c.ID = id
}

// GetID returns the ID of a comment.
func (c *Comment) GetPostID() uuid.UUID {
	return c.PostID
}

// SetID sets the ID of a comment.
func (c *Comment) SetPostID(id uuid.UUID) {
	c.PostID = id
}

// GetUserID returns the user ID associated with a comment.
func (c *Comment) GetUserID() uuid.UUID {
	return c.UserID
}

// SetUserID sets the user ID associated with a comment.
func (c *Comment) SetUserID(userID uuid.UUID) {
	c.UserID = userID
}

// GetContent returns the content of a comment.
func (c *Comment) GetContent() string {
	return c.Content
}

// SetContent sets the content of a comment.
func (c *Comment) SetContent(content string) {
	c.Content = content
}

// Display returns a formatted string containing the comment's details.
func (c *Comment) Display() string {
	return fmt.Sprintf("ID: %d, UserID: %d, Content: %s", c.GetID(), c.GetUserID(), c.GetContent())
}
