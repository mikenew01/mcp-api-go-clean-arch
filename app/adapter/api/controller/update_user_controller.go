package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UpdateUserController struct {
	usecase usecase.UpdateUserUseCase
}

func NewUpdateUserController(usecase usecase.UpdateUserUseCase) UpdateUserController {
	return UpdateUserController{
		usecase: usecase,
	}
}

func (controller UpdateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("userId")
	logger.StructuredLog("UpdateUserController", "Execute", "input", userId).Info("Init update user function")
	var input usecase.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()
	input.Id = userId
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logger.StructuredLog("UpdateUserController", "Execute", "input", userId).Info("Finish update user function")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
