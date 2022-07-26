package events

import (
	"fmt"
	"os"
	"strings"

	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type MessageEvent struct {
	Event
}

func (e MessageEvent) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		logchan = os.Getenv("DISCORD_LOGS_CHANNEL")
		artchan = os.Getenv("DISCORD_ART_CHANNEL")
		modrole = os.Getenv("DISCORD_MOD_ROLE")
		server  = os.Getenv("DISCORD_MAIN_SERVER")
	)

	if m.ChannelID == artchan {
		for i := range m.Attachments {
			encodedValue := utils.GetBase64FromUrl(m.Attachments[i].URL)
			uploads := models.GetAllImagesUploadedByHash(encodedValue)
			if len(uploads) > 0 {
				uploadList := make([]string, 0, len(uploads))

				for i := range uploads {
					var (
						url  = fmt.Sprintf("https://discord.com/channels/%s/%s/%s", server, artchan, uploads[i].MessageId)
						item = fmt.Sprintf("- `%s` ([Message Link](%s))", uploads[i].UserId, url)
					)

					uploadList = append(uploadList, item)
				}

				s.ChannelMessageSendComplex(logchan, &discordgo.MessageSend{
					Embeds: []*discordgo.MessageEmbed{{
						Title:       "Duplicate Image Detected",
						Color:       11648506,
						Description: fmt.Sprintf("This piece of art was uploaded by the following member(s): \n%v", strings.Join(uploadList, "\n")),
					}},
					Content: fmt.Sprintf("<@&%s>, An image that has been uploaded is a duplicate", modrole),
				})
			}

			models.AddImageUpload(m.Author.ID, m.ID, encodedValue)
		}
	}
}
