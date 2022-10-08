package usecases

import "MorselShogiew/Users-service-rest/models"

// i

func (r *Service) AddUser(id int, name, mail string) error {
	user := models.User{
		Id:   id,
		Name: name,
		Mail: mail,
	}

	return r.APIRepo.AddUser(user)
}
func (r *Service) DeleteUser(id int) error {
	return r.APIRepo.DeleteUser(id)
}

func (r *Service) GetUsers() (*[]models.User, error) {
	return r.APIRepo.GetUsers()
}
