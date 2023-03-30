package report

import (
	"fmt"
	"math/rand"
	"time"
)

// Names of the various topologies.
const (
	Endpoint          = "endpoint"
	Process           = "process"
	Container         = "container"
	Pod               = "pod"
	Service           = "service"
	Namespace         = "namespace"
	ContainerImage    = "container_image"
	CloudProvider     = "cloud_provider"
	CloudRegion       = "cloud_region"
	Host              = "host"
	Overlay           = "overlay"
	KubernetesCluster = "kubernetes_cluster"

	// Shapes used for different nodes
	Circle         = "circle"
	Triangle       = "triangle"
	Square         = "square"
	Pentagon       = "pentagon"
	Hexagon        = "hexagon"
	Heptagon       = "heptagon"
	Octagon        = "octagon"
	Cloud          = "cloud"
	Storage        = "storage"
	Cylinder       = "cylinder"
	DottedCylinder = "dottedcylinder"
	StorageSheet   = "sheet"
	Camera         = "camera"
	DottedTriangle = "dottedtriangle"
	// AWS
	EBS                    = "ebs"
	EFS                    = "elasticfilesystem"
	FSx                    = "fsx"
	S3                     = "s3"
	Backup                 = "backup"
	Snowball               = "snowball"
	StorageGateway         = "storagegateway"
	RDS                    = "rds"
	DynamoDB               = "dynamodb"
	ElastiCache            = "elasticache"
	Elastisearch           = "elastisearch"
	Keyspaces              = "cassandra"
	QLDB                   = "qldb"
	ApplicationAutoScaling = "application-autoscaling"
	EC2                    = "ec2"
	Lightsail              = "lightsail"
	AWSBatch               = "batch"
	ComputeOptimizer       = "compute-optimizer"
	ElasticBeanstalk       = "elasticbeanstalk"
	Fargate                = "ecs"
	Lambda                 = "lambda"
	Outposts               = "outposts"
	ServerlessRepo         = "serverlessrepo"
	CloudDirectory         = "clouddirectory"
	CloudFront             = "cloudfront"
	Route53                = "route53"
	AppMesh                = "appmesh"
	CloudMap               = "servicediscovery"
	DirectConnect          = "directconnect"
	GlobalAccelerator      = "globalaccelerator"
	ElasticLoadBalancing   = "elasticloadbalancing"
	Cognito                = "cognito"
	Detective              = "detective"
	GuardDuty              = "guardduty"
	Inspector              = "inspector"
	Macie                  = "macie"
	Artifact               = "codeartifact"
	AuditManager           = "auditmanager"
	CertificateManager     = "acm"
	CloudHSM               = "cloudhsm"
	DirectoryService       = "ds"
	FirewallManager        = "fms"
	AWSIAM                 = "aws-iam"
	KMS                    = "kms"
	NetworkManager         = "networkmanager"
	SecretsManager         = "secretsmanager"
	SecurityHub            = "securityhub"
	Shield                 = "shield"
	Signer                 = "signer"
	SSO                    = "sso"
	WAF                    = "waf"
	// GCP
	CloudSpanner          = "spanner"
	CloudStorage          = "cloudstorage"
	Filestore             = "filestore"
	CloudBigtable         = "cloudbigtable"
	CloudSQL              = "cloudsql"
	Datastore             = "datastore"
	Memorystore           = "memorystore"
	AppEngine             = "appengine"
	CloudFunctions        = "cloudfunctions"
	CloudDNS              = "clouddns"
	VirtualPrivateCloud   = "vpc"
	DataLossPreventionAPI = "dlp"
	GCPIAM                = "gcp-iam"
	IdentityAwareProxy    = "iap"
	IdentityPlatform      = "identitytoolkit"
	CloudKMS              = "cloudkms"
	SecurityCommandCenter = "securitycenter"
	WebSecurityScanner    = "websecurityscanner"
	// Azure
	ActiveDirectory         = "azure_active_directory"
	AzureApiManagement      = "azure_api_management"
	AzureBlobStorage        = "azure_blob_storage"
	AzureCDN                = "azure_cdn"
	AzureCloudServices      = "azure_cloud_services"
	AzureContainerRegistry  = "azure_container_registry"
	AzureFiles              = "azure_files"
	AzureFrontDoor          = "azure_front_door"
	AzureManagementServices = "azure_management_services"
	AzureMediaServices      = "azure_media_services"
	AzureMobileApps         = "azure_mobile_apps"
	AzureQueueStorage       = "azure_queue_storage"
	AzureServiceBus         = "azure_service_bus"
	AzureStackEdge          = "azure_stack_edge"
	AzureSQLDatabase        = "azure_sql_database"
	AzureTableStorage       = "azure_table_storage"
	AzureTrafficManager     = "azure_traffic_manager"
	AzureWebsites           = "azure_websites"
	VisualStudioCodespaces  = "visual_studio_codespaces"
	// Dev tools/clouds
	Atlassian     = "atlassian"
	Bootstrap     = "bootstrap"
	Bugsnag       = "bugsnag"
	Docker        = "docker"
	GitHub        = "github"
	Gitlab        = "gitlab"
	Heroku        = "heroku"
	Hostgator     = "hostgator"
	IANABlackhole = "ianablackhole"
	LetsEncrypt   = "letsencrypt"
	NodeJS        = "nodejs"
	NTP           = "ntp"
	ReverseDNS    = "reversedns"
	// File sharing
	Box         = "box"
	Dropbox     = "dropbox"
	FileFactory = "filefactory"
	FourShared  = "fourshared" // 4shared
	GoogleDrive = "googledrive"
	MSOnedrive  = "onedrive"
	Pastebin    = "pastebin"
	// APM/Monitoring/Third Party
	Amazon         = "amazon"
	AmazonAlexa    = "alexa"
	AppleAPIs      = "appleapis"
	ATT            = "att" // at&t
	Autodesk       = "autodesk"
	DataDog        = "datadog"
	FedexAPIs      = "fedexapis"
	FreshworksAPIs = "freshworks"
	GoogleCloud    = "gcloud"
	GoogleSuite    = "gsuite"
	Grafana        = "grafana"
	LogicMonitor   = "logicmonitor"
	NewRelic       = "newrelic"
	OpenDNS        = "opendns"
	Pingdom        = "pingdom"
	Pubnub         = "pubnub"
	Sentry         = "sentry"
	Servicenow     = "servicenow"
	Slack          = "slack"
	SolarWinds     = "solorwinds"
	Splunk         = "splunk"
	Trello         = "trello"
	Twilio         = "twilio"
	Wix            = "wix"
	Wordpress      = "wordpress"
	ZenDesk        = "zendesk"
	Zoho           = "zoho"
	// Databases
	Confluent = "confluent"
	Elastic   = "elastic"
	MongoDB   = "mongodb"
	Redis     = "redis"
	Snowflake = "snowflake"
	// CDN
	AdobeAds         = "adobeads"
	Akamai           = "akamai"
	AmazonAds        = "amazonads"
	AmazonCloudfront = "amazoncloudfront"
	AOL              = "aol"
	BellCanada       = "bellcanada"
	BranchIO         = "branchio"
	CDN77            = "cdn77"
	Changeip         = "changeip"
	CloudFlare       = "cloudflare"
	Discord          = "discord"
	Discuss          = "discuss"
	DYN              = "dyn"
	Fastly           = "fastly"
	FontAwesome      = "fontawesome"
	GoDaddy          = "godaddy"
	GoogleAds        = "googleads"
	GoogleAnalytics  = "googleanalytics"
	GoogleDomains    = "googledomains"
	Hubspot          = "hubspot"
	Intercom         = "intercom"
	Jquery           = "jquery"
	JSDeliver        = "jsdeliver"
	KeyCDN           = "keycdn"
	Linode           = "linode"
	LumenCDN         = "lumencdn"
	Mailchimp        = "mailchimp"
	MicrosoftTeams   = "microsoftteams"
	Mixpanel         = "mixpanel"
	NetlifyCDN       = "netlifycdn"
	Orange           = "orange"
	Outbrain         = "outbrain"
	Pubmatic         = "pubmatic"
	Salesforce       = "saleforce"
	Segment          = "segment"
	Sendgrid         = "sendgrid"
	StackPath        = "stackpath"
	Tmobile          = "tmobile"
	Verizon          = "verizon"
	Vodafone         = "vodafone"
	Yandex           = "yandex"

	// Used when counting the number of containers
	ContainersKey = "containers"
)

