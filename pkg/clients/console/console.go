package console

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"rover/pkg/controller"
	"rover/pkg/logger"
	"rover/pkg/rovers"
)

type ConsoleApp struct {
	rover rovers.Rover
	log   *logger.Logger
}

func New(rover rovers.Rover, log *logger.Logger) *ConsoleApp {
	return &ConsoleApp{
		rover: rover,
		log:   log,
	}
}

func (app *ConsoleApp) read(ctx context.Context, scanner *bufio.Scanner) {
	if scanner.Scan() {
		input := strings.ToUpper(strings.TrimSpace(scanner.Text()))

		err := controller.ChangePosition(ctx, app.rover, input)
		if err != nil {
			app.log.Warn("Ошибка при изменении позиции:", "msg", err)
			return
		}

		x, y, err := app.rover.Position(ctx)
		if err != nil {
			app.log.Warn("Ошибка при запросе позиции:", "msg", err)
			return
		}

		fmt.Printf("Текущая позиция x = %d, y = %d\n", x, y)
	}

	if err := scanner.Err(); err != nil {
		app.log.Warn("Ошибка при чтении ввода:", "msg", err)
	}
}

func (app *ConsoleApp) Run(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	// fmt, а не log тк это вывод именно для консольного приложения
	fmt.Println("Введите текст:")

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				app.read(ctx, scanner)
			}
		}
	}()
}
