package vanta

import "context"

type Vanta interface {
	PeopleService
	MonitoredComputersService
	VulnerabilitiesService
}

type PeopleService interface {
	ListPeople(ctx context.Context) (*ListPeopleOutput, error)
	GetPersonByID(ctx context.Context, id string) (*Person, error)
}

type MonitoredComputersService interface {
	ListMonitoredComputers(ctx context.Context) (*ListMonitoredComputersOutput, error)
	GetMonitoredComputerByID(ctx context.Context, id string) (*MonitoredComputer, error)
}

type VulnerabilitiesService interface {
	ListVulnerabilities(ctx context.Context, opts ...ListVulnerabilitiesOption) (*ListVulnerabilitiesOutput, error)
}
