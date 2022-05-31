package main

//func stream(ctx context.Context, logger *logrus.Logger, client *twitter.Client) error {
//	if ctx == nil {
//		ctx = context.Background()
//	}
//	if logger == nil {
//		logger = logrus.New()
//	}
//
//	opts := twitter.TweetSampleStreamOpts{}
//	//opts := twitter.TweetSearchStreamOpts{}
//
//	tweetStream, err := client.TweetSampleStream(context.Background(), opts)
//	//tweetStream, err := client.TweetSearchStream(ctx, opts)
//	if err != nil {
//		return errors.Errorf("tweet sample callout error: %v", err)
//	}
//	defer tweetStream.Close()
//	for {
//		select {
//		case tm := <-tweetStream.Tweets():
//			tmb, err := json.Marshal(tm)
//			if err != nil {
//				logger.Warnf("error decoding tweet message %v", err)
//
//			}
//			logger.Infof(fmt.Sprintf("tweet: %s\n\n", string(tmb)))
//
//		case sm := <-tweetStream.SystemMessages():
//			smb, err := json.Marshal(sm)
//			if err != nil {
//				logger.Warnf("error decoding system message %v", err)
//			}
//			logger.Infof("system: %s\n\n", string(smb))
//
//		case strErr := <-tweetStream.Err():
//			logger.Warnf("error in revceived stream %v", strErr)
//
//		default:
//		}
//		if tweetStream.Connection() == false {
//			return errors.New("connection lost")
//		}
//	}
//}
