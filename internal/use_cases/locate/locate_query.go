package locate

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/query"
	"context"
	"errors"
)

const LocateQueryType query.Type = "locate"

type LocateQuery struct {
	GroupID int
}

func NewLocateQuery(groupID int) LocateQuery {
	return LocateQuery{
		GroupID: groupID,
	}
}

func (c LocateQuery) Type() query.Type {
	return LocateQueryType
}

type LocationResponse struct {
	id    int
	seats int
}

func NewLocationResponse(params ...interface{}) interface{} {
	car, ok := params[0].(carpool.Car)
	if !ok {
		return errors.New("unexpected response")
	}
	return LocationResponse{
		id:    car.ID().Value(),
		seats: car.Seats().Value(),
	}
}

type LocateQueryHandler struct {
	useCase LocateUseCase
}

func NewLocateQueryHandler(useCase LocateUseCase) LocateQueryHandler {
	return LocateQueryHandler{useCase: useCase}
}

// Handle implements the query.Handler interface
func (h LocateQueryHandler) Handle(context context.Context, query query.Query) (interface{}, error) {
	locateQuery, ok := query.(LocateQuery)
	if !ok {
		return LocationResponse{}, errors.New("unexpected query")
	}
	carPool := context.Value("carPool").(*carpool.CarPool)
	car, err := h.useCase.Locate(carPool, locateQuery.GroupID)
	if err != nil {
		return LocationResponse{}, err
	}
	return NewLocationResponse(car), nil
}
