// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties of the image repository for `Source.fromAsset()`.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetProps struct {
	// Represents the docker image asset.
	// Experimental.
	Asset awsecrassets.DockerImageAsset `json:"asset"`
	// The image configuration for the image built from the asset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration"`
}

// Represents the source from local assets.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetSource interface {
	Source
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for AssetSource
type jsiiProxy_AssetSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewAssetSource(props *AssetProps) AssetSource {
	_init_.Initialize()

	j := jsiiProxy_AssetSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetSource_Override(a AssetSource, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		[]interface{}{props},
		a,
	)
}

// Source from local assets.
// Experimental.
func AssetSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func AssetSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func AssetSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func AssetSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the Job is initialized to allow this object to bind.
// Experimental.
func (a *jsiiProxy_AssetSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html
//
// Experimental.
type CodeConfiguration struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `json:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a apprunner.yaml file in the
	// source code repository (or ignoring the file if it exists).
	// Experimental.
	ConfigurationValues *CodeConfigurationValues `json:"configurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities,
// use a `apprunner.yaml` file in the source code repository.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeConfigurationValues struct {
	// The command App Runner runs to build your application.
	// Experimental.
	BuildCommand *string `json:"buildCommand"`
	// The environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *string `json:"port"`
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents
	// a programming language runtime.
	// Experimental.
	Runtime Runtime `json:"runtime"`
	// The command App Runner runs to start your application.
	// Experimental.
	StartCommand *string `json:"startCommand"`
}

// Properties of the CodeRepository.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeRepositoryProps struct {
	// Configuration for building and running the service from a source code repository.
	// Experimental.
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration"`
	// The App Runner connection for GitHub.
	// Experimental.
	Connection GitHubConnection `json:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// Experimental.
	SourceCodeVersion *SourceCodeVersion `json:"sourceCodeVersion"`
}

// The source of the App Runner configuration.
// Experimental.
type ConfigurationSourceType string

const (
	ConfigurationSourceType_API ConfigurationSourceType = "API"
	ConfigurationSourceType_REPOSITORY ConfigurationSourceType = "REPOSITORY"
)

// The number of CPU units reserved for each instance of your App Runner service.
//
// TODO: EXAMPLE
//
// Experimental.
type Cpu interface {
	Unit() *string
}

// The jsii proxy struct for Cpu
type jsiiProxy_Cpu struct {
	_ byte // padding
}

func (j *jsiiProxy_Cpu) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom CPU unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-cpu
//
// Experimental.
func Cpu_Of(unit *string) Cpu {
	_init_.Initialize()

	var returns Cpu

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Cpu_ONE_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"ONE_VCPU",
		&returns,
	)
	return returns
}

func Cpu_TWO_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"TWO_VCPU",
		&returns,
	)
	return returns
}

// Properties of the image repository for `Source.fromEcr()`.
//
// TODO: EXAMPLE
//
// Experimental.
type EcrProps struct {
	// The image configuration for the image from ECR.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration"`
	// Represents the ECR repository.
	// Experimental.
	Repository awsecr.IRepository `json:"repository"`
	// Image tag.
	// Experimental.
	Tag *string `json:"tag"`
}

// Properties of the image repository for `Source.fromEcrPublic()`.
//
// TODO: EXAMPLE
//
// Experimental.
type EcrPublicProps struct {
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration"`
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `json:"imageIdentifier"`
}

