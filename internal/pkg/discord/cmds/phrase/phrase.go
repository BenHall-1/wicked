package phrase

import (
	"fmt"
	"strings"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	interfaces.Command
}

func (c Command) Name() string {
	return "menu"
}

func (c Command) Description() string {
	return "Adds options for menus"
}

func (c Command) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "add",
			Description: "Adds a phrase to the blocked phrases list",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "phrase",
					Description: "The phrase to block",
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "list",
			Description: "Lists all currently added phrases",
		},
	}
}

func (c Command) AppCommand() discordgo.ApplicationCommand {
	var defaultPermissions int64 = discordgo.PermissionManageServer
	return discordgo.ApplicationCommand{
		Name:                     c.Name(),
		Description:              c.Description(),
		DefaultMemberPermissions: &defaultPermissions,
		Options:                  c.Options(),
	}
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		options = i.ApplicationCommandData().Options
		cmd     = options[0].Name
		cmdopts = options[0].Options
	)

	switch cmd {
	case "add":
		var (
			phrase = cmdopts[0].Value.(string)
		)

		if p := models.GetPhrase(phrase, i.GuildID); p != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Phrase already exists",
				},
			})
		} else {
			regexPhrase := "/"
			chars := strings.Split(phrase, "")

			for i := range chars {
				regexPhrase += fmt.Sprintf("[%s]+", utils.ConvertChar(chars[i]))
			}

			regexPhrase += "/gm"

			models.AddPhrase(i.GuildID, i.Member.User.ID, phrase, regexPhrase)
		}

	}
}
