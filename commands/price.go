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

// PriceCommand describes the price command
type PriceCommand struct {
	BaseCommand
}

// NewPriceCommand creates and returns a new PriceCommand
func NewPriceCommand() *PriceCommand {
	pc := PriceCommand{}
	pc.command = configuration.CurrentConfig.Commands.Price.Base.Command
	pc.description = configuration.CurrentConfig.Commands.Price.Base.Description

	return &pc
}

// Handle is the function which handles the incomming command
func (p *PriceCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
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

	s.ChannelMessageSend(m.ChannelID, constructPriceMessage(ogreTicker, ccPrice))
}

func constructPriceMessage(ogreTicker *tradeogre.Ticker, ccPrice cryptocompare.Price) string {
	var msg bytes.Buffer
	msg.WriteString("BTC-WAE market value from tradeogre\n")
	msg.WriteString("```\n")

	WaeBTC, _ := strconv.ParseFloat(ogreTicker.Price, 64)

	if configuration.CurrentConfig.Commands.Price.ShowBTC {
		msg.WriteString(fmt.Sprintf("BTC    %.8f\n", WaeBTC))
	}

	if configuration.CurrentConfig.Commands.Price.ShowUSD {
		btcUSDPrice := ccPrice["USD"]
		WaeUSD := btcUSDPrice / (1 / WaeBTC)
		msg.WriteString(fmt.Sprintf("USD    %.8f\n", WaeUSD))
	}

	if configuration.CurrentConfig.Commands.Price.ShowUGX {
		btcUGXPrice := ccPrice["UGX"]
		WaeUGX := btcUGXPrice / (1 / WaeBTC)
		msg.WriteString(fmt.Sprintf("UGX    %.8f\n", WaeUGX))
	}

	msg.WriteString(fmt.Sprintf("\nhigh: %s low: %s\n", ogreTicker.High, ogreTicker.Low))

	msg.WriteString("```")

	return msg.String()
}
