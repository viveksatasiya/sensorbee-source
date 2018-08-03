package tickers

import (
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type Ticker struct {
	interval time.Duration
}

func (t *Ticker) GenerateStream(ctx *core.Context, w core.Writer) error {
	var cnt int64
	for ; ; cnt++ {
		tuple := core.NewTuple(data.Map{"tick": data.Int(cnt)})
		if err := w.Write(ctx, tuple); err != nil {
			return err
		}
		time.Sleep(t.interval)
	}
	return nil
}

func (t *Ticker) Stop(ctx *core.Context) error {
	return nil
}

func CreateTicker(ctx *core.Context,
	ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 1 * time.Second
	if v, ok := params["interval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&Ticker{
		interval: interval,
	}), nil
}
