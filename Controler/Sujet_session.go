package forum

import "fmt"

type TopicsAndSession struct {
	Error       string
	SessionUser string
	Category    string
	Topics      []Topic `Topic`
}

type Topic struct {
	Id             int
	Name           string
	Likes          int
	CreationDate   string
	Owner          string
	Uuid           string
	FirstMessage   string
	NmbPosts       int
	LastPost       string
	IsLiked        int
	IsDisliked     int
	Category       string
	IsOwnerOrAdmin int
}

func (t *TopicsAndSession) AddTopic(newTopic Topic) {
	t.Topics = append(t.Topics, newTopic)
}

func (t *TopicsAndSession) RemoveTopic(uuid string) {
	for i, topic := range t.Topics {
		if topic.Uuid == uuid {
			t.Topics = append(t.Topics[:i], t.Topics[i+1:]...)
			return
		}
	}
}

func (t *TopicsAndSession) UpdateLikes(uuid string, likes int) error {
	for i, topic := range t.Topics {
		if topic.Uuid == uuid {
			t.Topics[i].Likes = likes
			return nil
		}
	}
	return fmt.Errorf("topic with UUID %s not found", uuid)
}
