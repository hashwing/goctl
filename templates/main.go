//cmd/{{.App}}/main.go
package main

import (
	"{{ .Mod }}/cmd/{{ .App }}/commond"
)

func main(){
	command.Execute()
}