package network

type Command byte

const (
	Connect Command = 1
	Move    Command = 2
)

func (cmd Command) MarshalBinary() ([]byte, error) {
	return []byte{byte(cmd)}, nil
}
