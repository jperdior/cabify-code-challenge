package post_journey

import (
	"cabify-code-challenge/kit/command"
	"errors"
)

const CreatingJourneyCommandType command.Type = "journey"

type CreatingJourneyCommand struct {
	groupID int
	people  int
}

func NewCreatingJourneyCommand(groupId int, people int) CreatingJourneyCommand {
	return CreatingJourneyCommand{
		groupID: groupId,
		people:  people,
	}
}

func (c CreatingJourneyCommand) Type() command.Type {
	return CreatingJourneyCommandType
}

type CreatingJourneyCommandHandler struct {
	service JourneyService
}

func NewCreatingJourneyCommandHandler(service JourneyService) CreatingJourneyCommandHandler {
	return CreatingJourneyCommandHandler{service: service}
}

// Handle implements the command.Handler interface
func (h CreatingJourneyCommandHandler) Handle(command command.Command) error {
	createJourneyCommand, ok := command.(CreatingJourneyCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	return h.service.Execute(createJourneyCommand.groupID, createJourneyCommand.people)

}
