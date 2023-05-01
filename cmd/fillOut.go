package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func FillOut(ctx context.Context, cancel context.CancelFunc) {
	fmt.Println("フォームの入力開始")

	// 最初のフォームを入力
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div/span`),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[1]/div[1]/label/div/div[2]/div/span`),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[2]/div/div/div[2]/div/div[1]/div[1]/div[1]/span`),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[2]/div/div/div[2]/div/div[2]/div[3]/span`),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div/span`),
	); err != nil {
		log.Fatal("err2@FillOut: Failed to operate the login button")
	}
	/*
		if err := chromedp.Run(ctx,
			chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[1]/div[1]/label/div/div[2]/div/span`, chromedp.NodeVisible),
			chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[2]/div/div/div[2]/div/div[1]/div[1]/div[1]/span`, chromedp.NodeVisible),
			chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[2]/div[2]/div/div/div[2]/div/div[2]/div[3]`, chromedp.NodeVisible),
			chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div/span/span`, chromedp.NodeVisible),
		); err != nil {
			log.Fatal("err2@FillOut: Failed to operate the login button")
		}
	*/
	fmt.Println("1st ok")

	var pageTransitionCheck string
	// フォーム遷移確認1
	time.Sleep(3 * time.Second)
	if err := chromedp.Run(ctx,
		chromedp.Text(`//*[@id="i1"]/span[1]`, &pageTransitionCheck, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err3@FillOut: Failed to select text to confirm page transition")
	}
	if pageTransitionCheck != "体調はどうですか？" {
		log.Fatal("err4@FillOut: Failed to load on 1st form")
	}
	fmt.Println("1st check ok")

	// 2つ目のフォームを入力
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div/span`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div[2]/span`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err5@FillOut: Failed to operate the login button")
	}
	fmt.Println("2nd ok")

	// フォーム遷移確認2
	time.Sleep(3 * time.Second)
	if err := chromedp.Run(ctx,
		chromedp.Text(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[2]/div[1]/div[2]/span/span`, &pageTransitionCheck, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err6@FillOut: Failed to select text to confirm page transition")
	}
	if pageTransitionCheck != "送信" {
		log.Fatal("err7@FillOut: Failed to load on 2nd form")
	}
	fmt.Println("2nd check ok")

	// 3つ目のフォームを入力
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[2]/div[1]/div[2]/span`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err8@FillOut: Failed to operate the login button")
	}

	// フォーム遷移確認3
	time.Sleep(3 * time.Second)
	if err := chromedp.Run(ctx,
		chromedp.Text(`body > div.Uc2NEf > div:nth-child(2) > div.RH5hzf.RLS9Fe > div > div.vHW8K`, &pageTransitionCheck, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal("err6@FillOut: Failed to select text to confirm page transition")
	}
	if pageTransitionCheck != "回答を記録しました。<br><br>引き続き、寮生はこちらへ。（寮生以外の学生は関係ありません）<br>　　　↓↓↓↓<br>*****  寮生へ（寮務主事室より）  *****<br>昨日の行動履歴について、Googleフォーム(下記URL)より報告してください。<br>■行動記録フォーム（回答受付 5:30～13:00）<br>  <a href=\"https://www.google.com/url?q=https://forms.gle/fHirT285dfDvZtLT9&amp;sa=D&amp;source=editors&amp;ust=1682829601701378&amp;usg=AOvVaw0SBYnSxQoJ8tBoLCmuamzj\">https://forms.gle/fHirT285dfDvZtLT9</a><br>" {
		fmt.Println(pageTransitionCheck)
		log.Fatal("err7@FillOut: Failed to load on 3rd form")
	}
	fmt.Println("3rd check ok")

	fmt.Println("入力完了")
}
