// btcmargin shows you a live view of your bitcoin margins
//
// Usage of btcmargin:
//  -amount=0: number of BTC
//  -key="": MtGox API key (will use GOX_KEY if not provided)
//  -secret="": MtGox secret key (will use GOX_SECRET if not provided)
//  -value=0: average price per bitcoin
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	mtgox "github.com/yanatan16/golang-mtgox"
)

var (
	buyValue = flag.Float64("value", 0.0, "average price per bitcoin")
	amount   = flag.Float64("amount", 0.0, "number of BTC")
	key      = flag.String("key", "", "MtGox API key (will use GOX_KEY if not provided)")
	secret   = flag.String("secret", "", "MtGox secret key (will use GOX_SECRET if not provided)")
)

func main() {
	flag.Parse()
	if *key == "" {
		*key = os.Getenv("GOX_KEY")
	}
	if *secret == "" {
		*secret = os.Getenv("GOX_SECRET")
	}
	if *key == "" || *secret == "" {
		fmt.Fprintln(os.Stderr, "Missing API Key and/or secret")
		flag.Usage()
		os.Exit(1)
	}
	if *buyValue == 0 || *amount == 0 {
		fmt.Println("Did you forget to supply -value or -amount?")
	}

	gox, err := mtgox.New(*key, *secret, "USD")
	if err != nil {
		log.Fatalln("Error creating api connection:", err)
	}

	c, err := gox.Subscribe("ticker")
	if err != nil {
		log.Fatalln("Error subscribing:", err)
	}

	for msg := range c {
		avg := msg["ticker"].(map[string]interface{})["last"].(map[string]interface{})
		if avg["value"] == nil {
			continue
		}
		var val float64
		fmt.Sscan(avg["value"].(string), &val)

		profitPer := val - *buyValue
		overallProfit := profitPer * *amount
		percentage := overallProfit / (*amount * *buyValue)

		fmt.Printf("$%.2f - %.2f%% - currently trading at $%f USD (spent:$%.2f, made:$%.2f)\n",
			overallProfit, percentage*100, val, *buyValue*(*amount), val*(*amount))
	}
}
