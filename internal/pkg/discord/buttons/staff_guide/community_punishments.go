package staff_guide

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type CommunityPunishments struct {
	interfaces.Button
}

func (b CommunityPunishments) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Community Punishments", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			The following community punishments apply:

			**Verbal / 1 Strike**:
			- Non English
			- Commands in chat
			- Minor Toxicity

			**3 Strikes**:
			- Copypasta
			- Toxicity (After 1 Strike)
			
			**6 Strikes**:
			- Soundboards / Mic Abuse
			- Asking for Nitro / Money
			- Self Promo
			- Moderate Toxicity
			- Slight NSFW

			**9 Strikes**:
			- Heavy Mic Abuse
			- Staff Disrespect
			- Moderate NSFW
			
			**Softban**:
			- Impersonation of a creator
			
			**Ban**:
			- Slurs in Chat/VC
			- Discord Crashing Media
			- DM Advertising
			- Underage
			- Better Discord / Client Alterations
			- Screensharing Movies / Copyrighted Content
			- Scam Links (Provide the following message)

			%s
		`, "``` \nScam links. If you recover your account, please change your password immediately! The ban can be appealed here: https://forms.gle/epaLCAiUG2WAwRQUA\n```"),
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
							Label:    "Art Punishments",
							CustomID: "guide_art_punishments",
							Emoji: discordgo.ComponentEmoji{
								Name: "🎨",
							},
						},
					},
				},
			},
		},
	})
}
