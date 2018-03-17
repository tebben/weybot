package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/commands"
	"github.com/tebben/weybot/configuration"
)

var (
	cfgFlag = flag.String("config", "config.json", "path of the config file")
)

func main() {
	flag.Parse()
	loadConfig()
	commands.Init()
	setupBot()
}

func loadConfig() {
	cfg := *cfgFlag
	conf, err := configuration.GetConfig(cfg)
	if err != nil {
		log.Fatal("config read error: ", err)
	}

	configuration.SetEnvironmentVariables(&conf)
}

func setupBot() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + configuration.CurrentConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening connection, %v", err)
	}

	fmt.Println("WeyBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

// gets called everytime a new message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself or no command prefix was found
	if m.Author.ID == s.State.User.ID || string(m.Content[0]) != configuration.CurrentConfig.CommandPrefix {
		return
	}

	inc := strings.Fields(m.Content)
	commands.HandleCommand(inc[0][1:], inc[1:], s, m)
}
