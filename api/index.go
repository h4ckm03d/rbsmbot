package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/subosito/gotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gotenv.Load()

	var (
		token = os.Getenv("TELEGRAM_TOKEN") // you must add it to your config vars
	)

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
		log.Println(strings.Fields(m.Payload))
		var buffer bytes.Buffer
		filter := strings.Fields(m.Payload)
		if len(filter) == 0 {
			_, _ = b.Send(m.Sender, "/gpx command perlu timezone nih, misal *GMT+2*")
			return
		}
		gpxFiles := getGPX(filter...)
		if len(gpxFiles) == 0 {
			_, _ = b.Send(m.Sender, fmt.Sprintf("Maaf bos data gpx buat %v, g ketemu.", filter))
			return
		}
		for k, v := range gpxFiles {
			buffer.WriteString(fmt.Sprintf(`*GPX %s*\n`, strings.ToUpper(k)))
			sort.Strings(k)
			for _, gpx := range v {
				buffer.WriteString(fmt.Sprintf(`-[%s](https://raw.githubusercontent.com/h4ckm03d/rbsmbot/master/static/gpx/%s)\n`, gpx, url.PathEscape(gpx)))
			}
			buffer.WriteString("\n")
		}
		b.Send(m.Sender, buffer.String(), &tb.SendOptions{
			ParseMode:             tb.ModeMarkdown,
			DisableWebPagePreview: true,
		})
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "Maaf bos, ga ngerti!")
	})

	var u tb.Update

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	if err = json.Unmarshal(body, &u); err == nil {
		b.ProcessUpdate(u)
	}
}

func getGPX(filters ...string) map[string][]string {
	res := map[string][]string{}
	for _, filter := range filters {
		if val, ok := generatedMap[strings.ToUpper(filter)]; ok {
			res[filter] = val
		}
	}

	return res
}

