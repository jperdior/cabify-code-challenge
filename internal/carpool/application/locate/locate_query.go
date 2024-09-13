package locate

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/query"
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
	Id    int
	Seats int
}

func NewLocationResponse(params ...interface{}) interface{} {
	car, ok := params[0].(domain.Car)
	if !ok {
		return errors.New("unexpected response")
	}
	return LocationResponse{
		Id:    car.ID().Value(),
		Seats: car.Seats().Value(),
	}
}

type LocateQueryHandler struct {
	service LocateService
}

func NewLocateQueryHandler(service LocateService) LocateQueryHandler {
	return LocateQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h LocateQueryHandler) Handle(query query.Query) (interface{}, error) {
	locateQuery, ok := query.(LocateQuery)
	if !ok {
		return LocationResponse{}, errors.New("unexpected query")
	}
	car, err := h.service.Execute(locateQuery.GroupID)
	if err != nil {
		return LocationResponse{}, err
	}
	return NewLocationResponse(car), nil
}
