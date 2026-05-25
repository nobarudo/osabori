package osabori

import (
	"log/slog"
	"time"

	"osabori/internal/logger"

	"github.com/k0kubun/go-ansi"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/schollz/progressbar/v3"
)

// Start executes the fake work simulation for the specified duration in minutes.
func Start(waitTime int) {
	logger := logger.New()
	slog.SetDefault(logger)

	text := "OSABORI"

	letters := putils.LettersFromString(text)

	pterm.DefaultBigText.WithLetters(letters).Render()

	slog.Info("Prompto moves around, please take a break. Have a good time.")
	slog.Debug("fake Download. Don't worry, not downloading anything.")
	slog.Info("=================================================")
	slog.Info("The start of the fake build tool.")
	slog.Info("=================================================")
	slog.Warn("Be careful, if someone carefully reviews your prompts, they'll give away your true identity.")

	for i := 1; i <= waitTime; i++ {
		for j := 0; j < 2; j++ {
			fakeDownloadBar()
			fakeInstallBar()
		}
		slog.Warn("%d minute has passed", i)
		pterm.DefaultBox.Println(i, "minute has passed")
	}

	slog.Info("Success!!!!")
	slog.Info("It's time. Get back to work.")
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
	slog.Info("Downlod completed!!")
	slog.Info("A meaningless message")
	slog.Info("A meaningless message")
	slog.Info("A meaningless message")
	slog.Error("When I download something, a warning message appears.")
	slog.Debug("You can take a little more rest.")
	slog.Error("I seriously don't want to work.")
}

func fakeInstallLog() {
	slog.Info("")
	slog.Info("Completed Install!!")
	slog.Info("I pretended to install various things.")
	slog.Warn("Different versions of life. Check your version.")
	slog.Warn("Different versions of salary. Check your version.")
	slog.Debug("Looks like the installation is complete.")
	slog.Debug("It's kind of a message like that.")
}
