// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) HealthGetLiveness() (checks GetLivenessResponse, err error) {
	input := P{}

	err = c.client.NewRequest("ext/health", "health.getLiveness", input, &checks)
	if err != nil {
		return checks, err
	}

	return checks, err
}
