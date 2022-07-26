package cmds

import "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Description() string
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Commands = map[string]Command{
	"createbreakrolemenu": BreakRoleCommand{},
}
