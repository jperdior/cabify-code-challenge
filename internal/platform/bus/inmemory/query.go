package inmemory

import (
	"cabify-code-challenge/kit/query"
	"fmt"
	"log"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Ask implements the query.Bus interface.
func (b *QueryBus) Ask(query query.Query) (interface{}, error) {
	handler, ok := b.handlers[query.Type()]
	if !ok {
		return nil, nil
	}
	fmt.Print("Asking a query\n")
	answer, err := handler.Handle(query)
	if err != nil {
		log.Printf("Error while handling %s - %s\n", query.Type(), err)
	}
	return answer, err
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlers[queryType] = handler
}
