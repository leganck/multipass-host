package serv

import (
	"context"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"net/http"
	"time"
)

var (
	s        service.Service
	logger   service.Logger
	httpServ *http.Server
)

type program struct{}

func (p *program) Start(service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	_ = logger.Info("StartServer, port = 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "httpServ time", time.Now())
	})
	httpServ = &http.Server{Addr: ":8080"}
	_ = httpServ.ListenAndServe()
}

func (p *program) Stop(service.Service) error {
	if httpServ != nil {
		_ = httpServ.Shutdown(context.Background()) // Go 1.8+
	}
	_ = logger.Info("StopServer")
	return nil
}
func GetServiceAndLogger() (service.Service, service.Logger) {
	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	options["DelayedAutoStart"] = true
	svcConfig := &service.Config{
		Name:        "Multipass Host",
		DisplayName: "Multipass Host Manager",
		Description: "This is Multipass Host Manager service.",
		Option:      options,
	}

	prg := &program{}
	var err error
	s, err = service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()
	return s, logger
}

func Interactive() bool {
	return service.Interactive()

}

func Control(s service.Service, action string) error {
	return service.Control(s, action)
}
