package main

import (
	"log"
	"multipass-host/serv"
	"os"
)

func main() {
	var s, logger = serv.GetServiceAndLogger()
	var err error
	if serv.Interactive() && len(os.Args) > 1 {
		err := serv.Control(s, os.Args[1])
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
