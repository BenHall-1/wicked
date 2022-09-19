package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type ArtPunishments struct {
	interfaces.Button
}

func (b ArtPunishments) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Art Punishments", utils.Emoji["blurple_guide"]),
		Description: `
			The following art punishments apply:

			**1 Strike**:
			- Gacha
			- PicCrew

			**3 Strikes**:
			- Gacha (Second Offence)
			- PicCrew (Second Offence)
			- Media
			- Edits (They do not count as original art)
			
			**6 Strikes**:
			- Tracing Art
			- Bases (With No Credit)

			**9 Strikes**:
			- Reposts (Other people's art)
			
			**Ban**:
			- Art Theft
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
						discordgo.Button{
							Style:    discordgo.SuccessButton,
							Label:    "Community Punishments",
							CustomID: "guide_community_punishments",
						},
					},
				},
			},
		},
	})
}
