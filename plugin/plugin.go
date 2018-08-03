package plugin

import (
	"nats-source/tickers"

	"gopkg.in/sensorbee/sensorbee.v0/bql"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("ticker",
		bql.SourceCreatorFunc(tickers.CreateTicker))
}
