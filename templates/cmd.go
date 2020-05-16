//cmd/{{.App}}/command/cmd.go
package command

import (
	"fmt"
	"os"

	"{{ .Mod }}/pkg/version"
	"github.com/spf13/cobra"
)


var cfgFile string

var rootCmd = &cobra.Command{
	Use: "{{ .App }}",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get())
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print app version",
	Long:  `print app version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get())
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfg", "", "config file path")
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
