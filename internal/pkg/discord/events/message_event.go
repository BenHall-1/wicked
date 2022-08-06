package events

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/benhall-1/wicked/internal/pkg/interfaces"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type MessageEvent struct {
	interfaces.Event
}

func (e MessageEvent) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		logchan           = os.Getenv("DISCORD_LOGS_CHANNEL")
		artchan           = os.Getenv("DISCORD_ART_CHANNEL")
		modrole           = os.Getenv("DISCORD_MOD_ROLE")
		server            = os.Getenv("DISCORD_MAIN_SERVER")
		currentMessageUrl = fmt.Sprintf("https://discord.com/channels/%s/%s/%s", m.GuildID, m.ChannelID, m.ID)
	)

	if m.ChannelID == artchan {
		for i := range m.Attachments {
			encodedValue := utils.GetBase64FromUrl(m.Attachments[i].URL)
			uploads := models.GetAllImagesUploadedByHash(utils.EscapeSpecialCharacters(encodedValue), m.Author.ID)
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
						Description: fmt.Sprintf("[%s](This) piece of art was uploaded by the following member(s): \n%v", currentMessageUrl, strings.Join(uploadList, "\n")),
					}},
					Content: fmt.Sprintf("<@&%s>, An image that has been uploaded is a duplicate", modrole),
				})
			}

			models.AddImageUpload(m.Author.ID, m.ID, utils.EscapeSpecialCharacters(encodedValue))
		}
	}

	content := m.Content
	phrases := models.GetAllPhrases(m.GuildID)
	matches := []string{}

	// Loop through all phrases and convert them to regex and then check whether the content string matches the regex
	for i := range phrases {
		r := regexp.MustCompile(phrases[i].RegexPhrase)
		match := r.Match([]byte(content))
		if match {
			matches = append(matches, phrases[i].RawPhrase)
		}
	}

	if len(matches) > 0 {
		s.ChannelMessageSendComplex(logchan, &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Title:       "Banned Phrase Detected",
				Color:       11648506,
				Description: fmt.Sprintf("The following banned phrases have been detected in [%s](this) message: \n%s", currentMessageUrl, strings.Join(matches, "\n- ")),
			}},
		})
	}

}
