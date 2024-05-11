package vanta

import "context"

type Vanta interface {
	PeopleService
}

type PeopleService interface {
	ListPeople(ctx context.Context) (*ListPeopleOutput, error)
	GetPersonByID(ctx context.Context, id string) (*Person, error)
}
