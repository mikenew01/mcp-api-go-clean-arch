package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"net/http"
)

type DeleteUserController struct {
	usecase usecase.DeleteUserUseCase
}

func NewDeleteUserController(usecase usecase.DeleteUserUseCase) DeleteUserController {
	return DeleteUserController{
		usecase: usecase,
	}
}

func (controller DeleteUserController) Execute(w http.ResponseWriter, r *http.Request) {
	logger.StructuredLog("DeleteUserController", "Execute", "", "").Info("Init delete user function")
	var userId = r.URL.Query().Get("userId")
	var input = usecase.DeleteUserInput{Id: userId}

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logger.StructuredLog("DeleteUserController", "Execute", "output", output).Info("Finish delete user function")

	response.NewSuccess(output, http.StatusNoContent).Send(w)
}
