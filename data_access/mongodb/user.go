package mongodb

import (
	"three_tiers/dto"

	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserDataAccess ...
	UserDataAccess struct {
		Session *mgo.Session
	}
)

// GetUser . Get user from database by given string id.
func (da UserDataAccess) GetUser(id string) (*dto.User, error) {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("id should be a mongodb object id format")
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := &dto.User{}

	// Fetch user
	if err := da.Session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
		return nil, err
	}

	return u, nil
}
