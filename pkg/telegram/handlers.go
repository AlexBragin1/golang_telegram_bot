package telegram

import (
	"encoding/json"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
)

func (b *Bot) handleCommand(message *tgbotapi.Message, UserName string, cnannelChat int64, userId int64) error {
	var channelAdmin int64 = -1002117831281
	//var chanelPrivate int64 = -1002111427878
	msg := tgbotapi.NewMessage(message.Chat.ID, "не правильная команда")

	switch message.Command() {
	case commandStart:
		msg.Text = startText + startMenu
	case commandPrice:
		msg.Text = startPrice
	case commandMenu:
		msg.Text = startMenu
	case commandBankDetails:
		msg.Text = startBankDetails
	case commandStartVideo:
		msg.Text = startVideo
	case commandPayFor:
		msg.Text = startPayFor
		b.bot.Send(msg)
		msg.Text = `Отправил оплату  
		/accept 
		/ignore`
		msg.ChatID = channelAdmin

	case commandAccept:
		if cnannelChat == channelAdmin {
			//s:=cnannelChat

			msg.ChatID = channelAdmin
			msg.Text = `приняли`
			s := tgbotapi.CreateChatInviteLinkConfig()
			msg1 := tgbotapi.NewMessage(userId, "приглашение")
			b.bot.Send(msg1)
			b.bot.Send()

		}

	case commandIgnore:
		if cnannelChat == channelAdmin {
			msg.ChatID = channelAdmin
			msg.Text = `отказать`
		}

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
func (user *User) InviteExistingUser(bot *tgbotapi.BotAPI) error {
	chatIds, err := user.GetChatIds()
	if err != nil {
		return err
	}
	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(user.TelegramUserID, "")
		unixTime24HoursFromNow := int(time.Now().Add(time.Duration(24 * time.Hour)).Unix())
		chatConfig := tgbotapi.ChatConfig{
			ChatID:             chatId,
			SuperGroupUsername: bot.Self.UserName,
		}
		createInviteLinkConfig := tgbotapi.CreateChatInviteLinkConfig{
			ChatConfig:         chatConfig,
			Name:               user.GetName(),
			ExpireDate:         unixTime24HoursFromNow,
			MemberLimit:        1,
			CreatesJoinRequest: false,
		}
		response, err := bot.Request(createInviteLinkConfig)
		if err != nil {
			return err
		}
		link, err := extractInviteLinkFromResponse(response)
		if err != nil {
			return err
		}
		msg.Text = link
		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

//func (b *Bot) handleMessage(message *tgbotapi.Message) {
//	log.Printf("[%s] %s", message.From.UserName, message.Text)

//	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
//	//msg.ReplyToMessageID = update.Message.MessageID

//	b.bot.Send(msg)
//}
