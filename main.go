/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/antiservicebot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/github.com/Anandpskerala/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package main

import (
	"fmt"
	"os"

	//"github.com/anandpskerala/antiservicebot/config"
	"github.com/anandpskerala/antiservicebot/service"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/parsemode"

	//"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.NewProductionEncoderConfig()

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))

	defer logger.Sync()

	l := logger.Sugar()

	token := os.Getenv("TOKEN")
	u, err := gotgbot.NewUpdater(logger, token)
	if err != nil {
		l.Fatalw("Updater failed starting", zap.Error(err))
		return
	}

	service.LoadService(u)

	u.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	err = u.StartPolling()
	if err != nil {
		l.Fatalw("Polling failed", zap.Error(err))
		return
	}
	bot := u.Bot.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Idle()

}

func start(b ext.Bot, u *gotgbot.Update) error {

	msg := b.NewSendableMessage(u.EffectiveChat.Id, "<b>Hello 👋, I Am An 𝗔𝗡𝗧𝗜𝗦𝗘𝗥𝗩𝗜𝗖𝗘 𝗠𝗘𝗦𝗦𝗔𝗚𝗘 𝗕𝗢𝗧 🔖.\n\nI'm A Bot Which Can <u>𝗗𝗘𝗟𝗘𝗧𝗘 𝗦𝗘𝗥𝗩𝗜𝗖𝗘 𝗠𝗘𝗦𝗦𝗔𝗚𝗘𝗦</u> Like When A User 𝗘𝗡𝗧𝗘𝗥𝗦 Or 𝗘𝗫𝗜𝗧𝗦 A Group.\n\nI'm Fully Written In GO Language.\n\n<u>Note:~</u> You Should Promote Me As An Administrator & Give Atleast Two Admins Rights Shown Below For Getting My Full Service.\n\n➨ Right To Delete Messages.\n➨ Right To Add Admins.\n\nFor Support & Bug Issues Contact @Iggie</b>")
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ParseMode = parsemode.Html
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
} 
