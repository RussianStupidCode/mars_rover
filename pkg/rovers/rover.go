package rovers

import (
	"context"
	"time"
)

type Direction int

const (
	Direction_N Direction = iota
	Direction_S
	Direction_W
	Direction_E
)

//go:generate mockery --name Rover
type Rover interface {
	MoveForward(ctx context.Context) error
	MoveBackward(ctx context.Context) error
	TurnLeft(ctx context.Context) error
	TurnRight(ctx context.Context) error
	Ping(ctx context.Context) (time.Duration, error)
	Position(ctx context.Context) (int, int, error)
	Direction(ctx context.Context) (Direction, error)
	Close()
}
