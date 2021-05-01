// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) KeystoreCreateUser(username string, password string) (success bool, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/keystore", "keystore.createUser", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) KeystoreDeleteUser(username string, password string) (success bool, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/keystore", "keystore.deleteUser", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) KeystoreExportUser(username string, password string) (user string, encoding string, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/keystore", "keystore.exportUser", input, &output)
	if err != nil {
		return user, encoding, err
	}
	err = output.OutStr("user", &user).OutStr("encoding", &encoding).Error()
	return user, encoding, err
}

func (c Calls) KeystoreImportUser(username string, password string, user string) (success bool, err error) {
	input := P{
		"username": username,
		"password": password,
		"user":     user,
	}
	output := P{}
	err = c.client.NewRequest("ext/keystore", "keystore.importUser", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) KeystoreListUsers(username string, password string, user string) (users ListUsersResponse, err error) {
	input := P{
		"username": username,
		"password": password,
		"user":     user,
	}

	err = c.client.NewRequest("ext/keystore", "keystore.listUsers", input, &users)
	if err != nil {
		return users, err
	}

	return users, err
}
