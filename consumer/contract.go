package consumer

type Consumer interface {
	Run() error
	Stop() error
}
