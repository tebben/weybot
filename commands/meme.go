package commands

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/bwmarrin/discordgo"
	"github.com/tebben/weybot/configuration"
	"github.com/tebben/weybot/utils"
)

var (
	client = &http.Client{Timeout: 30 * time.Second}
)

// MemeCommand describes the meme command
type MemeCommand struct {
	BaseCommand
}

// NewMemeCommand creates and returns a new NewMemeCommand
func NewMemeCommand() *MemeCommand {
	mc := MemeCommand{}
	mc.command = configuration.CurrentConfig.Commands.Meme.Base.Command
	mc.description = configuration.CurrentConfig.Commands.Meme.Base.Description

	return &mc
}

// Handle is the function which handles the incomming command
func (c *MemeCommand) Handle(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	title, img, err := c.getRandomImage()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", err))
	}

	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("%s - 9GAG\n", title))
	msg.WriteString(img)

	s.ChannelMessageSend(m.ChannelID, msg.String())
}

func (c *MemeCommand) getRandomImage() (string, string, error) {
	url := "https://9gag.com/random"

	resp, err := utils.GetDataFromEndpoint(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Unable to get data from %s", url)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	title, _ := c.getHTMLTitle(strings.NewReader(bodyString))
	realTitle := strings.Replace(title, " - 9GAG", "", -1)
	locationParts := strings.Split(resp.Request.URL.Path, "/")
	img := locationParts[len(locationParts)-1]

	if strings.Contains(bodyString, fmt.Sprintf("\"title\":\"%s\",\"type\":\"Photo\"", realTitle)) {
		return realTitle, fmt.Sprintf("https://img-9gag-fun.9cache.com/photo/%s_700b.jpg", img), nil
	}

	return c.getRandomImage()
}

func (c *MemeCommand) isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func (c *MemeCommand) traverse(n *html.Node) (string, bool) {
	if c.isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		result, ok := c.traverse(child)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func (c *MemeCommand) getHTMLTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return c.traverse(doc)
}
