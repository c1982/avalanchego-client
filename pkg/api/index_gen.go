// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) IndexGetLastAccepted(encoding string) (lastaccepted GetLastAcceptedResponse, err error) {
	input := P{
		"encoding": encoding,
	}

	err = c.client.NewRequest("ext/index/X/tx", "index.getLastAccepted", input, &lastaccepted)
	if err != nil {
		return lastaccepted, err
	}

	return lastaccepted, err
}

func (c Calls) IndexGetContainerByIndex(index uint64, encoding string) (container GetContainerByIndexResponse, err error) {
	input := P{
		"index":    index,
		"encoding": encoding,
	}

	err = c.client.NewRequest("ext/index/X/tx", "index.getContainerByIndex", input, &container)
	if err != nil {
		return container, err
	}

	return container, err
}

func (c Calls) IndexGetContainerRange(startIndex uint64, numToFetch uint64, encoding string) (container []GetContainerRangeResponse, err error) {
	input := P{
		"startIndex": startIndex,
		"numToFetch": numToFetch,
		"encoding":   encoding,
	}

	err = c.client.NewRequest("ext/index/X/tx", "index.getContainerRange", input, &container)
	if err != nil {
		return container, err
	}

	return container, err
}

func (c Calls) IndexGetIndex(containerID string, encoding string) (index string, err error) {
	input := P{
		"containerID": containerID,
		"encoding":    encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/index/X/tx", "index.getIndex", input, &output)
	if err != nil {
		return index, err
	}
	err = output.OutStr("index", &index).Error()
	return index, err
}

func (c Calls) IndexIsAccepted(containerID string, encoding string) (isAccepted bool, err error) {
	input := P{
		"containerID": containerID,
		"encoding":    encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/index/X/tx", "index.isAccepted", input, &output)
	if err != nil {
		return isAccepted, err
	}
	err = output.OutBool("isAccepted", &isAccepted).Error()
	return isAccepted, err
}
