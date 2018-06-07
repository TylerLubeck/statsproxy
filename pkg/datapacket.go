package server

type Message struct {
	Message string
}

func NewMessage(msg string) *Message {
	return &Message{Message: msg}
}

func (m *Message) ToString() string {
	return m.Message
}

func (m *Message) GetType() string {
	return "basic"
}

func (m *Message) ToFormat(format string) (string, error) {
	return m.Message, nil
}
