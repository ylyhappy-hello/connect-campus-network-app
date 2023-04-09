package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kardianos/service"
	"github.com/ylyhappy-hello/connect-campus-network-app/apis"
	"github.com/ylyhappy-hello/connect-campus-network-app/services"
)

var logger service.Logger

type CompusNetWok struct{}

func (p *CompusNetWok) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *CompusNetWok) run() {
	fmt.Println("start")
	// apis.GetQueryString()

}
func (p *CompusNetWok) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main(){
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	fmt.Println(services.FirstLogin(os.Args[1], os.Args[2]))	
}

func ssss(){
	apis.GetQueryString()
	// apis.ConnectCampusNetWork()
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	svcConfig := &service.Config{
			Name:        "auto-connect-campus-network",
			DisplayName: "auto-connect-campus-network dispaly name",
			Description: "auto-connect-campus-network description",
	}

	prg := &CompusNetWok{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Install()
	if err != nil {
		logger.Error(err)
	}

}
