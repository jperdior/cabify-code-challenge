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
		return carpool.Car{}, errors.New("unexpected query")
	}
	carPool := context.Value("carPool").(*carpool.CarPool)
	car, err := h.useCase.Locate(carPool, locateQuery.GroupID)
	if err != nil {
		return carpool.Car{}, err
	}
	return car, nil
}
