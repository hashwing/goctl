package main

//go:generate go run pgenerate.go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	content := []byte("package main\n")
	infos := []string{}
	tplFiles := getFiles("templates")
	for _, tplF := range tplFiles {
		data, err := ioutil.ReadFile(tplF.Path)
		if err != nil {
			panic(err)
		}
		line, err := bytes.NewBuffer(data).ReadString('\n')
		if err != nil {
			panic(err)
		}
		path := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line, "\n", ""), "\r", ""), "//", "")
		dataStr := strings.Replace(string(data), line, "", 1)
		key := tplF.Name
		info := fmt.Sprintf(`
			var %s_file = fileInfo{
				Path: "%s",
				Data: %s,
			}
			`, key, path, key)
		infos = append(infos, key+"_file")
		fmt.Println(key)
		v := info + "var " + key + " = []byte{" + hexdump([]byte(dataStr)) + "\n}\n"
		content = append(content, []byte(v)...)
	}
	fs := `
	type fileInfo struct {
		Path string
		Data []byte
	}
	var fs=[]fileInfo{
		`
	for _, info := range infos {
		fs += info + ",\n"
	}
	fs += "}\n"
	contentStr := string(content) + fs
	err := ioutil.WriteFile("tpl.go", []byte(contentStr), 0664)
	if err != nil {
		panic(err)
	}
	err = exec.Command("go", "fmt", "tpl.go").Run()
	if err != nil {
		panic(err)
	}
}

type tplInfo struct {
	Path string
	Name string
}

func getFiles(dir string) []tplInfo {
	res := make([]tplInfo, 0)
	tplFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, tplF := range tplFiles {
		if tplF.IsDir() {
			subFiles := getFiles(dir + "/" + tplF.Name())
			if err != nil {
				panic(err)
			}
			res = append(res, subFiles...)
			continue
		}
		path := dir + "/" + tplF.Name()
		name := strings.Replace(strings.Replace(path, ".", "_", -1), "/", "_", -1)
		res = append(res, tplInfo{Path: path, Name: name})
	}
	return res
}

// hexdump is a template function that creates a hux dump
// similar to xxd -i.
func hexdump(v interface{}) string {
	var data []byte
	switch vv := v.(type) {
	case []byte:
		data = vv
	case string:
		data = []byte(vv)
	default:
		return ""
	}
	var buf bytes.Buffer
	for i, b := range data {
		dst := make([]byte, 4)
		src := []byte{b}
		encode(dst, src, ldigits)
		buf.Write(dst)

		buf.WriteString(",")
		if (i+1)%cols == 0 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

// default number of columns
const cols = 12

// hex lookup table for hex encoding
const (
	ldigits = "0123456789abcdef"
	udigits = "0123456789ABCDEF"
)

func encode(dst, src []byte, hextable string) {
	dst[0] = '0'
	dst[1] = 'x'
	for i, v := range src {
		dst[i+1*2] = hextable[v>>4]
		dst[i+1*2+1] = hextable[v&0x0f]
	}
}
