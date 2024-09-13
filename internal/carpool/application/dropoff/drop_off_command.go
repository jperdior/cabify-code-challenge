package dropoff

import (
	"cabify-code-challenge/kit/command"
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
	dropOffService DropOffService
}

func NewDropOffCommandHandler(dropOffService DropOffService) DropOffCommandHandler {
	return DropOffCommandHandler{dropOffService: dropOffService}
}

// Handle implements the command.Handler interface
func (h DropOffCommandHandler) Handle(command command.Command) error {
	dropOffCommand, ok := command.(DropOffCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.dropOffService.Execute(dropOffCommand.groupID)
}
