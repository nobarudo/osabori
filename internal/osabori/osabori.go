package osabori

import (
	"log/slog"

	"osabori/internal/logger"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// Start executes the fake work simulation for the specified duration in minutes.
func Start(waitTime int) {
	logger := logger.New()
	pterm.DefaultCenter.Println("This is a tool that simply keeps the prompt running while you take a break.\nYou might see error messages, but don't worry. It's not doing anything.")

	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("OSABORI")).Srender()

	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s)

	// Print each line of the text separately centered in the terminal
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("This text is centered!\nBut each line is\ncentered\nseparately")

	slog.SetDefault(logger)

	slog.Info("Prompto moves around, please take a break. Have a good time.")
	slog.Debug("fake Download. Don't worry, not downloading anything.")
	slog.Info("=================================================")
	slog.Info("The start of the fake build tool.")
	slog.Info("=================================================")
	slog.Warn("Be careful, if someone carefully reviews your prompts, they'll give away your true identity.")

	for i := 1; i <= waitTime; i++ {
		fakeDownloadBar()
		fakeInstallBar()
		fakeTreeLog()
		pterm.DefaultBox.Println(i, "minute has passed")
	}

	slog.Info("Success!!!!")
	slog.Info("It's time. Get back to work.")
}
