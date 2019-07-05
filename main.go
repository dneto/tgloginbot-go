package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	pamType = os.Getenv("PAM_TYPE")
	user    = os.Getenv("PAM_USER")
	host    = os.Getenv("PAM_RHOST")
)

func main() {
	if pamType == "open_session" {
		token := flag.String("token", "", "Telegram Bot Token")
		chatID := flag.Int64("chatID", 0, "Telegram Chat ID")
		flag.Parse()

		if *chatID == 0 {
			log.Fatal("Chat ID must be a valid number")
		}

		bot, err := telegram.NewBotAPI(*token)
		if err != nil {
			log.Fatal("Could not connect to Telegram Bot API. Reason:", err)
		}

		locationInfo := host
		resp, err := http.Get(fmt.Sprintf("https://ipinfo.io/%s", host))
		if err != nil {
			log.Print("Could not retrieve IP information. Reason:", err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Print("Could read IP information. Reason:", err)
			} else {
				ipInfo := map[string]string{}
				if err := json.Unmarshal(body, &ipInfo); err == nil {
					locationInfo = fmt.Sprintf(
						"%s, %s ,%s (IP=%s)",
						ipInfo["city"], ipInfo["region"],
						ipInfo["country"], ipInfo["ip"],
					)
				}
			}
		}

		hostname, _ := os.Hostname()
		now := time.Now().Format(time.RFC822Z)
		text := fmt.Sprintf(
			"*New login!*\n"+
				"  💻 %s\n"+
				"  🕓 %s\n"+
				"  👤 %s\n"+
				"  🌎 %s\n",
			hostname, now, user, locationInfo,
		)

		msg := telegram.NewMessage(*chatID, text)
		msg.ParseMode = telegram.ModeMarkdown
		bot.Send(msg)
	}
}
