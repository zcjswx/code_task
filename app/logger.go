package app

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewProduction()
	defer l.Sync() // flushes buffer, if any
	sugar := l.Sugar()
	logger = sugar
}
