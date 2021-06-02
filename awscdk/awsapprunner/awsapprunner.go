package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AppRunner::Service`.
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

type CfnService_AuthenticationConfigurationProperty struct {
	// `CfnService.AuthenticationConfigurationProperty.AccessRoleArn`.
	AccessRoleArn *string `json:"accessRoleArn"`
	// `CfnService.AuthenticationConfigurationProperty.ConnectionArn`.
	ConnectionArn *string `json:"connectionArn"`
}

type CfnService_CodeConfigurationProperty struct {
	// `CfnService.CodeConfigurationProperty.ConfigurationSource`.
	ConfigurationSource *string `json:"configurationSource"`
	// `CfnService.CodeConfigurationProperty.CodeConfigurationValues`.
	CodeConfigurationValues interface{} `json:"codeConfigurationValues"`
}

type CfnService_CodeConfigurationValuesProperty struct {
	// `CfnService.CodeConfigurationValuesProperty.Runtime`.
	Runtime *string `json:"runtime"`
	// `CfnService.CodeConfigurationValuesProperty.BuildCommand`.
	BuildCommand *string `json:"buildCommand"`
	// `CfnService.CodeConfigurationValuesProperty.Port`.
	Port *string `json:"port"`
	// `CfnService.CodeConfigurationValuesProperty.RuntimeEnvironmentVariables`.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables"`
	// `CfnService.CodeConfigurationValuesProperty.StartCommand`.
	StartCommand *string `json:"startCommand"`
}

type CfnService_CodeRepositoryProperty struct {
	// `CfnService.CodeRepositoryProperty.RepositoryUrl`.
	RepositoryUrl *string `json:"repositoryUrl"`
	// `CfnService.CodeRepositoryProperty.SourceCodeVersion`.
	SourceCodeVersion interface{} `json:"sourceCodeVersion"`
	// `CfnService.CodeRepositoryProperty.CodeConfiguration`.
	CodeConfiguration interface{} `json:"codeConfiguration"`
}

type CfnService_EncryptionConfigurationProperty struct {
	// `CfnService.EncryptionConfigurationProperty.KmsKey`.
	KmsKey *string `json:"kmsKey"`
}

type CfnService_HealthCheckConfigurationProperty struct {
	// `CfnService.HealthCheckConfigurationProperty.HealthyThreshold`.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// `CfnService.HealthCheckConfigurationProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnService.HealthCheckConfigurationProperty.Path`.
	Path *string `json:"path"`
	// `CfnService.HealthCheckConfigurationProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnService.HealthCheckConfigurationProperty.Timeout`.
	Timeout *float64 `json:"timeout"`
	// `CfnService.HealthCheckConfigurationProperty.UnhealthyThreshold`.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

type CfnService_ImageConfigurationProperty struct {
	// `CfnService.ImageConfigurationProperty.Port`.
	Port *string `json:"port"`
	// `CfnService.ImageConfigurationProperty.RuntimeEnvironmentVariables`.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables"`
	// `CfnService.ImageConfigurationProperty.StartCommand`.
	StartCommand *string `json:"startCommand"`
}

type CfnService_ImageRepositoryProperty struct {
	// `CfnService.ImageRepositoryProperty.ImageIdentifier`.
	ImageIdentifier *string `json:"imageIdentifier"`
	// `CfnService.ImageRepositoryProperty.ImageRepositoryType`.
	ImageRepositoryType *string `json:"imageRepositoryType"`
	// `CfnService.ImageRepositoryProperty.ImageConfiguration`.
	ImageConfiguration interface{} `json:"imageConfiguration"`
}

type CfnService_InstanceConfigurationProperty struct {
	// `CfnService.InstanceConfigurationProperty.Cpu`.
	Cpu *string `json:"cpu"`
	// `CfnService.InstanceConfigurationProperty.InstanceRoleArn`.
	InstanceRoleArn *string `json:"instanceRoleArn"`
	// `CfnService.InstanceConfigurationProperty.Memory`.
	Memory *string `json:"memory"`
}

type CfnService_KeyValuePairProperty struct {
	// `CfnService.KeyValuePairProperty.Name`.
	Name *string `json:"name"`
	// `CfnService.KeyValuePairProperty.Value`.
	Value *string `json:"value"`
}

type CfnService_SourceCodeVersionProperty struct {
	// `CfnService.SourceCodeVersionProperty.Type`.
	Type *string `json:"type"`
	// `CfnService.SourceCodeVersionProperty.Value`.
	Value *string `json:"value"`
}

type CfnService_SourceConfigurationProperty struct {
	// `CfnService.SourceConfigurationProperty.AuthenticationConfiguration`.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration"`
	// `CfnService.SourceConfigurationProperty.AutoDeploymentsEnabled`.
	AutoDeploymentsEnabled interface{} `json:"autoDeploymentsEnabled"`
	// `CfnService.SourceConfigurationProperty.CodeRepository`.
	CodeRepository interface{} `json:"codeRepository"`
	// `CfnService.SourceConfigurationProperty.ImageRepository`.
	ImageRepository interface{} `json:"imageRepository"`
}

// Properties for defining a `AWS::AppRunner::Service`.
type CfnServiceProps struct {
	// `AWS::AppRunner::Service.SourceConfiguration`.
	SourceConfiguration interface{} `json:"sourceConfiguration"`
	// `AWS::AppRunner::Service.AutoScalingConfigurationArn`.
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn"`
	// `AWS::AppRunner::Service.EncryptionConfiguration`.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration"`
	// `AWS::AppRunner::Service.HealthCheckConfiguration`.
	HealthCheckConfiguration interface{} `json:"healthCheckConfiguration"`
	// `AWS::AppRunner::Service.InstanceConfiguration`.
	InstanceConfiguration interface{} `json:"instanceConfiguration"`
	// `AWS::AppRunner::Service.ServiceName`.
	ServiceName *string `json:"serviceName"`
	// `AWS::AppRunner::Service.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

