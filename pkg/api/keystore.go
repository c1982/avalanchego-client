package api

import (
	"avalanchego-client/pkg/network"
	"errors"
)

//CreateUser create a Keystore User
func (c Calls) CreateUser(username, password string) error {
	out := network.P{}
	err := c.client.NewRequestFor(&out, "ext/keystore", "keystore.createUser",
		network.P{
			"username": username,
			"password": password,
		})
	if err != nil {
		return err
	}

	if !out.Success() {
		return errors.New("user could not be created ")
	}

	return nil
}

//DeleteUser Delete a user.
func (c Calls) DeleteUser(username, password string) error {
	out := network.P{}
	err := c.client.NewRequestFor(&out, "ext/keystore", "keystore.deleteUser",
		network.P{
			"username": username,
			"password": password,
		})
	if err != nil {
		return err
	}

	if !out.Success() {
		return errors.New("user could not be created ")
	}

	return nil
}

//ExportUser Export a user. The user can be imported to another node with keystore.importUser. The user’s password remains encrypted.
func (c Calls) ExportUser(username, password string) (user, encoding string, err error) {
	encoding = "cb58" //Default encoding type
	out := network.P{}
	err = c.client.NewRequestFor(&out, "ext/keystore", "keystore.exportUser",
		network.P{
			"username": username,
			"password": password,
		})
	if err != nil {
		return user, encoding, err
	}

	user, isstring := (out["user"]).(string)
	if !isstring {
		return user, encoding, errors.New("invalid type of 'user' key")
	}

	encoding, isstring = (out["encoding"]).(string)
	if !isstring {
		return user, encoding, errors.New("invalid type of 'encoding' key")
	}

	return user, encoding, err
}

//TODO: wait implementation
func (c Calls) ImportUser(username, password, user, encoding string) (bool, error) {
	return false, nil
}

//TODO: wait implementation
func (c Calls) ListUsers(username, password, user, encoding string) (users []string, err error) {
	return users, nil
}
