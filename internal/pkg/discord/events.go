package discord

import (
	"strings"

	"github.com/benhall-1/wicked/internal/pkg/discord/cmds"
	"github.com/bwmarrin/discordgo"
)

func (db DiscordBot) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	db.RegisterCommands()
}

func (db DiscordBot) InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmds := cmds.Commands

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := cmds[i.ApplicationCommandData().Name]; ok {
			h.HandleCommand(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := cmds[strings.Split(i.MessageComponentData().CustomID, "-")[0]]; ok {
			h.HandleComponents(s, i)
		}
	case discordgo.InteractionModalSubmit:
		if h, ok := cmds[strings.Split(i.ModalSubmitData().CustomID, "-")[0]]; ok {
			h.HandleModals(s, i)
		}
	}
}
