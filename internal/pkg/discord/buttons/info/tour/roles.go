package tour

import (
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
		Title: "<:blurple_badge:1001796743068930129> Important Roles (1/3)",
		Description: `
			Roles identify key members of the community. If you want to find out more about what our main roles stand for and what they do exactly, keep on reading.

			<:badge_owner:1001800446580629565> <@&472898372571889676>

			<:badge_admin:1001798283724853298> <@&866451599270084639>
			The Barista team is here to manage the server. These are our Administrators.

			<:badge_mod:1001798282403663902> <@&934372109717737523>
			The Janitor team is here to keep the community a safer and comfortable space at any given time. These are our Moderators. 

			<:badge_artist:1001800233845526618> <@&872924363807158353>
			The Latte Artists are here to keep an eye on the art channels and make sure that art is not stolen. 
		`,
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
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
								ID: "1001785025185988708",
							},
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Previous Page",
							CustomID: "art_rules_1",
							Emoji: discordgo.ComponentEmoji{
								ID: "1001785027694178365",
							},
							Disabled: true,
						},
						discordgo.Button{
							Style:    discordgo.SecondaryButton,
							Label:    "Next Page",
							CustomID: "art_rules_2",
							Emoji: discordgo.ComponentEmoji{
								ID: "1001785026461061130",
							},
						},
					},
				},
			},
		},
	})
}
