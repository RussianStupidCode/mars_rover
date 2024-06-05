package imitation

import (
	"context"
	"testing"

	"rover/pkg/rovers"
)

// хелпер функция для упрощения тестов
func pathToMove(path string, r *roverImitation) error {
	var err error = nil

	ctx := context.Background()

	for _, p := range path {
		switch p {
		case 'F':
			err = r.MoveForward(ctx)
		case 'B':
			err = r.MoveBackward(ctx)
		case 'R':
			err = r.TurnRight(ctx)
		case 'L':
			err = r.TurnLeft(ctx)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func TestMoveImitation(t *testing.T) {
	type want struct {
		x       int
		y       int
		d       rovers.Direction
		isError bool
	}

	type args struct {
		rover *roverImitation
		path  string
	}

	cases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "move test 1",
			args: args{
				rover: &roverImitation{x: 0, y: 0, direction: rovers.Direction_N, conn: "test:123456", errPercent: 0},
				path:  "FFFFF",
			},
			want: want{
				x:       0,
				y:       5,
				d:       rovers.Direction_N,
				isError: false,
			},
		},
		{
			name: "move test 2",
			args: args{
				rover: &roverImitation{x: 0, y: 0, direction: rovers.Direction_N, conn: "test:123456", errPercent: 0},
				path:  "RFFF",
			},
			want: want{
				x:       3,
				y:       0,
				d:       rovers.Direction_E,
				isError: false,
			},
		},
		{
			name: "move test 3",
			args: args{
				rover: &roverImitation{x: 0, y: 0, direction: rovers.Direction_N, conn: "test:123456", errPercent: 1},
				path:  "RFFF",
			},
			want: want{
				x:       3,
				y:       0,
				d:       rovers.Direction_E,
				isError: true,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := c.args.rover

			err := pathToMove(c.args.path, r)
			if (err != nil) != c.want.isError {
				t.Errorf("err = %v, wantErr = %v", err, c.want.isError)
			}

			if c.want.isError {
				return
			}

			if r.x != c.want.x || r.y != c.want.y || r.direction != c.want.d {
				t.Errorf("X = %d wantX =  %d, Y = %d, wantY = %d, direction = %d wantD = %d", r.x, c.want.x, r.y, c.want.y, r.direction, c.want.d)
			}
		})
	}
}
