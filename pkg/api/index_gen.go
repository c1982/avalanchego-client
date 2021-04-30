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

func (c Calls) IndexGetContainerByIndex(index int, encoding string) (container GetContainerByIndexResponse, err error) {
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
