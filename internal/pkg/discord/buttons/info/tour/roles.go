package tour

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Roles struct {
	interfaces.Button
}

func (b Roles) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Important Roles (2/2)", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			Roles identify key members of the community. If you want to find out more about what our main roles stand for and what they do exactly, keep on reading.

			%s <@&472898372571889676>
			This is Tubbo.
			
			%s <@&866451599270084639>
			The Staff team is here to manage the server. This is our Moderation Team.

			%s <@&934372109717737523>
			The Trial Staff team is members of the team that have been newly added to the team. 
		`, utils.Emoji["badge_owner"], utils.Emoji["badge_admin"], utils.Emoji["badge_mod"]),
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
							Style:    discordgo.SecondaryButton,
							Label:    "Previous Page",
							CustomID: "get_started",
							Emoji: discordgo.ComponentEmoji{
								ID: utils.EmojiIds["previous_step"],
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "channels",
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
