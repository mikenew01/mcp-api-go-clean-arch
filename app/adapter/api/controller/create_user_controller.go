package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CreateUserController struct {
	usecase usecase.CreateUserUseCase
}

func NewCreateUserController(usecase usecase.CreateUserUseCase) CreateUserController {
	return CreateUserController{
		usecase: usecase,
	}
}

func (controller CreateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	logger.StructuredLog("CreateUserController", "Execute", "input", input).Info("Init create user function")
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
	logger.StructuredLog("CreateUserController", "Execute", "output", output).Info("Finish create user function")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
