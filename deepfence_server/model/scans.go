package model

import (
	"github.com/deepfence/ThreatMapper/deepfence_server/reporters"
	"github.com/deepfence/golang_deepfence_sdk/utils/utils"
)

type VulnerabilityScanConfigLanguage struct {
	Language string `json:"language" required:"true" enum:"all,base,ruby,python,javascript,php,golang,java,rust,dotnet"`
}

type VulnerabilityScanConfig struct {
	ScanConfigLanguages []VulnerabilityScanConfigLanguage `json:"scan_config" required:"true"`
}

type VulnerabilityScanTriggerReq struct {
	ScanTriggerCommon
	VulnerabilityScanConfig
}

type SecretScanTriggerReq struct {
	ScanTriggerCommon
}

type MalwareScanTriggerReq struct {
	ScanTriggerCommon
}

type ComplianceScanTriggerReq struct {
	ScanTriggerCommon
	ComplianceBenchmarkTypes
}

type ScanFilter struct {
	ImageScanFilter             reporters.ContainsFilter `json:"image_scan_filter" required:"true"`
	ContainerScanFilter         reporters.ContainsFilter `json:"container_scan_filter" required:"true"`
	HostScanFilter              reporters.ContainsFilter `json:"host_scan_filter" required:"true"`
	CloudAccountScanFilter      reporters.ContainsFilter `json:"cloud_account_scan_filter" required:"true"`
	KubernetesClusterScanFilter reporters.ContainsFilter `json:"kubernetes_cluster_scan_filter" required:"true"`
}

type ScanTriggerCommon struct {
	NodeIds []NodeIdentifier `json:"node_ids" required:"true"`
	Filters ScanFilter       `json:"filters" required:"true"`
}

type NodeIdentifier struct {
	NodeId   string `json:"node_id" required:"true"`
	NodeType string `json:"node_type" required:"true" enum:"image,host,container,cloud_account,cluster,registry"`
}

type ComplianceBenchmarkTypes struct {
	BenchmarkTypes []string `json:"benchmark_types" required:"true"`
}

type ScanStatus string

type ScanInfo struct {
	ScanId         string           `json:"scan_id" required:"true"`
	Status         string           `json:"status" required:"true"`
	UpdatedAt      int64            `json:"updated_at" required:"true" format:"int64"`
	NodeId         string           `json:"node_id" required:"true"`
	NodeType       string           `json:"node_type" required:"true"`
	SeverityCounts map[string]int32 `json:"severity_counts" required:"true"`
	NodeName       string           `json:"node_name" required:"true"`
}

type ComplianceScanInfo struct {
	ScanId         string           `json:"scan_id" required:"true"`
	BenchmarkType  string           `json:"benchmark_type" required:"true"`
	Status         string           `json:"status" required:"true"`
	UpdatedAt      int64            `json:"updated_at" required:"true" format:"int64"`
	NodeId         string           `json:"node_id" required:"true"`
	NodeType       string           `json:"node_type" required:"true"`
	SeverityCounts map[string]int32 `json:"severity_counts" required:"true"`
}

const (
	SCAN_STATUS_SUCCESS    = utils.SCAN_STATUS_SUCCESS
	SCAN_STATUS_STARTING   = utils.SCAN_STATUS_STARTING
	SCAN_STATUS_INPROGRESS = utils.SCAN_STATUS_INPROGRESS
)

type ScanTriggerResp struct {
	ScanIds    []string `json:"scan_ids" required:"true"`
	BulkScanId string   `json:"bulk_scan_id" required:"true"`
}

type ScanStatusReq struct {
	ScanIds    []string `json:"scan_ids" required:"true"`
	BulkScanId string   `json:"bulk_scan_id" required:"true"`
}

type ScanStatusResp struct {
	Statuses map[string]ScanInfo `json:"statuses" required:"true"`
}

type ComplianceScanStatusResp struct {
	Statuses []ComplianceScanInfo `json:"statuses" required:"true"`
}

type ScanListReq struct {
	NodeIds []NodeIdentifier `json:"node_ids" required:"true"`
	Window  FetchWindow      `json:"window"  required:"true"`
}

type ScanListResp struct {
	ScansInfo []ScanInfo `json:"scans_info" required:"true"`
}

type CloudComplianceScanListResp struct {
	ScansInfo []ComplianceScanInfo `json:"scans_info" required:"true"`
}

type ScanResultsMaskRequest struct {
	ScanID                   string   `json:"scan_id" validate:"required" required:"true"`
	DocIds                   []string `json:"doc_ids" validate:"required,gt=0,dive,min=3" required:"true"`
	ScanType                 string   `json:"scan_type" validate:"required,oneof=SecretScan VulnerabilityScan MalwareScan ComplianceScan CloudComplianceScan" required:"true" enum:"SecretScan,VulnerabilityScan,MalwareScan,ComplianceScan,CloudComplianceScan"`
	MaskAcrossHostsAndImages bool     `json:"mask_across_hosts_and_images"`
}

