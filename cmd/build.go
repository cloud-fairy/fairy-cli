package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var templateDockerCommand = "docker run --rm --mount type=bind,source=$(pwd),target=/app/input cfairy/assembler:dev run --allow-all --no-check assembler.js --project=../../input/"
var templatePostfix = " --out /app/input/CF_BUILD_OUTPUT"

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Try build a project from a cloudfairy JSON project",
	Example: "fairy build path/to/project.json",
	Aliases: []string{"b", "bild", "bld", "buld"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectFile := args[0]
		out, err := exec.Command("/bin/bash", "-c", templateDockerCommand+projectFile+templatePostfix).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
