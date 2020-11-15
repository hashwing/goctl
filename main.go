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
	EnableMysql bool
	EnableVue   bool
	Router      string
}

func main() {
	appName := flag.String("app", "server", "app name")
	mod := flag.String("mod", "server", "mod path")
	dir := flag.String("dir", "", "code dir")
	enableMongo := flag.Bool("mongo", false, "add mongo")
	enableMysql := flag.Bool("mysql", false, "add mysql")
	enableVue := flag.Bool("vue", false, "create vue ui")
	router := flag.String("router", "gin", "http router, gin or beego")
	flag.Parse()
	if *mod == "" {
		mod = appName
	}
	if *dir == "" {
		dir = appName
	}
	v := &Values{
		Mod:         *mod,
		App:         *appName,
		Dir:         *dir,
		EnableMongo: *enableMongo,
		EnableMysql: *enableMysql,
		EnableVue:   *enableVue,
		Router:      *router,
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
		if strings.HasPrefix(f.Key, "vuetify2") && !vars.EnableVue {
			continue
		}
		if !strings.HasPrefix(f.Key, "vuetify2") && vars.EnableVue {
			continue
		}
		if strings.HasPrefix(f.Key, "mongo") && !vars.EnableMongo {
			continue
		}
		if strings.HasPrefix(f.Key, "gorm/mysql") && !vars.EnableMysql {
			continue
		}
		if strings.HasPrefix(f.Key, "beego") && vars.Router != "beego" {
			continue
		}
		if strings.HasPrefix(f.Key, "gin") && vars.Router != "gin" {
			continue
		}
		err = os.MkdirAll(vars.Dir+"/"+filepath.Dir(path), 0755)
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
	if vars.EnableVue {
		return nil
	}
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

	if vars.EnableMysql {
		fmt.Println("install generate tools")
		cmd = exec.Command("go", "install", "github.com/hashwing/togo")
		cmd.Dir = vars.Dir + "/"
		err = cmd.Run()
		if err != nil {
			return err
		}
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
