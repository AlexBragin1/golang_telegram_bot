package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("accept"),
		tgbotapi.NewKeyboardButton("ignore"),
		tgbotapi.NewKeyboardButton("delete"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

const (
	commandStart       = `start`
	commandMenu        = `menu`
	commandPrice       = `price`
	commandBankDetails = `bankdetails`
	commandPayFor      = `payfor`
	commandStartVideo  = `startvideo`
	commandAccept      = `accept`
	commandIgnore      = `ignore`
	commandDeleteUsers = `delete`
)

func (b *Bot) handleCommand(message *tgbotapi.Message, UserName string, channelChat int64, userId int64) error {
	var channelAdmin int64 = -1002117831281
	//var chanelPrivate int64 = -1002111427878
	msg := tgbotapi.NewMessage(channelChat, "не правильная команда")

	switch message.Command() {
	case commandStart:
		msg.Text = startText + startMenu
	case commandPrice:
		msg.Text = startPrice
	case commandMenu:
		b.startMenuChat(msg, channelChat, UserName)
	case commandBankDetails:
		msg.Text = startBankDetails
	case commandStartVideo:
		msg.Text = startVideo
	case commandPayFor:
		msg.Text = `Отправил оплату`

		b.startMenuChat(msg, channelAdmin, UserName)
		b.bot.Send(msg)
		msg.Text = `Отправил оплату  
		/accept 
		/ignore`

	case commandAccept:
		if channelChat == channelAdmin {
			//s:=cnannelChat

			msg.ChatID = channelAdmin
			msg.Text = `приняли`
			b.InviteExistingUser(message, UserName, channelChat, userId)
			msg1 := tgbotapi.NewMessage(userId, "приглашение")
			b.bot.Send(msg1)
			b.bot.Send(msg)

		}

	case commandIgnore:
		if channelChat == channelAdmin {
			msg.ChatID = channelAdmin
			msg.Text = `отказать`
		}
	case commandDeleteUsers:
		b.deleteUsers(UserName, channelChat, userId)

		//	keyboard := tgbotapi.InlineKeyboardMarkup{}
		//var row []tgbotapi.InlineKeyboardButton
		//	btn := tgbotapi.NewInlineKeyboardButtonData("text", "text")
		//	row = append(row, btn)
		//	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	}
	_, err := b.bot.Send(msg)
	msg.ChatID = 1108447926
	return err
}
func (b *Bot) startMenuChat(msg tgbotapi.MessageConfig, channelChat int64, userName string) {

	if channelChat == -1002117831281 {

		msg.ChatID = channelChat
		msg.ReplyMarkup = numericKeyboard

		//msg.Text = `отправил оплату:` + userName + startMenuChatAdmin

		//msg.Text = startMenuChatAdmin
		//b.bot.Send(msg)
	}
	if channelChat == 1108447926 {
		msg.Text = startMenu
	}

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}

}

func (b *Bot) deleteUsers(UserName string, cnannelChat int64, userId int64) {
	b.bot.Send(tgbotapi.DeclineChatJoinRequest{
		UserID: userId,
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: cnannelChat,
		},
	})

}

func extractInviteLinkFromResponse(response *tgbotapi.APIResponse) (string, error) {
	var dat map[string]interface{}
	json.Unmarshal(response.Result, &dat)
	for k, v := range dat {
		if k == "invite_link" {
			return fmt.Sprint(v), nil
		}
	}
	return "", fmt.Errorf("Could not find invite link")
}
func (b *Bot) InviteExistingUser(message *tgbotapi.Message, UserName string, channelChat int64, userId int64) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "не правильная команда")
	unixTime24HoursFromNow := int(time.Now().Add(time.Duration(24 * time.Hour)).Unix())
	chatConfig := tgbotapi.ChatConfig{
		ChatID:             channelChat,
		SuperGroupUsername: b.bot.Self.UserName,
	}
	createInviteLinkConfig := tgbotapi.CreateChatInviteLinkConfig{
		ChatConfig:         chatConfig,
		Name:               UserName,
		ExpireDate:         unixTime24HoursFromNow,
		MemberLimit:        1,
		CreatesJoinRequest: false,
	}
	response, err := b.bot.Request(createInviteLinkConfig)
	if err != nil {
		return err
	}
	link, err := extractInviteLinkFromResponse(response)
	if err != nil {
		return err
	}
	msg.Text = link
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

//func (b *Bot) handleMessage(message *tgbotapi.Message) {
//	log.Printf("[%s] %s", message.From.UserName, message.Text)

//	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
//	//msg.ReplyToMessageID = update.Message.MessageID

//	b.bot.Send(msg)
//}
