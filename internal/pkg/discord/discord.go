package discord

import (
	"log"
	"os"

	"github.com/benhall-1/wicked/internal/pkg/discord/cmds"
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
	var (
		staffserv = os.Getenv("DISCORD_STAFF_SERVER")
	)
	for i := range cmds.Commands {
		c := cmds.Commands[i].Cmd()
		_, err := db.session.ApplicationCommandCreate(db.session.State.User.ID, staffserv, c)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", c.Name, err)
		}
	}
}

// func (db DiscordBot) RegisterCommands() {
// 	var (
// 		dperms = false
// 		cmds   = []*discordgo.ApplicationCommand{
// 			{
// 				Name:              "createbreakrolemenu",
// 				Description:       "Creates break role menu",
// 				DefaultPermission: &dperms,
// 			},
// 		}
// 		staffserv = os.Getenv("DISCORD_STAFF_SERVER")
// 	)

// 	registeredCommands := make([]*discordgo.ApplicationCommand, len(cmds))
// 	for i := range cmds {
// 		cmd, err := db.session.ApplicationCommandCreate(db.session.State.User.ID, staffserv, cmds[i])
// 		if err != nil {
// 			log.Panicf("Cannot create '%v' command: %v", cmds[i].Name, err)
// 		}
// 		registeredCommands[i] = cmd
// 	}
// }

func NewDiscordBot(token string, session *discordgo.Session, db *sqlx.DB) DiscordBot {
	return DiscordBot{token, session, db}
}
