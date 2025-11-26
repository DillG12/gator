package main

import (
	"errors"
)

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.list[cmd.name]; exists {
		err := handler(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("command not found")
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.list[name] = handler
}
