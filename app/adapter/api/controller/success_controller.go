package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SuccessController struct {
	usecase usecase.SaveSuccessUseCase
}

func NewSuccessController(usecase usecase.SaveSuccessUseCase) SuccessController {
	return SuccessController{
		usecase: usecase,
	}
}

func (controller SuccessController) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.SaveTransactionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	logger.StructuredLog("SuccessController", "Execute", "input", input).Info("Init create user function")

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		logger.Errorf("Error validating user", err)
		return
	}

	defer r.Body.Close()

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	logger.StructuredLog("SuccessController", "Execute", "output", output).Info("Finish create user function")

	response.NewSuccess(output, http.StatusCreated).Send(w)

}
