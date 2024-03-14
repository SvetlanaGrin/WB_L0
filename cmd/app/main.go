package main

import (
	"WB_L0"
	"WB_L0/Sabscribe"
	"WB_L0/cmd/initStart"
	"WB_L0/internal/handler"
	"WB_L0/internal/repository"
	"WB_L0/internal/service"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
)

//	@title			Order details API
//	@version		1.0
//	@description	API server for receiving order information

//	@host		localhost:8000
//	@BasePath	/

func main() {
	if err := initStart.InitConfig(); err != nil {
		log.Fatalf("Error init config: %v\n", err)
		return
	}
	c := cache.New(20*time.Minute, 30*time.Minute)
	db := initStart.InitDB()
	repos := repository.NewRepository(db, c)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	subs := Sabscribe.NewSubscribers(handlers)
	Sc := Sabscribe.ConnectStan("nats-example_2")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	stop := make(chan bool)

	subs.RunSubscribers(Sc, stop)

	srv := new(WB_L0.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		return
	}
	<-ch

	// Остановка горутины
	stop <- true
	// Ожидание до тех пор, пока не выполнится
	<-stop
}
