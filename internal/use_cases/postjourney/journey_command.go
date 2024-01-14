package postjourney

import (
	"cabify-code-challenge/kit/command"
	"context"
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
	useCase CreateJourneyUseCase
}

func NewCreatingJourneyCommandHandler(useCase CreateJourneyUseCase) CreatingJourneyCommandHandler {
	return CreatingJourneyCommandHandler{useCase: useCase}
}

// Handle implements the command.Handler interface
func (h CreatingJourneyCommandHandler) Handle(context context.Context, command command.Command) error {
	createJourneyCommand, ok := command.(CreatingJourneyCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	return h.useCase.CreateJourney(context, createJourneyCommand.groupID, createJourneyCommand.people)

}
