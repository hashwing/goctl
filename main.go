package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Values struct {
	Mod string
	App string
	Dir string
}

func main() {
	appName := flag.String("app", "server", "app name")
	mod := flag.String("mod", "server", "mod path")
	dir := flag.String("dir", ".", "code dir")
	flag.Parse()
	if *mod == "" {
		mod = appName
	}
	v := &Values{
		Mod: *mod,
		App: *appName,
		Dir: *dir,
	}
	err := create(v)
	if err != nil {
		panic(err)
	}
}

func create(vars *Values) error {
	for _, f := range fs {
		path, err := ParseTpl(f.Path, vars)
		if err != nil {
			return err
		}
		err = os.MkdirAll(vars.Dir+"/"+filepath.Dir(path), 0664)
		if err != nil {
			return err
		}
		data, err := ParseTpl(string(f.Data), vars)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(vars.Dir+"/"+path, []byte(data), 0664)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("go", "mod", "init", vars.Mod)
	cmd.Dir = vars.Dir + "/"
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = vars.Dir + "/"
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ParseTpl(tpl string, vars *Values) (string, error) {
	tmpl, err := newTpl().Parse(tpl)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, vars)
	return b.String(), err
}

func newTpl() *template.Template {
	tmpl := template.New("tpl")
	tmpl = tmpl.Funcs(funcMap)
	return tmpl
}

// funcMap provides extra functions for the templates.
var funcMap = template.FuncMap{
	"substr":  substr,
	"replace": replace,
}

func substr(s string, i int) string {
	return s[:i]
}

func replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}
