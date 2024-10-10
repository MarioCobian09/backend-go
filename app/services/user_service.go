package services

import (
	"backend-go/app/models"
	"backend-go/app/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) isUserDuplicated(usuario string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Usuario == usuario {
			return true, nil
		}
	}

	return false, nil
}

func (s *UserService) isNameDuplicated(nombre, apPaterno, apMaterno string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Nombre == nombre && u.ApPaterno == apPaterno && u.ApMaterno == apMaterno {
			return true, nil
		}
	}

	return false, nil
}

func (s *UserService) CreateUser(user models.Usuario) error {
	isDuplicated, err := s.isUserDuplicated(user.Usuario)
	if err != nil {
		return err
	}
	if isDuplicated {
		return errors.New("El usuario ya existe")
	}

	isNameDuplicated, err := s.isDuplicatedName(user.Nombre, user.ApPaterno, user.ApMaterno)
	if err != nil {
		return err
	}
	if isNameDuplicated {
		return errors.New("El nombre completo ya existe")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Imagen = "default.png"

	err = s.Repo.CreateUser(user)

	if err != nil {
		return err
	}
	return nil
}
