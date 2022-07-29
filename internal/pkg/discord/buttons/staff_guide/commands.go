package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Commands struct {
	interfaces.Button
}

func (b Commands) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Commands", utils.Emoji["blurple_guide"]),
		Description: `
			The following commands can be used for moderation:

			**Vortex**:
			- >>strike [amount] [user_id] [reason]
			*Allows you to add strikes to a user*
			- >>ban [user_id] [reason]
			*Bans a user*
			- >>check [user_id]
			*Checks a user's strikes*
			- >>userinfo [user_id]
			*Gets information about a user*
			- >>pardon [amount] [user_id] [reason]
			*Removes strikes from a user*
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
