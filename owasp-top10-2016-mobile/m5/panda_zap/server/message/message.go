package message

// Message is the struct that holds information on a message.
type Message struct {
	ID            string `bson:"id" json:"id"`
	Owner         string `bson:"owner" json:"owner"`
	EncryptedText string `bson:"text" json:"text"`
	Date          string `bson:"date" json:"date"`
}
