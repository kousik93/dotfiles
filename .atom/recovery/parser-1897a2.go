package parser

import (
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

/*
Main Block
*/

//ManifestData : DO NOT add omitempty tag to the following struct for certain structs. It has a bad side-effect of not marshaling "zero" value sections.
//E.g. if authorization->createOauthAccount = false, it'll not be visible in the final generated yaml.
type ManifestData struct {
	Metadata             Metadata                                `yaml:"metadata"`
	Dependencies         Dependencies                            `yaml:"dependencies,omitempty"`
	Authorization        Authorization                           `yaml:"authorization"`
	Config               Config                                  `yaml:"config"`
	Scheduling           Scheduling                              `yaml:"scheduling,omitempty"`
	CodeQuality          CodeQuality                             `yaml:"codeQuality"`
	DeploymentProfiles   map[string]DeploymentProfile            `yaml:"deploymentProfiles"`
	EnvProfiles          map[string]map[string]string            `yaml:"envProfiles,omitempty"`
	HealthProfiles       map[string]HealthProfile                `yaml:"healthProfiles"`
	ResourcesProfiles    map[string]ResourcesProfile             `yaml:"resourcesProfiles"`
	ArgumentsProfiles    map[string][]string                     `yaml:"argumentsProfiles"`
	Ports                map[string]Port                         `yaml:"ports"`
	Endpoints            map[string]Endpoint                     `yaml:"endpoints"`
	MonitoringProfiles   map[string]MonitoringProfile            `yaml:"monitoringProfiles"`
	AlertingProfiles     AlertingProfiles                        `yaml:"alertingProfiles"`
	ServiceComposition   ServiceComposition                      `yaml:"serviceComposition"`
	Environment          map[string]map[string]EnvironmentValues `yaml:"environments" json:"environments"`
	Rollout              interface{}                             `yaml:"rollout"`
	PipelineNotification PipelineNotification                    `yaml:"pipelineNotification,omitempty"`
}

/*
Metadata Block
*/

type Metadata struct {
	ManifestVersion string            `yaml:"manifestVersion"`
	MajorVersion    string            `yaml:"majorVersion,omitempty"`
	ServiceName     string            `yaml:"serviceName"`
	AppType         string            `yaml:"appType"`
	TeamNamespace   string            `yaml:"teamNamespace"`
	Tags            map[string]string `yaml:"tags,omitempty"`
	Organization    string            `yaml:"organization"`
}

/*
Dependencies Block
*/

type Dependencies struct {
	ServiceDependencies map[string]DependenciesInfo `yaml:"service,omitempty"`
	TestDependencies    map[string]DependenciesInfo `yaml:"tests,omitempty"`
}
type DependenciesInfo struct {
	Rights   []string `yaml:"rights,omitempty"`
	Location string   `yaml:"location,omitempty"`
}

/*
Authorization Block
*/

type Authorization struct {
	CreateOauthAccount         bool     `yaml:"createOauthAccount"`
	CreateOauthAccountForTests bool     `yaml:"createOauthAccountForTests"`
	DefineRights               []string `yaml:"defineRights,omitempty"`
}

/*
Config Block
*/

type Config struct {
	HasConfiguration bool `yaml:"hasConfiguration,omitempty"`
}

/*
Scheduling Block
*/
type Scheduling struct {
	InstancegroupName string `yaml:"instanceGroup,omitempty"`
}

/*
Third Party Integrations Block
*/

type Sealights struct {
	Enabled bool `yaml:"enabled"`
}
type Coverity struct {
	Enabled bool `yaml:"enabled"`
}
type Blackduck struct {
	Enabled bool `yaml:"enabled"`
}
type CodeQuality struct {
	Blackduck Blackduck `yaml:"blackduck"`
	Sealights Sealights `yaml:"sealights"`
	Coverity  Coverity  `yaml:"coverity"`
}

/*
Deployment Profile Block
*/

type DeploymentProfile struct {
	Timeout           string    `yaml:"timeout,omitempty"`
	Replicas          Replicas  `yaml:"replicas"`
	MaxUnavailability string    `yaml:"maxUnavailability"`
	MaxSurge          string    `yaml:"maxSurge"`
	Autoscale         Autoscale `yaml:"autoscale,omitempty"`
}

type Autoscale struct {
	Thresholds map[string]AutoscaleThresholdParams `yaml:"thresholds"`
}

type Replicas struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type AutoscaleThresholdParams struct {
	Percentage int    `yaml:"percentage,omitempty"`
	Value      string `yaml:"value,omitempty"`
}

/*
Health Profile
*/

type HealthProfile struct {
	LivenessProbe  LivenessProbe  `yaml:"livenessProbe,omitempty"`
	ReadinessProbe ReadinessProbe `yaml:"readinessProbe,omitempty"`
}

type ProbeCommon struct {
	Exec                Exec `yaml:"exec,omitempty"`
	InitialDelaySeconds int  `yaml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       int  `yaml:"periodSeconds,omitempty"`
	FailureThreshold    int  `yaml:"failureThreshold,omitempty"`
	TimeoutSeconds      int  `yaml:"timeoutSeconds,omitempty"`
}
type LivenessProbe struct {
	ProbeCommon `yaml:",inline"`
	HttpGet     HttpGet `yaml:"httpGet,omitempty"`
}
type ReadinessProbe struct {
	ProbeCommon `yaml:",inline"`
	HttpGet     HttpGet `yaml:"httpGet,omitempty"`
}
type Exec struct {
	Command []string `yaml:"command"`
}
type HttpGet struct {
	Path string `yaml:"path"`
	Port string `yaml:"port"`
}

/*
Resources Block
*/
type ResourcesProfile struct {
	Cpu    Cpu    `yaml:"cpu"`
	Memory Memory `yaml:"memory"`
}

type Cpu struct {
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

type Memory struct {
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

/*
Port Info
*/
type Port struct {
	Port   int    `yaml:"port"`
	Scheme string `yaml:"scheme"`
}

/*
Endpoint Info block
*/
type Endpoint struct {
	Prefix            string            `yaml:"prefix"`
	Domain            string            `yaml:"domain"`
	Paths             []Path            `yaml:"paths"`
	Access            string            `yaml:"access"`
	SessionAffinity   string            `yaml:"sessionAffinity"`
	Certificate       Certificate       `yaml:"certificate"`
	CustomAnnotations map[string]string `yaml:"customAnnotations"`
	Discovery         Discovery         `yaml:"discovery"`
}

type Path struct {
	Port        string `yaml:"port"`
	Path        string `yaml:"path"`
	DestService string `yaml:"destService"`
}

type Certificate struct {
	SecretName string `yaml:"secretName"`
}

type Discovery struct {
	Enabled     bool   `yaml:"enabled"`
	ServiceType string `yaml:"serviceType"`
	Pool        string `yaml:"pool"`
	Weight      int    `yaml:"weight"`
}

/*
Monitoring Block
*/

type MonitoringProfile struct {
	MonitoringEndpoints map[string]MonitoringEndpoint `yaml:"ports" json:"ports"`
}

type MonitoringEndpoint struct {
	ScrapeInterval string `yaml:"scrapeInterval"`
}

/*
Alerting Block
*/
type AlertingProfiles struct {
	CustomRules CustomRules             `yaml:"customRules,omitempty"`
	Profiles    map[string]AlertProfile `yaml:"profiles"`
}

type CustomRules struct {
	Prometheus []Prometheus `yaml:"prometheus"`
}

type Prometheus struct {
	Name        string `yaml:"name"`
	Condition   string `yaml:"condition"`
	PersistsFor string `yaml:"persistsFor,omitempty"`
	Dashboard   string `yaml:"dashboard,omitempty"`
	Runbook     string `yaml:"runbook,omitempty"`
	Channel     string `yaml:"channel,omitempty"`
	Severity    string `yaml:"severity,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type AlertProfile struct {
	DefaultAlerts DefaultAlerts `yaml:"defaultAlerts"`
	CustomAlerts  CustomAlerts  `yaml:"customAlerts"`
	Destinations  Destinations  `yaml:"destinations,omitempty"`
}

type DefaultAlerts struct {
	CpuWarningThreshold         float32 `yaml:"cpuWarningThreshold,omitempty"`
	CpuCriticalThreshold        float32 `yaml:"cpuCriticalThreshold,omitempty"`
	Ingress4xxWarningThreshold  int     `yaml:"ingress4xxWarningThreshold,omitempty"`
	Ingress4xxCriticalThreshold int     `yaml:"ingress4xxCriticalThreshold,omitempty"`
	Ingress5xxWarningThreshold  int     `yaml:"ingress5xxWarningThreshold,omitempty"`
	Ingress5xxCriticalThreshold int     `yaml:"ingress5xxCriticalThreshold,omitempty"`
}

type CustomAlerts struct {
	Prometheus CustomAlertsPrometheus `yaml:"prometheus,omitempty"`
}

type CustomAlertsPrometheus struct {
	CustomVariables map[string]string `yaml:"customVariables,omitempty"`
}

type Destinations struct {
	Hipchat Hipchat `yaml:"hipchat"`
	Email   Email   `yaml:"email"`
}

type Hipchat struct {
	Room int `yaml:"roomNumber"`
}
type Email struct {
	Recipients []string `yaml:"recipients"`
}

/*
ServiceComposition Section  Block
*/

type ServiceComposition struct {
	Tools          Tools          `yaml:"tools"`
	CustomPackages CustomPackages `yaml:"customPackages,omitempty"`
	Service        Service        `yaml:"service"`
	Tests          Tests          `yaml:"tests,omitempty"`
}

type Tools struct {
	Jdk string `yaml:"jdk,omitempty"`
}

type CustomPackages struct {
	Helm []CustomHelmPackage `yaml:"helm"`
}

type CustomHelmPackage struct {
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	Version string `yaml:"version"`
}

//NOTE: We are using json tags here since "github.com/miracl/conflate" Unmarshal function only supports json tags.
type Service struct {
	EfsVolumes        []EfsVolume         `yaml:"efsVolumes,omitempty"`
	Dependencies      ServiceDependencies `yaml:"dependencies,omitempty"`
	ServiceContainers []ServiceContainer  `yaml:"containers" json:"containers"`
	InitContainers    []InitContainer     `yaml:"init,omitempty" json:"init,omitempty"`
}

type EfsVolume struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
}

