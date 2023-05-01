package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func DebugURL(ctx context.Context) {
	var url string
	if err := chromedp.Run(ctx,
		chromedp.Location(&url),
	); err != nil {
		log.Fatal("err@debugURL: Failed to location url")
	}
	fmt.Printf("debugURL: %s\n", url)
}

func DebugPic(ctx context.Context) {
	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		log.Fatal("err1@debugPic: Failed to capture a screenshot")
	}

	if err := os.WriteFile(
		"fullScreenshot.png", buf, 0o644,
	); err != nil {
		log.Fatal("err2@debugPic: Failed to output the screenshot")
	}
}
