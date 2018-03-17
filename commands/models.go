package commands

import "github.com/bwmarrin/discordgo"

// ICommand interface defines the command interface
type ICommand interface {
	GetListener() string
	GetDescription() string
	Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string)
}

// BaseCommand describes a default command
type BaseCommand struct {
	command     string
	description string
}

// GetListener returns the command string to listen on
func (c *BaseCommand) GetListener() string {
	return c.command
}

// GetDescription returns the command description
func (c *BaseCommand) GetDescription() string {
	return c.description
}