type ServiceDependencies struct {
	Helm []HelmDependency `yaml:"helm"`
}

type HelmDependency struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
	Version    string `yaml:"version"`
}

type ServiceContainer struct {
	Name          string   `yaml:"name"`
	DockerfileDir string   `yaml:"dockerfileDir,omitempty"`
	Docker        Docker   `yaml:"docker,omitempty"`
	ContextPath   string   `yaml:"contextPath,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	Commands      []string `yaml:"commands,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	Health        string   `yaml:"health"`
	Logfiles      []string `yaml:"logfiles,omitempty"`
}

type InitContainer struct {
	Name          string   `yaml:"name"`
	DockerfileDir string   `yaml:"dockerfileDir,omitempty"`
	Docker        Docker   `yaml:"docker,omitempty"`
	ContextPath   string   `yaml:"contextPath,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	Commands      []string `yaml:"commands,omitempty"`
}

type Tests struct {
	TestSets []Test `yaml:"testSets,omitempty"`
}

type TestDependencies struct {
	Helm []HelmDependency `yaml:"helm"`
}

//NOTE: We are using json tags here since "github.com/miracl/conflate" Unmarshal function only supports json tags.
type Test struct {
	Name           string           `yaml:"name"`
	TestContainers []TestContainers `yaml:"containers" json:"containers"`
	InitContainers []InitContainer  `yaml:"init,omitempty" json:"init,omitempty"`
	Dependencies   TestDependencies `yaml:"dependencies,omitempty"`
}

type TestContainers struct {
	Name          string   `yaml:"name"`
	DockerfileDir string   `yaml:"dockerfileDir,omitempty"`
	Docker        Docker   `yaml:"docker,omitempty"`
	ContextPath   string   `yaml:"contextPath,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	Commands      []string `yaml:"commands,omitempty"`
}

