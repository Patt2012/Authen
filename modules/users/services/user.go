package services

import "context"

type RegisterRequest struct {
	Email		string 	`json:"email"`
	Password	string 	`json:"password"`
	FullName	string	`json:"full_name"`
}

type UserUpdate struct {
	Password	string 	`json:"password"`
	FullName	string	`json:"full_name"`
	Active 		bool	`json:"active"`
}

type LoginRequest struct {
	Email 		string	`json:"email"`
	Password	string	`json:"password"`
}

type UserResponse struct {
	ID			string 	`json:"id"`
	Email		string 	`json:"email"`
	Password	string 	`json:"password"`
	FullName	string	`json:"full_name"`
	Active		bool	`json:"active"`
	CreatedAt	string	`json:"created_at"`
	UpdatedAt	string	`json:"updated_at"`
}

type UserSevice interface {
	Insert(context.Context, RegisterRequest) error
	Update(context.Context, UserUpdate, string) error
	Delete(context.Context, string) error
	FindById(context.Context, string) (UserResponse, error)
	FindByEmail(context.Context, string) (UserResponse, error)
	FindAll(context.Context) ([]UserResponse, error)
}