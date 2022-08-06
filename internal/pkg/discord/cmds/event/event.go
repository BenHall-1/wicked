package event

import (
	"fmt"
	"os"
	"time"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	interfaces.Command
}

func (c Command) Name() string {
	return "event"
}

func (c Command) Description() string {
	return "Adds options for events"
}

func (c Command) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "create",
			Description: "Creates an event",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "date",
					Description: "The date of the event (DD-MM-YYYY)",
					Required:    true,
				}, {
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "time",
					Description: "The time of the event (HH:MM)",
					Required:    true,
				},
			},
		},
	}
}

func (c Command) AppCommand() discordgo.ApplicationCommand {
	var defaultPermissions int64 = discordgo.PermissionManageEvents
	return discordgo.ApplicationCommand{
		Name:                     c.Name(),
		Description:              c.Description(),
		DefaultMemberPermissions: &defaultPermissions,
		Options:                  c.Options(),
	}
}

func (c Command) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		options = i.ApplicationCommandData().Options
		cmd     = options[0].Name
		cmdopts = options[0].Options
		dopt    = cmdopts[0].Value
		topt    = cmdopts[1].Value
		pubchat = os.Getenv("DISCORD_PUBLIC_CHANNEL")
	)

	switch cmd {
	case "create":
		{
			startTime, err := time.Parse("02-01-2006 15:04", fmt.Sprintf("%s %s", dopt, topt))
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   1 << 6,
						Content: "Invalid Date/Time Format",
					},
				})
			}
			endTime := startTime.Add(time.Hour * 24)
			event, err := s.GuildScheduledEventCreate(i.GuildID, &discordgo.GuildScheduledEventParams{
				Name:               "Public TubChat!",
				ScheduledStartTime: &startTime,
				ScheduledEndTime:   &endTime,
				EntityType:         discordgo.GuildScheduledEventEntityTypeExternal,
				PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
				EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
					Location: fmt.Sprintf("<#%s>", pubchat),
				},
			})
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   1 << 6,
						Content: fmt.Sprintf("Unknown error whilst creating event: %v", err),
					},
				})
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: 1 << 6,
					Content: fmt.Sprintf("Created the new event! Use the below text to setup the sticky message: \n```\n %v \n```",
						fmt.Sprintf(`
						-sm -d 600s Want to be notified for the next public chat?
						[https://discord.com/events/%s/%s](Click here) to get notified!

						Why do some users have '#' and '##' in their name?
						Please see [https://discord.com/channels/472883504217063466/929655565364375602/1005430414384582737](this) message for more information!

						How can I contact a mod? I cannot see any other channels
						- Please DM any of the users with the <@&866451599270084639> or <@&934372109717737523> roles
						- The only other public channel is <#866444906142892033>
						`, i.GuildID, event.ID),
					),
				},
			})
		}
	}
}
