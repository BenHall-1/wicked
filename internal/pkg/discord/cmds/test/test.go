package test

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	interfaces.Command
}

func (c Command) Name() string {
	return "test"
}

func (c Command) Description() string {
	return "Test command for Ben"
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := utils.MessageEmbed(models.Embed{
		Title: "<:blurple_guide:1001796740787220602> Tubbo's Pastel Café",
		Description: `This is a place where you can share your art with the world
						and see what other people have shared.`,
	})
	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
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
						Style:    discordgo.SuccessButton,
						Label:    "Art Rules",
						CustomID: "art_rules_1",
						Emoji: discordgo.ComponentEmoji{
							ID: "1001785025185988708",
						},
					},
				},
			},
		},
	})
}
