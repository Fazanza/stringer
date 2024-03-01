package stringer

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "stringer",
	Version: version,
	Short:   "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding)
One can use string to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while exeucting your CLU: '%s'", err)
		os.Exit(1)
	}
}
