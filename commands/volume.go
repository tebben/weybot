package commands

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/configuration"
	"github.com/tebben/weybot/cryptocompare"
	"github.com/tebben/weybot/tradeogre"
)

// VolumeCommand describes the volume command
type VolumeCommand struct {
	BaseCommand
}

// NewVolumeCommand creates and returns a new VolumeCommand
func NewVolumeCommand() *VolumeCommand {
	vc := VolumeCommand{}
	vc.command = configuration.CurrentConfig.Commands.Volume.Base.Command
	vc.description = configuration.CurrentConfig.Commands.Volume.Base.Description

	return &vc
}

// Handle is the function which handles the incomming command
func (v *VolumeCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	ogreTicker, err := tradeogre.GetTicker(configuration.CurrentConfig.Commands.Price.TradeogreTicker)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", err))
		return
	}

	ccPrice, err := cryptocompare.GetPrice("BTC", []string{"USD", "UGX"}, "", "", false, false)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", err))
		return
	}

	s.ChannelMessageSend(m.ChannelID, constructVolumeMessage(ogreTicker, ccPrice))
}

func constructVolumeMessage(ogreTicker *tradeogre.Ticker, ccPrice cryptocompare.Price) string {
	var msg bytes.Buffer
	msg.WriteString("BTC-WAE market volume from tradeogre\n")
	msg.WriteString("```\n")

	volume, _ := strconv.ParseFloat(ogreTicker.Volume, 64)

	if configuration.CurrentConfig.Commands.Price.ShowBTC {
		msg.WriteString(fmt.Sprintf("BTC    %.8f\n", volume))
	}

	if configuration.CurrentConfig.Commands.Price.ShowUSD {
		btcUSDPrice := ccPrice["USD"]
		volumeUSD := btcUSDPrice / (1 / volume)
		msg.WriteString(fmt.Sprintf("USD    %.8f\n", volumeUSD))
	}

	if configuration.CurrentConfig.Commands.Price.ShowUGX {
		btcUGXPrice := ccPrice["UGX"]
		volumeUGX := btcUGXPrice / (1 / volume)
		msg.WriteString(fmt.Sprintf("UGX    %.8f\n", volumeUGX))
	}

	msg.WriteString("```")

	return msg.String()
}
