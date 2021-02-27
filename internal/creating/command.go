package creating

import (
	"context"
	"errors"
	"xmen-mutant/kit/command"
	"xmen-mutant/kit/utils"
)

const PersonCommandType command.Type = "command.creating.person"

// personCommand is the command dispatched to create a new person.
type PersonCommand struct {
	mutant bool
	dna    []string
}

// NewPersonCommand creates a new PersonCommand.
func NewPersonCommand(mutant bool, dna []string) PersonCommand {
	return PersonCommand{
		mutant: mutant,
		dna:    dna,
	}
}

func (c PersonCommand) Type() command.Type {
	return PersonCommandType
}

// PersonCommandHandler is the command handler
// responsible for creating persons.
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
func (h PersonCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createPersonCmd, ok := cmd.(PersonCommand)
	mutant := utils.IsMutant(createPersonCmd.dna)
	createPersonCmd.mutant = mutant

	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreatePerson(
		ctx,
		createPersonCmd.mutant,
		createPersonCmd.dna,
	)
}
