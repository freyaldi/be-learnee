package util

import "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"

func UserLevelBenefit(user *entity.User) float32 {
	var benefit float32
	if user.Level == entity.Newbie {
		benefit = 0
	} else if user.Level == entity.Junior {
		benefit = 0.05
	} else if user.Level == entity.Senior {
		benefit = 0.1
	} else if user.Level == entity.Master {
		benefit = 0.2
	}
	return benefit
}