type ScanResultsActionRequest struct {
	ScanID   string   `json:"scan_id" validate:"required" required:"true"`
	DocIds   []string `json:"doc_ids" validate:"required,gt=0,dive,min=3" required:"true"`
	ScanType string   `json:"scan_type" validate:"required,oneof=SecretScan VulnerabilityScan MalwareScan ComplianceScan CloudComplianceScan" required:"true" enum:"SecretScan,VulnerabilityScan,MalwareScan,ComplianceScan,CloudComplianceScan"`
	//utils.Neo4jScanType
}

type DownloadReportResponse struct {
	UrlLink string `json:"url_link"`
}

type ScanActionRequest struct {
	ScanID   string `path:"scan_id" validate:"required" required:"true"`
	ScanType string `path:"scan_type" validate:"required,oneof=SecretScan VulnerabilityScan MalwareScan ComplianceScan CloudComplianceScan" required:"true" enum:"SecretScan,VulnerabilityScan,MalwareScan,ComplianceScan,CloudComplianceScan"`
	//utils.Neo4jScanType
}

type ScanResultDocumentRequest struct {
	DocId    string `path:"doc_id" validate:"required" required:"true"`
	ScanID   string `path:"scan_id" validate:"required" required:"true"`
	ScanType string `path:"scan_type" validate:"required,oneof=SecretScan VulnerabilityScan MalwareScan ComplianceScan CloudComplianceScan" required:"true" enum:"SecretScan,VulnerabilityScan,MalwareScan,ComplianceScan,CloudComplianceScan"`
	//utils.Neo4jScanType
}

type SbomRequest struct {
	// either scan_id or node_id+node_type is required
	ScanID   string `json:"scan_id"`
	NodeID   string `json:"node_id"`
	NodeType string `json:"node_type"`
}

type SbomResponse struct {
	PackageName string   `json:"package_name"`
	Version     string   `json:"version"`
	Locations   []string `json:"locations"`
	Licenses    []string `json:"licenses"`
	CveID       string   `json:"cve_id"`
	Severity    string   `json:"severity"`
}

type ScanResultsReq struct {
	ScanId       string                  `json:"scan_id" required:"true"`
	FieldsFilter reporters.FieldsFilters `json:"fields_filter" required:"true"`
	Window       FetchWindow             `json:"window"  required:"true"`
}

type ScanResultsCommon struct {
	ContainerName         string `json:"docker_container_name" required:"true"`
	ImageName             string `json:"docker_image_name" required:"true"`
	HostName              string `json:"host_name" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name" required:"true"`
	NodeID                string `json:"node_id" required:"true"`
	NodeName              string `json:"node_name" required:"true"`
	NodeType              string `json:"node_type" required:"true"`
	ScanID                string `json:"scan_id" required:"true"`
	UpdatedAt             int64  `json:"updated_at" required:"true" format:"int64"`
}

type SecretScanResult struct {
	ScanResultsCommon
	Secrets        []Secret         `json:"secrets" required:"true"`
	Rules          []Rule           `json:"rules" required:"true"`
	RuleSecrets    map[int][]int32  `json:"rule_2_secrets" required:"true"`
	SeverityCounts map[string]int32 `json:"severity_counts" required:"true"`
}

type VulnerabilityScanResult struct {
	ScanResultsCommon
	Vulnerabilities []Vulnerability  `json:"vulnerabilities" required:"true" required:"true"`
	SeverityCounts  map[string]int32 `json:"severity_counts" required:"true"`
}

type MalwareScanResult struct {
	ScanResultsCommon
	Malwares       []Malware        `json:"malwares" required:"true"`
	SeverityCounts map[string]int32 `json:"severity_counts" required:"true"`
}

type ComplianceScanResult struct {
	ScanResultsCommon
	ComplianceAdditionalInfo
	Compliances []Compliance `json:"compliances" required:"true"`
}

type ComplianceAdditionalInfo struct {
	BenchmarkType        string           `json:"benchmark_type" required:"true"`
	StatusCounts         map[string]int32 `json:"status_counts" required:"true"`
	CompliancePercentage float64          `json:"compliance_percentage" required:"true"`
}

type CloudComplianceScanResult struct {
	ScanResultsCommon
	ComplianceAdditionalInfo
	Compliances []CloudCompliance `json:"compliances" required:"true"`
}

type Secret struct {
	StartingIndex         int32  `json:"starting_index" required:"true"`
	RelativeStartingIndex int32  `json:"relative_starting_index" required:"true"`
	RelativeEndingIndex   int32  `json:"relative_ending_index" required:"true"`
	FullFilename          string `json:"full_filename" required:"true"`
	MatchedContent        string `json:"matched_content" required:"true"`
}

