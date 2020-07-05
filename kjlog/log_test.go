package kjlog

import "testing"

func TestKjlog(t *testing.T) {
	t.Run("test output to stdout", testOutputStdout)
	t.Run("test output to file", testOutputFile)
}

func testOutputStdout(t *testing.T) {
	err := InitLog("", LevelDebug)
	if err != nil {
		t.Fatal(err)
	}
	var i int
	Debug("Debug test", i)
	i++
	Info("Info", i)
	i++
	Warning("Warning test", i)
	i++
	Error("Error test", i)
	i++

	SetLogLevel(LevelWarning)
	Debug("Debug test", i)
	i++
	Info("Info", i)
	i++
	Warning("Warning test", i)
	i++
	Error("Error test", i)
	i++
}

func testOutputFile(t *testing.T) {
	err := InitLog("./logtest.txt", LevelDebug)
	if err != nil {
		t.Fatal(err)
	}
	var i int
	Debug("Debug test", i)
	i++
	Info("Info", i)
	i++
	Warning("Warning test", i)
	i++
	Error("Error test", i)
	i++

	SetLogLevel(LevelWarning)
	Debug("Debug test", i)
	i++
	Info("Info", i)
	i++
	Warning("Warning test", i)
	i++
	Error("Error test", i)
	i++
}
