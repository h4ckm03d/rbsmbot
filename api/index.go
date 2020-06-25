package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/subosito/gotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gotenv.Load()

	var (
		token = os.Getenv("TELEGRAM_TOKEN") // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	b, err := tb.NewBot(tb.Settings{
		Token:       token,
		Synchronous: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/gpx", func(m *tb.Message) {
		fmt.Println(strings.Fields(m.Payload))
		resp := ""
		for _, v := range strings.Fields(m.Payload) {
			resp += fmt.Sprintf(`*GPX %s*
- [Tokyo](https://duckduckgo.com)
- [Malang](https://duckduckgo.com)
- [Tokyo](https://duckduckgo.com)
- [Malang](https://duckduckgo.com)
`, strings.ToUpper(v))
		}
		b.Send(m.Sender, resp, &tb.SendOptions{
			ParseMode:             tb.ModeMarkdown,
			DisableWebPagePreview: true,
		})
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "Maaf bos, ga ngerti!")
	})

	var u tb.Update
	if err = json.Unmarshal([]byte(r.Body), &u); err == nil {
		b.ProcessUpdate(u)
	}
}
