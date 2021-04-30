// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) KeystoreCreateUser(username string, password string) (success bool, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/info", "keystore.createUser", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}
