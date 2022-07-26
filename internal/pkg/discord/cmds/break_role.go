package cmds

import (
	"github.com/bwmarrin/discordgo"
)

type BreakRoleCommand struct {
	Command
}

func (c BreakRoleCommand) Name() string {
	return "createbreakrolemenu"
}

func (c BreakRoleCommand) Description() string {
	return "Creates break role menu"
}

func (c BreakRoleCommand) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
