package cmds

import "github.com/bwmarrin/discordgo"

type Command interface {
	Cmd() *discordgo.ApplicationCommand
	HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleModals(s *discordgo.Session, i *discordgo.InteractionCreate) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Commands = map[string]Command{
	"createbreakrolemenu": BreakRoleCommand{},
}
