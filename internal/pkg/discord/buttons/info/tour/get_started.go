package tour

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type GetStarted struct {
	FromMainMenu bool
	interfaces.Button
}

func (b GetStarted) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Getting Started in Tubbo's Pastel Café (1/2)", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			As a new member, you might come across a few things that may confuse you. To help you with getting started in the community, we created this "Getting Started" guide to walk you trough everything you have to know, step by step. 
			Any questions? Feel free to contact one of our team members!

			%s **Subscribing**
			Subscribing is a way of gaining access to the server. This will grant you the <@&1078013927931269291> role.
			You can subscribe by [Clicking Here](https://www.twitch.tv/subs/tubbo).

			%s **Boosting**
			Boosting is one way to get access to the rest of the server. This will grant you the <@&620374536625324042> role.
		`, utils.Emoji["badge_sub"], utils.Emoji["blurple_boosts"]),
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
