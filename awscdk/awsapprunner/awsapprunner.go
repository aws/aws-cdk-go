package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapprunner/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Properties of the image repository for `Source.fromAsset()`.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetProps struct {
	// Represents the docker image asset.
	// Experimental.
	Asset awsecrassets.DockerImageAsset `json:"asset" yaml:"asset"`
	// The image configuration for the image built from the asset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
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
		"monocdk.aws_apprunner.AssetSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetSource_Override(a AssetSource, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.AssetSource",
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
		"monocdk.aws_apprunner.AssetSource",
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
		"monocdk.aws_apprunner.AssetSource",
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
		"monocdk.aws_apprunner.AssetSource",
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
		"monocdk.aws_apprunner.AssetSource",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnService(scope awscdk.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::Service`.
func NewCfnService_Override(c CfnService, scope awscdk.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.CfnService",
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
// Experimental.
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.CfnService",
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
		"monocdk.aws_apprunner.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnService) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnService) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnService) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
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
	ConfigurationSource ConfigurationSourceType `json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a apprunner.yaml file in the
	// source code repository (or ignoring the file if it exists).
	// Experimental.
	ConfigurationValues *CodeConfigurationValues `json:"configurationValues" yaml:"configurationValues"`
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
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents
	// a programming language runtime.
	// Experimental.
	Runtime Runtime `json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	// Experimental.
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *string `json:"port" yaml:"port"`
	// The command App Runner runs to start your application.
	// Experimental.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Properties of the CodeRepository.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeRepositoryProps struct {
	// Configuration for building and running the service from a source code repository.
	// Experimental.
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration" yaml:"codeConfiguration"`
	// The App Runner connection for GitHub.
	// Experimental.
	Connection GitHubConnection `json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// Experimental.
	SourceCodeVersion *SourceCodeVersion `json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
}

// The source of the App Runner configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type ConfigurationSourceType string

const (
	ConfigurationSourceType_REPOSITORY ConfigurationSourceType = "REPOSITORY"
	ConfigurationSourceType_API ConfigurationSourceType = "API"
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
		"monocdk.aws_apprunner.Cpu",
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
		"monocdk.aws_apprunner.Cpu",
		"ONE_VCPU",
		&returns,
	)
	return returns
}

func Cpu_TWO_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Cpu",
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
	// Represents the ECR repository.
	// Experimental.
	Repository awsecr.IRepository `json:"repository" yaml:"repository"`
	// The image configuration for the image from ECR.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image tag.
	// Experimental.
	Tag *string `json:"tag" yaml:"tag"`
}

// Properties of the image repository for `Source.fromEcrPublic()`.
//
// TODO: EXAMPLE
//
// Experimental.
type EcrPublicProps struct {
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
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
		"monocdk.aws_apprunner.EcrPublicSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrPublicSource_Override(e EcrPublicSource, props *EcrPublicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrPublicSource",
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
		"monocdk.aws_apprunner.EcrPublicSource",
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
		"monocdk.aws_apprunner.EcrPublicSource",
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
		"monocdk.aws_apprunner.EcrPublicSource",
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
		"monocdk.aws_apprunner.EcrPublicSource",
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
		"monocdk.aws_apprunner.EcrSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrSource_Override(e EcrSource, props *EcrProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.EcrSource",
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
		"monocdk.aws_apprunner.EcrSource",
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
		"monocdk.aws_apprunner.EcrSource",
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
		"monocdk.aws_apprunner.EcrSource",
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
		"monocdk.aws_apprunner.EcrSource",
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
		"monocdk.aws_apprunner.GitHubConnection",
		[]interface{}{arn},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubConnection_Override(g GitHubConnection, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.GitHubConnection",
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
		"monocdk.aws_apprunner.GitHubConnection",
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
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `json:"configurationSource" yaml:"configurationSource"`
	// ARN of the connection to Github.
	//
	// Only required for Github source.
	// Experimental.
	Connection GitHubConnection `json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
	// The branch name that represents a specific version for the repository.
	// Experimental.
	Branch *string `json:"branch" yaml:"branch"`
	// The code configuration values.
	//
	// Will be ignored if configurationSource is `REPOSITORY`.
	// Experimental.
	CodeConfigurationValues *CodeConfigurationValues `json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
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
		"monocdk.aws_apprunner.GithubSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubSource_Override(g GithubSource, props *GithubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.GithubSource",
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
		"monocdk.aws_apprunner.GithubSource",
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
		"monocdk.aws_apprunner.GithubSource",
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
		"monocdk.aws_apprunner.GithubSource",
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
		"monocdk.aws_apprunner.GithubSource",
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
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// TODO: EXAMPLE
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
	ImageIdentifier *string `json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether
	// the repository is private or public.
	// Experimental.
	ImageRepositoryType ImageRepositoryType `json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `json:"imageConfiguration" yaml:"imageConfiguration"`
}

// The image repository types.
// Experimental.
type ImageRepositoryType string

const (
	ImageRepositoryType_ECR_PUBLIC ImageRepositoryType = "ECR_PUBLIC"
	ImageRepositoryType_ECR ImageRepositoryType = "ECR"
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
		"monocdk.aws_apprunner.Memory",
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
		"monocdk.aws_apprunner.Memory",
		"FOUR_GB",
		&returns,
	)
	return returns
}

