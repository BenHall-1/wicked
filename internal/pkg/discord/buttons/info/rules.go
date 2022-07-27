package info

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

type Rules struct {
	interfaces.Button
}

// func (b Rules) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Embeds: []*discordgo.MessageEmbed{&mainRulesEmbed, &artRulesEmbed},
// 			Flags:  1 << 6,
// 			Components: []discordgo.MessageComponent{
// 				discordgo.ActionsRow{
// 					Components: []discordgo.MessageComponent{
// 						discordgo.Button{
// 							Style: discordgo.LinkButton,
// 							Label: "Discord Terms of Service",
// 							URL:   "https://dis.gd/tos",
// 						},
// 						discordgo.Button{
// 							Style: discordgo.LinkButton,
// 							Label: "Discord Community Guidelines",
// 							URL:   "https://dis.gd/guidelines",
// 						},
// 						discordgo.Button{
// 							Style: discordgo.LinkButton,
// 							Label: "Ban Appeals",
// 							URL:   "https://forms.gle/epaLCAiUG2WAwRQUA",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	})
// }
