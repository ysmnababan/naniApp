package domain

type Conversation struct{
	ConvID string
	ConvName string
	ConvType string
	Participants []Participant
	Messages []Message
}