package main

import (
	"fmt"

	"github.com/sclevine/agouti"
)

func main() {
	err := Main()
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func Main() error {
	driver, err := startDriver()
	if err != nil {
		return err
	}
	defer driver.Stop()

	err = RunE2E(driver)

	return nil
}

func startDriver() (*agouti.WebDriver, error) {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--homepage=about:blank",
			"--disable-gpu",
			"--allow-insecure-localhost",
			"--no-first-run",
			"--no-default-browser-check",
			"--no-sandbox",
			// "--whitelisted-ips",
			"--window-size=1280,800",
		}),
		agouti.Debug,
	)
	err := driver.Start()
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func RunE2E(driver *agouti.WebDriver) error {
	page, err := driver.NewPage()
	if err != nil {
		return err
	}

	// サイトトップ
	err = page.Navigate("https://www.google.com/")
	if err != nil {
		return err
	}

	// 検索
	page.FindByName("q").Fill("Yappli")
	page.FindByName("q").Submit()

	// 検索結果
	page.Screenshot("./screenshot.png")

	return nil
}