// Represents the service source from ECR Public.
//
// TODO: EXAMPLE
//
// Experimental.
type EcrPublicSource interface {
	Source
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrPublicSource
type jsiiProxy_EcrPublicSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrPublicSource(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	j := jsiiProxy_EcrPublicSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrPublicSource_Override(e EcrPublicSource, props *EcrPublicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrPublicSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrPublicSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrPublicSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrPublicSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the Job is initialized to allow this object to bind.
// Experimental.
func (e *jsiiProxy_EcrPublicSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the service source from ECR.
//
// TODO: EXAMPLE
//
// Experimental.
type EcrSource interface {
	Source
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrSource
type jsiiProxy_EcrSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrSource(props *EcrProps) EcrSource {
	_init_.Initialize()

	j := jsiiProxy_EcrSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrSource_Override(e EcrSource, props *EcrProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the Job is initialized to allow this object to bind.
// Experimental.
func (e *jsiiProxy_EcrSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the App Runner connection that enables the App Runner service to connect to a source repository.
//
// It's required for GitHub code repositories.
//
// TODO: EXAMPLE
//
// Experimental.
type GitHubConnection interface {
	ConnectionArn() *string
}

// The jsii proxy struct for GitHubConnection
type jsiiProxy_GitHubConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_GitHubConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubConnection(arn *string) GitHubConnection {
	_init_.Initialize()

	j := jsiiProxy_GitHubConnection{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		[]interface{}{arn},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubConnection_Override(g GitHubConnection, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		[]interface{}{arn},
		g,
	)
}

// Using existing App Runner connection by specifying the connection ARN.
//
// Returns: Connection
// Experimental.
func GitHubConnection_FromConnectionArn(arn *string) GitHubConnection {
	_init_.Initialize()

	var returns GitHubConnection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		"fromConnectionArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Properties of the Github repository for `Source.fromGitHub()`.
//
// TODO: EXAMPLE
//
// Experimental.
type GithubRepositoryProps struct {
	// The branch name that represents a specific version for the repository.
	// Experimental.
	Branch *string `json:"branch"`
	// The code configuration values.
	//
	// Will be ignored if configurationSource is `REPOSITORY`.
	// Experimental.
	CodeConfigurationValues *CodeConfigurationValues `json:"codeConfigurationValues"`
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `json:"configurationSource"`
	// ARN of the connection to Github.
	//
	// Only required for Github source.
	// Experimental.
	Connection GitHubConnection `json:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl"`
}

// Represents the service source from a Github repository.
//
// TODO: EXAMPLE
//
// Experimental.
type GithubSource interface {
	Source
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for GithubSource
type jsiiProxy_GithubSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewGithubSource(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	j := jsiiProxy_GithubSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubSource_Override(g GithubSource, props *GithubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		[]interface{}{props},
		g,
	)
}

// Source from local assets.
// Experimental.
func GithubSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func GithubSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func GithubSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func GithubSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the Job is initialized to allow this object to bind.
// Experimental.
func (g *jsiiProxy_GithubSource) Bind(_scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Represents the App Runner Service.
// Experimental.
type IService interface {
	awscdk.IResource
	// The ARN of the service.
	// Experimental.
	ServiceArn() *string
	// The Name of the service.
	// Experimental.
	ServiceName() *string
}

// The jsii proxy for IService
type jsiiProxy_IService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html
//
// Experimental.
type ImageConfiguration struct {
	// Environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `json:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `json:"startCommand"`
}

// Describes a source image repository.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
//
// Experimental.
type ImageRepository struct {
	// Configuration for running the identified image.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration"`
	// The identifier of the image.
	//
	// For `ECR_PUBLIC` imageRepositoryType, the identifier domain should
	// always be `public.ecr.aws`. For `ECR`, the pattern should be
	// `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/.*)`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html for more details.
	//
	// Experimental.
	ImageIdentifier *string `json:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether
	// the repository is private or public.
	// Experimental.
	ImageRepositoryType ImageRepositoryType `json:"imageRepositoryType"`
}

// The image repository types.
// Experimental.
type ImageRepositoryType string

const (
	ImageRepositoryType_ECR ImageRepositoryType = "ECR"
	ImageRepositoryType_ECR_PUBLIC ImageRepositoryType = "ECR_PUBLIC"
)

// The amount of memory reserved for each instance of your App Runner service.
//
// TODO: EXAMPLE
//
// Experimental.
type Memory interface {
	Unit() *string
}

// The jsii proxy struct for Memory
type jsiiProxy_Memory struct {
	_ byte // padding
}

func (j *jsiiProxy_Memory) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom Memory unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-memory
//
// Experimental.
func Memory_Of(unit *string) Memory {
	_init_.Initialize()

	var returns Memory

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Memory",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Memory_FOUR_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Memory",
		"FOUR_GB",
		&returns,
	)
	return returns
}

func Memory_THREE_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Memory",
		"THREE_GB",
		&returns,
	)
	return returns
}

func Memory_TWO_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Memory",
		"TWO_GB",
		&returns,
	)
	return returns
}

// The code runtimes.
//
// TODO: EXAMPLE
//
// Experimental.
type Runtime interface {
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Other runtimes.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfigurationvalues.html#cfn-apprunner-service-codeconfigurationvalues-runtime for all available runtimes.
//
// Experimental.
func Runtime_Of(name *string) Runtime {
	_init_.Initialize()

	var returns Runtime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func Runtime_NODEJS_12() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"NODEJS_12",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"PYTHON_3",
		&returns,
	)
	return returns
}

// The App Runner Service.
//
// TODO: EXAMPLE
//
// Experimental.
type Service interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ServiceArn() *string
	ServiceId() *string
	ServiceName() *string
	ServiceStatus() *string
	ServiceUrl() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Service) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	j := jsiiProxy_Service{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import from service attributes.
// Experimental.
func Service_FromServiceAttributes(scope constructs.Construct, id *string, attrs *ServiceAttributes) IService {
	_init_.Initialize()

	var returns IService

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Service",
		"fromServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import from service name.
// Experimental.
func Service_FromServiceName(scope constructs.Construct, id *string, serviceName *string) IService {
	_init_.Initialize()

	var returns IService

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Service",
		"fromServiceName",
		[]interface{}{scope, id, serviceName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Service_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Service",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (s *jsiiProxy_Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_Service) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_Service) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for the App Runner Service.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceAttributes struct {
	// The ARN of the service.
	// Experimental.
	ServiceArn *string `json:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The status of the service.
	// Experimental.
	ServiceStatus *string `json:"serviceStatus"`
	// The URL of the service.
	// Experimental.
	ServiceUrl *string `json:"serviceUrl"`
}

// Properties of the AppRunner Service.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceProps struct {
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	// Experimental.
	AccessRole awsiam.IRole `json:"accessRole"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Experimental.
	Cpu Cpu `json:"cpu"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	// Experimental.
	InstanceRole awsiam.IRole `json:"instanceRole"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Experimental.
	Memory Memory `json:"memory"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The source of the repository for the service.
	// Experimental.
	Source Source `json:"source"`
}

// Represents the App Runner service source.
// Experimental.
type Source interface {
	Bind(scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Experimental.
func NewSource_Override(s Source) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.Source",
		nil, // no parameters
		s,
	)
}

// Source from local assets.
// Experimental.
func Source_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func Source_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func Source_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func Source_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the Job is initialized to allow this object to bind.
// Experimental.
func (s *jsiiProxy_Source) Bind(scope constructs.Construct) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html
//
// Experimental.
type SourceCodeVersion struct {
	// The type of version identifier.
	// Experimental.
	Type *string `json:"type"`
	// A source code version.
	// Experimental.
	Value *string `json:"value"`
}

// Result of binding `Source` into a `Service`.
//
// TODO: EXAMPLE
//
// Experimental.
type SourceConfig struct {
	// The code repository configuration (mutually exclusive  with `imageRepository`).
	// Experimental.
	CodeRepository *CodeRepositoryProps `json:"codeRepository"`
	// The ECR repository (required to grant the pull privileges for the iam role).
	// Experimental.
	EcrRepository awsecr.IRepository `json:"ecrRepository"`
	// The image repository configuration (mutually exclusive  with `codeRepository`).
	// Experimental.
	ImageRepository *ImageRepository `json:"imageRepository"`
}

