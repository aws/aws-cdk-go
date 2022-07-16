// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties of the image repository for `Source.fromAsset()`.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type AssetProps struct {
	// Represents the docker image asset.
	// Experimental.
	Asset awsecrassets.DockerImageAsset `field:"required" json:"asset" yaml:"asset"`
	// The image configuration for the image built from the asset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Represents the source from local assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImageAsset dockerImageAsset
//
//   assetSource := apprunner_alpha.NewAssetSource(&assetProps{
//   	asset: dockerImageAsset,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type AssetSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var runtime runtime
//
//   codeConfiguration := &codeConfiguration{
//   	configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
//
//   	// the properties below are optional
//   	configurationValues: &codeConfigurationValues{
//   		runtime: runtime,
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.String("port"),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html
//
// Experimental.
type CodeConfiguration struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a apprunner.yaml file in the
	// source code repository (or ignoring the file if it exists).
	// Experimental.
	ConfigurationValues *CodeConfigurationValues `field:"optional" json:"configurationValues" yaml:"configurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities,
// use a `apprunner.yaml` file in the source code repository.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_API,
//   		codeConfigurationValues: &codeConfigurationValues{
//   			runtime: apprunner.runtime_PYTHON_3(),
//   			port: jsii.String("8000"),
//   			startCommand: jsii.String("python app.py"),
//   			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type CodeConfigurationValues struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents
	// a programming language runtime.
	// Experimental.
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The command App Runner runs to start your application.
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

// Properties of the CodeRepository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//
//   codeRepositoryProps := &codeRepositoryProps{
//   	codeConfiguration: &codeConfiguration{
//   		configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
//
//   		// the properties below are optional
//   		configurationValues: &codeConfigurationValues{
//   			runtime: runtime,
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.String("port"),
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   	connection: gitHubConnection,
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersion{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
//   	},
//   }
//
// Experimental.
type CodeRepositoryProps struct {
	// Configuration for building and running the service from a source code repository.
	// Experimental.
	CodeConfiguration *CodeConfiguration `field:"required" json:"codeConfiguration" yaml:"codeConfiguration"`
	// The App Runner connection for GitHub.
	// Experimental.
	Connection GitHubConnection `field:"required" json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// Experimental.
	SourceCodeVersion *SourceCodeVersion `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
}

// The source of the App Runner configuration.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type ConfigurationSourceType string

const (
	// App Runner reads configuration values from `the apprunner.yaml` file in the source code repository and ignores `configurationValues`.
	// Experimental.
	ConfigurationSourceType_REPOSITORY ConfigurationSourceType = "REPOSITORY"
	// App Runner uses configuration values provided in `configurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	// Experimental.
	ConfigurationSourceType_API ConfigurationSourceType = "API"
)

