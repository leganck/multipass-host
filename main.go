package main

import (
	"github.com/kardianos/service"
	"log"
	"multipass-host/serv"
	"os"
)

var (
	s      service.Service
	logger service.Logger
)

func init() {
	s = serv.S
	logger = serv.Logger
}

func main() {
	var err error
	if service.Interactive() && len(os.Args) > 1 {
		err := service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = s.Run()
		if err != nil {
			_ = logger.Error(err)
		}

	}

}
