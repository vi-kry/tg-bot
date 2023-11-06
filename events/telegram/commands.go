package telegram

import (
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("text: %s || chatID: %d || username: @%s", text, chatID, username)
	//log.Printf("User @%s: %s", username, text)

	switch text {
	case StartCmd:
		return p.sendHello(chatID)
	case HelpCmd:
		return p.sendHelp(chatID)
	default:
		return p.tg.SendMessage(chatID, text)
	}

}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
