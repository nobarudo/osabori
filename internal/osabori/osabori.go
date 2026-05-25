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

func fakeTreeLog() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
	sleep()
	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	sleep()
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	sleep()
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	sleep()
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	sleep()
	logger.Error("Oh no, this process is getting killed!", logger.Args("fatal", false))
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))
	sleep()

	// Define a map with interesting stuff
	interstingStuff = map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
	sleep()
	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	sleep()
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	sleep()
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	sleep()
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
}

// Function to pause the execution for 3 seconds
func sleep() {
	time.Sleep(time.Second * 3)
}
