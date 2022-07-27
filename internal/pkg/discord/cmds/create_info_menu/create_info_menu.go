package create_info_menu

import (
	"fmt"
	"os"

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
	var (
		infid = os.Getenv("DISCORD_INFO_CHANNEL")
	)
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Tubbo's Pastel Café", utils.Emoji["blurple_guide"]),
		Description: `This is a place where you can share your art with the world
						and see what other people have shared.`,
	})
	s.ChannelMessageSendComplex(infid, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{&embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Style:    discordgo.PrimaryButton,
						Label:    "Community Rules",
						CustomID: "community_rules_button",
						Emoji: discordgo.ComponentEmoji{
							ID: utils.EmojiIds["rules"],
						},
					},
					discordgo.Button{
						Style:    discordgo.PrimaryButton,
						Label:    "Art Rules",
						CustomID: "art_rules_button",
						Emoji: discordgo.ComponentEmoji{
							ID: utils.EmojiIds["rules"],
						},
					},
					discordgo.Button{
						Style:    discordgo.SuccessButton,
						Label:    "Get Started",
						CustomID: "get_started",
						Emoji: discordgo.ComponentEmoji{
							ID: utils.EmojiIds["guide"],
						},
					},
				},
			},
		},
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: fmt.Sprintf("Posted the information menu in <#%s>", infid),
		},
	})
}
