package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// set timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to target URL
	var price string
	err := chromedp.Run(ctx, chromedp.Navigate("https://www.coindesk.com/price/bitcoin"))
	if err != nil {
		log.Fatal(err)
	}

	// wait for target element to load
	err = chromedp.Run(ctx, chromedp.WaitVisible("#export-chart-element > div > div > span"))
	if err != nil {
		log.Fatal(err)
	}

	// extract target element from the page
	var nodeID cdp.NodeID
	err = chromedp.Run(ctx, dom.QuerySelector("#export-chart-element > div > div > span", &nodeID))
	if err != nil {
		log.Fatal(err)
	}

	// get the value of the target element
	err = chromedp.Run(ctx, dom.GetOuterHTML(nodeID, &price))
	if err != nil {
		log.Fatal(err)
	}

	// print the value to the console
	fmt.Println(price)
}
