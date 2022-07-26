package discord

import (
	"log"
	"os"

	"github.com/benhall-1/wicked/internal/pkg/discord/cmds"
	"github.com/benhall-1/wicked/internal/pkg/discord/events"
	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"
)

type DiscordBot struct {
	token    string
	session  *discordgo.Session
	database *sqlx.DB
}

func (db DiscordBot) SetupHandlers() {
	db.session.AddHandler(events.MessageEvent{}.Handle)
	db.session.AddHandler(events.InteractionEvent{}.Handle)
	db.session.AddHandler(db.OnReady)
}

func (db DiscordBot) RegisterCommands() {
	var (
		staffserv = os.Getenv("DISCORD_STAFF_SERVER")
	)
	for i := range cmds.Commands {
		c := &discordgo.ApplicationCommand{
			Name:        cmds.Commands[i].Name(),
			Description: cmds.Commands[i].Description(),
		}
		_, err := db.session.ApplicationCommandCreate(db.session.State.User.ID, staffserv, c)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", c.Name, err)
		}
	}
}

func NewDiscordBot(token string, session *discordgo.Session, db *sqlx.DB) DiscordBot {
	return DiscordBot{token, session, db}
}
