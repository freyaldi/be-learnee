package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/db"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepository := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})

	jwt := util.NewAuth(&util.AuthConfig{})

	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepository: userRepository,
		JWT: jwt,
	})

	return NewRouter(&RouterConfig{
		UserUsecase: userUsecase,
	})
}

func Init()  {
	r := createRouter()
	err :=r.Run(":8080")
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}