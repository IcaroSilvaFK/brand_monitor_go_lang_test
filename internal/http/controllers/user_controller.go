package controllers

import (
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/http/dtos"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/smtp"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	createUserUseCase     use_cases.CreateUserUseCaseInterface
	finByEmailUserUseCase use_cases.FindUserByEmailUseCaseInterface
	deleteUserUseCase     use_cases.DeleteUserUseCaseInterface
}

type UserControllerInterface interface {
	CreateUser(ctx *gin.Context)
	FindUserByEmail(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(
	createUserUseCase use_cases.CreateUserUseCaseInterface,
	finByEmailUserUseCase use_cases.FindUserByEmailUseCaseInterface,
	deleteUserUseCase use_cases.DeleteUserUseCaseInterface,
) UserControllerInterface {
	return &UserController{
		createUserUseCase:     createUserUseCase,
		finByEmailUserUseCase: finByEmailUserUseCase,
		deleteUserUseCase:     deleteUserUseCase,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var u dtos.CreateUserInput

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	user, err := uc.createUserUseCase.Execute(u.Email, u.Username, u.Password)

	smtpClient := smtp.NewSMTPClient()

	go smtpClient.CreateAccountSendEmail(user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) FindUserByEmail(ctx *gin.Context) {

	var u dtos.UserLogin

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	user, err := uc.finByEmailUserUseCase.Execute(u.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	if err := user.VerifyPassword(u.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Email or password is invalid",
			"timestamp": time.Now(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)

}
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	_, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	if err = uc.deleteUserUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
