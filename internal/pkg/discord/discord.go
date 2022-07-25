package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"
)

type DiscordBot struct {
	token    string
	session  *discordgo.Session
	database *sqlx.DB
}

func (db DiscordBot) SetupHandlers() {
	db.session.AddHandler(db.OnMessageCreate)
	db.session.AddHandler(db.InteractionCreate)
	db.session.AddHandler(db.OnReady)
}

func (db DiscordBot) RegisterCommands() {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "createbreakrolemenu",
			Description: "Creates break role menu",
		},
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := db.session.ApplicationCommandCreate(db.session.State.User.ID, os.Getenv("DISCORD_STAFF_SERVER"), v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func NewDiscordBot(token string, session *discordgo.Session, db *sqlx.DB) DiscordBot {
	return DiscordBot{token, session, db}
}
