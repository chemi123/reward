package main

import (
	"chemi123/reward/internal/domain/service"
	"chemi123/reward/internal/infra/db"
	"chemi123/reward/internal/infra/repository_impl"
	"chemi123/reward/internal/presentation/endpoint"
	"log"
	"net"

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

func RunGrpcServer(rewardApiServer *endpoint.RewardApiServer) {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	endpoint.RegisterRewardApiServer(s, rewardApiServer)

	// TODO: graceful shutdown
	if err := s.Serve(listenPort); err != nil {
		log.Fatal(err)
	}
}

func main() {
	RunGrpcServer(inject())
}
