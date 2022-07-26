package art_rules

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type ArtRules2 struct {
	interfaces.Button
}

func (b ArtRules2) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Art Rules (2/2)", utils.Emoji["blurple_rules"]),
		Description: `
			**No Offensive or negative/inappropriate emotes**
			We do not allow the use of offensive emotes (emotes making fun of others’ art, emotes saying racial slurs, etc.) in our art channel. This also pertains to any inappropriate emotes (emotes that are of pornographic body parts, emotes depicting murder, etc.).
			Finally, we do not allow any negative emotes on any art piece (including but not limited to: ❌, 🤮, or any emotes spelling out negative words). 

			**Clarifications**
			Posting art and saying it is not yours is not considered stealing. Please credit the original artist when posting artwork that isn’t yours (as well as clarifying said art is not yours).

			**Shipping Content Creators & Misc**
			Shipping creators, members, etc., is not allowed under any circumstance. Please follow all our rules in our server when posting artwork. Shipping, gore, etc., are not allowed under any circumstance. 
			Excessive lights will need a CW and a spoiler over the image. 
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
							Label:    "Community Rules",
							CustomID: "community_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["rules"],
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Previous Page",
							CustomID: "art_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["previous_step"],
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "art_rules_2",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["next_step"],
							},
							Disabled: true,
						},
					},
				},
			},
		},
	})
}
