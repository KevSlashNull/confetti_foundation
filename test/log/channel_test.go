package log

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogToDefaultChannel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log().Info("the message")

	assert.Len(t, openAndReadFile(testFile), 0)
	assert.Len(t, openAndReadFile(testFileSecond), 1)
}

func TestLogToNonExistingChannel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	assert.PanicsWithError(t, "can not log to channel. Channel 'fake' does not exist", func() {
		app.Log("fake").Info("the message")
	})
}

func TestLogToMultipleNonExistingChannels(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	assert.PanicsWithError(t, "can not log to one of the channels. Channel 'fake1' does not exist", func() {
		app.Log("fake1", "fake2").Info("the message")
	})
}

func TestLogToSpecificChannel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log("first").Info("the message")

	assert.Len(t, openAndReadFile(testFile), 1)
	assert.Len(t, openAndReadFile(testFileSecond), 0)
}

func TestLogToMultipleChannels(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log("first", "second").Info("the message")

	assert.Len(t, openAndReadFile(testFile), 1)
	assert.Len(t, openAndReadFile(testFileSecond), 1)
}

func getAppWithChannels() inter.App {
	first := loggers.Syslog{Path: testFile, MinLevel: syslog.INFO}
	second := loggers.Syslog{Path: testFileSecond, MinLevel: syslog.INFO}
	allLoggers := map[string]interface{}{"first": first, "second": second}

	app := newTestApp()
	app.Bind("config.Logging.Default", "second")
	app.Bind("config.Logging.Channels", allLoggers)

	return app
}
