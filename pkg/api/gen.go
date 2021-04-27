// +build ignore

package main

import (
	"html/template"
	"log"
	"os"
)

type CallLogic struct {
	BasePath     string
	FunctionName string
	MethodName   string
	Args         string
	Returns      string
	Outs         []string
}

func main() {
	f, err := os.Create("./avm_calls_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	packageTemplate.Execute(f, AVMCalls())
}

func AVMCalls() []CallLogic {
	logics := []CallLogic{
		{
			BasePath:     "ext/bc/X",
			FunctionName: "AVMExportAVAX",
			MethodName:   "avm.exportAVAX",
			Args:         "args ExportAVAXArgs",
			Returns:      "txID, changeAddr string, err error",
			Outs:         []string{"txID", "changeAddr"},
		},
		{
			BasePath:     "ext/bc/X",
			FunctionName: "AVMExportKey",
			MethodName:   "avm.exportKey",
			Args:         "username, password, address string",
			Returns:      "privateKey string, err error",
			Outs:         []string{"privateKey"},
		},
	}
}

var packageTemplate = template.Must(template.New("").Parse(`
package api

func (c Calls) {{ .FunctionName }}({{ .Args }}) ({{ .Returns }}){
	rsp, err := c.client.NewRequestStruct("{{ .BasePath }}", "{{ .MethodName }}", args)
	if err != nil {
		return txID, changeAddr, err
	}

	err = rsp.
	OutStr("bytes", &bytes).
	OutStr("encoding", &encoding).
	Error()


	return ,err
}
`))
