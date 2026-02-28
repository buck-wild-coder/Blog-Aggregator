package main

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if commandName, exists := c.registeredCommands[cmd.name]; exists {
		err := commandName(s, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
