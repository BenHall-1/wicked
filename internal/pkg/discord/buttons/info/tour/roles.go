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
		Title: fmt.Sprintf("%s Important Roles (2/3)", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			Roles identify key members of the community. If you want to find out more about what our main roles stand for and what they do exactly, keep on reading.

			%s <@&472898372571889676>
			This is Tubbo.
			
			%s <@&866451599270084639>
			The Barista team is here to manage the server. These are our Administrators.

			%s <@&934372109717737523>
			The Janitor team is here to keep the community a safer and comfortable space at any given time. These are our Moderators. 

			%s <@&872924363807158353>
			The Latte Artists are here to keep an eye on the art channels and make sure that art is not stolen. 
		`, utils.Emoji["badge_owner"], utils.Emoji["badge_admin"], utils.Emoji["badge_mod"], utils.Emoji["badge_artist"]),
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
