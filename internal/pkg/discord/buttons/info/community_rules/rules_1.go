package community_rules

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CommunityRules1 struct {
	FromMainMenu bool
	interfaces.Button
}

func (b CommunityRules1) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Community Rules (1/2)", utils.Emoji["blurple_rules"]),
		Description: `
			**No Hate Speech**
			Hate speech, homophobia, discrimination, and controversial topics will not be tolerated. This includes general toxic behavior and drama.

			**No Self Promotion**
			Self promotion of any kind is not allowed.

			**No Media in Art**
			Media is not allowed in <#866444906142892033>. Cosplays & Photography is fine.

			**No Pinging Creators**
			Do not ping Tubbo or any other creators you see in the server. Refrain from pinging any staff unless immediate help is required. If you need to report a message, react with <:report:868243118808530964>.
		`,
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
							Disabled: true,
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "community_rules_2",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["next_step"],
							},
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
