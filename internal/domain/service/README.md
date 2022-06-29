### 補足
GetRewardsServiceだけpresentation層から呼ばれているため、presentation層からinterfaceが提供されている。  
以下のserviceはdomain層内で完結しているため、特にinterfaceを用意していない。  
- BaseRewardService
- IncentiveService