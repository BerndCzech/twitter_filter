package main

import (
	"context"
	"encoding/json"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type rule struct {
	name string
	// for query examples,
	// see: https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule#examples
	query string
}

func rules(ctx context.Context,
	logger *logrus.Logger,
	client *twitter.Client,
	r rule,
	dryRun bool,
) (string, error) {

	streamRule := twitter.TweetSearchStreamRule{
		Value: r.query,
		Tag:   r.name,
	}

	// check if tag is already there
	var (
		resp any
		err  error
	)
	resp, err = client.TweetSearchStreamRules(context.Background(), []twitter.TweetSearchStreamRuleID{})
	if err != nil {
		return "", errors.Errorf("tweet search stream rule callout error: %v", err)
	}

	ruleID := findRuleID(resp, r)
	if ruleID == "" {
		resp, err = client.TweetSearchStreamAddRule(ctx, []twitter.TweetSearchStreamRule{streamRule}, dryRun)
		if err != nil {
			return "", err
		}
		ruleID = findRuleID(resp, r)
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		//return errors.Wrap(err,"could not marshall twitter rule response")
		logger.Warnf("could not marshall twitter rule response %+v", err)
		return ruleID, nil
	}

	logger.Infof("applied rules %v", string(enc))

	return ruleID, nil
}

func findRuleID(resp any, r rule) string {

	var rules []*twitter.TweetSearchStreamRuleEntity
	switch v := resp.(type) {
	case *twitter.TweetSearchStreamAddRuleResponse:
		rules = v.Rules
	case *twitter.TweetSearchStreamRulesResponse:
		rules = v.Rules
	default:
		// should not happen
		return ""
	}

	var ruleID string
	for _, rule := range rules {
		if rule.Tag == r.name || rule.Value == r.query {
			ruleID = string(rule.ID)
			break
		}
	}
	return ruleID
}
