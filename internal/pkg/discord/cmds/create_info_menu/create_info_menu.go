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
	return "createinfomenu"
}

func (c Command) Description() string {
	return "Creates the main info menu"
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		infid = os.Getenv("DISCORD_INFO_CHANNEL")
	)
	embed := utils.MessageEmbed(models.Embed{
		Title: fmt.Sprintf("%s Tubbo's Pastel Café", utils.Emoji["blurple_guide"]),
		Description: fmt.Sprintf(`
			Welcome to Tubbo's Discord! Please use the buttons below to navigate information about the server.

			**Socials**:
			%[1]s Instagram [@TubboLive](https://instagram.com/TubboLive)
			%[2]s Merch [ShopTubbo.com](https://shoptubbo.com)
			%[3]s Reddit [@Tubbo](https://www.reddit.com/r/Tubbo)
			%[4]s Spotify [Tubbo](https://open.spotify.com/user/p04neko25tv0jw7p24si4pcis)
			%[4]s Spotify (Artist) [Tubbo](https://open.spotify.com/artist/4B1kkhDbhSyiS3VDgbKV2T)
			%[6]s Twitch [@Tubbo](https://twitch.tv/Tubbo)
			%[6]s Twitch Alt [@TubboLIVE](https://twitch.tv/TubboLIVE)
			%[5]s Twitter [@TubboLive](https://twitter.com/TubboLive)
			%[5]s Twitter Alt [@TubboTWO](https://twitter.com/TubboTWO)
			%[7]s YouTube [@Tubbo](https://www.youtube.com/Tubbo)

			**Contact**:
			- Email (Business) [tubbobusiness@gmail.com](mailto:tubbobusiness@gmail.com)
			- Email (Merch) [support@shoptubbo.com](mailto:support@shoptubbo.com)
		`, utils.Emoji["blurple_instagram"], utils.Emoji["blurple_merch"], utils.Emoji["blurple_reddit"], utils.Emoji["blurple_spotify"], utils.Emoji["blurple_twitch"], utils.Emoji["blurple_twitter"], utils.Emoji["blurple_youtube"]),
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
						CustomID: "get_started_button",
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
