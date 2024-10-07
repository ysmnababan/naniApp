package domain

type Message struct{
	MessageID string
	SenderID string
	Content string
	MediaURL *string
	Status []Status
}