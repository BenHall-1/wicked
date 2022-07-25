package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/benhall-1/wicked/internal/pkg/db"
	"github.com/benhall-1/wicked/internal/pkg/discord"
	"github.com/bwmarrin/discordgo"
)

func Start() {

	discordBot := setupDiscordBot()

	fmt.Println(fmt.Sprintf("Logged in as %s#%s", discordBot.State.User.Username, discordBot.State.User.Discriminator))
	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChannel

	fmt.Println("Shutting down discord bot...")
	discordBot.Close()
}

func setupDiscordBot() *discordgo.Session {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error whilst creating session: %v", err.Error()))
		os.Exit(1)
	}

	discordBot := discord.NewDiscordBot(os.Getenv("DISCORD_BOT_TOKEN"), dg, db.Database)
	discordBot.SetupHandlers()

	if err := dg.Open(); err != nil {
		fmt.Println(fmt.Sprintf("Error whilst opening connection: %v", err.Error()))
		os.Exit(1)
	}

	return dg
}
