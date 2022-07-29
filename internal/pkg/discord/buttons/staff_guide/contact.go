package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Contact struct {
	interfaces.Button
}

func (b Contact) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Contact Information", utils.Emoji["blurple_guide"]),
		Description: `
			When users request to contact Tubbo, you can send them the following contact information:

			- Business: tubbobusiness@gmail.com
			- Merch: support@shoptubbo.com
			- Tubnet: *N/A - Discord Link Coming Soon*
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
							Label:    "Main Menu",
							CustomID: "guide_start",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["guide"],
							},
						},
					},
				},
			},
		},
	})
}
