package info

import (
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type Rules struct {
	interfaces.Button
}

func (b Rules) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Respond to interaction with a private embed
	mainRulesEmbed := utils.MessageEmbed(models.Embed{
		Title: "Community Rules",
		Description: `
			**No Hate Speech**
			Hate speech, homophobia, discrimination, and controversial topics will not be tolerated. This includes general toxic behavior and drama.

			**No Self Promotion**
			Self promotion of any kind is not allowed.

			**No Media in Art**
			Media is not allowed in <#931829057895677984>. Cosplays & Photography is fine.

			**No Pinging Creators**
			Do not ping Tubbo or any other creators you see in the server. Refrain from pinging any staff unless immediate help is required. If you need to report a message, react with <:report:810042563008135169>.

			**No NSFW**
			No NSFW or profane content. This includes heavy gore, sensitive topics, etc.

			**No Impersonation**
			Do not impersonate content creators, staff, or even other members.

			**No Spamming**
			Spamming by sending repetitive messages, emojis, images, mentions or reactions is not tolerated in this server or in DMs.

			**English Only**
			We require all chats to be english only for our team to be able to moderate the server safely.
		`,
	})
	artRulesEmbed := utils.MessageEmbed(models.Embed{
		Title: "Community Rules",
		Description: `
			**No Stealing Art**
			Theft of any art is not allowed in our art channel. This includes stealing a full image (or partial) and claiming it as your own, tracing over an image and claiming it as your own, and using any type of base without credit to the original creator. These are non-negotiable. 

			**Usage of Apps**
			Apps such as piccrew and gacha are not allowed as it's not drawn yourself. 

			**No Joke/Meme Posts**
			Joke posts (“shitposts”) are also not allowed in our art channel. This can range from images that are memes, images making fun of others, etc. Staff have the final say on shitpost images. 

			**No Offensive or negative/inappropriate emotes**
			We do not allow the use of offensive emotes (emotes making fun of others’ art, emotes saying racial slurs, etc.) in our art channel. This also pertains to any inappropriate emotes (emotes that are of pornographic body parts, emotes depicting murder, etc.).
			Finally, we do not allow any negative emotes on any art piece (including but not limited to: ❌, 🤮, or any emotes spelling out negative words). 

			**Clarifications**
			Posting art and saying it is not yours is not considered stealing. Please credit the original artist when posting artwork that isn’t yours (as well as clarifying said art is not yours).

			**Shipping Content Creators & Misc**
			Shipping creators, members, etc., is not allowed under any circumstance. Please follow all our rules in our server when posting artwork. Shipping, gore, etc., are not allowed under any circumstance. 
			Excessive lights will need a CW and a spoiler over the image. 
		`,
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&mainRulesEmbed, &artRulesEmbed},
			Flags:  1 << 6,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Discord Terms of Service",
							URL:   "https://dis.gd/tos",
						},
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Discord Community Guidelines",
							URL:   "https://dis.gd/guidelines",
						},
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Ban Appeals",
							URL:   "https://forms.gle/epaLCAiUG2WAwRQUA",
						},
					},
				},
			},
		},
	})
}
