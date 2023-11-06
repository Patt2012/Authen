package services

import (
	"context"

	"github.com/Patt2012/authen/modules/users/repositories"
	"github.com/google/uuid"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserSevice {
	return userService{repo}
}

func (s userService) Insert(ctx context.Context, register RegisterRequest) error {
	user := repositories.User{
		Email: register.Email,
		Password: register.Password,
		FullName: register.FullName,
		Active: true,
	}

	err := s.repo.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil	
}

func (s userService) Update(ctx context.Context, update UserUpdate, id string) error {

	//update := UserUpdate{}

	user := repositories.User{
		ID: 		uuid.MustParse(id),
		Password: 	update.Password,
		FullName: 	update.FullName,
		Active: 	update.Active,
	}

	err := s.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s userService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s userService) FindById(ctx context.Context, id string) (response UserResponse, err error) {
	
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return 
	}

	userResponse := UserResponse{
		ID: user.ID.String(),
		Email: user.Email,
		Password: user.Password,
		FullName: user.FullName,
		Active: user.Active,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
	
	return userResponse, nil
}

func (s userService) FindByEmail(ctx context.Context, email string) (response UserResponse, err error) {
	user, err := s.repo.FindByEmail(ctx, email)

	response = UserResponse{
		ID: user.ID.String(),
		Email: user.Email,
		Password: user.Password,
		FullName: user.FullName,
		Active: user.Active,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
	
	return response, nil
}

func (s userService) FindAll(ctx context.Context) (responses []UserResponse, err error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return 
	}

	for _, user := range users{
		responses = append(responses, UserResponse{
			ID: user.ID.String(),
			Email: user.Email,
			Password: user.Password,
			FullName: user.FullName,
			Active: user.Active,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}

	return responses, nil
}