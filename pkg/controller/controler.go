package controller

import (
	"context"

	"rover/pkg/rovers"
)

// TODO сделать оптимизацию/минимизацию пути перед реальным вызовом машины (пример "RL" -> "") + возможно
func pathOptimazer(path string) (string, error) {
	return path, nil
}

func ChangePosition(ctx context.Context, rover rovers.Rover, path string) error {
	path, err := pathOptimazer(path)
	if err != nil {
		return err
	}

	for _, p := range path {
		switch p {
		case 'F':
			err = rover.MoveForward(ctx)
		case 'B':
			err = rover.MoveBackward(ctx)
		case 'R':
			err = rover.TurnRight(ctx)
		case 'L':
			err = rover.TurnLeft(ctx)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
