package telegram

import (
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *Processor) doCmd(text string, chatID int, username string, messageID int) error {
	text = strings.TrimSpace(text)
	log.Printf("text: %s || chatID: %d || username: @%s || messageID: %d", text, chatID, username, messageID)
	//log.Printf("User @%s: %s", username, text)

	switch text {
	case StartCmd:
		return p.sendHello(chatID, messageID)
	case HelpCmd:
		return p.sendHelp(chatID, messageID)
	default:
		return p.tg.SendMessage(chatID, text, messageID)
	}

}

func (p *Processor) sendHelp(chatID int, messageID int) error {
	return p.tg.SendMessage(chatID, msgHelp, messageID)
}

func (p *Processor) sendHello(chatID int, messageID int) error {
	return p.tg.SendMessage(chatID, msgHello, messageID)
}
