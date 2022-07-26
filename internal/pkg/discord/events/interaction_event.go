package events

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds"
	"github.com/benhall-1/wicked/internal/pkg/discord/modals"
	"github.com/bwmarrin/discordgo"
)

type InteractionEvent struct {
	Event
}

func (e InteractionEvent) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		cmds = cmds.Commands
		btns = buttons.Buttons
		mdls = modals.Modals
	)

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := cmds[i.ApplicationCommandData().Name]; ok {
			h.Handle(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := btns[i.MessageComponentData().CustomID]; ok {
			h.Handle(s, i)
		}
	case discordgo.InteractionModalSubmit:
		if h, ok := mdls[i.ModalSubmitData().CustomID]; ok {
			h.Handle(s, i)
		}
	}
}
