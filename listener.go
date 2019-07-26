package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ============================================================
// Types and globals
// ============================================================

var instances map[string]*instance
var commands map[string]func(*discordgo.Message)

// ============================================================
// Init
// ============================================================

func init() {
	instances = make(map[string]*instance)
	commands = map[string]func(*discordgo.Message){
		"help": help,
		"h":    help,
		"run":  run,
		"r":    run,
	}
}

// ============================================================
// Commands
// ============================================================

func run(message *discordgo.Message) {
	if _, ok := instances[message.Author.ID]; ok {
		discord.ChannelMessageSend(message.ChannelID, "You already have a running VM!")
		return
	}
	vmInstance := newInstance(message.Author.ID, message.ChannelID)
	vmInstance.vm.LoadProg([]byte(removeCommandName(message.Content)))
	instances[message.Author.ID] = vmInstance

	go func() {
		for {
			out := vmInstance.vm.Read()
			discord.ChannelMessageSend(message.ChannelID, string(out))
		}
	}()

	if err := vmInstance.vm.Exec(); err != nil {
		discord.ChannelMessageSend(message.ChannelID, err.Error())
	}
	delete(instances, message.Author.ID)
}

func help(message *discordgo.Message) {
	discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Brainfuck VM Help",
		Description: "`>r/run <program>` - Runs a given brainfuck program\n`>h/help` - Displays this help",
	})
}

// ============================================================
// Handlers
// ============================================================

func commandHandler(discord *discordgo.Session, mc *discordgo.MessageCreate) {
	message := mc.Message
	if !strings.HasPrefix(message.Content, conf.Prefix) {
		return
	}
	tokens := strings.Split(message.Content, " ")
	if cmd, ok := commands[strings.TrimPrefix(tokens[0], conf.Prefix)]; ok {
		cmd(message)
	}
}

func inputHandler(discord *discordgo.Session, mc *discordgo.MessageCreate) {
	message := mc.Message
	if strings.HasPrefix(message.Content, conf.Prefix) {
		return
	}
	if inst, ok := instances[message.Author.ID]; ok {
		inst.vm.Write(message.Content[0])
	}
}

// ============================================================
// Helper functions
// ============================================================

func removeCommandName(message string) string {
	for cmd := range commands {
		message = strings.TrimPrefix(message, conf.Prefix+cmd)
	}
	return strings.TrimSpace(message)
}
