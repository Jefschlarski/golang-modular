package user

import (
	"errors"

	"github.com/Jefschlarski/golang-modular/pkg/validate"
)

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (dto *CreateUserDto) Validate() error {
	if dto.Name == "" {
		return errors.New("name is required")
	}

	if dto.Email == "" {
		return errors.New("email is required")
	} else {
		err := validate.Email(dto.Email)
		if err != nil {
			return err
		}
	}

	if dto.Phone == "" {
		return errors.New("phone is required")
	} else {
		err := validate.Phone(dto.Phone)
		if err != nil {
			return err
		}
	}

	if dto.Password == "" {
		return errors.New("password is required")
	} else {
		err := validate.Password(dto.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReturnUserDto struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserType  int    `json:"usertype"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType int    `json:"usertype"`
}

func (dto *UpdateUserDto) Validate() error {
	if dto.Email != "" {
		err := validate.Email(dto.Email)
		if err != nil {
			return err
		}
	}

	if dto.Phone != "" {
		err := validate.Phone(dto.Phone)
		if err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserPasswordDto struct {
	NewPassword string `json:"new_password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
}

func (dto *UpdateUserPasswordDto) Validate() error {
	if dto.NewPassword == "" {
		return errors.New("new password is required")
	} else {
		err := validate.Password(dto.NewPassword)
		if err != nil {
			return err
		}
	}

	return nil
}
