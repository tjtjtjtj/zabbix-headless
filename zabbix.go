package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/sclevine/agouti"
)

type Zabbix struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
	URL      string `json:"stressurl"`
	Page     *agouti.Page
}

func (z *Zabbix) SetupEnv() {
	file, err := os.Open("./conf/zabbix_env.json")
	if err != nil { // エラー処理
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(z)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*z)
}

func (z *Zabbix) Login() {
	fmt.Println("login start")
	fmt.Printf("page:%v", z.Page)

	err := z.Page.Navigate(z.URL)
	if err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	/*
		page_html, _ := page.HTML()
				if strings.Contains(page_html, "Unlock Jenkins") {
							}
	*/

	fmt.Println("login start2")
	html, _ := z.Page.HTML()
	fmt.Printf("html:%s", html)
	z.Page.Screenshot("/tmp/outputs/zabbix1.jpg")
	fmt.Println("login start3")

	userid := z.Page.FindByID("name")
	password := z.Page.FindByID("password")
	userid.Fill(z.UserID)
	password.Fill(z.Password)
	z.Page.Screenshot("/tmp/outputs/zabbix2.jpg")
	if err := z.Page.FindByID("enter").Click(); err != nil {
		log.Fatal("Failed to set password", err)
	}

	if err := z.Page.FindByLink("スクリーン").Click(); err != nil {
		log.Fatal("Failed to click", err)
	}

	if err := z.Page.FindByID("elementid").Select("testscreen"); err != nil {
		log.Fatal("Failed to click", err)
	}

	s := z.Page.FindByClass("screen_view").AllByName("img").At(0).String()

	fmt.Printf("imgstring %s", s)

	time.Sleep(10 * time.Second)

	z.Page.Screenshot("/tmp/outputs/zabbix3.jpg")

}
