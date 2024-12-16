package models

type MessageType string

const (
	UserM MessageType = "user"
	SystemM MessageType = "system"
)