// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) AdminAlias(alias string, endpoint string) (success bool, err error) {
	input := P{
		"alias":    alias,
		"endpoint": endpoint,
	}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.alias", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AdminAliasChain(chain string, alias string) (success bool, err error) {
	input := P{
		"chain": chain,
		"alias": alias,
	}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.aliasChain", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AdminGetChainAliases(chain string) (aliases GetChainAliasesResponse, err error) {
	input := P{
		"chain": chain,
	}

	err = c.client.NewRequest("ext/admin", "admin.getChainAliases", input, &aliases)
	if err != nil {
		return aliases, err
	}

	return aliases, err
}

func (c Calls) AdminLockProfile() (success bool, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.lockProfile", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AdminMemoryProfile() (success bool, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.memoryProfile", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AdminStartCPUProfiler() (success bool, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.startCPUProfiler", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}

func (c Calls) AdminStopCPUProfiler() (success bool, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/admin", "admin.stopCPUProfiler", input, &output)
	if err != nil {
		return success, err
	}
	err = output.OutBool("success", &success).Error()
	return success, err
}
