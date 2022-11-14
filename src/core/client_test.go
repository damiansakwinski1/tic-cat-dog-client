package core_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net"
	"testing"
	"tic-cat-dog-client/src/core"
	"tic-cat-dog-client/src/network"
)

func TestClientConnection(t *testing.T) {
	mockRead, mockWrite := net.Pipe()
	clientUUID, _ := uuid.NewRandom()
	client := core.Client{Connection: mockWrite, Id: clientUUID}
	go func() {
		client.Connect()
		mockWrite.Close()
		mockRead.Close()
	}()

	response, err := io.ReadAll(mockRead)
	require.NoError(t, err)

	connectAsBytes, _ := network.Command.MarshalBinary(network.Connect)
	clientUUIDBytes, _ := clientUUID.MarshalBinary()
	expected := append(connectAsBytes, clientUUIDBytes...)

	assert.Equal(t, expected, response)
}
