package arsenal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"strings"
	"time"
)



// Found is a channel for posts
// parameters contain a key word or a collection of strings you want to trigger on
type Bot struct {
	bot reddit.Bot
	discordBot *discordgo.Session
	parameters []string
	Found chan *reddit.Post
}


// Post is a handler method for graw PostHandler interface
// Post will trigger on every new post make to the subreddit you have configured
func (b *Bot) Post(p *reddit.Post) error {
	for _, param := range b.parameters {
		if strings.Contains(p.Title, param) {

			// Discord Bot
			if b.discordBot != nil {
				b.discordBot.ChannelMessageSend("780955861463859220", fmt.Sprintf("Post found\n %s \n %s", p.Title, p.URL))
			} else {
				// if we hit a match ship post to channel
				log.Infof("Post found", p.Title, p.URL)
			}

		}
	}
	return nil
}

// NewArsenalBot creates a pointer to the bot
func NewArsenalBot(bot reddit.Bot, parameters []string, discordBot *discordgo.Session) *Bot {
	found := make(chan *reddit.Post)
	return &Bot{
		bot:bot,
		discordBot: discordBot,
		parameters:parameters,
		Found: found,
	}
}


// Run executed the Bot
// TODO add twillo access

func Run(config string, payload []string, discordBot *discordgo.Session){
	log.Info("Creating Reddit connection from config file")
	bot, err := reddit.NewBotFromAgentFile(config, 2*time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Configuring Arsenal Bot")
	ab := NewArsenalBot(bot, payload, discordBot)
	cfg := graw.Config{
		Subreddits:        []string{"gundeals"},
	}

	// run bot
	log.Infof("Running Arsenal ... searching for %s", strings.Join(payload, ", "))
	if _, wait, err := graw.Run(ab, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}