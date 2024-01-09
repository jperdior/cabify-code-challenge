package creating_journey

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/command"
	"context"
	"errors"
)

const CreatingJourneyCommandType command.Type = "creating_journey"

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
	carPool, ok := context.Value("carPool").(*carpool.CarPool)
	if !ok {
		return errors.New("carPool not found in context")
	}
	createJourneyCommand, ok := command.(CreatingJourneyCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	return h.useCase.CreateJourney(carPool, createJourneyCommand.groupID, createJourneyCommand.people)

}
