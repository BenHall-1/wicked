package modals

import "github.com/bwmarrin/discordgo"

type Modal interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Modals = map[string]Modal{
	"break_role_modal": BreakRoleModal{},
}
