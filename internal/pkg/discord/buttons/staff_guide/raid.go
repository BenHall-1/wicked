package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Raid struct {
	interfaces.Button
}

func (b Raid) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Raid Protocol", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			In the event that the server is raided, you should do the following steps:
			- Lockdown the server using %s command.
			*If this failed, use %s on every public channel.
			- Enable Anti-Raid mode via %s for **at least 30 minutes**.
			- Delete any messages that were not auto-deleted by the bots.
			- Collect IDs from the logs channel and start list of IDs to be banned.
			- Ban all of the IDs using the massban command %s.
			- Speak with the team and agree a time to re-open the server.
			- Re-open the server
			- Report the list of IDs to discord via https://dis.gd/report -> Report Spam -> Spam from a user or bot
		`, "`?lockdown`", "`?lock #channel`", "`>>antiraid on`", "`>>ban [id1] [id2] etc`"),
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
