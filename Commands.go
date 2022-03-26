package netgo

type Command func(string) string
type Commands map[string]Command

func NewCommands() Commands {
	return make(Commands, 32)
}

func (c Commands) SetCommand(name string, cmd Command) {
	c[name] = cmd
}

func (c Commands) SetInfo(name, value string) {
	c[name] = func(s string) string {
		return value
	}
}
