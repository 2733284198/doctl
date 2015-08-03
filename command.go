package doit

import (
	"os"
	"os/exec"
)

// Command runs commands.
type Command interface {
	Run(args ...string) ([]byte, error)
	Start(args ...string) error
	Stop() error
}

// LiveCommand is a live implementation of Command.
type LiveCommand struct {
	path string
	cmd  *exec.Cmd
}

// NewLiveCommand creates a LiveCommand.
func NewLiveCommand(path string) *LiveCommand {
	return &LiveCommand{
		path: path,
	}
}

var _ Command = &LiveCommand{}

// Run runs a LiveCommand with args and returns stdout and an error if there was one.
func (c *LiveCommand) Run(args ...string) ([]byte, error) {
	return exec.Command(c.path, args...).Output()
}

// Start runs a LiveCommand with args and starts it. This would most likely block,
// so you should call it in a goroutine.
func (c *LiveCommand) Start(args ...string) error {
	c.cmd = exec.Command(c.path, args...)
	c.cmd.Stderr = os.Stderr
	return c.cmd.Start()
}

// Stop stops an existing LiveCommand.
func (c *LiveCommand) Stop() error {
	return c.cmd.Process.Kill()
}

type MockCommand struct {
	path    string
	running bool
	runFn   func() error
	startFn func() error
	stopFn  func() error
}

var _ Command = &MockCommand{}

func NewMockCommand(path string) *MockCommand {
	return &MockCommand{
		path: path,
	}
}

func (c *MockCommand) Run(args ...string) ([]byte, error) {
	return nil, c.runFn()
}

func (c *MockCommand) Start(args ...string) error {
	c.running = true
	return c.startFn()
}

func (c *MockCommand) Stop() error {
	c.running = false
	return c.stopFn()
}
