package main

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

//const zabbixprdurl = "http://172.27.49.133/zabbix/index.php"
//const zabbixstressurl = "http://172.27.100.42/zabbix/index.php"

func main() {
	zabbix := new(Zabbix)
	zabbix.SetupEnv()
	driver := agouti.ChromeDriver(agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{
				"headless",
				"disable-gpu",
				"no-sandbox",
			},
		},
	}))

	err := driver.Start()
	if err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	zabbix.Page, err = driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	fmt.Printf("sssss:%v", zabbix.Page)

	zabbix.Login()

}
