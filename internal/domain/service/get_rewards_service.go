package service

import (
	repository "chemi123/reward/internal/domain/repository_inerface"
	rewardapi "chemi123/reward/proto/rewardapipb"
	"fmt"
)

type GetRewardsService struct {
	baseRewardService *BaseRewardService
}

func NewGetRewardsService(baseRewardRepository repository.BaseRewardRepositoryInterface) GetRewardsService {
	return GetRewardsService{
		baseRewardService: NewBaseRewardService(
			baseRewardRepository,
		),
		// initialize with other service used in GetRewrdsService
	}
}

// injection等の確認をしたいだけなのでbaserewardをgetしても空のレスポンスを返す
func (s *GetRewardsService) GetRewards() (*rewardapi.Response, error) {
	// TODO: goroutineを使って並列化してjoinして返す
	basereward, err := s.baseRewardService.GetBaseReward()
	if err != nil {
		return nil, err
	}

	fmt.Println(basereward)
	return &rewardapi.Response{}, nil
}
