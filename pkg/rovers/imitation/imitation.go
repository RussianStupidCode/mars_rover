package imitation

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"rover/pkg/config"
	"rover/pkg/rovers"
)

const (
	N = rovers.Direction_N
	S = rovers.Direction_S
	W = rovers.Direction_W
	E = rovers.Direction_E

	MaxPingMS = 10000
)

type roverImitation struct {
	x          int
	y          int
	conn       string
	direction  rovers.Direction
	errPercent float32 // вероятность получить ошибку число от 0 до 1
}

func New(cfg config.Rover) (*roverImitation, error) {
	// стартовая позиция по умолчанию
	x := 1
	y := 1
	d := N

	if cfg.ConnectionURL == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	return &roverImitation{conn: cfg.ConnectionURL, x: x, y: y, direction: d, errPercent: float32(cfg.ErrPercent)}, nil
}

func (r *roverImitation) Close() {
	// pass
}

// вероятность любая, но чтобы имитировать случайные ошибки
func (r *roverImitation) createRandomError() error {
	n := rand.Float32()

	if n < r.errPercent {
		return fmt.Errorf("unexpected error")
	}

	return nil
}

func (r *roverImitation) MoveForward(ctx context.Context) error {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return randomErr
	}

	switch r.direction {
	case N:
		r.y = r.y + 1
	case S:
		r.y = r.y - 1
	case W:
		r.x = r.x - 1
	case E:
		r.x = r.x + 1
	}

	return nil
}

func (r *roverImitation) MoveBackward(ctx context.Context) error {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return randomErr
	}

	switch r.direction {
	case N:
		r.y = r.y - 1
	case S:
		r.y = r.y + 1
	case W:
		r.x = r.x + 1
	case E:
		r.x = r.x - 1
	}
	return nil
}

func (r *roverImitation) TurnLeft(ctx context.Context) error {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return randomErr
	}

	switch r.direction {
	case N:
		r.direction = W
	case S:
		r.direction = E
	case W:
		r.direction = S
	case E:
		r.direction = N
	}

	return nil
}

func (r *roverImitation) TurnRight(ctx context.Context) error {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return randomErr
	}

	switch r.direction {
	case N:
		r.direction = E
	case S:
		r.direction = W
	case W:
		r.direction = N
	case E:
		r.direction = S
	}

	return nil
}

func (r *roverImitation) Position(ctx context.Context) (int, int, error) {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return 0, 0, randomErr
	}

	return r.x, r.y, nil
}

func (r *roverImitation) Direction(ctx context.Context) (rovers.Direction, error) {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return rovers.Direction_N, randomErr
	}

	return r.direction, nil
}

func (r *roverImitation) Ping(ctx context.Context) (time.Duration, error) {
	randomErr := r.createRandomError()
	if randomErr != nil {
		return time.Duration(0), randomErr
	}

	n := rand.Intn(MaxPingMS)
	return time.Duration(n) * time.Millisecond, nil
}
