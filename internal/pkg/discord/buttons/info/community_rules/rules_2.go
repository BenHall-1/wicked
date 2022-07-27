package community_rules

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CommunityRules2 struct {
	interfaces.Button
}

func (b CommunityRules2) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Community Rules (2/2)", utils.Emoji["blurple_rules"]),
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
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
			Flags:  1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style:    discordgo.SuccessButton,
							Label:    "Art Rules",
							CustomID: "art_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["rules"],
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Previous Page",
							CustomID: "community_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["previous_step"],
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "community_rules_2",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["next_step"],
							},
							Disabled: true,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
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
