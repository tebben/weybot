package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/configuration"
)

// WalletCommand describes the wallet command
type WalletCommand struct {
	BaseCommand
}

// NewWalletCommand creates and returns a new WalletCommand
func NewWalletCommand() *WalletCommand {
	wc := WalletCommand{}
	wc.command = configuration.CurrentConfig.Commands.Wallet.Base.Command
	wc.description = configuration.CurrentConfig.Commands.Wallet.Base.Description

	return &wc
}

// Handle is the function which handles the incomming command
func (v *WalletCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	s.ChannelMessageSend(m.ChannelID, "Not yet implemented")
}
