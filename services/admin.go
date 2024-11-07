package services

import (
	"errors"
	"golang-beginner-chap24/repositories"
)

type AdminService struct {
	AdminRepo repositories.AdminRepository
}

func NewAdminService(adminRepo repositories.AdminRepository) *AdminService {
	return &AdminService{AdminRepo: adminRepo}
}

func (s *AdminService) LoginAdmin(username, password, token string) error {
	if username == "" {
		return errors.New("username is required")
	} else if password == "" {
		return errors.New("password is required")
	}

	err := s.AdminRepo.Login(username, password, token)
	if err != nil {
		return err
	}
	return nil
}

// func (s *AdminService) LogoutAdmin(token string) error {
// 	return s.AdminRepo.Logout(token)
// }

func (s *AdminService) GetAdminByToken(token string) (string, error) {
	return s.AdminRepo.GetByToken(token)
}
