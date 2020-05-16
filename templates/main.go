//cmd/{{.App}}/main.go
package main

import (
	"{{ .Mod }}/cmd/{{ .App }}/command"
)

func main(){
	command.Execute()
}