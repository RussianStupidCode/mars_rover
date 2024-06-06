package controller

import (
	"context"
	"fmt"
	"strings"

	"rover/pkg/rovers"
)

var AllowedRunes = map[rune]struct{}{
	'F':{},
	'B':{},
	'R':{},
	'L':{},
}


// TODO сделать оптимизацию/минимизацию пути перед реальным вызовом машины
func pathOptimazer(path string) (string, error) {
	path = strings.ToUpper(strings.TrimSpace(path))
	
	refined := make([]rune, 0)
	for _, r  := range path {
		if _, ok := AllowedRunes[r]; ok  {
			refined = append(refined, r)
		}
	}

	if len(refined) == 0 {
		return "", fmt.Errorf("empty path")
	}

	return string(refined), nil
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
