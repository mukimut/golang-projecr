
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// buildexecuteCmd represents the buildexecute command
var buildexecuteCmd = &cobra.Command{
	Use:   "buildexecute",
	Short: "short line",
	Long: `Loooooong line`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := ""
		buildDir := ""

		for index, command := range args {
			if command == "copydir" {
				sourceDir = args[index + 1]
			}else if command == "builddir" {
				buildDir = args[index + 1]
			}
		}

		if len(sourceDir) == 0 {
			fmt.Println("Error: Source Directory Needed")
			return
		}

		if len(buildDir) == 0 {
			buildDir = "./buildpath/"
		}
		//fmt.Println("cp " + sourceDir + " " + buildDir)
		out, err := exec.Command("cp", "-a", sourceDir, buildDir ).Output()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(buildexecuteCmd)
	//fmt.Println("Build and Excecute starts here")
}
