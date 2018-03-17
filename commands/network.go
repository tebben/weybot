package commands

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/configuration"
	"github.com/tebben/weybot/utils"
)

// NetworkCommand describes the network command
type NetworkCommand struct {
	BaseCommand
}

// NewNetworkCommand creates and returns a new NetworkCommandCommand
func NewNetworkCommand() *NetworkCommand {
	nc := NetworkCommand{}
	nc.command = configuration.CurrentConfig.Commands.Network.Base.Command
	nc.description = configuration.CurrentConfig.Commands.Network.Base.Description

	return &nc
}

// Handle is the function which handles the incomming command
func (c *NetworkCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	apiEndpoint := fmt.Sprintf("%s/api", configuration.CurrentConfig.Commands.Network.URL)

	diffString, _ := utils.GetStringFromEndpoint(fmt.Sprintf("%s/getdifficulty", apiEndpoint))
	blocksString, _ := utils.GetStringFromEndpoint(fmt.Sprintf("%s/getblockcount", apiEndpoint))
	hashString, _ := utils.GetStringFromEndpoint(fmt.Sprintf("%s/getnetworkhashps", apiEndpoint))
	connectionsString, _ := utils.GetStringFromEndpoint(fmt.Sprintf("%s/getconnectioncount", apiEndpoint))

	hash, _ := strconv.ParseFloat(hashString, 64)
	diff, _ := strconv.ParseFloat(diffString, 64)

	var msg bytes.Buffer

	msg.WriteString(fmt.Sprintf("Weycoin network data from %s\n", configuration.CurrentConfig.Commands.Network.URL))
	msg.WriteString("```\n")
	msg.WriteString(fmt.Sprintf("Connections  %s \n", connectionsString))
	msg.WriteString(fmt.Sprintf("Block index  %s \n", blocksString))
	msg.WriteString(fmt.Sprintf("Hashrate     %.4f(GH/s) \n", hash/1000000000))
	msg.WriteString(fmt.Sprintf("Difficulty   %.2f \n", diff))
	msg.WriteString("```")
	s.ChannelMessageSend(m.ChannelID, msg.String())
}
