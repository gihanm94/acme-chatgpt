package main

import (
	"fmt"
	"log"
	"github.com/gihanm94/acme-chatgpt/internal/config"
	"github.com/gihanm94/acme-chatgpt/internal/openai"
	"github.com/gihanm94/acme-chatgpt/internal/larkbot"
)

func main() {
	v := config.LoadConfig()

	g := openai.New(v.GetString("open_ai_key"))
	bot := larkbot.New(larkbot.Config{
		AppID:             v.GetString("app_id"),
		AppSecret:         v.GetString("app_secret"),
		VerificationToken: v.GetString("verification_token"),
		EventEncryptKey:   v.GetString("event_encrypt_key"),
		Name:              v.GetString("bot_name"),
		BaseUrl:           v.GetString("lark_base_url"),
		Port:              v.GetInt("port"),
	})

	handler := func(msg larkbot.Message) {
		if msg.Content == "" {
			return
		}
		log.Printf("receive message: %s %s\n", msg.ID, msg.Content)
		if msg.Type == larkbot.GroupChat && !msg.MentionMe {
			return
		}

		res, err := g.Handle(fmt.Sprintf("Q:%s\nA:", msg.Content))
		if err != nil {
			log.Printf("gpt error: %s\n", err)
			return
		}
		log.Printf("gpt: %s %s\n", msg.Content, res)
		err = bot.Reply(msg.ID, res)
		if err != nil {
			log.Printf("reply error: %s\n", err)
			return
		}
	}

	log.Fatal(bot.Run(handler))
}
