package server

type HandlerConfig struct {
	Name    string
	Host    string
	Port    string
	Options map[string]interface{}
}

type Handler interface {
	Handle(*Message) error
	Close() error
}
