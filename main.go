package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Specification struct {
	Token string `envconfig:"TOKEN"`
	//Filter string `envconfig:"FILTER_STRING"`
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

/**
	In order to run, the user will need to provide the bearer token and the list of tweet ids.
**/
func main() {

	log := logrus.New()
	log.SetFormatter(
		&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		})
	log.SetOutput(os.Stdout)

	var cfg Specification
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("could not process env vars %v", err)
	}

	// construction
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := &twitter.Client{
		Authorizer: authorize{
			Token: cfg.Token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	ch := make(chan os.Signal)
	go func() {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ch:

			// shutdown on ctrl-c
			cancel()
		case <-ctx.Done():
			// shutdown for app reason
		}
	}()
	// app logic
	if err := run(ctx, log, client); err != nil {
		log.Fatal("shut down", err)
	}
}

func run(ctx context.Context, log *logrus.Logger, client *twitter.Client) error {
	// logic
	//  * Reset Rules
	//  * Stream
	r := rule{name: "sarstedt",
		// rules are case insensitive
		query: "#sarstedt"}
	_, err := rules(ctx, log, client, r, false)
	if err != nil {
		return errors.Errorf("could not add rules %+v", err)
	}

	if err := stream(ctx, log, client); err != nil {
		return errors.WithMessage(err, "terminated streaming")
	}

	return nil
}
