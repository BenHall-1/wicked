package break_role

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/bwmarrin/discordgo"
)

type Add struct {
	interfaces.Button
}

func (b Add) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "break_role_modal",
			Title:    "Get The Break Role",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "reason",
							Label:       "Why do you need the break role?",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Just a brief reason is fine",
							Required:    true,
							MaxLength:   2000,
							MinLength:   10,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:  "date_end",
							Label:     "When will you return from break?",
							Style:     discordgo.TextInputShort,
							Required:  true,
							MaxLength: 300,
						},
					},
				},
			},
		},
	})
}
