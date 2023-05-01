package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
	"github.com/manifoldco/promptui"
	"log"
	"strings"
	"time"
)

func Login(ctx context.Context, cancel context.CancelFunc) {
	fmt.Println("ログイン開始………")

	// ログイン処理
	fmt.Printf("sを含めた学籍番号: ")
	var mailAddress string
	fmt.Scan(&mailAddress)
	mailAddress = mailAddress + "@ga.ariake-nct.ac.jp"
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="identifierId"]`, chromedp.NodeVisible),
		input.InsertText(mailAddress),
		chromedp.Click(`//*[@id="identifierNext"]/div/button/div[3]`, chromedp.NodeVisible),
	); err != nil {
		cancel()
		log.Fatal("err1@Login: Failed to enter email address")
	}

	// メールアドレスの検証
	time.Sleep(1 * time.Second)
	var pageTransitionCheck string
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`//*[@id="headingText"]/span`),
		chromedp.Location(&pageTransitionCheck),
	); err != nil {
		cancel()
		log.Fatal("err2@Login: Failed in page transition confirmation process")
	}
	if !strings.Contains(pageTransitionCheck, "https://accounts.google.com/v3/signin/challenge/pwd?TL=") {
		cancel()
		log.Fatal("err3@Login: Email address is incorrect")
	}

	// パスワード入力、ログイン実行
	mailPasswd := passwdInputer("統合認証のパスワード")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="password"]/div[1]/div/div[1]/input`, chromedp.NodeVisible),
		input.InsertText(mailPasswd),
		chromedp.Click(`//*[@id="passwordNext"]/div/button/div[3]`, chromedp.NodeVisible),
	); err != nil {
		cancel()
		log.Fatal("err4@Login: Failed to operate the login button")
	}

	// パスワードの検証
	time.Sleep(3 * time.Second)
	if err := chromedp.Run(ctx,
		chromedp.Location(&pageTransitionCheck),
	); err != nil {
		cancel()
		log.Fatal("err5@Login: Failed in page transition confirmation process")
	}
	if strings.Contains(pageTransitionCheck, "https://accounts.google.com/v3/signin/challenge/pwd?TL=") {
		cancel()
		log.Fatal("err6@Login: Password is incorrect")
	}

	// ページ遷移の確認
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`body > div.Uc2NEf > div:nth-child(2) > div:nth-child(2) > div.I3zNcc.yF4pU > a > img`),
		chromedp.Location(&pageTransitionCheck),
	); err != nil {
		cancel()
		log.Fatal("err7@Login: Failed in page transition confirmation process")
	}
	if !strings.Contains(pageTransitionCheck, "https://docs.google.com/forms/") {
		cancel()
		log.Fatal("err8@Login: Failed to load on password input page")
	}

	// ついでに、受け付けいているかも確認
	if strings.Contains(pageTransitionCheck, "/closedform") {
		fmt.Println("本日の報告の受け付けは終了しました。\n回答受付時間は、5時00分～13時00分です。明日は遅れずに報告してください。")
		cancel()
		log.Fatal("err9@Login: Form is closed")
	}

	// 念のため1秒待つ
	time.Sleep(1 * time.Second)
	fmt.Println("ログイン完了")
}

func passwdInputer(labelMessage string) string {
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
		log.Fatal("err0@passwdInputer: Failed to run prompt")
	}

	return passwd
}
