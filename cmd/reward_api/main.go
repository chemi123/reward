package main

import (
	"chemi123/reward/internal/domain/service"
	"chemi123/reward/internal/infra/db"
	"chemi123/reward/internal/infra/repository_impl"
	"chemi123/reward/internal/presentation/endpoint"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func inject() *endpoint.RewardApiServer {
	db, err := db.OpenDB("reward")
	if err != nil {
		log.Fatal(err)
	}

	baseRewardRepository := repository_impl.NewBaseRewardRepository(db)
	incentiveRepository := repository_impl.NewIncentiveRepository(db)
	getRewardsService := service.NewGetRewardsService(baseRewardRepository, incentiveRepository)
	rewardApiServer := endpoint.NewRewardApiServer(getRewardsService)
	return rewardApiServer
}

func Run(rewardApiServer *endpoint.RewardApiServer) {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	defer listenPort.Close()

	s := grpc.NewServer()
	endpoint.RegisterRewardApiServer(s, rewardApiServer)
	defer s.GracefulStop()

	go func() {
		if err := s.Serve(listenPort); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("shutdown gracefully")
}

func main() {
	Run(inject())
}
