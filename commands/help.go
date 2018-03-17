package commands

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/configuration"
)

// HelpCommand describes the help command
type HelpCommand struct {
	BaseCommand
}

// NewHelpCommand creates and returns a new HelpCommand
func NewHelpCommand() *HelpCommand {
	hc := HelpCommand{}
	hc.command = configuration.CurrentConfig.Commands.Help.Base.Command
	hc.description = configuration.CurrentConfig.Commands.Help.Base.Description

	return &hc
}

// Handle is the function which handles the incomming command
func (c *HelpCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	var helpString bytes.Buffer
	helpString.WriteString(fmt.Sprintf("Showing u da wae to tha WeyBot %s\n", configuration.Version))
	helpString.WriteString("```\n")

	for k, v := range Commands {
		helpString.WriteString(fmt.Sprintf("%s%s - %s\n", configuration.CurrentConfig.CommandPrefix, k, v.GetDescription()))
	}

	helpString.WriteString("```\n")
	helpString.WriteString("Ideas for weybot? add them to the issue list on github.com/tebben/weybot\n")

	s.ChannelMessageSend(m.ChannelID, helpString.String())
}
