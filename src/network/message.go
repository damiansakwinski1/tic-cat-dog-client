package network

import (
	"github.com/google/uuid"
	"os"
)

type ConnectMessage struct {
	Message
}

func NewConnectMessage(clientUUID uuid.UUID) ConnectMessage {
	return ConnectMessage{Message: Message{Type: Connect, ClientId: clientUUID}}
}

type Message struct {
	Type     Command
	ClientId uuid.UUID
}

func (m Message) MarshalBinary() ([]byte, error) {
	clientUUID, err := m.ClientId.MarshalBinary()

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	result := make([]byte, 1)
	result[0] = byte(m.Type)
	result = append(result, clientUUID...)

	return result, nil
}
