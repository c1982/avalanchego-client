// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) AuthNewToken(newtoken NewTokenArgs) (token string, err error) {
	input := newtoken
	output := P{}
	err = c.client.NewRequest("ext/auth", "auth.newToken", input, &output)
	if err != nil {
		return token, err
	}
	err = output.OutStr("token", &token).Error()
	return token, err
}

func (c Calls) AuthRevokeToken(password string, token string) (success bool, err error) {
	input := P{
		"password": password,
		"token":    token,
	}
	output := P{}
	err = c.client.NewRequest("ext/auth", "auth.revokeToken", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AuthChangePassword(oldPassword string, newPassword string) (success bool, err error) {
	input := P{
		"oldPassword": oldPassword,
		"newPassword": newPassword,
	}
	output := P{}
	err = c.client.NewRequest("ext/auth", "auth.changePassword", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}
