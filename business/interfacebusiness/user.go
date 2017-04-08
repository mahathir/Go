package interfacebusiness

import (
	"three_tiers/dto"
)

type IUserService interface {
	GetUser(id string) (*dto.User, error)
}
