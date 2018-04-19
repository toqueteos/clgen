package clgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"shortHash": func(s string) string { return s[:9] },
	"contains":  strings.Contains,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
}

func WriteTemplate(pathIn, pathOut string, commits []Commit) error {
	body, err := ioutil.ReadFile(pathIn)
	if os.IsNotExist(err) {
		return fmt.Errorf("%q doesn't exist", pathIn)
	} else {
		if err != nil {
			return err
		}
	}

	tpl := template.New("changelog-markdown")
	t := template.Must(tpl.Funcs(funcMap).Parse(string(body)))

	f, err := os.Create(pathOut)
	if err != nil {
		return err
	}

	return t.Execute(f, commits)
}
