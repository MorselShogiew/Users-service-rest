package usecases

import "MorselShogiew/Users-service-rest/models"

// i

func (r *Service) AddUser(name, mail string) error {
	user := models.User{
		Name: name,
		Mail: mail,
	}

	return r.APIRepo.AddUser(user)
}
func (r *Service) DeleteUser(id int) error {
	return r.APIRepo.DeleteUser(id)
}

func (r *Service) GetUsers() (*[]models.User, error) {
	var err error
	res := r.cache.GetUsersList()
	if res != nil {
		return res, nil
	}
	res, err = r.APIRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	r.cache.SetUsersList(res)
	return res, nil

}
