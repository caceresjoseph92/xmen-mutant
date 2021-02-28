package consulting

import (
	"context"
	"errors"
	"xmen-mutant/kit/command"
)

const PersonCommandType command.Type = "command.consulting.person"

// personCommand is the command dispatched to create a new person.
type PersonCommand struct {
	id     int
	mutant bool
	dna    []string
}

// NewPersonCommand search PersonCommand.
func NewPersonCommand(id int, mutant bool, dna []string) PersonCommand {
	return PersonCommand{
		id:     id,
		mutant: mutant,
		dna:    dna,
	}
}

func (c PersonCommand) Type() command.Type {
	return PersonCommandType
}

// PersonCommandHandler is the command handler
// responsible for search persons.
type PersonCommandHandler struct {
	service PersonService
}

// NewPersonCommandHandler initializes a new PersonCommandHandler.
func NewPersonCommandHandler(service PersonService) PersonCommandHandler {
	return PersonCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h PersonCommandHandler) Handle(ctx context.Context, cmd command.Command) (stats map[string]interface{}, err error) {
	consultPersonCmd, ok := cmd.(PersonCommand)

	args := map[string]interface{}{
		"id":     consultPersonCmd.id,
		"mutant": consultPersonCmd.mutant,
		"dna":    consultPersonCmd.dna,
	}

	if !ok {
		return nil, errors.New("unexpected command")
	}
	stats, errs := h.service.ConsultPerson(ctx, args)
	if errs != nil {
		return nil, errors.New("unexpected command")
	}

	return stats, nil
}
