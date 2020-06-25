package main

import (
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	
	gotenv.Load()

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
	})

	b.han("/gpx", func(m *tb.Message) {
		m.Chat.Type == tb.OnText
		b.Send(m.Sender, `*List GPX GMT+1*
- [Tokyo](https://duckduckgo.com)
- [Malang](https://duckduckgo.com)
- [Tokyo](https://duckduckgo.com)
- [Malang](https://duckduckgo.com)
`,&tb.SendOptions{
	ParseMode: tb.ModeMarkdown,
	DisableWebPagePreview: true,
})
	})

	b.Start()
}