package main

import (
	"context"
	//"errors"
	"fmt"
	"log"
	"os"
	//"strconv"
	"time"

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

	// level1: Access the page
	fmt.Println("Ctrl+C で強制停止できます(Windows)。全く動かない時など、お試しあれ。")
	fmt.Println("アクセス開始(時間がかかります)………")
	if err := chromedp.Run(ctx, chromedp.Navigate("https://forms.gle/2iPTW6X4XjHCu4ar7")); err != nil {
		log.Fatal("err1: Failed login")
	}
	fmt.Println("アクセス完了")

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
	var mailAddress, checkPageTransition string
	fmt.Printf("sを含めた学籍番号: ")
	fmt.Scan(&mailAddress)
	mailAddress = mailAddress + "@ga.ariake-nct.ac.jp"
	mailPasswd := passwdInputer("統合認証のパスワード")

	fmt.Println("ログイン処理開始(時間がかかります)………")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="identifierId"]`, chromedp.NodeVisible),
		input.InsertText(mailAddress),
		chromedp.Click(`#identifierNext > div > button > div.VfPpkd-RLmnJb`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err1@login: Failed login")
	}
	time.Sleep(5 * time.Second)

	// メールアドレスの入力が正しいか(遷移しているか)確認
	if err := chromedp.Run(ctx,
		chromedp.Text(`#selectionc1`, &checkPageTransition, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err2@login: Failed in page transition confirmation process")
	}
	if checkPageTransition != "パスワードを表示する" {
		log.Fatal("err3@login: Failed to load on email address input page")
	}

	// パスワード入力、ログイン実行
	if err := chromedp.Run(ctx,
		input.InsertText(mailPasswd),
		chromedp.Click(`#passwordNext > div > button > div.VfPpkd-RLmnJb`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err4@login: Failed to operate the login button")
	}
	time.Sleep(5 * time.Second)
	fmt.Println("ok")

	if err := chromedp.Run(ctx, chromedp.Navigate("https://forms.gle/2iPTW6X4XjHCu4ar7")); err != nil {
		log.Fatal("err5@login: 遷移できんで")
	}
	/* // ログインしてページ遷移をしているか確認1
	if err := chromedp.Run(ctx,
		chromedp.Text(`//*[@id="yDmH0d"]/c-wiz/div/div[2]/div/div[1]/div/form/span/section[2]/div/div/div[1]/div[2]/div[2]/span`, &checkPageTransition, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err5@login: Failed in page transition confirmation process")
	}
	if checkPageTransition == "パスワードが正しくありません。入力し直してください。[パスワードをお忘れの場合] をクリックすると、再設定できます。" {
		log.Fatal("err6@login: Failed to load on password input page")
	}
	fmt.Println("ok")
	*/

	// ログインしてページ遷移をしているか確認2
	if err := chromedp.Run(ctx,
		chromedp.Text(`head > title`, &checkPageTransition, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err5@login: Failed in page transition confirmation process")
	}
	if checkPageTransition != "健康チェック報告" {
		log.Fatal("err6@login: Failed to load on password input page")
	}
	fmt.Println("ok")

	// 時間外かどうか調べる
	if err := chromedp.Run(ctx,
		chromedp.Text(`//*[@id="i1"]/span[1]`, &checkPageTransition, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err7@login: Failed to select text to confirm page transition")
	}
	if checkPageTransition != "今朝の体温   (平熱：普段の健康な時の体温，過去1か月程度の平均体温を目安に)" {
		log.Fatal("err8@login: Failed to load on 1st form")
	}
	fmt.Println("ok")

	// 最初のフォームを入力
	if err := chromedp.Run(ctx,
		input.InsertText(mailPasswd),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div/div/div/div[2]/div/div[1]/div[1]/div[1]`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div/div/div/div[2]/div/div[2]/div[3]/span`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div/span`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err9@login: Failed to operate the login button")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("ok")

	// フォーム遷移確認
	if err := chromedp.Run(ctx,
		chromedp.Text(`//*[@id="i1"]/span[1]`, &checkPageTransition, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err10@login: Failed to select text to confirm page transition")
	}
	if checkPageTransition != "体調はどうですか？" {
		log.Fatal("err11@login: Failed to load on 2nd form")
	}
	fmt.Println("ok")

	// 2つ目のフォームを入力
	if err := chromedp.Run(ctx,
		input.InsertText(mailPasswd),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[2]/div/div/div[2]/div/div/span/div/div[1]/label`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div[2]/span`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err9@login: Failed to operate the login button")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("ok")

	// 3つ目のフォームを入力
	if err := chromedp.Run(ctx,
		input.InsertText(mailPasswd),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[2]/div[1]/div[2]/span`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err9@login: Failed to operate the login button")
	}
	time.Sleep(3 * time.Second)

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
		log.Fatal("err1@debugPic: Failed to capture a screenshot")
	}
	// debug level2: スクショ出力
	if err := os.WriteFile(
		"fullScreenshot.png", buf, 0o644,
	); err != nil {
		log.Fatal("err2@debugPic: Failed to output the screenshot")
	}
}
