package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (db DiscordBot) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	db.RegisterCommands()
}
