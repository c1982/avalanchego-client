//go:generate go run gen.go
//This file generate all api calls in to the pkg/api directory.
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clls := Calls{
		Api: []CallLogic{},
	}

	chains := map[string]string{
		"platform": "",
		"avax":     "",
		"avm":      "",
		"admin":    "",
		"auth":     "",
		"health":   "",
		"index":    "",
		"info":     "",
		"ipc":      "",
		"keystore": "",
		"metrics":  "",
	}

	for k := range chains {
		inputfile := fmt.Sprintf("./config/%s.api.yaml", k)
		outputfile := fmt.Sprintf("./pkg/api/%s_gen.go", k)

		content, err := ioutil.ReadFile(inputfile)
		err = yaml.Unmarshal(content, &clls)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(outputfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		packageTemplate.Execute(f, clls)
	}

	_, err := exec.Command("go", "fmt ./...").Output()
	if err != nil {
		log.Fatal(err)
	}

}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

{{ range .Api }}func (c Calls) {{ .FunctionName }}({{ .FunctionSignature }}) ({{ .ReturnSignature }}){
	{{ if .StructArgs }}input := {{ .ArgsFirstItem }} {{- else }}input := P{ 
		{{ range $v:= .Args }}"{{ index $v 0 }}":{{ index $v 0 }}, 
	{{ end }}}{{- end }}
	{{ if not .StructReturns }}output := P{}{{- end }}
	err = c.client.NewRequest("{{ .BasePath }}", "{{ .MethodName }}", input, {{ if .StructReturns }} &{{ .ReturnsFirstItem }} {{- else}} &output{{- end}})
	if err != nil {
		return {{ .ReturnAll }}
	}
	{{ if not .StructReturns }}err = output.{{ range $v:= .Returns }}{{ if eq (index $v 1) "bool" }}OutBool("{{ index $v 0 }}", &{{ index $v 0 }}). {{- end}}{{ if eq (index $v 1) "string" }} OutStr("{{ index $v 0 }}", &{{ index $v 0 }}). {{- end}}{{- end }}Error(){{- end}}
	return {{ .ReturnAll }}
}

{{ end }}
`))

type Calls struct {
	Api []CallLogic `yaml:"Api"`
}

type CallLogic struct {
	BasePath      string     `yaml:"path"`
	MethodName    string     `yaml:"name"`
	StructArgs    bool       `yaml:"struct-args"`
	Args          [][]string `yaml:"args"`
	StructReturns bool       `yaml:"struct-returns"`
	Returns       [][]string `yaml:"returns"`
}

func (cl CallLogic) FunctionName() string {
	formatstr := strings.Replace(cl.MethodName, ".", " ", -1)
	formatstr = strings.Title(formatstr)
	formatstr = strings.Replace(formatstr, " ", "", -1)
	return formatstr
}

func (cl CallLogic) FunctionSignature() string {
	args := ""
	for i := 0; i < len(cl.Args); i++ {
		args += fmt.Sprintf("%s %s", cl.Args[i][0], cl.Args[i][1])
		if i+1 < len(cl.Args) {
			args += ", "
		}
	}

	return args
}

func (cl CallLogic) ReturnSignature() string {
	rtrns := ""
	for i := 0; i < len(cl.Returns); i++ {
		rtrns += fmt.Sprintf("%s %s", cl.Returns[i][0], cl.Returns[i][1])
		if i+1 < len(cl.Returns) {
			rtrns += ", "
		}
	}

	return rtrns
}

func (cl CallLogic) ReturnAll() string {
	args := ""
	for i := 0; i < len(cl.Returns); i++ {
		args += fmt.Sprintf("%s", cl.Returns[i][0])
		if i+1 < len(cl.Returns) {
			args += ", "
		}
	}

	return args
}

func (cl CallLogic) ArgsFirstItem() string {
	return cl.Args[0][0]
}

func (cl CallLogic) ReturnsFirstItem() string {
	return cl.Returns[0][0]
}

func (cl CallLogic) PassByRef() []string {
	outs := []string{}
	for i := 0; i < len(cl.Returns); i++ {
		v := cl.Returns[i]
		if v[1] == "error" {
			continue
		}
		outs = append(outs, v[0])
	}
	return outs
}
