package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/spoc/internal/collector"
	"github.com/charmingruby/spoc/internal/shared/port"
	"github.com/charmingruby/spoc/internal/tm1"
	"github.com/charmingruby/spoc/internal/tm2"
	"github.com/charmingruby/spoc/internal/tm2/usecase"
	"github.com/robfig/cron/v3"
)

func main() {
	fetchers := []port.Fetcher{}

	tm1 := tm1.NewFetcher(tm1.Config{
		APIKey:                      "2",
		ShouldSimulateAuthError:     false,
		ShouldSimulateRelatoryError: false,
	})

	tm2 := tm2.NewFetcher(usecase.Config{
		ShouldSimulateRelatoryError: false,
	})

	fetchers = append(fetchers, tm1, tm2)

	collector := collector.New()

	job, err := collectorJob(collector, fetchers)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	job.Stop()
}

func collectorJob(collector *collector.Collector, fetchers []port.Fetcher) (*cron.Cron, error) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	c := cron.New(cron.WithSeconds(), cron.WithLocation(location))

	_, err = c.AddFunc("*/5 * * * * *", func() {
		if errs := collector.Run(fetchers); errs != nil {
			for _, err := range errs {
				println(err.Error())
			}
		}
	})

	if err != nil {
		return nil, err
	}

	c.Start()

	return c, nil
}