// The number of CPU units reserved for each instance of your App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   cpu := apprunner_alpha.cpu.of(jsii.String("unit"))
//
// Experimental.
type Cpu interface {
	// The unit of CPU.
	// Experimental.
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
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcr(&ecrProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(80),
//   		},
//   		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		tagOrDigest: jsii.String("latest"),
//   	}),
//   })
//
// Experimental.
type EcrProps struct {
	// Represents the ECR repository.
	// Experimental.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The image configuration for the image from ECR.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image tag.
	// Deprecated: use `tagOrDigest`.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Image tag or digest (digests must start with `sha256:`).
	// Experimental.
	TagOrDigest *string `field:"optional" json:"tagOrDigest" yaml:"tagOrDigest"`
}

// Properties of the image repository for `Source.fromEcrPublic()`.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
// Experimental.
type EcrPublicProps struct {
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Represents the service source from ECR Public.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   ecrPublicSource := apprunner_alpha.NewEcrPublicSource(&ecrPublicProps{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type EcrPublicSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repository repository
//
//   ecrSource := apprunner_alpha.NewEcrSource(&ecrProps{
//   	repository: repository,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   	tag: jsii.String("tag"),
//   	tagOrDigest: jsii.String("tagOrDigest"),
//   })
//
// Experimental.
type EcrSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
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
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type GitHubConnection interface {
	// The ARN of the Connection for App Runner service to connect to the repository.
	// Experimental.
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
// Returns: Connection.
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
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_REPOSITORY,
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type GithubRepositoryProps struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// ARN of the connection to Github.
	//
	// Only required for Github source.
	// Experimental.
	Connection GitHubConnection `field:"required" json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The branch name that represents a specific version for the repository.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The code configuration values.
	//
	// Will be ignored if configurationSource is `REPOSITORY`.
	// Experimental.
	CodeConfigurationValues *CodeConfigurationValues `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

// Represents the service source from a Github repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//
//   githubSource := apprunner_alpha.NewGithubSource(&githubRepositoryProps{
//   	configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
//   	connection: gitHubConnection,
//   	repositoryUrl: jsii.String("repositoryUrl"),
//
//   	// the properties below are optional
//   	branch: jsii.String("branch"),
//   	codeConfigurationValues: &codeConfigurationValues{
//   		runtime: runtime,
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.String("port"),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type GithubSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
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

// Represents the App Runner VPC Connector.
// Experimental.
type IVpcConnector interface {
	awsec2.IConnectable
	awscdk.IResource
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn() *string
	// The Name of the VPC connector.
	// Experimental.
	VpcConnectorName() *string
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision() *float64
}

// The jsii proxy for IVpcConnector
type jsiiProxy_IVpcConnector struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IVpcConnector) VpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) VpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpcConnectorRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html
//
// Experimental.
type ImageConfiguration struct {
	// Environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   imageRepository := &imageRepository{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//   	imageRepositoryType: apprunner_alpha.imageRepositoryType_ECR_PUBLIC,
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfiguration{
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.Number(123),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
//
// Experimental.
type ImageRepository struct {
	// The identifier of the image.
	//
	// For `ECR_PUBLIC` imageRepositoryType, the identifier domain should
	// always be `public.ecr.aws`. For `ECR`, the pattern should be
	// `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/.*)`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html for more details.
	//
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether
	// the repository is private or public.
	// Experimental.
	ImageRepositoryType ImageRepositoryType `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

// The image repository types.
// Experimental.
type ImageRepositoryType string

const (
	// Amazon ECR Public.
	// Experimental.
	ImageRepositoryType_ECR_PUBLIC ImageRepositoryType = "ECR_PUBLIC"
	// Amazon ECR.
	// Experimental.
	ImageRepositoryType_ECR ImageRepositoryType = "ECR"
)

// The amount of memory reserved for each instance of your App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   memory := apprunner_alpha.memory_FOUR_GB()
//
// Experimental.
type Memory interface {
	// The unit of memory.
	// Experimental.
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
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_API,
//   		codeConfigurationValues: &codeConfigurationValues{
//   			runtime: apprunner.runtime_PYTHON_3(),
//   			port: jsii.String("8000"),
//   			startCommand: jsii.String("python app.py"),
//   			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type Runtime interface {
	// The runtime name.
	// Experimental.
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
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type Service interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the Service.
	// Experimental.
	ServiceArn() *string
	// The ID of the Service.
	// Experimental.
	ServiceId() *string
	// The name of the service.
	// Experimental.
	ServiceName() *string
	// The status of the Service.
	// Experimental.
	ServiceStatus() *string
	// The URL of the Service.
	// Experimental.
	ServiceUrl() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
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

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Service_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Service",
		"isOwnedResource",
		[]interface{}{construct},
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

func (s *jsiiProxy_Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   serviceAttributes := &serviceAttributes{
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceStatus: jsii.String("serviceStatus"),
//   	serviceUrl: jsii.String("serviceUrl"),
//   }
//
// Experimental.
type ServiceAttributes struct {
	// The ARN of the service.
	// Experimental.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The status of the service.
	// Experimental.
	ServiceStatus *string `field:"required" json:"serviceStatus" yaml:"serviceStatus"`
	// The URL of the service.
	// Experimental.
	ServiceUrl *string `field:"required" json:"serviceUrl" yaml:"serviceUrl"`
}

// Properties of the AppRunner Service.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type ServiceProps struct {
	// The source of the repository for the service.
	// Experimental.
	Source Source `field:"required" json:"source" yaml:"source"`
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	//
	// The role must be assumable by the 'build.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.access
	//
	// Experimental.
	AccessRole awsiam.IRole `field:"optional" json:"accessRole" yaml:"accessRole"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Experimental.
	Cpu Cpu `field:"optional" json:"cpu" yaml:"cpu"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	//
	// The role must be assumable by the 'tasks.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.instance
	//
	// Experimental.
	InstanceRole awsiam.IRole `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Experimental.
	Memory Memory `field:"optional" json:"memory" yaml:"memory"`
	// Name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Settings for an App Runner VPC connector to associate with the service.
	// Experimental.
	VpcConnector IVpcConnector `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
}

// Represents the App Runner service source.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type Source interface {
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   sourceCodeVersion := &sourceCodeVersion{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html
//
// Experimental.
type SourceCodeVersion struct {
	// The type of version identifier.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A source code version.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Result of binding `Source` into a `Service`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gitHubConnection gitHubConnection
//   var repository repository
//   var runtime runtime
//
//   sourceConfig := &sourceConfig{
//   	codeRepository: &codeRepositoryProps{
//   		codeConfiguration: &codeConfiguration{
//   			configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
//
//   			// the properties below are optional
//   			configurationValues: &codeConfigurationValues{
//   				runtime: runtime,
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				port: jsii.String("port"),
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   		connection: gitHubConnection,
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersion{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrRepository: repository,
//   	imageRepository: &imageRepository{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: apprunner_alpha.imageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfiguration{
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.Number(123),
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The code repository configuration (mutually exclusive  with `imageRepository`).
	// Experimental.
	CodeRepository *CodeRepositoryProps `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// The ECR repository (required to grant the pull privileges for the iam role).
	// Experimental.
	EcrRepository awsecr.IRepository `field:"optional" json:"ecrRepository" yaml:"ecrRepository"`
	// The image repository configuration (mutually exclusive  with `codeRepository`).
	// Experimental.
	ImageRepository *ImageRepository `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

// The App Runner VPC Connector.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	cidr: jsii.String("10.0.0.0/16"),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &vpcConnectorProps{
//   	vpc: vpc,
//   	vpcSubnets: vpc.selectSubnets(&subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	}),
//   	vpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	vpcConnector: vpcConnector,
//   })
//
// Experimental.
type VpcConnector interface {
	awscdk.Resource
	IVpcConnector
	// Allows specifying security group connections for the VPC connector.
	// Experimental.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn() *string
	// The name of the VPC connector.
	// Experimental.
	VpcConnectorName() *string
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision() *float64
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcConnector
type jsiiProxy_VpcConnector struct {
	internal.Type__awscdkResource
	jsiiProxy_IVpcConnector
}

func (j *jsiiProxy_VpcConnector) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpcConnectorRevision",
		&returns,
	)
	return returns
}


// Experimental.
func NewVpcConnector(scope constructs.Construct, id *string, props *VpcConnectorProps) VpcConnector {
	_init_.Initialize()

	j := jsiiProxy_VpcConnector{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVpcConnector_Override(v VpcConnector, scope constructs.Construct, id *string, props *VpcConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import from VPC connector attributes.
// Experimental.
func VpcConnector_FromVpcConnectorAttributes(scope constructs.Construct, id *string, attrs *VpcConnectorAttributes) IVpcConnector {
	_init_.Initialize()

	var returns IVpcConnector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"fromVpcConnectorAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func VpcConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func VpcConnector_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VpcConnector_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_VpcConnector) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for the App Runner VPC Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   vpcConnectorAttributes := &vpcConnectorAttributes{
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   	vpcConnectorRevision: jsii.Number(123),
//   }
//
// Experimental.
type VpcConnectorAttributes struct {
	// The security groups associated with the VPC connector.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn *string `field:"required" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
	// The name of the VPC connector.
	// Experimental.
	VpcConnectorName *string `field:"required" json:"vpcConnectorName" yaml:"vpcConnectorName"`
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision *float64 `field:"required" json:"vpcConnectorRevision" yaml:"vpcConnectorRevision"`
}

// Properties of the AppRunner VPC Connector.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	cidr: jsii.String("10.0.0.0/16"),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &vpcConnectorProps{
//   	vpc: vpc,
//   	vpcSubnets: vpc.selectSubnets(&subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	}),
//   	vpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	vpcConnector: vpcConnector,
//   })
//
// Experimental.
type VpcConnectorProps struct {
	// The VPC for the VPC Connector.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The name for the VpcConnector.
	// Experimental.
	VpcConnectorName *string `field:"optional" json:"vpcConnectorName" yaml:"vpcConnectorName"`
	// Where to place the VPC Connector within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

