package main

type iCommand interface {
	execute()
}

type Light interface {
	On()
	Off()
}

type lightOnCommand struct {
	light *Light
}

func execute(l *Light) {
}
func newLightOnCommand(l *Light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

func (l *lightOnCommand) execute() {
}
