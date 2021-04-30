// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) Metrics() (metrics string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/metrics", "metrics", input, &output)
	if err != nil {
		return metrics, err
	}
	err = output.OutStr("metrics", &metrics).Error()
	return metrics, err
}
