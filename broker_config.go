package gocelery

const (
	defaultQueue = "celery"
)

type BrokerConfig struct {
	Queue string
}

func NewBrokerConfig() *BrokerConfig {
	return &BrokerConfig{Queue: defaultQueue}
}
