package commands

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("command name is needed")
	}
	s.PointerToConfig.SetUser(cmd.Args[0])
	fmt.Println("User was set")
	return nil
}
