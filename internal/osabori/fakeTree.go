package osabori

import (
	"time"

	"github.com/pterm/pterm"
)

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
