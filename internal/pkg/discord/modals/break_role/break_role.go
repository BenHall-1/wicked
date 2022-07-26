package break_role

import (
	"fmt"
	"os"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/bwmarrin/discordgo"
)

type Modal struct {
	interfaces.Modal
}

func (b Modal) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		data = i.ModalSubmitData()
		ssid = os.Getenv("DISCORD_STAFF_SERVER")
		rid  = os.Getenv("DISCORD_BREAK_ROLE")
		bid  = os.Getenv("DISCORD_BREAK_LOG_CHANNEL_ID")
	)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Thank you for submitting this, You should now have the break role",
			Flags:   1 << 6,
		},
	})

	embed := &discordgo.MessageEmbed{
		Title: "Break Role Request",
		Color: 11648506,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "User",
				Value:  fmt.Sprintf("<@%s>", i.Interaction.Member.User.ID),
				Inline: true,
			},
			{
				Name:   "Reason",
				Value:  data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
				Inline: true,
			},
			{
				Name:   "Date End",
				Value:  data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
				Inline: true,
			},
		},
	}
	s.GuildMemberRoleAdd(ssid, i.Interaction.Member.User.ID, rid)

	s.ChannelMessageSendComplex(bid, &discordgo.MessageSend{
		Embeds:  []*discordgo.MessageEmbed{embed},
		Content: fmt.Sprintf("@here, A new break role request has been made by <@%s>", i.Interaction.Member.User.ID),
	})
}
