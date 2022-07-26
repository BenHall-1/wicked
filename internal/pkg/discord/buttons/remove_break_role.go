package buttons

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type RemoveBreakRoleButton struct {
	Button
}

func (b RemoveBreakRoleButton) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		sid = os.Getenv("DISCORD_STAFF_SERVER")
		rid = os.Getenv("DISCORD_BREAK_ROLE")
		msg = "Your break role has been removed"
	)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
			Flags:   1 << 6,
		},
	})
	s.GuildMemberRoleRemove(sid, i.Interaction.Member.User.ID, rid)
}
