package poller

import (
	"github.com/coherentopensource/go-service-framework/constants"
	"time"
)

type Config struct {
	Blockchain  constants.Blockchain `env:"BLOCKCHAIN,required"`
	BatchSize   int                  `env:"BATCH_SIZE" envDefault:"100"`
	ReorgDepth  int                  `env:"REORG_DEPTH" envDefault:"8"`
	HttpRetries int                  `env:"HTTP_RETRIES" envDefault:"10"`
	SleepTime   time.Duration        `env:"POLLER_SLEEP_TIME" envDefault:"12s"`
	Tick        time.Duration        `env:"POLLER_TICK_DURATION" envDefault:"1s"`
	AutoStart   bool                 `env:"POLLER_AUTO_START" envDefault:"false"`
}
