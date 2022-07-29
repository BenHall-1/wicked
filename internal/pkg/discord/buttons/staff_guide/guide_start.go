package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type GuideStart struct {
	FromMainMenu bool
	interfaces.Button
}

func (b GuideStart) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title:       fmt.Sprintf("%s Staff Guide", utils.Emoji["blurple_guide"]),
		Description: "You can use the buttons below to navigate the staff guide",
	})

	t := discordgo.InteractionResponseUpdateMessage

	if b.FromMainMenu {
		t = discordgo.InteractionResponseChannelMessageWithSource
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: t,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
			Flags:  1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style:    discordgo.DangerButton,
							Label:    "Raid Protocol",
							CustomID: "guide_raid",
							Emoji: discordgo.ComponentEmoji{
								Name: "⚠️",
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Contact Information",
							CustomID: "guide_contact",
							Emoji: discordgo.ComponentEmoji{
								Name: "📧",
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Commands",
							CustomID: "guide_commands",
							Emoji: discordgo.ComponentEmoji{
								Name: "💻",
							},
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Punishments (Community)",
							CustomID: "guide_community_punishments",
							Emoji: discordgo.ComponentEmoji{
								Name: "⚔️",
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Punishments (Art)",
							CustomID: "guide_art_punishments",
							Emoji: discordgo.ComponentEmoji{
								Name: "🎨",
							},
						},
					},
				},
			},
		},
	})
}
