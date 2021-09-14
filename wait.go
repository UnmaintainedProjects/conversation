package conversation

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type ContextBool func(ctx *ext.Context) bool

// Waits for the next update in the chat of the provided context.
// Note that the context must have the `EffectiveChat` set in order to work.
// The validate parameter tells if the received update is valid or not,
// it can be `nil`. If the validate function is set and it returned `false`,
// the received update will not be counted and will be waiting for the next.
func (c *Conversation) Wait(ctx *ext.Context, validate ContextBool) *ext.Context {
	c.mu.Lock()
	_, ok := c.channels[ctx.EffectiveChat.Id]
	if !ok {
		c.channels[ctx.EffectiveChat.Id] = make(chan *ext.Context)
	}

	channel := c.channels[ctx.EffectiveChat.Id]
	c.mu.Unlock()

	answer := <-channel
	if validate != nil && !validate(answer) {
		return c.Wait(ctx, validate)
	}

	return answer
}

// Cancels the update listener for the chat of the provided context.
// Returns bool telling if it was canceled or not.
// It cancels it by sending `nil` to the channel.
func (c *Conversation) Cancel(ctx *ext.Context) bool {
	c.mu.Lock()
	channel, ok := c.channels[ctx.EffectiveChat.Id]
	c.mu.Unlock()
	if ok {
		channel <- nil
		delete(c.channels, ctx.EffectiveChat.Id)
		return true
	}

	return false
}

// Waits for a next message to be received in the chat of the provided context.
func (c *Conversation) WaitForMessage(ctx *ext.Context, validate ContextBool) *ext.Context {
	return c.Wait(
		ctx,
		func(ctx *ext.Context) bool {
			validateResult := true
			if validate != nil {
				validateResult = validate(ctx)
			}

			return ctx.Message != nil && validateResult
		},
	)
}

// Waits for a message to be received or edited in the chat of the provided context.
func (c *Conversation) WaitForEffectiveMessage(ctx *ext.Context, validate ContextBool) *ext.Context {
	return c.Wait(
		ctx,
		func(ctx *ext.Context) bool {
			validateResult := true
			if validate != nil {
				validateResult = validate(ctx)
			}

			return ctx.EffectiveMessage != nil && validateResult
		},
	)
}

// Waits for a message to be edited in the chat of the provided context.
func (c *Conversation) WaitForEditedMessage(ctx *ext.Context, validate ContextBool) *ext.Context {
	return c.Wait(
		ctx,
		func(ctx *ext.Context) bool {
			validateResult := true
			if validate != nil {
				validateResult = validate(ctx)
			}

			return ctx.EditedMessage != nil && validateResult
		},
	)
}

// Waits for a callback query to be received in the chat of the provided context.
// `message` is the message that we are waiting for the callback of its buttons,
// it should be provided so there are no conflicts with older messages with buttons.
func (c *Conversation) WaitForCallback(ctx *ext.Context, message *gotgbot.Message) *ext.Context {
	return c.Wait(
		ctx,
		func(ctx *ext.Context) bool {
			return ctx.CallbackQuery != nil && ctx.CallbackQuery.Message.MessageId == message.MessageId
		},
	)
}
