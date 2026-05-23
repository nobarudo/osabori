package cmd

import (
	"os"
	"osabori/internal/osabori"

	"github.com/spf13/cobra"
)

var (
	waitTime int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "osabori",
	Short: "This command will output a prompt like you are working",
	Long: `This command will output a prompt like you are working.
Option:
-t --time int You can specify the run time in minutes.(it defaults to 5 minutes.)`,
	Run: func(cmd *cobra.Command, args []string) {
		osabori.Start(waitTime)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(
		&waitTime,
		"time",
		"t",
		5,
		"You can specify the run time in minutes.",
	)
}

