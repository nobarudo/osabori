/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/kpango/glg"
	"github.com/schollz/progressbar/v3"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		OutputLog()
	},
}


func OutputLog() {
	glg.Success("Prompto moves around, please take a break. Have a good time.")
	glg.Info("fake Download. Don't worry, not downloading anything.")
	glg.Println("=================================================")
	glg.Println("The start of the fake build tool.")
	glg.Println("=================================================")
	glg.Warn("Be careful, if someone carefully reviews your prompts, they'll give away your true identity.")
	for i := 1; i <= 5; i++ {
		for j := 0; j < 2; j++ {
			fakeDownloadBar()
			fakeInstallBar()
		}
		glg.Warn(i,"minute has passed")	
	}
	glg.Success("Success!!!!")
	glg.Success("It's time. Get back to work.")
}

func fakeDownloadBar() {
	bar := progressbar.Default(100,"downloading")
	for i := 0; i < 100; i++ {
		time.Sleep(200 * time.Millisecond)
    	bar.Add(1)	
	}
	bar.Finish()
	fakeDownloadLog()
}

func fakeInstallBar() {
	bar := progressbar.NewOptions(100,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        " ",
			AltSaucerHead: "[yellow]<[reset]",
			SaucerHead:    "[yellow]-[reset]",
			SaucerPadding: "[white]•",
			BarStart:      "[blue]|[reset]",
			BarEnd:        "[blue]|[reset]",
		}),
	)

	
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
    	bar.Add(1)	
	}
	fakeInstallLog()
}

func fakeDownloadLog() {
	glg.Success("Downlod completed!!")
	glg.Info("A meaningless message")
	glg.Warn("When I download something, a warning message appears.")
	glg.Log("You can take a little more rest.")
	glg.Fail("I seriously don't want to work.")
}

func fakeInstallLog() {
	glg.Println("")
	glg.Success("Completed Install!!")
	glg.Info("I pretended to install various things.")
	glg.Warn("Different versions of life. Check your version.")
	glg.Warn("Different versions of salary. Check your version.")
	glg.Log("Looks like the installation is complete.")
	glg.Log("It's kind of a message like that.")
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
