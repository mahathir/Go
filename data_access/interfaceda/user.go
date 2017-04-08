package interfaceda

import (
	"three_tiers/dto"
)

type (
	IUserDataAccess interface {
		GetUser(id string) (*dto.User, error)
	}
)
