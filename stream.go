package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func stream(ctx context.Context, logger *logrus.Logger, client *twitter.Client) error {
	if logger == nil {
		logger = logrus.New()
	}

	opts := twitter.TweetSearchStreamOpts{}

	tweetStream, err := client.TweetSearchStream(ctx, opts)
	if err != nil {
		return errors.Errorf("tweet sample callout error: %v", err)
	}
	defer tweetStream.Close()

	for {
		if tweetStream.Connection() == false {
			logger.Warn("client got disconnected")
			time.Sleep(time.Second)
			//tweetStream, err = client.TweetSearchStream(ctx, opts)
			//if err != nil {
			//	return errors.Errorf("tweet sample callout error: %v", err)
			//}
		}
		select {
		case tm := <-tweetStream.Tweets():
			tmb, err := json.Marshal(tm)
			if err != nil {
				logger.Warnf("error decoding tweet message %v", err)

			}
			logger.Infof(fmt.Sprintf("tweet: %s\n\n", string(tmb)))

		case sm := <-tweetStream.SystemMessages():
			smb, err := json.Marshal(sm)
			if err != nil {
				logger.Warnf("error decoding system message %v", err)
			}
			logger.Infof("system: %s\n\n", string(smb))

		case strErr := <-tweetStream.Err():
			logger.Warnf("error in revceived stream %v", strErr)

		case <-ctx.Done():
			return ctx.Err()

		default:
			//	continue
		}
	}
}
