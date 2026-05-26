package osabori

import (
	"log/slog"
	"time"

	"github.com/schollz/progressbar/v3"
)

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

func fakeDownloadLog() {
	slog.Info("Downlod completed!!")
	slog.Info("A meaningless message")
	slog.Info("A meaningless message")
	slog.Info("A meaningless message")
	slog.Error("When I download something, a warning message appears.")
	slog.Debug("You can take a little more rest.")
	slog.Error("I seriously don't want to work.")
}
