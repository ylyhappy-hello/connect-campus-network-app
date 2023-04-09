package services

import (
	"fmt"
	"strings"

	"github.com/ylyhappy-hello/connect-campus-network-app/apis"
)

var cache apis.CampusLogin

func FirstLogin(userId string, password string) (string, error) {
	fmt.Printf("userId = %s, password = %s", userId, password)
	queryString, err := apis.GetQueryString()
	if (err != nil){
		return  "", err
	}

	service, err := apis.GetCampusNetService(userId, queryString)

	if (err != nil){
		return "", err
	}

	service, err = formatCampusNetService(service)

	if (err != nil){
		return "", err
	}

	cache = apis.CampusLogin{
		Service: service,UserId:userId,Password:password,QueryString:queryString}

	_, err = apis.ConnectCampusNetWork(&cache)
	if (err != nil){
		return "", err
	}
	return "ok", nil
}

func GetUserIndex() string {
	if (cache != apis.CampusLogin{}){
		res, _:= apis.ConnectCampusNetWork(&cache)
		return res.UserIndex
	}
	return ""
}

func LogoutCampusNetWork() {
	userIndex := GetUserIndex()
	apis.LogoutCampusNetWork(userIndex)
}
func userOnline() bool {
	// 这里利用校园网， 如果已经登陆上， 会返回百度的页面， 这GetQueryStirng会返回错误， err 则不为空
	_, err := apis.GetQueryString()
	return err != nil
}

func formatCampusNetService(service string) (string, error){
	fmt.Printf("service = %s", service)
	if (len(service) > 9){
		ser := strings.Split(service, "@")
		if (ser[0] == "校园网"){
			return ser[1], nil
		}
		return  ser[0], nil
	} else {
		if (service != "教师组"){
			return "", &FormatCampusNetServiceError{"service string is invalid, can not format", "ok"}
		}
		return service, nil
	}
}

func FastLogin(){
	if (cache != apis.CampusLogin{}) {
		apis.ConnectCampusNetWork(&cache)
	}
}