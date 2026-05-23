package osabori

import (
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/kpango/glg"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/schollz/progressbar/v3"
)

// Start executes the fake work simulation for the specified duration in minutes.
func Start(waitTime int) {
	text := "OSABORI"

	letters := putils.LettersFromString(text)

	pterm.DefaultBigText.WithLetters(letters).Render()

	glg.Success("Prompto moves around, please take a break. Have a good time.")
	glg.Info("fake Download. Don't worry, not downloading anything.")
	glg.Println("=================================================")
	glg.Println("The start of the fake build tool.")
	glg.Println("=================================================")
	glg.Warn("Be careful, if someone carefully reviews your prompts, they'll give away your true identity.")

	for i := 1; i <= waitTime; i++ {
		for j := 0; j < 2; j++ {
			fakeDownloadBar()
			fakeInstallBar()
		}
		glg.Warn(i, " minute has passed")
		pterm.DefaultBox.Println(i, "minute has passed")
	}

	glg.Success("Success!!!!")
	glg.Success("It's time. Get back to work.")
}

func fakeDownloadBar() {
	for j := 0; j < 4; j++ {
		bar := progressbar.Default(100, "downloading")
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			_ = bar.Add(2)
		}
		bar.Clear()
	}
	fakeDownloadLog()
}

func fakeInstallBar() {
	bar := progressbar.NewOptions(
		100,
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
		_ = bar.Add(1)
	}
	fakeInstallLog()
}

func fakeDownloadLog() {
	glg.Success("Downlod completed!!")
	glg.Info("A meaningless message")
	glg.Info("A meaningless message")
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
