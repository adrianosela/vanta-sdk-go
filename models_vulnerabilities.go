package vanta

import (
	"net/url"
	"strconv"
	"time"
)

type ListVulnerabilitiesOption func(m map[string]string)

func WithSearchQuery(q string) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["q"] = url.QueryEscape(q) }
}

func WithPageSize(pageSize uint32) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["pageSize"] = strconv.Itoa(int(pageSize)) }
}

func WithIsDeactivated(isDeactivated bool) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["isDeactivated"] = strconv.FormatBool(isDeactivated) }
}

func WithExternalVulnerabilityID(externalVulnerabilityID string) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["externalVulnerabilityId"] = externalVulnerabilityID }
}

func WithIsFixAvailable(isFixAvailable bool) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["isFixAvailable"] = strconv.FormatBool(isFixAvailable) }
}

func WithPackageIdentifier(packageIdentifier string) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["packageIdentifier"] = packageIdentifier }
}

func WithSLADeadlineAfterDate(date time.Time) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["slaDeadlineAfterDate"] = date.String() }
}

func WithSLADeadlineBeforeDate(date time.Time) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["slaDeadlineBeforeDate"] = date.String() }
}

func WithSeverity(severity VulnerabilitySeverity) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["severity"] = string(severity) }
}

func WithIntegrationID(integrationID string) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["integrationId"] = integrationID }
}

func WithIncludeVulnerabilitiesWithoutSLAs(include bool) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["includeVulnerabilitiesWithoutSlas"] = strconv.FormatBool(include) }
}

func WithVulnerableAssetID(vulnerableAssetID string) ListVulnerabilitiesOption {
	return func(m map[string]string) { m["vulnerableAssetId"] = vulnerableAssetID }
}

type ListVulnerabilitiesOutput struct {
	Results ListVulnerabilitiesResults `json:"results"`
}

type ListVulnerabilitiesResults struct {
	PageInfo PageInfo        `json:"pageInfo"`
	Data     []Vulnerability `json:"data"`
}

type Vulnerability struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	IntegrationID      string                 `json:"integrationId"`
	PackageIdentifier  *string                `json:"packageIdentifier,omitempty"`
	VulnerabilityType  VulnerabilityType      `json:"vulnerabilityType"`
	TargetID           string                 `json:"targetId"`
	FirstDetectedDate  time.Time              `json:"firstDetectedDate"`
	LastDetectedDate   *time.Time             `json:"lastDetectedDate,omitempty"`
	Severity           *VulnerabilitySeverity `json:"severity,omitempty"`
	CVSSSeverityScore  *float64               `json:"cvssSeverityScore,omitempty"`
	ScannerScore       *float64               `json:"scannerScore,omitempty"`
	IsFixable          bool                   `json:"isFixable"`
	RemediateByDate    *time.Time             `json:"remediateByDate,omitempty"`
	RelatedVulns       []string               `json:"relatedVulns,omitempty"`
	RelatedURLs        []string               `json:"relatedUrls,omitempty"`
	ExternalURL        string                 `json:"externalURL"`
	DeactivateMetadata *DeactivateMetadata    `json:"deactivateMetadata,omitempty"`
}

type DeactivateMetadata struct {
	IsVulnDeactivatedIndefinitely bool       `json:"isVulnDeactivatedIndefinitely"`
	DeactivatedUntilDate          *time.Time `json:"deactivatedUntilDate,omitempty"`
	DeactivationReason            string     `json:"deactivationReason"`
	DeactivatedOnDate             time.Time  `json:"deactivatedOnDate"`
	DeactivatedBy                 string     `json:"deactivatedBy,omitempty"`
}

type VulnerabilityType string
type VulnerabilitySeverity string

const (
	VulnerabilityTypeConfiguration VulnerabilityType = "CONFIGURATION"
	VulnerabilityTypeCommon        VulnerabilityType = "COMMON"
	VulnerabilityTypeGrouped       VulnerabilityType = "GROUPED"

	VulnerabilitySeverityCritical VulnerabilitySeverity = "CRITICAL"
	VulnerabilitySeverityHigh     VulnerabilitySeverity = "HIGH"
	VulnerabilitySeverityMedium   VulnerabilitySeverity = "MEDIUM"
	VulnerabilitySeverityLow      VulnerabilitySeverity = "LOW"
)
