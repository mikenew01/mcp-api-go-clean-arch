package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/domain"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"net/http"
)

type GetUserController struct {
	usecase usecase.GetUserUseCase
}

func NewGetUserController(usecase usecase.GetUserUseCase) GetUserController {
	return GetUserController{
		usecase: usecase,
	}
}

func (controller GetUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("userId")
	var input = usecase.GetUserInput{Id: userId}
	logger.StructuredLog("GetUserController", "Execute", "input", input).Info("Init get user function")

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		controller.handleErrors(w, err)
		return
	}
	logger.StructuredLog("GetUserController", "Execute", "output", output).Info("Finish get user function")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

func (controller GetUserController) handleErrors(w http.ResponseWriter, err error) {
	logger.StructuredLog("GetUserController", "handleErrors", "", "").Errorf("msg=%s", err.Error())

	var status int
	switch err {
	case domain.ErrUserNotFound:
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	response.NewError(err, status).Send(w)
}
