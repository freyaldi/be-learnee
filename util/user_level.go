package util

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/constant"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
)

func UserLevelBenefit(user *entity.User) float32 {
	var benefit float32
	if user.Level == &entity.Newbie {
		benefit = constant.LevelBenefitNewbie
	} else if user.Level == &entity.Junior {
		benefit = constant.LevelBenefitJunior
	} else if user.Level == &entity.Senior {
		benefit = constant.LevelBenefitSenior
	} else if user.Level == &entity.Master {
		benefit = constant.LevelBenefitMaster
	}
	return benefit
}
func UpdateLevel(level string, totalTransaction int) string {
	if level == string(entity.Newbie){
		if totalTransaction >= constant.MinTotalTransactionJunior {
			level = string(entity.Junior)
		}
	} 
	if level == string(entity.Junior){
		if totalTransaction >= constant.MinTotalTransactionSenior {
			level = string(entity.Senior)
		}
	} else if level == string(entity.Senior){
		if totalTransaction >= constant.MinTotalTransactionMaster {
			level = string(entity.Master)
		}
	}
	return level
}

