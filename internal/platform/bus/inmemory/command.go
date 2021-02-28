package inmemory

import (
	"context"

	"xmen-mutant/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) (reponse map[string]interface{}, err error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil, nil
	}
	result, errs := handler.Handle(ctx, cmd)
	if errs != nil {
		return nil, nil
	}
	return result, nil
}

// Register implements the command.Bus interface.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
