package conversation

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

// Requires context to have `EffectiveMessage`.
func EffectiveMessage(ctx *ext.Context) bool {
	return ctx.EffectiveMessage != nil
}

// Requires `EffectiveMessage` in context to have `Text`.
func Text(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Text != ""
}

// Requires `EffectiveMessage` in context to have `Photo`.
func Photo(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Photo != nil
}

// Requires `EffectiveMessage` in context to have `Video`.
func Video(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Video != nil
}

// Requires `EffectiveMessage` in context to have `Animation`.
func Animation(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Animation != nil
}

// Requires `EffectiveMessage` in context to have `Audio`.
func Audio(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Audio != nil
}

// Requires `EffectiveMessage` in context to have `Voice`.
func Voice(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Voice != nil
}

// Requires `EffectiveMessage` in context to have `Document`.
func Document(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Document != nil
}

// Requires `EffectiveMessage` in context to have `Caption`.
func Caption(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Caption != ""
}

// Requires `EffectiveMessage` in context to have `Dice`.
func Dice(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Dice != nil
}

// Requires `EffectiveMessage` in context to have `Location`.
func Location(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Location != nil
}

// Requires `EffectiveMessage` in context to have `Contact`.
func Contact(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.Contact != nil
}

// Requires `EffectiveMessage` in context to have `VideoNote`.
func VideoNote(ctx *ext.Context) bool {
	return EffectiveMessage(ctx) && ctx.EffectiveMessage.VideoNote != nil
}

// Requires `EffectiveMessage` in context to have `Audio` or `Voice`.
func AnyAudio(ctx *ext.Context) bool {
	return Audio(ctx) || Voice(ctx)
}

// Requires `EffectiveMessage` in context to have `Video` or `VideoNote`.
func AnyVideo(ctx *ext.Context) bool {
	return Video(ctx) || VideoNote(ctx)
}
