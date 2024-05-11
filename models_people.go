package vanta

import (
	"time"
)

type ListPeopleOutput struct {
	Results ListPeopleResults `json:"results"`
}

type ListPeopleResults struct {
	PageInfo PageInfo `json:"pageInfo"`
	Data     []Person `json:"data"`
}

type Person struct {
	ID           string        `json:"id"`
	EmailAddress string        `json:"emailAddress"`
	Employment   *Employment   `json:"employment,omitempty"`
	Name         *Name         `json:"name,omitempty"`
	GroupIDs     []string      `json:"groupIds,omitempty"`
	Sources      *Sources      `json:"sources,omitempty"`
	TasksSummary *TasksSummary `json:"tasksSummary,omitempty"`
}

type Employment struct {
	EndDate   *time.Time        `json:"endDate,omitempty"`
	JobTitle  *string           `json:"jobTitle,omitempty"`
	StartDate *time.Time        `json:"startDate,omitempty"`
	Status    *EmploymentStatus `json:"status,omitempty"`
}

type Name struct {
	Display *string `json:"display,omitempty"`
	Fist    *string `json:"first,omitempty"`
	Last    *string `json:"last,omitempty"`
}

type Sources struct {
	EmailAddress *GenericDataSource    `json:"emailAddress,omitempty"`
	Employment   *EmploymentDataSource `json:"employment,omitempty"`
}

type EmploymentDataSource struct {
	StartDate *GenericDataSource `json:"startDate,omitempty"`
	EndDate   *GenericDataSource `json:"endDate,omitempty"`
}

type GenericDataSource struct {
	IntegrationID *string `json:"intgrationId,omitempty"`
	ResourceID    *string `json:"resourceId,omitempty"`
	Type          *string `json:"type,omitempty"`
}

type TasksSummary struct {
	CompletionDate *time.Time   `json:"completionDate,omitempty"`
	DueDate        *time.Time   `json:"dueDate,omitempty"`
	Status         *TaskStatus  `json:"status,omitempty"`
	Details        *TaskDetails `json:"details,omitempty"`
}

type TaskDetails struct {
	CompleteTrainings              *CompleteTrainings              `json:"completeTrainings,omitempty"`
	CompleteCustomTasks            *CompleteCustomTasks            `json:"completeCustomTasks,omitempty"`
	CompleteOffboardingCustomTasks *CompleteOffboardingCustomTasks `json:"completeOffboardingCustomTasks,omitempty"`
	CompleteBackgroundChecks       *CompleteBackgroundChecks       `json:"completeBackgroundChecks,omitempty"`
	AcceptPolicies                 *AcceptPolicies                 `json:"acceptPolicies,omitempty"`
	InstallDeviceMonitoring        *InstallDeviceMonitoring        `json:"installDeviceMonitoring,omitempty"`
}

type BaseTaskDetails struct {
	TaskType       *string     `json:"taskType,omitempty"`
	Status         *TaskStatus `json:"status,omitempty"`
	DueDate        *time.Time  `json:"dueDate,omitempty"`
	CompletionDate *time.Time  `json:"completionDate,omitempty"`
	Disabled       *bool       `json:"disabled,omitempty"`
}

type CompleteTrainings struct {
	BaseTaskDetails

	IncompleteTrainings []GenericNamedItem `json:"incompleteTrainings,omitempty"`
	CompletedTrainings  []GenericNamedItem `json:"completedTrainings,omitempty"`
}

type CompleteCustomTasks struct {
	BaseTaskDetails

	IncompleteCustomTasks []GenericNamedItem `json:"incompleteCustomTasks,omitempty"`
	CompletedCustomTasks  []GenericNamedItem `json:"completedCustomTasks,omitempty"`
}

type CompleteOffboardingCustomTasks struct {
	BaseTaskDetails

	IncompleteCustomOffboardingTasks []GenericNamedItem `json:"incompleteCustomOffboardingTasks,omitempty"`
	CompletedCustomOffboardingTasks  []GenericNamedItem `json:"completedCustomOffboardingTasks,omitempty"`
}

type CompleteBackgroundChecks struct {
	BaseTaskDetails
}

type AcceptPolicies struct {
	BaseTaskDetails

	UnacceptedPolicies []GenericNamedItem `json:"unacceptedPolicies,omitempty"`
	AcceptedPolicies   []GenericNamedItem `json:"acceptedPolicies,omitempty"`
}

type InstallDeviceMonitoring struct {
	BaseTaskDetails
}

type TaskStatus string

type EmploymentStatus string

const (
	TaskStatusComplete                      TaskStatus = "COMPLETE"
	TaskStatusDueSoon                       TaskStatus = "DUE_SOON"
	TaskStatusNone                          TaskStatus = "NONE"
	TaskStatusOverdue                       TaskStatus = "OVERDUE"
	TaskStatusPaused                        TaskStatus = "PAUSED"
	TaskStatusOffboardingComplete           TaskStatus = "OFFBOARDING_COMPLETE"
	TaskStatusOffboardingDueSoon            TaskStatus = "OFFBOARDING_DUE_SOON"
	TaskStatusOffboardingOffboardingOverdue TaskStatus = "OFFBOARDING_OVERDUE"

	EmploymentStatusUpcoming EmploymentStatus = "UPCOMING"
	EmploymentStatusCurrent  EmploymentStatus = "CURRENT"
	EmploymentStatusOnLeave  EmploymentStatus = "ON_LEAVE"
	EmploymentStatusInactive EmploymentStatus = "INACTIVE"
	EmploymentStatusFormer   EmploymentStatus = "FORMER"
)
