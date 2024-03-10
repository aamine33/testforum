package forum

type User struct {
	Username        string
	Email           string
	CreationDate    string
	BirthDate       string
	Uuid            string
	Admin           int
	TopicsCreated   []string
	MessagesSent    []MessageSend
	UuidOfTopics    []string
	PrivateMessages []PrivateMessage
}

type PrivateMessage struct {
	PrivateMessage        string
	PrivateMessage2ndUser string
}

type MessageSend struct {
	MessageSendByUser string
	TopicSentInName   string
}

func (u *User) AddMessageSent(message MessageSend) {
	u.MessagesSent = append(u.MessagesSent, message)
}

func (u *User) AddTopicCreated(topicName string) {
	u.TopicsCreated = append(u.TopicsCreated, topicName)
}

func (u *User) GetTopicsCreated() []string {
	return u.TopicsCreated
}

func (u *User) AddPrivateMessage(message PrivateMessage) {
	u.PrivateMessages = append(u.PrivateMessages, message)
}