type Docker struct {
	Directory string `yaml:"directory"`
	Arguments string `yaml:"arguments,omitempty"`
}

/*
Environment Block
*/

type EnvironmentValues struct {
	Deployment            string                 `yaml:"deployment"`
	EndpointSuffix        string                 `yaml:"endpointSuffix"`
	DisableEndpoints      []string               `yaml:"disableEndpoints"`
	DisableIngressOnPorts []string               `yaml:"disableIngressOnPorts"`
	Resources             map[string]string      `yaml:"resources"`
	Arguments             map[string]string      `yaml:"arguments"`
	Env                   string                 `yaml:"env"`
	RunTests              []RunTest              `yaml:"runTests"`
	PackageValues         map[string]interface{} `yaml:"packageValues"`
	Monitoring            string                 `yaml:"monitoring"`
	Alerting              string                 `yaml:"alerting"`
	DisableAwsSpot        bool                   `yaml:"DisableAwsSpot,omitempty"`
}

type RunTest struct {
	Name    string `yaml:"name"`
	Timeout string `yaml:"timeout,omitempty"`
}

type PipelineNotification struct {
	Email   []string `yaml:"email"`
	Hipchat []uint   `yaml:"hipchat"`
}

//Parse expects the Absolute path of the cloud15.yaml file
func Parse(cloud15File string) (*ManifestData, error) {
	//Get the cloud15.yaml file
	var manifestData ManifestData

	file, err := os.Open(cloud15File)
	if err != nil {
		return nil, errors.Errorf("Error opening file: %s", cloud15File)
	}
	err = yaml.NewDecoder(file).Decode(&manifestData)
	if err != nil {
		return nil, errors.Errorf("Error in decoding file: %s : %v", cloud15File, err)
	}

	return &manifestData, nil
}

func Write(file string, manifestData *ManifestData) error {
	f, err := os.Create(file)
	if err != nil {
		return errors.Errorf("Error creating file: %s : %v", file, err)
	}
	defer f.Close()
	err = yaml.NewEncoder(f).Encode(&manifestData)
	if err != nil {
		return errors.Errorf("Error in encoding file: %s : %v", file, err)
	}
	return nil
}
