package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CampusLogin struct {
	UserId      string
	Password    string
	Service     string
	QueryString string
}

func GetQueryString() (string, error) {
	url := "http://61.135.169.121/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fomrmatRes := strings.Split(strings.Split(string(body), "?")[1], "'</")[0]
	if len(fomrmatRes) > 511 {
		return "", &GetQueryStringError{"already connect to campus network", "ok"}
	}
	return fomrmatRes, nil
}

func GetCampusNetService(userId string, queryString string) (string, error) {

	url_ := "http://172.30.0.11/eportal/userV2.do?method=getServices"
	method := "POST"

	// 这个校园网是真的奇怪
	// payload := strings.NewReader(url.QueryEscape("username=" + userId + "&search=?" + queryString))
	payload := strings.NewReader("username=" + userId + "&search=" + url.QueryEscape("?"+queryString))

	client := &http.Client{}
	req, err := http.NewRequest(method, url_, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Origin", "http://172.30.0.11")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "http://172.30.0.11/eportal/index.jsp?"+queryString)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.47")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func LogoutCampusNetWork(userIndex string) (string, error) {
	url := "http://172.30.0.11/eportal/InterFace.do?method=logout"
	method := "POST"

	payload := strings.NewReader("userIndex="+userIndex)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func ConnectCampusNetWork(cl *CampusLogin) (ConnectCampusNetWorkResp, error) {

	url_ := "http://172.30.0.11/eportal/InterFace.do?method=login"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("userId=%s&password=%s&service=%s&queryString=%s&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false",
		cl.UserId,
		cl.Password,
		url.QueryEscape(url.QueryEscape(cl.Service)),
		url.QueryEscape(cl.QueryString),
	))

	client := &http.Client{}
	req, err := http.NewRequest(method, url_, payload)

	if err != nil {
		return ConnectCampusNetWorkResp{}, err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Origin", "http://172.30.0.11")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "http://172.30.0.11/eportal/index.jsp?"+cl.QueryString)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.47")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		return ConnectCampusNetWorkResp{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ConnectCampusNetWorkResp{}, err
	}
	var resp ConnectCampusNetWorkResp
	err = json.Unmarshal(body, &resp)
	if (err != nil){
		return ConnectCampusNetWorkResp{}, nil
	}
	return resp, nil
}
