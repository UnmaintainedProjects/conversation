package conversation

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Conversation struct {
	mu       *sync.Mutex
	channels map[int64]chan *ext.Context
}

func New() *Conversation {
	return &Conversation{
		mu:       &sync.Mutex{},
		channels: make(map[int64]chan *ext.Context),
	}
}
