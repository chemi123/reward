package service

import (
	repository "chemi123/reward/internal/domain/repository_inerface"
	rewardapi "chemi123/reward/proto/rewardapipb"
	"fmt"
)

type GetRewardsService struct {
	baseRewardService *BaseRewardService
	incentiveService  *IncentiveService
}

func NewGetRewardsService(
	baseRewardRepository repository.BaseRewardRepositoryInterface,
	incentiveRepository repository.IncentiveRepositoryInterface,
) *GetRewardsService {
	return &GetRewardsService{
		baseRewardService: NewBaseRewardService(
			baseRewardRepository,
		),
		incentiveService: NewIncentiveService(
			incentiveRepository,
		),
	}
}

// injection等の確認をしたいだけなのでbaserewardをgetしても空のレスポンスを返す
func (s *GetRewardsService) GetRewards() (*rewardapi.Response, error) {
	// TODO: goroutineを使って並列化してjoinして返す
	basereward, baseRewardErr := s.baseRewardService.GetBaseReward()
	if baseRewardErr != nil {
		return nil, baseRewardErr
	}

	incentive, incentiveErr := s.incentiveService.GetIncentive()
	if incentiveErr != nil {
		return nil, incentiveErr
	}

	fmt.Println(basereward)
	fmt.Println(incentive)
	return &rewardapi.Response{}, nil
}
