package discord

import (
	"fmt"
	"os"
	"strings"

	"github.com/benhall-1/wicked/internal/pkg/discord/cmds"
	"github.com/benhall-1/wicked/internal/pkg/models"
	"github.com/benhall-1/wicked/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

func (db DiscordBot) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	db.RegisterCommands()
}

func (db DiscordBot) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func (db DiscordBot) InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmds := cmds.Commands

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := cmds[i.ApplicationCommandData().Name]; ok {
			h.HandleCommand(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := cmds[strings.Split(i.MessageComponentData().CustomID, "-")[0]]; ok {
			h.HandleComponents(s, i)
		}
	case discordgo.InteractionModalSubmit:
		if h, ok := cmds[strings.Split(i.ModalSubmitData().CustomID, "-")[0]]; ok {
			h.HandleModals(s, i)
		}
	}
}
