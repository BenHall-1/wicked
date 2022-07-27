package tour

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type GetStarted struct {
	interfaces.Button
}

func (b GetStarted) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Getting Started in Tubbo's Pastel Café (1/3)", utils.Emoji["blurple_guide"]),
		Description: `
			As a new member, you might come across a few things that may confuse you. To help you with getting started in the community, we created this "Getting Started" guide to walk you trough everything you have to know, step by step. 
			Any questions? Feel free to contact one of our team!

			<:badge_sub:1001803508590321664> **Subscribing**
			Subscribing is a way of gaining access to the server. This will grant you the <@&790351230865506325> role.
			You can subscribe by [Clicking Here](https://www.twitch.tv/products/orphictubbo).

			<a:blurple_boosts:1001802001304272917> **Boosting**
			Boosting is one way to get access to the rest of the server. This will grant you the <@&620374536625324042> role.
		`,
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
			Flags:  1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Previous Page",
							CustomID: "fake",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["previous_step"],
							},
							Disabled: true,
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "roles",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["next_step"],
							},
						},
					},
				},
			},
		},
	})
}
