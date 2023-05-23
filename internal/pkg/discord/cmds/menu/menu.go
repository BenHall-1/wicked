package menu

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	interfaces.Command
}

func (c Command) Name() string {
	return "menu"
}

func (c Command) Description() string {
	return "Adds options for menus"
}

func (c Command) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "create",
			Description: "Creates a menu",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "The name of the menu",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Information Panel",
							Value: "info",
						},
						{
							Name:  "Staff Guide",
							Value: "staff_guide",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "The channel for the menu",
					Required:    true,
				},
			},
		},
	}
}

func (c Command) AppCommand() discordgo.ApplicationCommand {
	var defaultPermissions int64 = discordgo.PermissionManageServer
	return discordgo.ApplicationCommand{
		Name:                     c.Name(),
		Description:              c.Description(),
		DefaultMemberPermissions: &defaultPermissions,
		Options:                  c.Options(),
	}
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		options  = i.ApplicationCommandData().Options
		cmd      = options[0].Name
		cmdopts  = options[0].Options
		menu     = cmdopts[0].Value
		channels = i.ApplicationCommandData().Resolved.Channels
		channel  *discordgo.Channel
	)

	for i := range channels {
		channel = channels[i]
		break
	}
	switch cmd {
	case "create":
		{
			var embed discordgo.MessageEmbed
			var components []discordgo.MessageComponent
			switch menu {
			case "info":
				{
					embed = utils.MessageEmbed(models.Embed{
						Title: fmt.Sprintf("%s Tubbo's Discord", utils.Emoji["blurple_guide"]),
						Description: fmt.Sprintf(`
							Welcome to Tubbo's Discord! Please use the buttons below to navigate information about the server.
				
							**Socials**:
							%[1]s Instagram [@TubboLive](https://instagram.com/TubboLive)
							%[2]s Reddit [@Tubbo_](https://www.reddit.com/r/Tubbo_)
							%[3]s Spotify [Tubbo](https://open.spotify.com/user/p04neko25tv0jw7p24si4pcis)
							%[3]s Spotify (Artist) [Tubbo](https://open.spotify.com/artist/4B1kkhDbhSyiS3VDgbKV2T)
							%[5]s Twitch [@Tubbo](https://twitch.tv/Tubbo)
							%[5]s Twitch Alt [@TubboLIVE](https://twitch.tv/TubboLIVE)
							%[4]s Twitter [@TubboLive](https://twitter.com/TubboLive)
							%[4]s Twitter Alt [@TubboTWO](https://twitter.com/TubboTWO)
							%[6]s YouTube [@Tubbo](https://www.youtube.com/Tubbo)
							
							TubNet: [discord.gg/tubnet](https://discord.gg/tubnet)
				
							**Contact**:
							- Email (Business) __tubbobusiness@gmail.com__
						`, utils.Emoji["blurple_instagram"], utils.Emoji["blurple_reddit"], utils.Emoji["blurple_spotify"], utils.Emoji["blurple_twitch"], utils.Emoji["blurple_twitter"], utils.Emoji["blurple_youtube"]),
					})
					components = []discordgo.MessageComponent{
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
					}
				}
			case "staff_guide":
				{
					embed = utils.MessageEmbed(models.Embed{
						Title:       fmt.Sprintf("%s Staff Guide", utils.Emoji["blurple_guide"]),
						Description: "Welcome to the Server Staff Guide. Please use the buttons below to navigate your way through the guide.",
					})
					components = []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Style:    discordgo.SuccessButton,
									Label:    "Get Started",
									CustomID: "guide_start_button",
									Emoji: discordgo.ComponentEmoji{
										ID: utils.EmojiIds["guide"],
									},
								},
							},
						},
					}
				}
			}
			s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
				Embeds:     []*discordgo.MessageEmbed{&embed},
				Components: components,
			})
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: fmt.Sprintf("Created the menu in %s", channel.Mention()),
				},
			})
		}
	}
}
