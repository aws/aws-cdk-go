package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppRunner::Service`.
//
// Specify an AWS App Runner service by using the `AWS::AppRunner::Service` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::Service` resource is an AWS App Runner resource type that specifies an App Runner service.
//
// TODO: EXAMPLE
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrServiceArn() *string
	AttrServiceId() *string
	AttrServiceUrl() *string
	AttrStatus() *string
	AutoScalingConfigurationArn() *string
	SetAutoScalingConfigurationArn(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	HealthCheckConfiguration() interface{}
	SetHealthCheckConfiguration(val interface{})
	InstanceConfiguration() interface{}
	SetInstanceConfiguration(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ServiceName() *string
	SetServiceName(val *string)
	SourceConfiguration() interface{}
	SetSourceConfiguration(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AutoScalingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) InstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) SourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::Service`.
func NewCfnService(scope constructs.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::Service`.
func NewCfnService_Override(c CfnService, scope constructs.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService) SetAutoScalingConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"autoScalingConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetInstanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"instanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sourceConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnService) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes resources needed to authenticate access to some source repositories.
//
// The specific resource depends on the repository provider.
//
// TODO: EXAMPLE
//
type CfnService_AuthenticationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	AccessRoleArn *string `json:"accessRoleArn" yaml:"accessRoleArn"`
	// The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository.
	//
	// It's required for GitHub code repositories.
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
}

// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// TODO: EXAMPLE
//
type CfnService_CodeConfigurationProperty struct {
	// The source of the App Runner configuration. Values are interpreted as follows:.
	//
	// - `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and ignores `CodeConfigurationValues` .
	// - `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	ConfigurationSource *string `json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a `apprunner.yaml` file in the source code repository (or ignoring the file if it exists).
	CodeConfigurationValues interface{} `json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.
//
// TODO: EXAMPLE
//
type CfnService_CodeConfigurationValuesProperty struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents a programming language runtime.
	Runtime *string `json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The port that your application listens to in the container.
	//
	// Default: `8080`
	Port *string `json:"port" yaml:"port"`
	// The environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// The command App Runner runs to start your application.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source code repository.
//
// TODO: EXAMPLE
//
type CfnService_CodeRepositoryProperty struct {
	// The location of the repository that contains the source code.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	SourceCodeVersion interface{} `json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Configuration for building and running the service from a source code repository.
	CodeConfiguration interface{} `json:"codeConfiguration" yaml:"codeConfiguration"`
}

// Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.
//
// TODO: EXAMPLE
//
type CfnService_EncryptionConfigurationProperty struct {
	// The ARN of the KMS key that's used for encryption.
	KmsKey *string `json:"kmsKey" yaml:"kmsKey"`
}

// Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.
//
// TODO: EXAMPLE
//
type CfnService_HealthCheckConfigurationProperty struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	//
	// Default: `1`
	HealthyThreshold *float64 `json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	//
	// Default: `5`
	Interval *float64 `json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	//
	// `Path` is only applicable when you set `Protocol` to `HTTP` .
	//
	// Default: `"/"`
	Path *string `json:"path" yaml:"path"`
	// The IP protocol that App Runner uses to perform health checks for your service.
	//
	// If you set `Protocol` to `HTTP` , App Runner sends health check requests to the HTTP path specified by `Path` .
	//
	// Default: `TCP`
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	//
	// Default: `2`
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	//
	// Default: `5`
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// TODO: EXAMPLE
//
type CfnService_ImageConfigurationProperty struct {
	// The port that your application listens to in the container.
	//
	// Default: `8080`
	Port *string `json:"port" yaml:"port"`
	// Environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// TODO: EXAMPLE
//
type CfnService_ImageRepositoryProperty struct {
	// The identifier of an image.
	//
	// For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html) in the *Amazon ECR User Guide* .
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether the repository is private or public.
	ImageRepositoryType *string `json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	ImageConfiguration interface{} `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Describes the runtime configuration of an AWS App Runner service instance (scaling unit).
//
// TODO: EXAMPLE
//
type CfnService_InstanceConfigurationProperty struct {
	// The number of CPU units reserved for each instance of your App Runner service.
	//
	// Default: `1 vCPU`
	Cpu *string `json:"cpu" yaml:"cpu"`
	// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	InstanceRoleArn *string `json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The amount of memory, in MB or GB, reserved for each instance of your App Runner service.
	//
	// Default: `2 GB`
	Memory *string `json:"memory" yaml:"memory"`
}

// Describes a key-value pair, which is a string-to-string mapping.
//
// TODO: EXAMPLE
//
type CfnService_KeyValuePairProperty struct {
	// The key name string to map to a value.
	Name *string `json:"name" yaml:"name"`
	// The value string to which the key name is mapped.
	Value *string `json:"value" yaml:"value"`
}

// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// TODO: EXAMPLE
//
type CfnService_SourceCodeVersionProperty struct {
	// The type of version identifier.
	//
	// For a git-based repository, branches represent versions.
	Type *string `json:"type" yaml:"type"`
	// A source code version.
	//
	// For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	Value *string `json:"value" yaml:"value"`
}

// Describes the source deployed to an AWS App Runner service.
//
// It can be a code or an image repository.
//
// TODO: EXAMPLE
//
type CfnService_SourceConfigurationProperty struct {
	// Describes the resources that are needed to authenticate access to some source repositories.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// If `true` , continuous integration from the source repository is enabled for the App Runner service.
	//
	// Each repository change (including any source code commit or new image version) starts a deployment.
	//
	// Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	AutoDeploymentsEnabled interface{} `json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The description of a source code repository.
	//
	// You must provide either this member or `ImageRepository` (but not both).
	CodeRepository interface{} `json:"codeRepository" yaml:"codeRepository"`
	// The description of a source image repository.
	//
	// You must provide either this member or `CodeRepository` (but not both).
	ImageRepository interface{} `json:"imageRepository" yaml:"imageRepository"`
}

// Properties for defining a `CfnService`.
//
// TODO: EXAMPLE
//
type CfnServiceProps struct {
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration interface{} `json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The settings for the health check that AWS App Runner performs to monitor the health of your service.
	HealthCheckConfiguration interface{} `json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of the App Runner service.
	InstanceConfiguration interface{} `json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// A name for the new service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// An optional list of metadata items that you can associate with your service resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

