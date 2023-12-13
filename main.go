package main

import (
	"flag"
	"log"
	"ytdlp_mult_download/yt_dlp"
)

func main() {
	var url, CookieFromBrowser string
	flag.StringVar(&url, "url", "", "Enter the url")
	flag.StringVar(&CookieFromBrowser, "cookie_browser", "", "Enter where browser to get cookies from")

	//TODO ... You can add the rest yourself

	flag.Parse()
	yt := &yt_dlp.Ytdlp{
		Url:               url,
		CookieFromBrowser: CookieFromBrowser,
	}
	yt.Init()
	err := yt.MultiThreadedDownload()
	if err != nil {
		// hand err
		log.Fatal(err)
		return
	}
}
