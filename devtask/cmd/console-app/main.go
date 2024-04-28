package main

import (
	"context"
	"devtask/cmd/console-app/cli"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	"devtask/internal/service/order"
	"devtask/internal/service/pack"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/file"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	if len(os.Args) == 1 {
		fmt.Println("Необходимо указать команду, для получения списка команд воспользуйтесь --help")
		return
	}
	defer cancel()

	database, err := db.NewDb(ctx, "config/config.json")
	if err != nil {
		log.Fatal("error while creating db", err)
	}

	defer database.GetPool(ctx).Close()

	args := os.Args[1:]
	orderStorage, err := file.NewStorageOrder("storage.json")
	if err != nil {
		fmt.Println("Не удалось подключиться к хранилищу данных ", err)
		return
	}

	packageVariants := map[model.PackageType]pack.CheckVariant{
		model.Packet: pack.PacketVariant{},
		model.Box:    pack.BoxVariant{},
		model.Tape:   pack.TapeVariant{},
	}
	packService := pack.NewService(packageVariants)
	orderService := order.NewService(&orderStorage, packService)

	pvzStorage, err := file.NewStoragePVZ("pvz_data.json")
	if err != nil {
		fmt.Println("Не удалось подключиться к хранилищу данных ", err)
		return
	}

	inMemoryCache := in_memory.NewInMemoryCache[model.PVZ]()

	pvzService := pvz.NewService(&pvzStorage, inMemoryCache, database)

	// TODO: Make an interface for a service and methods
	err = cli.ServicesCases(ctx, args, orderService, pvzService)
	fmt.Println("Работа завершена.")
	if err != nil {
		fmt.Println(err)
		return
	}
}
