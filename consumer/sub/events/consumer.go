package events

type Consumer interface {
	Subscribe(topics []string) error
}
