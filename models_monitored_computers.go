package vanta

import "time"

type ListMonitoredComputersOutput struct {
	Results ListMonitoredComputersResults `json:"results"`
}

type ListMonitoredComputersResults struct {
	PageInfo PageInfo            `json:"pageInfo"`
	Data     []MonitoredComputer `json:"data"`
}

type MonitoredComputer struct {
	ID                    string                 `json:"id"`
	IntegrationID         string                 `json:"integrationId"`
	LastCheckDate         *time.Time             `json:"lastCheckDate,omitempty"`
	ScreenLock            MonitoredComputerCheck `json:"screenlock"`
	DiskEncryption        MonitoredComputerCheck `json:"diskEncryption"`
	PasswordManager       MonitoredComputerCheck `json:"passwordManager"`
	AntivirusInstallation MonitoredComputerCheck `json:"antivirusInstallation"`
	OperatingSystem       *OperatingSystem       `json:"operatingSystem,omitempty"`
	Owner                 *Owner                 `json:"owner,omitempty"`
	SerialNumber          *string                `json:"serialNumber,omitempty"`
	UDID                  *string                `json:"udid,omitempty"`
}

type MonitoredComputerCheck struct {
	Outcome MonitoredComputerCheckOutcome `json:"outcome"`
}

type OperatingSystem struct {
	Type    MonitoredComputerOSType `json:"type"`
	Version *string                 `json:"version,omitempty"`
}

type Owner struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
}

type MonitoredComputerOSType string
type MonitoredComputerCheckOutcome string

const (
	MonitoredComputerOSTypeMacOS   MonitoredComputerOSType = "macOS"
	MonitoredComputerOSTypeLinux   MonitoredComputerOSType = "linux"
	MonitoredComputerOSTypeWindows MonitoredComputerOSType = "windows"

	MonitoredComputerCheckOutcomeFail       MonitoredComputerCheckOutcome = "FAIL"
	MonitoredComputerCheckOutcomeInProgress MonitoredComputerCheckOutcome = "IN_PROGRESS"
	MonitoredComputerCheckOutcomeNA         MonitoredComputerCheckOutcome = "NA"
	MonitoredComputerCheckOutcomePass       MonitoredComputerCheckOutcome = "PASS"
)
