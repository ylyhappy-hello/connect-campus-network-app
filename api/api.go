package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getQueryString() {
	client := &http.Client{}
	var data = strings.NewReader(`userId=%3CuserId%3E&password=%3Cpassword%3E&service=%3Cservice%3E&queryString=%3CqueryString%3E&operatorPwd=%3CoperatorPwd%3E&operatorUserId=%3CoperatorUserId%3E&validcode=%3Cvalidcode%3E&passwordEncrypt=%3CpasswordEncrypt%3E`)
	req, err := http.NewRequest("POST", "http://172.30.0.11/eportal/InterFace.do?method=<method>", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "<Accept>")
	req.Header.Set("Accept-Language", "<Accept-Language>")
	req.Header.Set("Cache-Control", "<Cache-Control>")
	req.Header.Set("Connection", "<Connection>")
	req.Header.Set("Origin", "<Origin>")
	req.Header.Set("Pragma", "<Pragma>")
	req.Header.Set("Referer", "<Referer>")
	req.Header.Set("User-Agent", "<User-Agent>")
	req.Header.Set("Cookie", "<Cookie>")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
