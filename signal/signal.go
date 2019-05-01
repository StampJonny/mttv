package signal

import "github.com/stampjohnny/mttv/context"

type Signal interface {
	IsBuySignal() bool
}

type signalImpl struct {
	Signal
}

type BuySignal struct {
	signalImpl
}

func (s *signalImpl) IsBuySignal() bool {
	return true
}

func Find(context context.Context) (Signal, error) {
	return &signalImpl{}, nil
}
