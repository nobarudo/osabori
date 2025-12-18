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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "osabori",
	Short: "This command will output a prompt like you are working",
	Long: `This command will output a prompt like you are working`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.osabori.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


