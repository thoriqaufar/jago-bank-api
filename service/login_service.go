package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/jago-bank-api/entity"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/repository"
)

type LoginService interface {
	Register(request *model.RegisterRequest) error
	Login(request *model.LoginRequest) (*model.LoginResponse, error)
}

type loginService struct {
	repository repository.LoginRepository
	validate   *validator.Validate
}

func NewLoginService(r repository.LoginRepository, v *validator.Validate) *loginService {
	return &loginService{
		repository: r,
		validate:   v,
	}
}

func (s *loginService) Register(request *model.RegisterRequest) error {
	if err := s.validate.Struct(request); err != nil {
		return &helper.BadRequestError{Message: err.Error()}
	}

	if emailExist := s.repository.EmailExists(request.Email); emailExist {
		return &helper.BadRequestError{Message: "Email already exists"}
	}

	if request.Password != request.PasswordConfirmation {
		return &helper.BadRequestError{Message: "Password does not match"}
	}

	passwordHash, err := helper.HashPassword(request.Password)
	if err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    passwordHash,
		PIN:         request.PIN,
		Address:     request.Address,
		Province:    request.Province,
		City:        request.City,
		PostalCode:  request.PostalCode,
	}

	if err = s.repository.Register(&user); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *loginService) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, &helper.BadRequestError{Message: err.Error()}
	}

	user, err := s.repository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, &helper.NotFoundError{Message: "Wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &helper.NotFoundError{Message: "Wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &helper.InternalServerError{Message: err.Error()}
	}

	response := model.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &response, nil
}
