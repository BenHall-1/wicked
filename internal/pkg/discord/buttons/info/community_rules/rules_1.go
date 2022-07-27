package community_rules

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CommunityRules1 struct {
	interfaces.Button
}

func (b CommunityRules1) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: "<:blurple_rules:1001796744822136852> Community Rules (1/2)",
		Description: `
			**No Hate Speech**
			Hate speech, homophobia, discrimination, and controversial topics will not be tolerated. This includes general toxic behavior and drama.

			**No Self Promotion**
			Self promotion of any kind is not allowed.

			**No Media in Art**
			Media is not allowed in <#931829057895677984>. Cosplays & Photography is fine.

			**No Pinging Creators**
			Do not ping Tubbo or any other creators you see in the server. Refrain from pinging any staff unless immediate help is required. If you need to report a message, react with <:report:810042563008135169>.
		`,
	})
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: []*discordgo.MessageEmbed{&embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Style:    discordgo.SuccessButton,
						Label:    "Art Rules",
						CustomID: "art_rules_1",
						Emoji: discordgo.ComponentEmoji{
							ID: "1001785025185988708",
						},
					},
					discordgo.Button{
						Style:    discordgo.SecondaryButton,
						Label:    "Previous Page",
						CustomID: "community_rules_1",
						Emoji: discordgo.ComponentEmoji{
							ID: "1001785027694178365",
						},
						Disabled: true,
					},
					discordgo.Button{
						Style:    discordgo.SecondaryButton,
						Label:    "Next Page",
						CustomID: "community_rules_2",
						Emoji: discordgo.ComponentEmoji{
							ID: "1001785026461061130",
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
	})
}
