package core

import (
	"github.com/google/uuid"
	"net"
	"tic-cat-dog-client/src/network"
)

type Client struct {
	Connection net.Conn
	Id         uuid.UUID
}

func (c *Client) Connect() {
	message, err := network.NewConnectMessage(c.Id).MarshalBinary()

	if err != nil {
		HandleError(err)
	}

	if _, err := c.Connection.Write(message); err != nil {
		HandleError(err)
	}
}