// TODO: need tobe generated from static/gpx
var generatedMap = map[string][]string{
	"GMT+1": {
		"GMT+1 London 1-100.gpx",
		"GMT+1 London 1.gpx",
		"GMT+1 London 100-200.gpx",
		"GMT+1 London 2.gpx",
		"GMT+1 London 200-300.gpx",
		"GMT+1 London 300-400.gpx",
		"GMT+1 London 400-500.gpx",
		"GMT+1 London 500-600.gpx",
		"GMT+1 London 600-700.gpx",
		"GMT+1 London 700-800.gpx",
		"GMT+1 London 800-900.gpx",
		"GMT+1 London 900-1000.gpx",
		"GMT+1 London1 - 526 pokestops.gpx",
		"GMT+1 London1-998stops.gpx",
		"GMT+1 London2 - 593 pokestops.gpx",
		"GMT+1 London2-490stops.gpx",
	},
	"GMT+10": {
		"GMT+10 Melbourne 1-100.gpx",
		"GMT+10 Melbourne 1.gpx",
		"GMT+10 Melbourne 100-200.gpx",
		"GMT+10 Melbourne 2.gpx",
		"GMT+10 Melbourne 200-300.gpx",
		"GMT+10 Melbourne 300-400.gpx",
		"GMT+10 Melbourne 400-500.gpx",
		"GMT+10 Melbourne 500-600.gpx",
		"GMT+10 Melbourne 600-700.gpx",
		"GMT+10 Melbourne 700-800.gpx",
		"GMT+10 Melbourne1 - 492 pokestops.gpx",
		"GMT+10 Melbourne2 - 309 pokestops.gpx",
		"GMT+10 Sydney 1-100.gpx",
		"GMT+10 Sydney 1.gpx",
		"GMT+10 Sydney 100-200.gpx",
		"GMT+10 Sydney 2.gpx",
		"GMT+10 Sydney 200-300.gpx",
		"GMT+10 Sydney 300-400.gpx",
		"GMT+10 Sydney 400-500.gpx",
		"GMT+10 Sydney 500-600.gpx",
		"GMT+10 Sydney 600-700.gpx",
		"GMT+10 Sydney 700-800.gpx",
		"GMT+10 Sydney 800-900.gpx",
		"GMT+10 Sydney1 - 473 pokestops.gpx",
		"GMT+10 Sydney2 - 399 pokestops.gpx",
		"GMT+10 Sydney3-643stops.gpx",
		"GMT+10 Sydney4-852stops.gpx",
	},
	"GMT+12": {
		"GMT+12 Wellington.gpx",
	},
	"GMT+2": {
		"GMT+2 Amsterdam - 484 pokestops.gpx",
		"GMT+2 Amsterdam 1-108.gpx",
		"GMT+2 Amsterdam 109-220.gpx",
		"GMT+2 Amsterdam 200.gpx",
		"GMT+2 Amsterdam 221-349.gpx",
		"GMT+2 Amsterdam 350-485.gpx",
		"GMT+2 Amsterdam.gpx",
		"GMT+2 Anchorage-257stops.gpx",
		"GMT+2 Brussels - 494 pokestops.gpx",
		"GMT+2 Brussels 1-100.gpx",
		"GMT+2 Brussels 100-200.gpx",
		"GMT+2 Brussels 200-300.gpx",
		"GMT+2 Brussels 300-400.gpx",
		"GMT+2 Brussels 400-500.gpx",
		"GMT+2 Brussels.gpx",
		"GMT+2 Budapest.gpx",
		"GMT+2 CapeTown-346stops.gpx",
		"GMT+2 Capetown 1-100.gpx",
		"GMT+2 Capetown 100-200.gpx",
		"GMT+2 Capetown.gpx",
		"GMT+2 CopenhagenMall.gpx",
		"GMT+2 Essen-355stops.gpx",
		"GMT+2 GMT+2 Barcelona - 577 pokestops.gpx",
		"GMT+2 GMT+2 Barcelona 1-100.gpx",
		"GMT+2 GMT+2 Barcelona 100-200.gpx",
		"GMT+2 GMT+2 Barcelona 200-300.gpx",
		"GMT+2 GMT+2 Barcelona 300-400.gpx",
		"GMT+2 GMT+2 Barcelona 400-500.gpx",
		"GMT+2 GMT+2 Barcelona 500-600.gpx",
		"GMT+2 GMT+2 Barcelona.gpx",
		"GMT+2 Kobenhavn5-350stops.gpx",
		"GMT+2 Madrid - 472 pokestops.gpx",
		"GMT+2 Madrid 1-100.gpx",
		"GMT+2 Madrid 100-200.gpx",
		"GMT+2 Madrid 200-300.gpx",
		"GMT+2 Madrid 300-400.gpx",
		"GMT+2 Madrid 400-500.gpx",
		"GMT+2 Madrid-366stops.gpx",
		"GMT+2 Madrid.gpx",
		"GMT+2 Paris - 610 pokestops.gpx",
		"GMT+2 Paris 1-100.gpx",
		"GMT+2 Paris 100-200.gpx",
		"GMT+2 Paris 200-300.gpx",
		"GMT+2 Paris 300-400.gpx",
		"GMT+2 Paris 400-500.gpx",
		"GMT+2 Paris 500-600.gpx",
		"GMT+2 Paris-438stops.gpx",
		"GMT+2 Paris.gpx",
		"GMT+2 Rome - 438 pokestops.gpx",
		"GMT+2 Rome 1-100.gpx",
		"GMT+2 Rome 100-200.gpx",
		"GMT+2 Rome 200-300.gpx",
		"GMT+2 Rome.gpx",
		"GMT+2 Rome300-400.gpx",
		"GMT+2 Venice-417stops.gpx",
		"GMT+2 Zaragoza.gpx",
	},
	"GMT+8": {
		"GMT+8 Metropolitan_Park.gpx",
		"GMT+8 Taipei Safari Zone.gpx",
		"GMT+8 Taiwan - 504 pokestops.gpx",
		"GMT+8 Taiwan 1-100.gpx",
		"GMT+8 Taiwan 100-200.gpx",
		"GMT+8 Taiwan 200-300.gpx",
		"GMT+8 Taiwan 300-400.gpx",
		"GMT+8 Taiwan 400-500.gpx",
		"GMT+8 Taiwan.gpx",
	},
	"GMT+9": {
		"GMT+9 Busan 1-100.gpx",
		"GMT+9 Busan 100-200.gpx",
		"GMT+9 Busan 2.gpx",
		"GMT+9 Busan 200-300.gpx",
		"GMT+9 Busan 3.gpx",
		"GMT+9 Busan 300-400.gpx",
		"GMT+9 Busan 400-500.gpx",
		"GMT+9 Busan2-279stops.gpx",
		"GMT+9 Busan3-171stops.gpx",
		"GMT+9 Kyoto 1-100.gpx",
		"GMT+9 Kyoto 100-200.gpx",
		"GMT+9 Kyoto 1000-1100.gpx",
		"GMT+9 Kyoto 1100-1200.gpx",
		"GMT+9 Kyoto 1200-1300.gpx",
		"GMT+9 Kyoto 2.gpx",
		"GMT+9 Kyoto 200-300.gpx",
		"GMT+9 Kyoto 3.gpx",
		"GMT+9 Kyoto 300-400.gpx",
		"GMT+9 Kyoto 400-500.gpx",
		"GMT+9 Kyoto 500-600.gpx",
		"GMT+9 Kyoto 600-700.gpx",
		"GMT+9 Kyoto 700-800.gpx",
		"GMT+9 Kyoto 800-900.gpx",
		"GMT+9 Kyoto 900-1000.gpx",
		"GMT+9 Kyoto.gpx",
		"GMT+9 Kyoto1 - 406 pokestops.gpx",
		"GMT+9 Kyoto2 - 548 pokestops.gpx",
		"GMT+9 Kyoto3 - 384 pokestops.gpx",
		"GMT+9 Osaka - 441 pokestops.gpx",
		"GMT+9 Seoul - 554 pokestops.gpx",
		"GMT+9 Songnam1-568stops.gpx",
		"GMT+9 Tokyo 1-100.gpx",
		"GMT+9 Tokyo 1.gpx",
		"GMT+9 Tokyo 100-200.gpx",
		"GMT+9 Tokyo 1000-1100.gpx",
		"GMT+9 Tokyo 1100-1200.gpx",
		"GMT+9 Tokyo 1200-1300.gpx",
		"GMT+9 Tokyo 1300-1400.gpx",
		"GMT+9 Tokyo 1400-1500.gpx",
		"GMT+9 Tokyo 1500-1600.gpx",
		"GMT+9 Tokyo 1600-1700.gpx",
		"GMT+9 Tokyo 1700-1800.gpx",
		"GMT+9 Tokyo 1800-1900.gpx",
		"GMT+9 Tokyo 1900-2000.gpx",
		"GMT+9 Tokyo 2.gpx",
		"GMT+9 Tokyo 200-300.gpx",
		"GMT+9 Tokyo 2000-2100.gpx",
		"GMT+9 Tokyo 2100-2200.gpx",
		"GMT+9 Tokyo 2200-2300.gpx",
		"GMT+9 Tokyo 2300-2400.gpx",
		"GMT+9 Tokyo 2400-2500.gpx",
		"GMT+9 Tokyo 2500-2600.gpx",
		"GMT+9 Tokyo 2600-2700.gpx",
		"GMT+9 Tokyo 2700-2800.gpx",
		"GMT+9 Tokyo 2800-2900.gpx",
		"GMT+9 Tokyo 2900-3000.gpx",
		"GMT+9 Tokyo 3.gpx",
		"GMT+9 Tokyo 300-400.gpx",
		"GMT+9 Tokyo 3000-3100.gpx",
		"GMT+9 Tokyo 3100-3200.gpx",
		"GMT+9 Tokyo 3200-3300.gpx",
		"GMT+9 Tokyo 3300-3400.gpx",
		"GMT+9 Tokyo 3400-3500.gpx",
		"GMT+9 Tokyo 3500-3600.gpx",
		"GMT+9 Tokyo 3600-3700.gpx",
		"GMT+9 Tokyo 3700-3800.gpx",
		"GMT+9 Tokyo 3800-3900.gpx",
		"GMT+9 Tokyo 3900-4000.gpx",
		"GMT+9 Tokyo 4.gpx",
		"GMT+9 Tokyo 400-500.gpx",
		"GMT+9 Tokyo 4000-4100.gpx",
		"GMT+9 Tokyo 4100-4200.gpx",
		"GMT+9 Tokyo 4200-4300.gpx",
		"GMT+9 Tokyo 4300-4400.gpx",
		"GMT+9 Tokyo 5.gpx",
		"GMT+9 Tokyo 500-600.gpx",
		"GMT+9 Tokyo 6.gpx",
		"GMT+9 Tokyo 600-700.gpx",
		"GMT+9 Tokyo 7.gpx",
		"GMT+9 Tokyo 700-800.gpx",
		"GMT+9 Tokyo 8.gpx",
		"GMT+9 Tokyo 800-900.gpx",
		"GMT+9 Tokyo 9.gpx",
		"GMT+9 Tokyo 900-1000.gpx",
		"GMT+9 Tokyo1 - 556 pokestops.gpx",
		"GMT+9 Tokyo2 - 507 pokestops.gpx",
		"GMT+9 Tokyo3 - 645 pokestops.gpx",
		"GMT+9 Tokyo4 - 451 pokestops.gpx",
		"GMT+9 Tokyo5 - 547 pokestops.gpx",
		"GMT+9 Tokyo6 - 492 pokestops.gpx",
		"GMT+9 Tokyo7 - 573 pokestops.gpx",
		"GMT+9 Tokyo8 - 400 pokestops.gpx",
		"GMT+9 Tokyo9 - 384 pokestops.gpx",
		"GMT+9 Yokohama - 470 pokestops.gpx",
	},
	"GMT-10": {
		"GMT-10 Honolulu.gpx",
	},
	"GMT-4": {
		"GMT-4 Boston - 522 pokestops.gpx",
		"GMT-4 Boston 1-100.gpx",
		"GMT-4 Boston 100-200.gpx",
		"GMT-4 Boston 200-300.gpx",
		"GMT-4 Boston 400-525.gpx",
		"GMT-4 Boston.gpx",
		"GMT-4 Boston300-400.gpx",
		"GMT-4 DC 1-100.gpx",
		"GMT-4 DC 100-200.gpx",
		"GMT-4 DC 200-300.gpx",
		"GMT-4 DC 300-400.gpx",
		"GMT-4 DC.gpx",
		"GMT-4 MiamiFL-259stops.gpx",
		"GMT-4 NY 1-100.gpx",
		"GMT-4 NY 1.gpx",
		"GMT-4 NY 100-200.gpx",
		"GMT-4 NY 1000-1100.gpx",
		"GMT-4 NY 1100-1200.gpx",
		"GMT-4 NY 1200-1300.gpx",
		"GMT-4 NY 1300-1400.gpx",
		"GMT-4 NY 1400-1500.gpx",
		"GMT-4 NY 1500-1600.gpx",
		"GMT-4 NY 1600-1700.gpx",
		"GMT-4 NY 1700-1800.gpx",
		"GMT-4 NY 1800-1900.gpx",
		"GMT-4 NY 1900-2000.gpx",
		"GMT-4 NY 2.gpx",
		"GMT-4 NY 200-300.gpx",
		"GMT-4 NY 2000-2100.gpx",
		"GMT-4 NY 2100-2200.gpx",
		"GMT-4 NY 2200-2300.gpx",
		"GMT-4 NY 2300-2400.gpx",
		"GMT-4 NY 2400-2500.gpx",
		"GMT-4 NY 3.gpx",
		"GMT-4 NY 300-400.gpx",
		"GMT-4 NY 4.gpx",
		"GMT-4 NY 400-500.gpx",
		"GMT-4 NY 5.gpx",
		"GMT-4 NY 500-600.gpx",
		"GMT-4 NY 6.gpx",
		"GMT-4 NY 600-700.gpx",
		"GMT-4 NY 700-800.gpx",
		"GMT-4 NY 800-900.gpx",
		"GMT-4 NY 900-1000.gpx",
		"GMT-4 New York1 - 602 pokestops.gpx",
		"GMT-4 New York1-532stops.gpx",
		"GMT-4 New York2 - 528 pokestops.gpx",
		"GMT-4 New York2-508stops.gpx",
		"GMT-4 New York3 - 542 pokestops.gpx",
		"GMT-4 New York3-570stops.gpx",
		"GMT-4 New York4 - 345 pokestops.gpx",
		"GMT-4 New York5 - 310 pokestops.gpx",
		"GMT-4 New York6 - 320 pokestops.gpx",
		"GMT-4 Philadelphia - 473 pokestops.gpx",
		"GMT-4 Philly 1-100.gpx",
		"GMT-4 Philly 100-200.gpx",
		"GMT-4 Philly 200-300.gpx",
		"GMT-4 Philly 300-400.gpx",
		"GMT-4 Philly 400-500.gpx",
		"GMT-4 Philly.gpx",
		"GMT-4 Washington - 463 pokestops.gpx",
		"GMT-4 Washington.gpx",
	}, "GMT-5": {
		"GMT-5 Austin 1-100.gpx",
		"GMT-5 Austin 100-200.gpx",
		"GMT-5 Austin 200-300.gpx",
		"GMT-5 Austin 300-400.gpx",
		"GMT-5 Austin 400-500.gpx",
		"GMT-5 Austin 500-573.gpx",
		"GMT-5 Austin.gpx",
		"GMT-5 AustinTX1-573stops.gpx",
		"GMT-5 Chicago 1-100.gpx",
		"GMT-5 Chicago 1.gpx",
		"GMT-5 Chicago 100-200.gpx",
		"GMT-5 Chicago 2.gpx",
		"GMT-5 Chicago 200-300.gpx",
		"GMT-5 Chicago 300-400.gpx",
		"GMT-5 Chicago 400-500.gpx",
		"GMT-5 Chicago 500-600.gpx",
		"GMT-5 Chicago 600-700.gpx",
		"GMT-5 Chicago 700-800.gpx",
		"GMT-5 Chicago All 2.gpx",
		"GMT-5 Chicago1 - 391 pokestops.gpx",
		"GMT-5 Chicago2 - 445 pokestops.gpx",
		"GMT-5 Ecuador CLuster.gpx",
		"GMT-5 HoustonTX1-398stops.gpx",
		"GMT-5 Mexico City-437stops.gpx",
		"GMT-5 OklahomaCity-404stops.gpx",
	},
	"GMT-6": {
		"GMT-6 Alberquerque 1-110.gpx",
		"GMT-6 Alberquerque 111-220.gpx",
		"GMT-6 Alberquerque 221-337.gpx",
		"GMT-6 Alberquerque.gpx",
		"GMT-6 Albuquerque NewMexico-337stops.gpx",
		"GMT-6 Colorado Springs 1-100.gpx",
		"GMT-6 Colorado Springs 100-200.gpx",
		"GMT-6 Colorado Springs 200-300.gpx",
		"GMT-6 Colorado Springs 300-400.gpx",
		"GMT-6 Colorado Springs.gpx",
		"GMT-6 ColoradoSprings1-380stops.gpx",
		"GMT-6 DenverCO-696stops.gpx",
		"GMT-6 Edmonton-356stops.gpx",
		"GMT-6 PuebloCOriverwalk-132stops.gpx",
	},
	"GMT-7": {
		"GMT-7 Caesar’sPalaceVegasNevada-457stops.gpx",
		"GMT-7 DisneyCA-331stops.gpx",
		"GMT-7 Disneyland.gpx",
		"GMT-7 GoldenGateParkSanFran-447stops.gpx",
		"GMT-7 Long BeachCA-534stops.gpx",
		"GMT-7 NianticHQSanFran-421stops.gpx",
		"GMT-7 P39SanFran-421stops.gpx",
		"GMT-7 Portland - 454 pokestops.gpx",
		"GMT-7 Portland 1-100.gpx",
		"GMT-7 Portland 100-200.gpx",
		"GMT-7 Portland 200-300.gpx",
		"GMT-7 Portland 300-400.gpx",
		"GMT-7 Portland.gpx",
		"GMT-7 RinconParkSanFran-407stops.gpx",
		"GMT-7 San Francisco - 588 pokestops.gpx",
		"GMT-7 San Francisco 1-100.gpx",
		"GMT-7 San Francisco 100-200.gpx",
		"GMT-7 San Francisco 200-300.gpx",
		"GMT-7 San Francisco 300-400.gpx",
		"GMT-7 San Francisco 500-600.gpx",
		"GMT-7 San Francisco.gpx",
		"GMT-7 Seattle - 475 pokestops.gpx",
		"GMT-7 Seattle 1-100.gpx",
		"GMT-7 Seattle 100-200.gpx",
		"GMT-7 Seattle 200-300.gpx",
		"GMT-7 Seattle 300-400.gpx",
		"GMT-7 Seattle 400-500.gpx",
		"GMT-7 Seattle.gpx",
		"GMT-7 UnionSqSanFran-500stops.gpx",
	},
}
