package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session

func main() {
	var err error
	if err = loadConfig(); err != nil {
		log.Fatal(err)
	}

	discord, err = discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(commandHandler)
	discord.AddHandler(inputHandler)

	if err = discord.Open(); err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	<-make(chan struct{})
}
