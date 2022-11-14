package network_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"tic-cat-dog-client/src/network"
)

func clientUUID(t *testing.T) uuid.UUID {
	result, err := uuid.NewRandom()
	require.NoError(t, err)

	return result
}

func TestToBytesCommand(t *testing.T) {
	message := network.NewConnectMessage(clientUUID(t))
	bytes, _ := message.MarshalBinary()
	connectAsByte, _ := network.Command.MarshalBinary(network.Connect)

	assert.Equal(t, connectAsByte[0], bytes[0])
}

func TestToBytesMessage(t *testing.T) {
	expectedClientUUID := clientUUID(t)
	message := network.NewConnectMessage(expectedClientUUID)
	bytes, _ := message.MarshalBinary()
	actualClientUUIDBytes := bytes[1:]
	actualClientUUID, err := uuid.FromBytes(actualClientUUIDBytes)
	require.NoError(t, err)

	assert.Equal(t, expectedClientUUID, actualClientUUID)
}
