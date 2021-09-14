package conversation

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (c *Conversation) CheckUpdate(b *gotgbot.Bot, u *gotgbot.Update) bool {
	return u.Message != nil || u.CallbackQuery != nil
}

func (c *Conversation) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	c.mu.Lock()
	channnel, ok := c.channels[ctx.EffectiveChat.Id]
	c.mu.Unlock()
	if !ok {
		return nil
	}

	channnel <- ctx
	return nil
}

func (c *Conversation) Name() string {
	return fmt.Sprintf("conversation_%p", c.HandleUpdate)
}
