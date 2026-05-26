package osabori

import (
	"log/slog"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

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

func fakeInstallLog() {
	slog.Info("")
	slog.Info("Completed Install!!")
	slog.Info("I pretended to install various things.")
	slog.Warn("Different versions of life. Check your version.")
	slog.Warn("Different versions of salary. Check your version.")
	slog.Debug("Looks like the installation is complete.")
	slog.Debug("It's kind of a message like that.")
}