// topologyNames are the names of all report topologies.
var topologyNames = []string{
	Endpoint,
	Process,
	Container,
	ContainerImage,
	CloudProvider,
	CloudRegion,
	KubernetesCluster,
	Pod,
	Service,
	Namespace,
	Host,
	Overlay,
}

var topologyAdjacencyNames = []string{
	Endpoint,
}

var topologyParentNames = []string{
	Endpoint,
	Process,
	Container,
	ContainerImage,
	CloudRegion,
	KubernetesCluster,
	Pod,
	Service,
	Namespace,
	Host,
	Overlay,
}

var topologySetNames = []string{
	Container,
	Overlay,
}

type TopologyAdjacency map[string]IDList

func (t TopologyAdjacency) Copy() TopologyAdjacency {
	newTopologyAdjacency := MakeTopologyAdjacency()
	for k, v := range t {
		newTopologyAdjacency[k] = v
	}
	return newTopologyAdjacency
}

func (t TopologyAdjacency) UnsafeMerge(o TopologyAdjacency) {
	for k, v := range o {
		t[k] = v
	}
}

func (t TopologyAdjacency) UnsafeUnMerge(o TopologyAdjacency) {

}

func (t TopologyAdjacency) AddAdjacency(nodeId string, id string) {
	if _, ok := t[nodeId]; !ok {
		t[nodeId] = MakeIDList(id)
	} else {
		t[nodeId].Add(id)
	}
}

