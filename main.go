package main

import (
	"context"
	//"errors"
	"fmt"
	"log"
	"os"
	//"strconv"
	//"time"

	"github.com/chromedp/cdproto/input"
	//"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"

	"github.com/manifoldco/promptui"
)

func main() {
	// 上位行をコメントアウトすることで下位行を有効化できます

	// level0: chromeのインスタンス作成
	//ctx, _ := chromedp.NewContext(context.Background()) /*
	// level0-debug1: ログあり でインスタンス作成
	//ctx, _ := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf)) /*
	// level0-debug2: no headless でインスタンス作成
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	//*/
	//defer cancel()

	// level1: Access the page
	fmt.Println("Ctrl+C で強制停止できます(Windows)。全く動かない時など、お試しあれ。")
	if err := chromedp.Run(ctx, chromedp.Navigate("https://forms.gle/2iPTW6X4XjHCu4ar7")); err != nil {
		log.Fatal("err1: Failed login")
	}

	// level2: login
	login(ctx)

	// level3: 遷移チェック
	/* fmt.Println("ロード中………")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtMaRePassword"]/font`, chromedp.AtLeast(0)),
	); err == nil {
		log.Fatal("err2: Transition check fails")
	}
	fmt.Println("ロード完了")
	*/
}

func login(ctx context.Context) {
	var mailAddress string
	fmt.Printf("sを含めた学籍番号: ")
	fmt.Scan(&mailAddress)
	mailAddress = mailAddress + "@ga.ariake-nct.ac.jp"
	mailPasswd := passwdInputer("統合認証のパスワード: ")

	fmt.Println("ログイン処理開始………")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="identifierId"]`, chromedp.NodeVisible),
		input.InsertText(mailAddress),
		chromedp.Click(`#identifierNext > div > button > div.VfPpkd-RLmnJb`, chromedp.NodeVisible),
		chromedp.Sleep(1000000), // 1000000 nano sec == 1sec
		chromedp.Click(`//*[@id="password"]`, chromedp.NodeVisible),
		input.InsertText(mailPasswd),
		chromedp.Click(`#passwordNext > div > button > div.VfPpkd-RLmnJb`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err1: Failed login")
	}
	fmt.Println("ログイン完了")
}

func passwdInputer(labelMessage string) string {
	fmt.Println("赤色の x が付く場合がありますが、気にしないでください。")
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    labelMessage,
		Validate: validate,
		Mask:     '*',
	}

	passwd, err := prompt.Run()
	if err != nil {
		log.Fatal("err@passwdInputer: Failed to run prompt")
	}

	return passwd
}

func debugURL(ctx context.Context) {
	var url string
	if err := chromedp.Run(ctx,
		chromedp.Location(&url),
	); err != nil {
		log.Fatal("err@debugURL: Failed to location url")
	}
	fmt.Printf("debugURL: %s\n", url)
}

func debugPic(ctx context.Context) {
	var buf []byte
	// debug level1: スクショ撮影
	if err := chromedp.Run(ctx,
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		log.Fatal("err@debugPic-1: Failed to capture a screenshot")
	}
	// debug level2: スクショ出力
	if err := os.WriteFile(
		"fullScreenshot.png", buf, 0o644,
	); err != nil {
		log.Fatal("err@debugPic-2: Failed to output the screenshot")
	}
}
