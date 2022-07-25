package main

import (
	"fmt"

	"github.com/benhall-1/wicked/internal/pkg/bot"
	"github.com/benhall-1/wicked/internal/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Setting up environment...")
	_ = godotenv.Load()
	fmt.Println("Opening connection to the database...")
	db.OpenDatabase()
	fmt.Println("Starting the bot...")
	bot.Start()
}
