package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/subosito/gotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func bot() {
	gotenv.Load()

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

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

	b.Start()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from rbsm!</h1>")
	bot()
}
