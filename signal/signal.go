package signal

import "github.com/stampjohnny/mttv/context"

type ResponseStatus int

const (
	StatusOK ResponseStatus = iota
)

type Response interface {
	GetStatus() ResponseStatus
}

type responseImpl struct {
	Response

	Status ResponseStatus
}

func (r *responseImpl) GetStatus() ResponseStatus {
	return StatusOK
}

type Signal interface {
	IsBuySignal() bool
	GetResponse(context.Context) Response
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

func GetResponse(ctx context.Context) Response {
	return &responseImpl{
		Status: StatusOK,
	}
}

func Find(context context.Context) (Signal, error) {
	return &signalImpl{}, nil
}
