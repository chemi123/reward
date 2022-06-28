package service_interface

import rewardapi "chemi123/reward/proto/rewardapipb"

type GetRewardsServiceInterface interface {
	GetRewards() (*rewardapi.Response, error)
}
