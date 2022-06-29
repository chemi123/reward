package service

import (
	"chemi123/reward/internal/domain/entity"
	repository "chemi123/reward/internal/domain/repository_inerface"
	rewardapi "chemi123/reward/proto/rewardapipb"
	"errors"
	"fmt"
)

const SERVICE_NUM = 2

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
	baseRewardCh := make(chan entity.BaseReward)
	incentiveCh := make(chan entity.Incenntive)
	errsCh := make(chan error, SERVICE_NUM)

	go func() {
		baseReward, err := s.baseRewardService.GetBaseReward()
		baseRewardCh <- baseReward
		errsCh <- err
	}()

	go func() {
		incentive, err := s.incentiveService.GetIncentive()
		incentiveCh <- incentive
		errsCh <- err
	}()

	var baseReward entity.BaseReward
	var incentive entity.Incenntive
	errs := make([]error, 0, SERVICE_NUM)

	// errorもあるのでservice数×2になる
	for i := 0; i < SERVICE_NUM*2; i++ {
		select {
		case baseReward = <-baseRewardCh:
		case incentive = <-incentiveCh:
		case err := <-errsCh:
			if err != nil {
				fmt.Println(err)
				errs = append(errs, err)
			}
		}
	}

	if len(errs) != 0 {
		return nil, errors.New("error occurred")
	}

	fmt.Println(baseReward)
	fmt.Println(incentive)
	return &rewardapi.Response{}, nil
}
