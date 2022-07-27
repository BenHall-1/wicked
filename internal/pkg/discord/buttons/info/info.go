package info

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Info struct {
	interfaces.Button
}

func (b Info) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Respond to interaction with a private embed
	embed := utils.MessageEmbed(models.Embed{
		Title: "Welcome to Tubbo's Pastel Café",
		Description: `This is a place where you can share your art with the world
						and see what other people have shared.`,
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
							Style:    discordgo.PrimaryButton,
							Label:    "Rules",
							CustomID: "rules_button",
						},
					},
				},
			},
		},
	})
}
