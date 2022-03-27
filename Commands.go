package netgo

type Command func(id int, message string) string
type Commands map[string]Command

func NewCommands() Commands {
	return make(Commands, 32)
}

func (c Commands) SetCommand(name string, cmd Command) {
	c[name] = cmd
}

func (c Commands) SetInfo(name, value string) {
	c[name] = func(id int, s string) string {
		return value
	}
}
