package crons

import (
	"context"
	"time"

	"rover/pkg/config"
	"rover/pkg/logger"
)

type Pinger interface {
	Ping(ctx context.Context) (time.Duration, error)
}

type HealthChecker struct {
	pinger Pinger
	cfg    config.HealthChecker
	log    *logger.Logger
}

func NewHealthChecker(cfg config.HealthChecker, log *logger.Logger, pinger Pinger) *HealthChecker {
	return &HealthChecker{
		pinger: pinger,
		cfg:    cfg,
		log:    log,
	}
}

func (c *HealthChecker) ping(ctx context.Context) {
	t, err := c.pinger.Ping(ctx)
	if err != nil {
		c.log.Error("ping error", err)
		return
	}

	c.log.Info("ping", "time", t)
}

func (c *HealthChecker) RunJob(ctx context.Context) {
	ticker := time.NewTicker(c.cfg.Period)

	go func() {
		// не ждать первого тика
		c.ping(ctx)

		for {
			select {
			case <-ticker.C:
				c.ping(ctx)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}
