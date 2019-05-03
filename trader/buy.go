package trader

import (
	"github.com/google/uuid"
	"github.com/stampjohnny/mttv/exchange"
)

func Buy() (uuid.UUID, error) {
	tx, err := exchange.Buy()
	return tx, err
}
