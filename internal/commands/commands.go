package commands

import "fmt"

type Commands struct {
	Mapped map[string]func(*State, Command) error
}

func NewCommandsStruct() *Commands {
	return &Commands{
		Mapped: make(map[string]func(*State, Command) error),
	}
}

func (c *Commands) Run(s *State, cmd Command) error {
	if s == nil {
		return fmt.Errorf("no source to read from")
	}
	if fn, exists := c.Mapped[cmd.Name]; exists {
		return fn(s, cmd)
	}
	return fmt.Errorf("command %q not found", cmd.Name)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if name == "" || f == nil {
		panic("invalid command registration")
	}
	c.Mapped[name] = f
}
