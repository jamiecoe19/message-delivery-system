package message

type Message interface {
	Send() error
}
