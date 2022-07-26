package break_role

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	interfaces.Command
}

func (c Command) Name() string {
	return "createbreakrolemenu"
}

func (c Command) Description() string {
	return "Creates break role menu"
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Break Role Menu",
					Description: "Click the button below to get the On Break role",
					Color:       11648506,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.PrimaryButton,
							Label: "Get Break Role",
							Emoji: discordgo.ComponentEmoji{
								ID: "981140032658223124",
							},
							CustomID: "break_role_button",
						},
						discordgo.Button{
							Style: discordgo.DangerButton,
							Label: "Remove Break Role",
							Emoji: discordgo.ComponentEmoji{
								ID: "981140032658223124",
							},
							CustomID: "remove_break_role_button",
						},
					},
				},
			},
		},
	})
}
