package signal

import "github.com/stampjohnny/mttv/context"

type Signal interface{}

type signalImpl struct {
	Signal
}

func Find(context context.Context) (Signal, error) {
	return &signalImpl{}, nil
}