func Memory_THREE_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"THREE_GB",
		&returns,
	)
	return returns
}

func Memory_TWO_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
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
		"monocdk.aws_apprunner.Runtime",
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
		"monocdk.aws_apprunner.Runtime",
		"NODEJS_12",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Runtime",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_Service) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_apprunner.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apprunner.Service",
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
		"monocdk.aws_apprunner.Service",
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
		"monocdk.aws_apprunner.Service",
		"fromServiceName",
		[]interface{}{scope, id, serviceName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Service_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Service",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_Service) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_Service) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_Service) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_Service) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_Service) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
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
	ServiceArn *string `json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// The status of the service.
	// Experimental.
	ServiceStatus *string `json:"serviceStatus" yaml:"serviceStatus"`
	// The URL of the service.
	// Experimental.
	ServiceUrl *string `json:"serviceUrl" yaml:"serviceUrl"`
}

// Properties of the AppRunner Service.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceProps struct {
	// The source of the repository for the service.
	// Experimental.
	Source Source `json:"source" yaml:"source"`
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	//
	// The role must be assumable by the 'build.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.access
	//
	// Experimental.
	AccessRole awsiam.IRole `json:"accessRole" yaml:"accessRole"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Experimental.
	Cpu Cpu `json:"cpu" yaml:"cpu"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	//
	// The role must be assumable by the 'tasks.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.instance
	//
	// Experimental.
	InstanceRole awsiam.IRole `json:"instanceRole" yaml:"instanceRole"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Experimental.
	Memory Memory `json:"memory" yaml:"memory"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// Represents the App Runner service source.
//
// TODO: EXAMPLE
//
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
		"monocdk.aws_apprunner.Source",
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
		"monocdk.aws_apprunner.Source",
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
		"monocdk.aws_apprunner.Source",
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
		"monocdk.aws_apprunner.Source",
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
		"monocdk.aws_apprunner.Source",
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
	Type *string `json:"type" yaml:"type"`
	// A source code version.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}

// Result of binding `Source` into a `Service`.
//
// TODO: EXAMPLE
//
// Experimental.
type SourceConfig struct {
	// The code repository configuration (mutually exclusive  with `imageRepository`).
	// Experimental.
	CodeRepository *CodeRepositoryProps `json:"codeRepository" yaml:"codeRepository"`
	// The ECR repository (required to grant the pull privileges for the iam role).
	// Experimental.
	EcrRepository awsecr.IRepository `json:"ecrRepository" yaml:"ecrRepository"`
	// The image repository configuration (mutually exclusive  with `codeRepository`).
	// Experimental.
	ImageRepository *ImageRepository `json:"imageRepository" yaml:"imageRepository"`
}

