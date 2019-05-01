package trader_test

import (
	"testing"

	"github.com/stampjohnny/mttv/context"
	"github.com/stampjohnny/mttv/signal"
	"github.com/stretchr/testify/assert"
)

func TestBuy(t *testing.T) {
	ctx := context.New()

	signal, err := signal.Find(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, signal)
}
