package cli

import (
	"bufio"
	"context"
	"devtask/internal/model"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func Run(ctx context.Context, service storagePVZ) error {
	wg := sync.WaitGroup{}
	commands := make(chan string)
	logs := make(chan string)
	go logger(logs, &wg)

	fmt.Println("Введите команду (add или show):")
	go publish(commands)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err := subscribe(ctx, service, commands, logs, done, &wg)
	wg.Wait()
	close(logs)
	close(commands)
	return err
}

func logger(logs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range logs {
		fmt.Println(msg)
	}
}

func publish(commands chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		commands <- input
	}
}

func subscribe(ctx context.Context, service storagePVZ, commands <-chan string, logs chan<- string, done <-chan os.Signal, wg *sync.WaitGroup) error {
	for {
		select {
		case command := <-commands:
			parts := strings.Split(command, " ")
			if len(parts) < 2 {
				fmt.Println("Неверный формат команды. Используйте формат add name address contact или show name.")
				break
			}

			switch command := parts[0]; command {
			case "show":
				if len(parts) != 2 {
					fmt.Println("Неверный формат команды. Используйте формат show id.")
					break
				}
				id, _ := strconv.ParseInt(parts[1], 10, 64)
				wg.Add(1)
				go func() {
					defer wg.Done()
					logs <- "Goroutine1: Пришла команда на чтение."
					pvzInfo, err := service.GetInfo(ctx, id)
					if err != nil {
						logs <- "ПВЗ не найден."
					} else {
						logs <- "Информация о ПВЗ:"
						logs <- fmt.Sprintf("%+v", pvzInfo)
					}
					logs <- "Goroutine1: Завершил команду."
				}()
			case "add":
				if len(parts) != 5 {
					fmt.Println("Неверный формат команды. Используйте формат show id name address add.")
					break
				}
				id, _ := strconv.ParseInt(parts[1], 10, 64)
				pvzModel := model.PVZ{
					ID:      id,
					Name:    parts[2],
					Address: parts[3],
					Contact: parts[4],
				}
				wg.Add(1)
				go func() {
					defer wg.Done()
					logs <- "Goroutine2: Пришла команда на добавление."
					_, err := service.AddInfo(ctx, pvzModel)
					if err != nil {
						logs <- err.Error()
					} else {
						logs <- "Информация о ПВЗ успешно добавлена."
					}
					logs <- "Goroutine2: Завершил команду."
				}()
			default:
				fmt.Println("Неверная команда. Введите show или add.")
			}
		case <-done:
			fmt.Println("Завершение работы всех горутин!")
			return nil
		}
	}
}
