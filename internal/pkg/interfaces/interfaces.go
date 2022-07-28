package interfaces

import "github.com/bwmarrin/discordgo"

type Button interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type Command interface {
	Name() string
	Description() string
	Options() []*discordgo.ApplicationCommandOption
	AppCommand() discordgo.ApplicationCommand
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type Event interface {
}

type Modal interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}
