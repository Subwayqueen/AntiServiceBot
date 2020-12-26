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
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/anandpskerala/antiservicebot/service"
)

func main() {

	token := os.Getenv("TOKEN")
	b, err := gotgbot.NewBot(token)
	if err != nil {
		log.Fatalf("New bot creation failed", err.Error())
		return
	}

	u := ext.NewUpdater(b, nil)

	service.LoadService(u)
	u.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	err = u.StartPolling(b, &ext.PollingOpts{Clean: true})
	if err != nil {
		log.Fatalf("Polling failed", err.Error())
		return
	}
	bot := u.Bot.User.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Idle()

}

func start(b *ext.Context) error {

	_, err := b.Bot.SendMessage(b.EffectiveChat.Id, "<b>Hello 👋, I Am An 𝗔𝗡𝗧𝗜𝗦𝗘𝗥𝗩𝗜𝗖𝗘 𝗠𝗘𝗦𝗦𝗔𝗚𝗘 𝗕𝗢𝗧 🔖.\n\nI'm A Bot Which Can <u>𝗗𝗘𝗟𝗘𝗧𝗘 𝗦𝗘𝗥𝗩𝗜𝗖𝗘 𝗠𝗘𝗦𝗦𝗔𝗚𝗘𝗦</u> Like When A User 𝗘𝗡𝗧𝗘𝗥𝗦 Or 𝗘𝗫𝗜𝗧𝗦 A Group.\n\nI'm Fully Written In GO Language.\n\n<u>Note:~</u> You Should Promote Me As An Administrator & Give Atleast Two Admins Rights Shown Below For Getting My Full Service.\n\n➨ Right To Delete Messages.\n➨ Right To Add Admins.\n\nFor Support & Bug Issues Contact @Iggie</b>", &gotgbot.SendMessageOpts{ReplyToMessageId: b.EffectiveMessage.MessageId, ParseMode: "HTML"})
	if err != nil {
		log.Printf("Error in sending", err.Error())
	}
	return err
}
