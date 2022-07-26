package buttons

import "github.com/bwmarrin/discordgo"

type Button interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Buttons = map[string]Button{
	"add_break_role_button":    AddBreakRoleButton{},
	"remove_break_role_button": RemoveBreakRoleButton{},
}
