package dropoff

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/command"
	"context"
	"errors"
)

const DropOffCommandType command.Type = "dropoff"

type DropOffCommand struct {
	groupID int
}

func NewDropOffCommand(groupId int) DropOffCommand {
	return DropOffCommand{
		groupID: groupId,
	}
}

func (c DropOffCommand) Type() command.Type {
	return DropOffCommandType
}

type DropOffCommandHandler struct {
	useCase DropOffUseCase
}

func NewDropOffCommandHandler(useCase DropOffUseCase) DropOffCommandHandler {
	return DropOffCommandHandler{useCase: useCase}
}

// Handle implements the command.Handler interface
func (h DropOffCommandHandler) Handle(context context.Context, command command.Command) error {
	dropOffCommand, ok := command.(DropOffCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	carPool := context.Value("carPool").(*carpool.CarPool)

	return h.useCase.DropOff(carPool, dropOffCommand.groupID)
}
