package endpoint

import (
	"chemi123/reward/internal/domain/service"
	service_interface "chemi123/reward/internal/presentation/service_interface"
	rewardapi "chemi123/reward/proto/rewardapipb"
	"context"
	"log"

	"google.golang.org/grpc"
)

type RewardApiServer struct {
	rewardapi.UnimplementedRewardApiServer
	getRewardsService service_interface.GetRewardsServiceInterface
}

func NewRewardApiServer(getRewardsService *service.GetRewardsService) *RewardApiServer {
	return &RewardApiServer{
		getRewardsService: getRewardsService,
	}
}

func (s *RewardApiServer) GetRewards(ctx context.Context, request *rewardapi.Request) (*rewardapi.Response, error) {
	log.Printf("Received %v", request.RequestId)

	// TODO: 値の詰め替えw
	return s.getRewardsService.GetRewards()
}

func RegisterRewardApiServer(s *grpc.Server, rewardApiserver *RewardApiServer) {
	rewardapi.RegisterRewardApiServer(s, rewardApiserver)
}
