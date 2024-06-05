package controller

import (
	"context"
	"fmt"
	"testing"

	"rover/pkg/rovers/mocks"

	"github.com/stretchr/testify/mock"
)

func createError(flag bool, msg string) error {
	if flag {
		return fmt.Errorf(msg)
	}
	return nil
}

func TestChangePosition(t *testing.T) {
	type args struct {
		countR int
		countL int
		countB int
		countF int
		path   string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test changeposition 1",
			args:    args{countL: 1, countR: 1, countB: 1, countF: 1, path: "RLBF"},
			wantErr: false,
		},
		{
			name:    "Test changeposition 2",
			args:    args{countL: 0, countR: 1, countB: 0, countF: 0, path: "RLBF"},
			wantErr: true,
		},
		{
			name:    "Test changeposition 3",
			args:    args{countL: 1, countR: 1, countB: 1, countF: 3, path: "RLBFFF"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := createError(tt.wantErr, tt.name)
			ctx := context.Background()

			mockRover := mocks.NewRover(t)
			mockRover.On("MoveForward", mock.Anything).Return(err).Times(tt.args.countF).Maybe()
			mockRover.On("MoveBackward", mock.Anything).Return(err).Times(tt.args.countB).Maybe()
			mockRover.On("TurnLeft", mock.Anything).Return(err).Times(tt.args.countL).Maybe()
			mockRover.On("TurnRight", mock.Anything).Return(err).Times(tt.args.countR).Maybe()

			err = ChangePosition(ctx, mockRover, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
