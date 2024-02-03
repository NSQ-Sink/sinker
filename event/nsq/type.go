package nsq

type Event struct {
	topic  string // topic name
	source []string
}