func (Secret) NodeType() string {
	return "Secret"
}

func (v Secret) GetCategory() string {
	return ""
}

type Rule struct {
	ID               int32   `json:"id" required:"true" required:"true"`
	Name             string  `json:"name" required:"true"`
	Part             string  `json:"part" required:"true"`
	SignatureToMatch string  `json:"signature_to_match" required:"true"`
	Level            string  `json:"level" required:"true"`
	Score            float64 `json:"score" required:"true"`
}

type Vulnerability struct {
	Cve_id                     string   `json:"cve_id" required:"true"`
	Cve_type                   string   `json:"cve_type" required:"true"`
	Cve_severity               string   `json:"cve_severity" required:"true"`
	Cve_caused_by_package      string   `json:"cve_caused_by_package" required:"true"`
	Cve_caused_by_package_path string   `json:"cve_caused_by_package_path" required:"true"`
	Cve_container_layer        string   `json:"cve_container_layer" required:"true"`
	Cve_fixed_in               string   `json:"cve_fixed_in" required:"true"`
	Cve_link                   string   `json:"cve_link" required:"true"`
	Cve_description            string   `json:"cve_description" required:"true"`
	Cve_cvss_score             float64  `json:"cve_cvss_score" required:"true"`
	Cve_overall_score          float64  `json:"cve_overall_score" required:"true"`
	Cve_attack_vector          string   `json:"cve_attack_vector" required:"true"`
	URLs                       []string `json:"urls" required:"true"`
	ExploitPOC                 string   `json:"exploit_poc" required:"true"`
}

func (Vulnerability) NodeType() string {
	return "Vulnerability"
}

func (v Vulnerability) GetCategory() string {
	return v.Cve_severity
}

type Malware struct {
	ImageLayerID     string  `json:"image_layer_id" required:"true"`
	Class            string  `json:"class" required:"true"`
	CompleteFilename string  `json:"complete_filename" required:"true"`
	FileSevScore     float64 `json:"file_sevScore" required:"true"`
	FileSeverity     string  `json:"file_severity" required:"true"`
	SeverityScore    float64 `json:"severity_score" required:"true"`
	Summary          string  `json:"summary" required:"true"`
	RuleName         string  `json:"rule_name" required:"true"`
}

func (Malware) NodeType() string {
	return "Malware"
}

func (v Malware) GetCategory() string {
	return v.Class
}

type Compliance struct {
	TestCategory        string `json:"test_category" required:"true"`
	TestNumber          string `json:"test_number" required:"true"`
	TestInfo            string `json:"description" required:"true"`
	RemediationScript   string `json:"remediation_script,omitempty" required:"true"`
	RemediationAnsible  string `json:"remediation_ansible,omitempty" required:"true"`
	RemediationPuppet   string `json:"remediation_puppet,omitempty" required:"true"`
	Resource            string `json:"resource" required:"true"`
	TestRationale       string `json:"test_rationale" required:"true"`
	TestSeverity        string `json:"test_severity" required:"true"`
	TestDesc            string `json:"test_desc" required:"true"`
	Status              string `json:"status" required:"true"`
	ComplianceCheckType string `json:"compliance_check_type" required:"true"`
	ComplianceNodeType  string `json:"compliance_node_type" required:"true"`
}

func (Compliance) NodeType() string {
	return "Compliance"
}

func (v Compliance) GetCategory() string {
	return v.TestSeverity
}

type CloudCompliance struct {
	Timestamp           string `json:"@timestamp" required:"true"`
	Count               int32  `json:"count,omitempty" required:"true"`
	Reason              string `json:"reason" required:"true"`
	Resource            string `json:"resource" required:"true"`
	Status              string `json:"status" required:"true"`
	Region              string `json:"region" required:"true"`
	AccountID           string `json:"account_id" required:"true"`
	Group               string `json:"group" required:"true"`
	Service             string `json:"service" required:"true"`
	Title               string `json:"title" required:"true"`
	ComplianceCheckType string `json:"compliance_check_type" required:"true"`
	CloudProvider       string `json:"cloud_provider" required:"true"`
	NodeName            string `json:"node_name" required:"true"`
	NodeID              string `json:"node_id" required:"true"`
	ScanID              string `json:"scan_id" required:"true"`
	Masked              string `json:"masked" required:"true"`
	Type                string `json:"type" required:"true"`
	ControlID           string `json:"control_id" required:"true"`
	Description         string `json:"description" required:"true"`
	Severity            string `json:"severity" required:"true"`
}

func (CloudCompliance) NodeType() string {
	return "CloudCompliance"
}

func (v CloudCompliance) GetCategory() string {
	return v.Severity
}
