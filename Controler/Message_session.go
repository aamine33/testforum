package forum

import (
	"fmt"
)

type Messages struct {
	Id           int
	Name         string
	Likes        int
	Dislikes     int
	CreationDate string
	Owner        string
	Uuid         string
	UuidPath     string
	SessionUser  string
	Category     string
	Error        string
	Messages     []Message `Message`
}

type Message struct {
	Message        string
	CreationDate   string
	Owner          string
	Report         int
	Uuid           string
	Id             int
	Like           int
	Edited         int
	IsLiked        int
	IsDisliked     int
	IsOwnerOrAdmin int
}

func (m *Messages) AddMessage(newMessage Message) {
	m.Messages = append(m.Messages, newMessage)
}

func (m *Messages) RemoveMessage(uuid string) {
	for i, msg := range m.Messages {
		if msg.Uuid == uuid {
			m.Messages = append(m.Messages[:i], m.Messages[i+1:]...)
			return
		}
	}
}

func (m *Messages) UpdateLikes(uuid string, likes int) error {
	for i, msg := range m.Messages {
		if msg.Uuid == uuid {
			m.Messages[i].Like = likes
			return nil
		}
	}
	return fmt.Errorf("message with UUID %s not found", uuid)
}
