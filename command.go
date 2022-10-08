package caddy_webhook

import (
	"os/exec"
)

type Cmd struct {
	Command string
	Args    []string
	Path    string
}

func (c *Cmd) AddCommand(command []string, path string) {
	c.Command = command[0]
	c.Args = command[1:]
	c.Path = path
}

func (c *Cmd) Run() {
	cmd := exec.Command(c.Command, c.Args...)
	cmd.Dir = c.Path
	cmd.Start()
	cmd.Wait()
}
