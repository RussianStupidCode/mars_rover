package application

import (
	"context"
	"fmt"
	"os"

	"rover/pkg/clients/console"
	"rover/pkg/config"
	"rover/pkg/crons"
	"rover/pkg/logger"
	"rover/pkg/rovers/imitation"
)

func Start(ctx context.Context) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	log, err := logger.New(cfg.Logger, os.Stdout)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}

	// при надобности поменять на реальный объект
	rover, err := imitation.New(cfg.Rover)
	if err != nil {
		return fmt.Errorf("create rover: %w", err)
	}

	defer rover.Close()

	// пингует для проверки жизнеспособности
	checker := crons.NewHealthChecker(cfg.HealthChecker, log, rover)
	checker.RunJob(ctx)

	// консольный клиент
	app := console.New(rover, log)
	app.Run(ctx)

	<-ctx.Done()
	return nil
}
