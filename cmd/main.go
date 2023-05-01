package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	// 上位行をコメントアウトすることで下位行を有効化できます

	// level0: chromeのインスタンス作成
	//ctx, cancel := chromedp.NewContext(context.Background()) /*
	// level0-debug1: ログあり でインスタンス作成
	//ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf)) /*
	// level0-debug2: no headless でインスタンス作成
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	// level1: ページにアクセス
	fmt.Println("Ctrl+C で強制停止できます(Windows)。全く動かない時など、お試しあれ。")
	fmt.Println("なお、どの工程も読み込みに時間がかかります。最長1分程度です。")
	fmt.Println("アクセス開始………")
	if err := chromedp.Run(ctx, chromedp.Navigate("https://forms.gle/2iPTW6X4XjHCu4ar7")); err != nil {
		cancel()
		log.Fatal("err1: Failed login")
	}
	fmt.Println("アクセス完了")

	Login(ctx, cancel)
	FillOut(ctx, cancel)

	fmt.Println("fin!")
}
