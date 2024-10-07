package domain

type Message struct{
	SenderID string
	Content string
	MediaURL *string
	Status []Status
}