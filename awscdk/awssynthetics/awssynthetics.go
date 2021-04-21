package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Synthetics::Canary`.
type CfnCanary interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	AttrId() *string
	AttrState() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Code() interface{}
	SetCode(val interface{})
	CreationStack() *[]*string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RunConfig() interface{}
	SetRunConfig(val interface{})
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	Schedule() interface{}
	SetSchedule(val interface{})
	Stack() awscdk.Stack
	StartCanaryAfterCreation() interface{}
	SetStartCanaryAfterCreation(val interface{})
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnCanary
type jsiiProxy_CfnCanary struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RunConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) StartCanaryAfterCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryAfterCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary(scope constructs.Construct, id *string, props *CfnCanaryProps) CfnCanary {
	_init_.Initialize()

	j := jsiiProxy_CfnCanary{}

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary_Override(c CfnCanary, scope constructs.Construct, id *string, props *CfnCanaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCanary) SetArtifactS3Location(val *string) {
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetFailureRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRunConfig(val interface{}) {
	_jsii_.Set(
		j,
		"runConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetStartCanaryAfterCreation(val interface{}) {
	_jsii_.Set(
		j,
		"startCanaryAfterCreation",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSuccessRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetVpcConfig(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConfig",
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
func CfnCanary_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCanary_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
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
func CfnCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCanary_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCanary) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCanary) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCanary) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCanary) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCanary) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnCanary) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCanary) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCanary) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCanary) GetMetadata(key *string) interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCanary) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCanary) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCanary) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCanary) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCanary) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCanary_CodeProperty struct {
	// `CfnCanary.CodeProperty.Handler`.
	Handler *string `json:"handler"`
	// `CfnCanary.CodeProperty.S3Bucket`.
	S3Bucket *string `json:"s3Bucket"`
	// `CfnCanary.CodeProperty.S3Key`.
	S3Key *string `json:"s3Key"`
	// `CfnCanary.CodeProperty.S3ObjectVersion`.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
	// `CfnCanary.CodeProperty.Script`.
	Script *string `json:"script"`
}

type CfnCanary_RunConfigProperty struct {
	// `CfnCanary.RunConfigProperty.ActiveTracing`.
	ActiveTracing interface{} `json:"activeTracing"`
	// `CfnCanary.RunConfigProperty.EnvironmentVariables`.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// `CfnCanary.RunConfigProperty.MemoryInMB`.
	MemoryInMb *float64 `json:"memoryInMb"`
	// `CfnCanary.RunConfigProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

type CfnCanary_ScheduleProperty struct {
	// `CfnCanary.ScheduleProperty.Expression`.
	Expression *string `json:"expression"`
	// `CfnCanary.ScheduleProperty.DurationInSeconds`.
	DurationInSeconds *string `json:"durationInSeconds"`
}

type CfnCanary_VPCConfigProperty struct {
	// `CfnCanary.VPCConfigProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnCanary.VPCConfigProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `CfnCanary.VPCConfigProperty.VpcId`.
	VpcId *string `json:"vpcId"`
}

// Properties for defining a `AWS::Synthetics::Canary`.
type CfnCanaryProps struct {
	// `AWS::Synthetics::Canary.ArtifactS3Location`.
	ArtifactS3Location *string `json:"artifactS3Location"`
	// `AWS::Synthetics::Canary.Code`.
	Code interface{} `json:"code"`
	// `AWS::Synthetics::Canary.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// `AWS::Synthetics::Canary.Name`.
	Name *string `json:"name"`
	// `AWS::Synthetics::Canary.RuntimeVersion`.
	RuntimeVersion *string `json:"runtimeVersion"`
	// `AWS::Synthetics::Canary.Schedule`.
	Schedule interface{} `json:"schedule"`
	// `AWS::Synthetics::Canary.StartCanaryAfterCreation`.
	StartCanaryAfterCreation interface{} `json:"startCanaryAfterCreation"`
	// `AWS::Synthetics::Canary.FailureRetentionPeriod`.
	FailureRetentionPeriod *float64 `json:"failureRetentionPeriod"`
	// `AWS::Synthetics::Canary.RunConfig`.
	RunConfig interface{} `json:"runConfig"`
	// `AWS::Synthetics::Canary.SuccessRetentionPeriod`.
	SuccessRetentionPeriod *float64 `json:"successRetentionPeriod"`
	// `AWS::Synthetics::Canary.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::Synthetics::Canary.VPCConfig`.
	VpcConfig interface{} `json:"vpcConfig"`
}

