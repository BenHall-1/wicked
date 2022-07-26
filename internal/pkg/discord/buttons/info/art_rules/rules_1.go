package art_rules

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type ArtRules1 struct {
	FromMainMenu bool
	interfaces.Button
}

func (b ArtRules1) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Art Rules (1/2)", utils.Emoji["blurple_rules"]),
		Description: `
			**No Stealing Art**
			Theft of any art is not allowed in our art channel. This includes stealing a full image (or partial) and claiming it as your own, tracing over an image and claiming it as your own, and using any type of base without credit to the original creator. These are non-negotiable. 

			**Usage of Apps**
			Apps such as piccrew and gacha are not allowed as it's not drawn yourself. 

			**No Joke/Meme Posts**
			Joke posts (“shitposts”) are also not allowed in our art channel. This can range from images that are memes, images making fun of others, etc. Staff have the final say on shitpost images. 
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
							Disabled: true,
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "art_rules_2",
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
