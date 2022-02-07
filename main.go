package main

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/utils/bot"
	"github.com/diamondburned/arikawa/v3/discord"
)

func main() {
	bot.Run("xxxxx", &Bot{},
		func(ctx *bot.Context) error {
			ctx.HasPrefix = bot.NewPrefix("!")
			return nil
		},
	)
}

type Bot struct {
	Ctx *bot.Context
}

func (b *Bot) Ping(*gateway.MessageCreateEvent) (string, error) {
	return "Pong!", nil
}

func (bot *Bot) Source(message *gateway.MessageCreateEvent) (error) {

	color := 0x2ECC70
	TheTitle := "This bot's source code"
	TheDescription := "The source code for this bot is open source, and is available on Github.\nhttps://www.github.com/KanchiMoe/AngelPlayer"

	embed := discord.Embed{
		Title:       TheTitle,
		Description: TheDescription,
		Color:       discord.Color(color),
	}

	bot.Ctx.SendEmbeds(message.ChannelID, embed)

	return nil
}
