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

	err = out.
		OutStr("user", &user).
		OutStr("encoding", &encoding).
		Error()

	return user, encoding, err
}

//ImportUser import a user. password must match the user’s password. username doesn’t have to match the username user had when it was exported.
func (c Calls) ImportUser(username, password, user string) (bool, error) {
	out := network.P{}
	err := c.client.NewRequestFor(&out, "ext/keystore", "keystore.importUser",
		network.P{
			"username": username,
			"password": password,
			"user":     user,
		})
	if err != nil {
		return false, err
	}

	if !out.Success() {
		return false, errors.New("user could not be import")
	}

	return true, nil
}

//ListUsers list the users in this keystore.
func (c Calls) ListUsers() (users []string, err error) {
	out := network.P{}
	err = c.client.NewRequestFor(&out, "ext/keystore", "keystore.listUsers", nil)
	if err != nil {
		return users, err
	}

	values, ok := (out["users"]).([]interface{})
	if !ok {
		return users, errors.New("not valid type 'users'")
	}

	if len(values) < 1 {
		return []string{}, nil
	}

	for _, v := range values {
		users = append(users, v.(string))
	}

	return users, nil
}
