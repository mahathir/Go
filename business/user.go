package business

import (
	"three_tiers/data_access/interfaceda"
	"three_tiers/dto"
)

type (
	// UserService ...
	UserService struct {
		Da interfaceda.IUserDataAccess
	}
)

// GetUser get user from DB
func (service UserService) GetUser(id string) (*dto.User, error) {
	return service.Da.GetUser(id)
}
