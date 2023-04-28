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

func LoginProcess(ctx context.Context, cancel context.CancelFunc) {
	fmt.Println("ログイン処理開始(時間がかかります)………")

	// クレデンシャルの入力
	fmt.Printf("sを含めた学籍番号: ")
	var mailAddress string
	fmt.Scan(&mailAddress)
	mailAddress = mailAddress + "@ga.ariake-nct.ac.jp"
	mailPasswd := passwdInputer("統合認証のパスワード")

	// ログイン処理
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="identifierId"]`, chromedp.NodeVisible),
		input.InsertText(mailAddress),
		chromedp.Click(`//*[@id="identifierNext"]/div/button/div[3]`, chromedp.NodeVisible),
	); err != nil {
		cancel()
		log.Fatal("err1@login: Failed login")
	}

	// メールアドレスの入力が正しいか(遷移しているか)確認
	time.Sleep(1 * time.Second)
	var url2CheckTransition string
	if err := chromedp.Run(ctx,
		//chromedp.WaitVisible(`body > div > div > div > div  > div`),
		chromedp.WaitVisible(`//*[@id="headingText"]/span`),
		chromedp.Location(&url2CheckTransition),
	); err != nil {
		cancel()
		log.Fatal("err2@login: Failed in page transition confirmation process")
	}
	if !strings.Contains(url2CheckTransition, "https://accounts.google.com/v3/signin/challenge/pwd?TL=") {
		cancel()
		log.Fatal("err3@login: Failed to load on email address input page")
	}

	// パスワード入力、ログイン実行
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="password"]/div[1]/div/div[1]/input`, chromedp.NodeVisible),
		input.InsertText(mailPasswd),
		chromedp.Click(`//*[@id="passwordNext"]/div/button/div[3]`, chromedp.NodeVisible),
	); err != nil {
		cancel()
		log.Fatal("err4@login: Failed to operate the login button")
	}

	// ログインしてページ遷移をしているか確認
	time.Sleep(1 * time.Second)
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`body > div.Uc2NEf > div:nth-child(2) > div.RH5hzf.RLS9Fe > div > div.pdLVYe.LgNcQe`),
		chromedp.Location(&url2CheckTransition),
	); err != nil {
		cancel()
		log.Fatal("err5@login: Failed in page transition confirmation process")
	}
	if !strings.Contains(url2CheckTransition, "https://docs.google.com/forms/") {
		cancel()
		log.Fatal("err6@login: Failed to load on password input page")
	}

	fmt.Println("ログイン処理終了")
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