func MakeTopologyAdjacency() TopologyAdjacency {
	return make(map[string]IDList)
}

type TopologySets map[string]Sets

func MakeTopologySets() TopologySets {
	return make(map[string]Sets)
}

func (p TopologySets) AddSet(nodeId string, sets Sets) {
	p[nodeId] = sets
}

func (t TopologySets) Copy() TopologySets {
	newTopologySets := MakeTopologySets()
	for k, v := range t {
		newTopologySets[k] = v
	}
	return newTopologySets
}

func (t TopologySets) UnsafeMerge(o TopologySets) {
	for k, v := range o {
		t[k] = v
	}
}

func (t TopologySets) UnsafeUnMerge(o TopologySets) {

}

type Parents map[string]Parent

func MakeParents() Parents {
	return make(map[string]Parent)
}

func (p Parents) AddParent(nodeId string, parents Parent) {
	p[nodeId] = parents
}

func (t Parents) Merge(o Parents) {
	for k, v := range o {
		t[k] = v
	}
}

func (t Parents) Copy() Parents {
	newParents := MakeParents()
	for k, v := range t {
		newParents[k] = v
	}
	return newParents
}

func (t Parents) UnsafeMerge(o Parents) {
	for k, v := range o {
		t[k] = v
	}
}

func (t Parents) UnsafeUnMerge(o Parents) {

}

