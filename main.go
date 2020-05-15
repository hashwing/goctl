package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Values struct {
	Mod         string
	App         string
	Dir         string
	EnableMongo bool
}

func main() {
	appName := flag.String("app", "server", "app name")
	mod := flag.String("mod", "server", "mod path")
	dir := flag.String("dir", ".", "code dir")
	enableMongo := flag.Bool("mongo", false, "add mongo")
	flag.Parse()
	if *mod == "" {
		mod = appName
	}
	v := &Values{
		Mod:         *mod,
		App:         *appName,
		Dir:         *dir,
		EnableMongo: *enableMongo,
	}
	err := create(v)
	if err != nil {
		panic(err)
	}
}

func create(vars *Values) error {
	fmt.Println("start parse code...")
	for _, f := range fs {
		path, err := ParseTpl(f.Path, vars)
		if err != nil {
			return err
		}
		if strings.HasPrefix(path, "pkg/store/mongo") && !vars.EnableMongo {
			continue
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
	fmt.Println("finish parse code")
	fmt.Println("start init mod...")
	cmd := exec.Command("go", "mod", "init", vars.Mod)
	cmd.Dir = vars.Dir + "/"
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("finish init mod")
	fmt.Println("start go get packages...")
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = vars.Dir + "/"
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("finish go get packages")
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
	"substr":    substr,
	"replace":   replace,
	"randomStr": randomStr,
}

func substr(s string, i int) string {
	return s[:i]
}

func replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func randomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
