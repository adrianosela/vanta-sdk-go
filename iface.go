package vanta

import "context"

type Vanta interface {
	PeopleService
	MonitoredComputersService
}

type PeopleService interface {
	ListPeople(ctx context.Context) (*ListPeopleOutput, error)
	GetPersonByID(ctx context.Context, id string) (*Person, error)
}

type MonitoredComputersService interface {
	ListMonitoredComputers(ctx context.Context) (*ListMonitoredComputersOutput, error)
	GetMonitoredComputerByID(ctx context.Context, id string) (*MonitoredComputer, error)
}
