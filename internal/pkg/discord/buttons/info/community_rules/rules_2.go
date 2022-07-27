package community_rules

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CommunityRules2 struct {
	interfaces.Button
}

func (b CommunityRules2) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Respond to interaction with a private embed
	mainRulesEmbed := utils.MessageEmbed(models.Embed{
		Title: "Community Rules",
		Description: `
			**No NSFW**
			No NSFW or profane content. This includes heavy gore, sensitive topics, etc.

			**No Impersonation**
			Do not impersonate content creators, staff, or even other members.

			**No Spamming**
			Spamming by sending repetitive messages, emojis, images, mentions or reactions is not tolerated in this server or in DMs.

			**English Only**
			We require all chats to be english only for our team to be able to moderate the server safely.
		`,
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&mainRulesEmbed},
			Flags:  1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style:    discordgo.PrimaryButton,
							Label:    "Art Rules",
							CustomID: "art_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: "1001785025185988708",
							},
						},
						discordgo.Button{
							Style:    discordgo.PrimaryButton,
							Label:    "Previous Page",
							CustomID: "community_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: "1001785027694178365",
							},
						},
						discordgo.Button{
							Style:    discordgo.PrimaryButton,
							Label:    "Next Page",
							CustomID: "community_rules_2",
							Emoji: discordgo.ComponentEmoji{
								ID: "1001785026461061130",
							},
							Disabled: true,
						},
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Discord Terms of Service",
							URL:   "https://dis.gd/tos",
						},
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Discord Community Guidelines",
							URL:   "https://dis.gd/guidelines",
						},
					},
				},
			},
		},
	})
}
