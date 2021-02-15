package main

import (
	"go.uber.org/zap"
)

// init zap
var (
	log *zap.Logger
)

func main() {

	log, _ = zap.NewProduction()
	defer log.Sync()

	Router()
}