type Parent struct {
	CloudProvider     string `json:"cloud_provider,omitempty"`
	CloudRegion       string `json:"cloud_region,omitempty"`
	KubernetesCluster string `json:"kubernetes_cluster,omitempty"`
	Host              string `json:"host,omitempty"`
	Container         string `json:"container,omitempty"`
	ContainerImage    string `json:"container_image,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	Pod               string `json:"pod,omitempty"`
}

// Report is the core data type. It's produced by probes, and consumed and
// stored by apps. It's composed of multiple topologies, each representing
// a different (related, but not equivalent) view of the network.
type Report struct {
	// TS is the time this report was generated
	TS time.Time

	// Endpoint nodes are individual (address, port) tuples on each host.
	// They come from inspecting active connections and can (theoretically)
	// be traced back to a process. Edges are present.
	Endpoint          Topology
	EndpointAdjacency TopologyAdjacency
	EndpointParents   Parents

	// Process nodes are processes on each host. Edges are not present.
	Process        Topology
	ProcessParents Parents

	// Container nodes represent all Docker containers on hosts running probes.
	// Metadata includes things like containter id, name, image id etc.
	// Edges are not present.
	Container        Topology
	ContainerParents Parents
	ContainerSets    TopologySets

	// CloudProvider nodes represent all cloud providers.
	// Metadata includes things like name etc. Edges are not
	// present.
	CloudProvider Topology

	// CloudRegion nodes represent all cloud regions.
	// Metadata includes things like name etc. Edges are not
	// present.
	CloudRegion        Topology
	CloudRegionParents Parents

	// KubernetesCluster nodes represent all Kubernetes clusters.
	// Metadata includes things like cluster id, name etc. Edges are not
	// present.
	KubernetesCluster        Topology
	KubernetesClusterParents Parents

	// Pod nodes represent all Kubernetes pods running on hosts running probes.
	// Metadata includes things like pod id, name etc. Edges are not
	// present.
	Pod        Topology
	PodParents Parents

	// Service nodes represent all Kubernetes services running on hosts running probes.
	// Metadata includes things like service id, name etc. Edges are not
	// present.
	Service        Topology
	ServiceParents Parents

	// Namespace nodes represent all Kubernetes Namespaces running on hosts running probes.
	// Metadata includes things like Namespace id, name, etc. Edges are not
	// present.
	Namespace        Topology
	NamespaceParents Parents

	// ContainerImages nodes represent all Docker containers images on
	// hosts running probes. Metadata includes things like image id, name etc.
	// Edges are not present.
	ContainerImage        Topology
	ContainerImageParents Parents

	// Host nodes are physical hosts that run probes. Metadata includes things
	// like operating system, load, etc. The information is scraped by the
	// probes with each published report. Edges are not present.
	Host        Topology
	HostParents Parents

	// Overlay nodes are active peers in any software-defined network that's
	// overlaid on the infrastructure. The information is scraped by polling
	// their status endpoints. Edges are present.
	Overlay        Topology
	OverlayParents Parents
	OverlaySets    TopologySets

	DNS DNSRecords `json:"DNS,omitempty" deepequal:"nil==empty"`
	// Backwards-compatibility for an accident in commit 951629a / release 1.11.6.
	BugDNS DNSRecords `json:"nodes,omitempty"`

	// Window is the amount of time that this report purports to represent.
	// Windows must be carefully merged. They should only be added when
	// reports cover non-overlapping periods of time. By default, we assume
	// that's true, and add windows in merge operations. When that's not true,
	// such as in the app, we expect the component to overwrite the window
	// before serving it to consumers.
	Window time.Duration

	// Shortcut reports should be propagated to the UI as quickly as possible,
	// bypassing the usual spy interval, publish interval and app ws interval.
	Shortcut bool

	// ID a random identifier for this report, used when caching
	// rendered views of the report.  Reports with the same id
	// must be equal, but we don't require that equal reports have
	// the same id.
	ID string `deepequal:"skip"`
}

// MakeReport makes a clean report, ready to Merge() other reports into.
func MakeReport() Report {
	return Report{
		Endpoint:          MakeTopology(),
		EndpointAdjacency: MakeTopologyAdjacency(),
		EndpointParents:   MakeParents(),

		Process:        MakeTopology(),
		ProcessParents: MakeParents(),

		Container:        MakeTopology(),
		ContainerParents: MakeParents(),
		ContainerSets:    MakeTopologySets(),

		CloudProvider: MakeTopology(),

		CloudRegion:        MakeTopology(),
		CloudRegionParents: MakeParents(),

		KubernetesCluster:        MakeTopology(),
		KubernetesClusterParents: MakeParents(),

		Pod:        MakeTopology(),
		PodParents: MakeParents(),

		Service:        MakeTopology(),
		ServiceParents: MakeParents(),

		Namespace:        MakeTopology(),
		NamespaceParents: MakeParents(),

		ContainerImage:        MakeTopology(),
		ContainerImageParents: MakeParents(),

		Host:        MakeTopology(),
		HostParents: MakeParents(),

		Overlay:        MakeTopology(),
		OverlayParents: MakeParents(),
		OverlaySets:    MakeTopologySets(),

		DNS:    DNSRecords{},
		Window: 0,
		ID:     fmt.Sprintf("%d", rand.Int63()),
	}
}

// Copy returns a value copy of the report.
func (r Report) Copy() Report {
	newReport := Report{
		TS:       r.TS,
		DNS:      r.DNS.Copy(),
		Window:   r.Window,
		Shortcut: r.Shortcut,
		ID:       fmt.Sprintf("%d", rand.Int63()),
	}
	newReport.WalkPairedTopologies(&r, func(newTopology, oldTopology *Topology) {
		*newTopology = oldTopology.Copy()
	})
	newReport.WalkPairedAdjacencies(&r, func(newAdjacency *TopologyAdjacency, oldAdjacency *TopologyAdjacency) {
		*newAdjacency = oldAdjacency.Copy()
	})
	newReport.WalkPairedParents(&r, func(newParents *Parents, oldParents *Parents) {
		*newParents = oldParents.Copy()
	})
	newReport.WalkPairedSets(&r, func(newSets *TopologySets, oldSets *TopologySets) {
		*newSets = oldSets.Copy()
	})
	return newReport
}

// UnsafeMerge merges another Report into the receiver. The original is modified.
func (r *Report) UnsafeMerge(other Report) {
	// Merged report has the earliest non-zero timestamp
	if !other.TS.IsZero() && (r.TS.IsZero() || other.TS.Before(r.TS)) {
		r.TS = other.TS
	}
	r.DNS = r.DNS.Merge(other.DNS)
	r.Window = r.Window + other.Window
	r.WalkPairedTopologies(&other, func(ourTopology, theirTopology *Topology) {
		ourTopology.UnsafeMerge(*theirTopology)
	})
	r.WalkPairedAdjacencies(&other, func(ourAdjacency *TopologyAdjacency, theirAdjacency *TopologyAdjacency) {
		ourAdjacency.UnsafeMerge(*theirAdjacency)
	})
	r.WalkPairedParents(&other, func(ourParents *Parents, theirParents *Parents) {
		ourParents.UnsafeMerge(*theirParents)
	})
	r.WalkPairedSets(&other, func(ourSets *TopologySets, theirSets *TopologySets) {
		ourSets.UnsafeMerge(*theirSets)
	})
}

// UnsafeUnMerge removes any information from r that would be added by merging other.
// The original is modified.
func (r *Report) UnsafeUnMerge(other Report) {
	// TODO: DNS, Sampling, Plugins
	r.Window = r.Window - other.Window
	r.WalkPairedTopologies(&other, func(ourTopology, theirTopology *Topology) {
		ourTopology.UnsafeUnMerge(*theirTopology)
	})
	r.WalkPairedAdjacencies(&other, func(ourAdjacency *TopologyAdjacency, theirAdjacency *TopologyAdjacency) {
		ourAdjacency.UnsafeUnMerge(*theirAdjacency)
	})
	r.WalkPairedParents(&other, func(ourParents *Parents, theirParents *Parents) {
		ourParents.UnsafeUnMerge(*theirParents)
	})
	r.WalkPairedSets(&other, func(ourSets *TopologySets, theirSets *TopologySets) {
		ourSets.UnsafeUnMerge(*theirSets)
	})
}

// WalkTopologies iterates through the Topologies of the report,
// potentially modifying them
func (r *Report) WalkTopologies(f func(*Topology)) {
	for _, name := range topologyNames {
		f(r.topology(name))
	}
}

// WalkNamedTopologies iterates through the Topologies of the report,
// potentially modifying them.
func (r *Report) WalkNamedTopologies(f func(string, *Topology)) {
	for _, name := range topologyNames {
		f(name, r.topology(name))
	}
}

// WalkPairedTopologies iterates through the Topologies of this and another report,
// potentially modifying one or both.
func (r *Report) WalkPairedTopologies(o *Report, f func(*Topology, *Topology)) {
	for _, name := range topologyNames {
		f(r.topology(name), o.topology(name))
	}
}

// WalkPairedAdjacencies iterates through the TopologyAdjacency of this and another report,
// potentially modifying one or both.
func (r *Report) WalkPairedAdjacencies(o *Report, f func(*TopologyAdjacency, *TopologyAdjacency)) {
	for _, name := range topologyAdjacencyNames {
		f(r.topologyAdjacency(name), o.topologyAdjacency(name))
	}
}

// WalkPairedParents iterates through the Parents of this and another report,
// potentially modifying one or both.
func (r *Report) WalkPairedParents(o *Report, f func(*Parents, *Parents)) {
	for _, name := range topologyParentNames {
		f(r.topologyParent(name), o.topologyParent(name))
	}
}

// WalkPairedSets iterates through the TopologySets of this and another report,
// potentially modifying one or both.
func (r *Report) WalkPairedSets(o *Report, f func(*TopologySets, *TopologySets)) {
	for _, name := range topologySetNames {
		f(r.topologySet(name), o.topologySet(name))
	}
}

// topology returns a reference to one of the report's topologies,
// selected by name.
func (r *Report) topology(name string) *Topology {
	switch name {
	case Endpoint:
		return &r.Endpoint
	case Process:
		return &r.Process
	case Container:
		return &r.Container
	case ContainerImage:
		return &r.ContainerImage
	case CloudProvider:
		return &r.CloudProvider
	case CloudRegion:
		return &r.CloudRegion
	case KubernetesCluster:
		return &r.KubernetesCluster
	case Pod:
		return &r.Pod
	case Service:
		return &r.Service
	case Namespace:
		return &r.Namespace
	case Host:
		return &r.Host
	case Overlay:
		return &r.Overlay
	}
	return nil
}

func (r *Report) topologyAdjacency(name string) *TopologyAdjacency {
	switch name {
	case Endpoint:
		return &r.EndpointAdjacency
	}
	return nil
}

func (r *Report) topologyParent(name string) *Parents {
	switch name {
	case Endpoint:
		return &r.EndpointParents
	case Process:
		return &r.ProcessParents
	case Container:
		return &r.ContainerParents
	case ContainerImage:
		return &r.ContainerImageParents
	case CloudRegion:
		return &r.CloudRegionParents
	case KubernetesCluster:
		return &r.KubernetesClusterParents
	case Pod:
		return &r.PodParents
	case Service:
		return &r.ServiceParents
	case Namespace:
		return &r.NamespaceParents
	case Host:
		return &r.HostParents
	case Overlay:
		return &r.OverlayParents
	}
	return nil
}

func (r *Report) topologySet(name string) *TopologySets {
	switch name {
	case Container:
		return &r.ContainerSets
	case Overlay:
		return &r.OverlaySets
	}
	return nil
}

// Topology returns one of the report's topologies, selected by name.
func (r Report) Topology(name string) (Topology, bool) {
	if t := r.topology(name); t != nil {
		return *t, true
	}
	return Topology{}, false
}

// Sampling describes how the packet data sources for this report were
// sampled. It can be used to calculate effective sample rates. We can't
// just put the rate here, because that can't be accurately merged. Counts
// in e.g. edge metadata structures have already been adjusted to
// compensate for the sample rate.
type Sampling struct {
	Count uint64 // observed and processed
	Total uint64 // observed overall
}

// Rate returns the effective sampling rate.
func (s Sampling) Rate() float64 {
	if s.Total <= 0 {
		return 1.0
	}
	return float64(s.Count) / float64(s.Total)
}

// Merge combines two sampling structures via simple addition and returns the
// result. The original is not modified.
func (s Sampling) Merge(other Sampling) Sampling {
	return Sampling{
		Count: s.Count + other.Count,
		Total: s.Total + other.Total,
	}
}

const (
	// HostNodeID is a metadata foreign key, linking a node in any topology to
	// a node in the host topology. That host node is the origin host, where
	// the node was originally detected.
	HostNodeID = "host_node_id"
	// ControlProbeID is the random ID of the probe which controls the specific node.
	ControlProbeID = "control_probe_id"
)
