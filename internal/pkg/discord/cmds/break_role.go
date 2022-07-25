package cmds

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type BreakRoleCommand struct {
	Command
}

func (c BreakRoleCommand) Cmd() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "createbreakrolemenu",
		Description: "Creates break role menu",
	}
}

func (c BreakRoleCommand) HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
							CustomID: "createbreakrolemenu-break_role_button",
						},
						discordgo.Button{
							Style: discordgo.DangerButton,
							Label: "Remove Break Role",
							Emoji: discordgo.ComponentEmoji{
								ID: "981140032658223124",
							},
							CustomID: "createbreakrolemenu-remove_break_role_button",
						},
					},
				},
			},
		},
	})
}

func (c BreakRoleCommand) HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"createbreakrolemenu-break_role_button": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: &discordgo.InteractionResponseData{
					CustomID: "createbreakrolemenu-break_role_modal",
					Title:    "Get The Break Role",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:    "reason",
									Label:       "Why do you need the break role?",
									Style:       discordgo.TextInputParagraph,
									Placeholder: "Just a brief reason is fine",
									Required:    true,
									MaxLength:   2000,
									MinLength:   10,
								},
							},
						},
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:  "date_end",
									Label:     "When will you return from break?",
									Style:     discordgo.TextInputShort,
									Required:  true,
									MaxLength: 300,
								},
							},
						},
					},
				},
			})
		},
		"createbreakrolemenu-remove_break_role_button": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Your break role has been removed",
					Flags:   1 << 6,
				},
			})
			s.GuildMemberRoleRemove(os.Getenv("DISCORD_STAFF_SERVER"), i.Interaction.Member.User.ID, os.Getenv("DISCORD_BREAK_ROLE"))
		},
	}
}

func (c BreakRoleCommand) HandleModals(s *discordgo.Session, i *discordgo.InteractionCreate) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"createbreakrolemenu-break_role_modal": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Thank you for submitting this, You should now have the break role",
					Flags:   1 << 6,
				},
			})

			data := i.ModalSubmitData()

			embed := &discordgo.MessageEmbed{
				Title: "Break Role Request",
				Color: 11648506,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "User",
						Value:  fmt.Sprintf("<@%s>", i.Interaction.Member.User.ID),
						Inline: true,
					},
					{
						Name:   "Reason",
						Value:  data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
						Inline: true,
					},
					{
						Name:   "Date End",
						Value:  data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
						Inline: true,
					},
				},
			}
			s.GuildMemberRoleAdd(os.Getenv("DISCORD_STAFF_SERVER"), i.Interaction.Member.User.ID, os.Getenv("DISCORD_BREAK_ROLE"))

			s.ChannelMessageSendComplex(os.Getenv("DISCORD_BREAK_LOG_CHANNEL_ID"), &discordgo.MessageSend{
				Embeds:  []*discordgo.MessageEmbed{embed},
				Content: fmt.Sprintf("@here, A new break role request has been made by <@%s>", i.Interaction.Member.User.ID),
			})

		},
	}
}
