package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Batch::ComputeEnvironment`.
type CfnComputeEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	ComputeResources() interface{}
	SetComputeResources(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	Stack() awscdk.Stack
	State() *string
	SetState(val *string)
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnComputeEnvironment
type jsiiProxy_CfnComputeEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ComputeResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Batch::ComputeEnvironment`.
func NewCfnComputeEnvironment(scope constructs.Construct, id *string, props *CfnComputeEnvironmentProps) CfnComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_CfnComputeEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::ComputeEnvironment`.
func NewCfnComputeEnvironment_Override(c CfnComputeEnvironment, scope constructs.Construct, id *string, props *CfnComputeEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment) SetComputeEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"computeEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment) SetComputeResources(val interface{}) {
	_jsii_.Set(
		j,
		"computeResources",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnComputeEnvironment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnComputeEnvironment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
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
func CfnComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputeEnvironment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnComputeEnvironment) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnComputeEnvironment) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnComputeEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnComputeEnvironment) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnComputeEnvironment) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnComputeEnvironment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnComputeEnvironment) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnComputeEnvironment) ToString() *string {
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
func (c *jsiiProxy_CfnComputeEnvironment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnComputeEnvironment_ComputeResourcesProperty struct {
	// `CfnComputeEnvironment.ComputeResourcesProperty.MaxvCpus`.
	MaxvCpus *float64 `json:"maxvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Subnets`.
	Subnets *[]*string `json:"subnets"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Type`.
	Type *string `json:"type"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.BidPercentage`.
	BidPercentage *float64 `json:"bidPercentage"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.DesiredvCpus`.
	DesiredvCpus *float64 `json:"desiredvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Ec2Configuration`.
	Ec2Configuration interface{} `json:"ec2Configuration"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Ec2KeyPair`.
	Ec2KeyPair *string `json:"ec2KeyPair"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.ImageId`.
	ImageId *string `json:"imageId"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.InstanceRole`.
	InstanceRole *string `json:"instanceRole"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.InstanceTypes`.
	InstanceTypes *[]*string `json:"instanceTypes"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.LaunchTemplate`.
	LaunchTemplate interface{} `json:"launchTemplate"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.MinvCpus`.
	MinvCpus *float64 `json:"minvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.PlacementGroup`.
	PlacementGroup *string `json:"placementGroup"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.SpotIamFleetRole`.
	SpotIamFleetRole *string `json:"spotIamFleetRole"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Tags`.
	Tags interface{} `json:"tags"`
}

type CfnComputeEnvironment_Ec2ConfigurationObjectProperty struct {
	// `CfnComputeEnvironment.Ec2ConfigurationObjectProperty.ImageType`.
	ImageType *string `json:"imageType"`
	// `CfnComputeEnvironment.Ec2ConfigurationObjectProperty.ImageIdOverride`.
	ImageIdOverride *string `json:"imageIdOverride"`
}

type CfnComputeEnvironment_LaunchTemplateSpecificationProperty struct {
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.LaunchTemplateId`.
	LaunchTemplateId *string `json:"launchTemplateId"`
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.LaunchTemplateName`.
	LaunchTemplateName *string `json:"launchTemplateName"`
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.Version`.
	Version *string `json:"version"`
}

// Properties for defining a `AWS::Batch::ComputeEnvironment`.
type CfnComputeEnvironmentProps struct {
	// `AWS::Batch::ComputeEnvironment.Type`.
	Type *string `json:"type"`
	// `AWS::Batch::ComputeEnvironment.ComputeEnvironmentName`.
	ComputeEnvironmentName *string `json:"computeEnvironmentName"`
	// `AWS::Batch::ComputeEnvironment.ComputeResources`.
	ComputeResources interface{} `json:"computeResources"`
	// `AWS::Batch::ComputeEnvironment.ServiceRole`.
	ServiceRole *string `json:"serviceRole"`
	// `AWS::Batch::ComputeEnvironment.State`.
	State *string `json:"state"`
	// `AWS::Batch::ComputeEnvironment.Tags`.
	Tags interface{} `json:"tags"`
}

// A CloudFormation `AWS::Batch::JobDefinition`.
type CfnJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContainerProperties() interface{}
	SetContainerProperties(val interface{})
	CreationStack() *[]*string
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	LogicalId() *string
	Node() constructs.Node
	NodeProperties() interface{}
	SetNodeProperties(val interface{})
	Parameters() interface{}
	SetParameters(val interface{})
	PlatformCapabilities() *[]*string
	SetPlatformCapabilities(val *[]*string)
	PropagateTags() interface{}
	SetPropagateTags(val interface{})
	Ref() *string
	RetryStrategy() interface{}
	SetRetryStrategy(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Timeout() interface{}
	SetTimeout(val interface{})
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnJobDefinition
type jsiiProxy_CfnJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) ContainerProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) NodeProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) PlatformCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) PropagateTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) RetryStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Timeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Batch::JobDefinition`.
func NewCfnJobDefinition(scope constructs.Construct, id *string, props *CfnJobDefinitionProps) CfnJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnJobDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::JobDefinition`.
func NewCfnJobDefinition_Override(c CfnJobDefinition, scope constructs.Construct, id *string, props *CfnJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetContainerProperties(val interface{}) {
	_jsii_.Set(
		j,
		"containerProperties",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetNodeProperties(val interface{}) {
	_jsii_.Set(
		j,
		"nodeProperties",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetPlatformCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"platformCapabilities",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetPropagateTags(val interface{}) {
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetRetryStrategy(val interface{}) {
	_jsii_.Set(
		j,
		"retryStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetTimeout(val interface{}) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJobDefinition) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
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
func CfnJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnJobDefinition) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnJobDefinition) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnJobDefinition) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnJobDefinition) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnJobDefinition) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnJobDefinition) ToString() *string {
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
func (c *jsiiProxy_CfnJobDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnJobDefinition_AuthorizationConfigProperty struct {
	// `CfnJobDefinition.AuthorizationConfigProperty.AccessPointId`.
	AccessPointId *string `json:"accessPointId"`
	// `CfnJobDefinition.AuthorizationConfigProperty.Iam`.
	Iam *string `json:"iam"`
}

type CfnJobDefinition_ContainerPropertiesProperty struct {
	// `CfnJobDefinition.ContainerPropertiesProperty.Image`.
	Image *string `json:"image"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Command`.
	Command *[]*string `json:"command"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Environment`.
	Environment interface{} `json:"environment"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// `CfnJobDefinition.ContainerPropertiesProperty.FargatePlatformConfiguration`.
	FargatePlatformConfiguration interface{} `json:"fargatePlatformConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnJobDefinition.ContainerPropertiesProperty.JobRoleArn`.
	JobRoleArn *string `json:"jobRoleArn"`
	// `CfnJobDefinition.ContainerPropertiesProperty.LinuxParameters`.
	LinuxParameters interface{} `json:"linuxParameters"`
	// `CfnJobDefinition.ContainerPropertiesProperty.LogConfiguration`.
	LogConfiguration interface{} `json:"logConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Memory`.
	Memory *float64 `json:"memory"`
	// `CfnJobDefinition.ContainerPropertiesProperty.MountPoints`.
	MountPoints interface{} `json:"mountPoints"`
	// `CfnJobDefinition.ContainerPropertiesProperty.NetworkConfiguration`.
	NetworkConfiguration interface{} `json:"networkConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Privileged`.
	Privileged interface{} `json:"privileged"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ReadonlyRootFilesystem`.
	ReadonlyRootFilesystem interface{} `json:"readonlyRootFilesystem"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ResourceRequirements`.
	ResourceRequirements interface{} `json:"resourceRequirements"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Secrets`.
	Secrets interface{} `json:"secrets"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Ulimits`.
	Ulimits interface{} `json:"ulimits"`
	// `CfnJobDefinition.ContainerPropertiesProperty.User`.
	User *string `json:"user"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Vcpus`.
	Vcpus *float64 `json:"vcpus"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Volumes`.
	Volumes interface{} `json:"volumes"`
}

type CfnJobDefinition_DeviceProperty struct {
	// `CfnJobDefinition.DeviceProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath"`
	// `CfnJobDefinition.DeviceProperty.HostPath`.
	HostPath *string `json:"hostPath"`
	// `CfnJobDefinition.DeviceProperty.Permissions`.
	Permissions *[]*string `json:"permissions"`
}

type CfnJobDefinition_EfsVolumeConfigurationProperty struct {
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.FileSystemId`.
	FileSystemId *string `json:"fileSystemId"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.AuthorizationConfig`.
	AuthorizationConfig interface{} `json:"authorizationConfig"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.RootDirectory`.
	RootDirectory *string `json:"rootDirectory"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.TransitEncryption`.
	TransitEncryption *string `json:"transitEncryption"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.TransitEncryptionPort`.
	TransitEncryptionPort *float64 `json:"transitEncryptionPort"`
}

type CfnJobDefinition_EnvironmentProperty struct {
	// `CfnJobDefinition.EnvironmentProperty.Name`.
	Name *string `json:"name"`
	// `CfnJobDefinition.EnvironmentProperty.Value`.
	Value *string `json:"value"`
}

type CfnJobDefinition_EvaluateOnExitProperty struct {
	// `CfnJobDefinition.EvaluateOnExitProperty.Action`.
	Action *string `json:"action"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnExitCode`.
	OnExitCode *string `json:"onExitCode"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnReason`.
	OnReason *string `json:"onReason"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnStatusReason`.
	OnStatusReason *string `json:"onStatusReason"`
}

type CfnJobDefinition_FargatePlatformConfigurationProperty struct {
	// `CfnJobDefinition.FargatePlatformConfigurationProperty.PlatformVersion`.
	PlatformVersion *string `json:"platformVersion"`
}

type CfnJobDefinition_LinuxParametersProperty struct {
	// `CfnJobDefinition.LinuxParametersProperty.Devices`.
	Devices interface{} `json:"devices"`
	// `CfnJobDefinition.LinuxParametersProperty.InitProcessEnabled`.
	InitProcessEnabled interface{} `json:"initProcessEnabled"`
	// `CfnJobDefinition.LinuxParametersProperty.MaxSwap`.
	MaxSwap *float64 `json:"maxSwap"`
	// `CfnJobDefinition.LinuxParametersProperty.SharedMemorySize`.
	SharedMemorySize *float64 `json:"sharedMemorySize"`
	// `CfnJobDefinition.LinuxParametersProperty.Swappiness`.
	Swappiness *float64 `json:"swappiness"`
	// `CfnJobDefinition.LinuxParametersProperty.Tmpfs`.
	Tmpfs interface{} `json:"tmpfs"`
}

type CfnJobDefinition_LogConfigurationProperty struct {
	// `CfnJobDefinition.LogConfigurationProperty.LogDriver`.
	LogDriver *string `json:"logDriver"`
	// `CfnJobDefinition.LogConfigurationProperty.Options`.
	Options interface{} `json:"options"`
	// `CfnJobDefinition.LogConfigurationProperty.SecretOptions`.
	SecretOptions interface{} `json:"secretOptions"`
}

type CfnJobDefinition_MountPointsProperty struct {
	// `CfnJobDefinition.MountPointsProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath"`
	// `CfnJobDefinition.MountPointsProperty.ReadOnly`.
	ReadOnly interface{} `json:"readOnly"`
	// `CfnJobDefinition.MountPointsProperty.SourceVolume`.
	SourceVolume *string `json:"sourceVolume"`
}

type CfnJobDefinition_NetworkConfigurationProperty struct {
	// `CfnJobDefinition.NetworkConfigurationProperty.AssignPublicIp`.
	AssignPublicIp *string `json:"assignPublicIp"`
}

type CfnJobDefinition_NodePropertiesProperty struct {
	// `CfnJobDefinition.NodePropertiesProperty.MainNode`.
	MainNode *float64 `json:"mainNode"`
	// `CfnJobDefinition.NodePropertiesProperty.NodeRangeProperties`.
	NodeRangeProperties interface{} `json:"nodeRangeProperties"`
	// `CfnJobDefinition.NodePropertiesProperty.NumNodes`.
	NumNodes *float64 `json:"numNodes"`
}

type CfnJobDefinition_NodeRangePropertyProperty struct {
	// `CfnJobDefinition.NodeRangePropertyProperty.TargetNodes`.
	TargetNodes *string `json:"targetNodes"`
	// `CfnJobDefinition.NodeRangePropertyProperty.Container`.
	Container interface{} `json:"container"`
}

type CfnJobDefinition_ResourceRequirementProperty struct {
	// `CfnJobDefinition.ResourceRequirementProperty.Type`.
	Type *string `json:"type"`
	// `CfnJobDefinition.ResourceRequirementProperty.Value`.
	Value *string `json:"value"`
}

type CfnJobDefinition_RetryStrategyProperty struct {
	// `CfnJobDefinition.RetryStrategyProperty.Attempts`.
	Attempts *float64 `json:"attempts"`
	// `CfnJobDefinition.RetryStrategyProperty.EvaluateOnExit`.
	EvaluateOnExit interface{} `json:"evaluateOnExit"`
}

type CfnJobDefinition_SecretProperty struct {
	// `CfnJobDefinition.SecretProperty.Name`.
	Name *string `json:"name"`
	// `CfnJobDefinition.SecretProperty.ValueFrom`.
	ValueFrom *string `json:"valueFrom"`
}

type CfnJobDefinition_TimeoutProperty struct {
	// `CfnJobDefinition.TimeoutProperty.AttemptDurationSeconds`.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds"`
}

type CfnJobDefinition_TmpfsProperty struct {
	// `CfnJobDefinition.TmpfsProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath"`
	// `CfnJobDefinition.TmpfsProperty.Size`.
	Size *float64 `json:"size"`
	// `CfnJobDefinition.TmpfsProperty.MountOptions`.
	MountOptions *[]*string `json:"mountOptions"`
}

type CfnJobDefinition_UlimitProperty struct {
	// `CfnJobDefinition.UlimitProperty.HardLimit`.
	HardLimit *float64 `json:"hardLimit"`
	// `CfnJobDefinition.UlimitProperty.Name`.
	Name *string `json:"name"`
	// `CfnJobDefinition.UlimitProperty.SoftLimit`.
	SoftLimit *float64 `json:"softLimit"`
}

type CfnJobDefinition_VolumesHostProperty struct {
	// `CfnJobDefinition.VolumesHostProperty.SourcePath`.
	SourcePath *string `json:"sourcePath"`
}

type CfnJobDefinition_VolumesProperty struct {
	// `CfnJobDefinition.VolumesProperty.EfsVolumeConfiguration`.
	EfsVolumeConfiguration interface{} `json:"efsVolumeConfiguration"`
	// `CfnJobDefinition.VolumesProperty.Host`.
	Host interface{} `json:"host"`
	// `CfnJobDefinition.VolumesProperty.Name`.
	Name *string `json:"name"`
}

// Properties for defining a `AWS::Batch::JobDefinition`.
type CfnJobDefinitionProps struct {
	// `AWS::Batch::JobDefinition.Type`.
	Type *string `json:"type"`
	// `AWS::Batch::JobDefinition.ContainerProperties`.
	ContainerProperties interface{} `json:"containerProperties"`
	// `AWS::Batch::JobDefinition.JobDefinitionName`.
	JobDefinitionName *string `json:"jobDefinitionName"`
	// `AWS::Batch::JobDefinition.NodeProperties`.
	NodeProperties interface{} `json:"nodeProperties"`
	// `AWS::Batch::JobDefinition.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `AWS::Batch::JobDefinition.PlatformCapabilities`.
	PlatformCapabilities *[]*string `json:"platformCapabilities"`
	// `AWS::Batch::JobDefinition.PropagateTags`.
	PropagateTags interface{} `json:"propagateTags"`
	// `AWS::Batch::JobDefinition.RetryStrategy`.
	RetryStrategy interface{} `json:"retryStrategy"`
	// `AWS::Batch::JobDefinition.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::Batch::JobDefinition.Timeout`.
	Timeout interface{} `json:"timeout"`
}

// A CloudFormation `AWS::Batch::JobQueue`.
type CfnJobQueue interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ComputeEnvironmentOrder() interface{}
	SetComputeEnvironmentOrder(val interface{})
	CreationStack() *[]*string
	JobQueueName() *string
	SetJobQueueName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	Ref() *string
	Stack() awscdk.Stack
	State() *string
	SetState(val *string)
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

// The jsii proxy struct for CfnJobQueue
type jsiiProxy_CfnJobQueue struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobQueue) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) ComputeEnvironmentOrder() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeEnvironmentOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) JobQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueue) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Batch::JobQueue`.
func NewCfnJobQueue(scope constructs.Construct, id *string, props *CfnJobQueueProps) CfnJobQueue {
	_init_.Initialize()

	j := jsiiProxy_CfnJobQueue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::JobQueue`.
func NewCfnJobQueue_Override(c CfnJobQueue, scope constructs.Construct, id *string, props *CfnJobQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJobQueue) SetComputeEnvironmentOrder(val interface{}) {
	_jsii_.Set(
		j,
		"computeEnvironmentOrder",
		val,
	)
}

func (j *jsiiProxy_CfnJobQueue) SetJobQueueName(val *string) {
	_jsii_.Set(
		j,
		"jobQueueName",
		val,
	)
}

func (j *jsiiProxy_CfnJobQueue) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnJobQueue) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
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
func CfnJobQueue_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnJobQueue_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
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
func CfnJobQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobQueue_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnJobQueue) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnJobQueue) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnJobQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnJobQueue) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnJobQueue) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnJobQueue) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnJobQueue) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJobQueue) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnJobQueue) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnJobQueue) ToString() *string {
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
func (c *jsiiProxy_CfnJobQueue) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnJobQueue_ComputeEnvironmentOrderProperty struct {
	// `CfnJobQueue.ComputeEnvironmentOrderProperty.ComputeEnvironment`.
	ComputeEnvironment *string `json:"computeEnvironment"`
	// `CfnJobQueue.ComputeEnvironmentOrderProperty.Order`.
	Order *float64 `json:"order"`
}

// Properties for defining a `AWS::Batch::JobQueue`.
type CfnJobQueueProps struct {
	// `AWS::Batch::JobQueue.ComputeEnvironmentOrder`.
	ComputeEnvironmentOrder interface{} `json:"computeEnvironmentOrder"`
	// `AWS::Batch::JobQueue.Priority`.
	Priority *float64 `json:"priority"`
	// `AWS::Batch::JobQueue.JobQueueName`.
	JobQueueName *string `json:"jobQueueName"`
	// `AWS::Batch::JobQueue.State`.
	State *string `json:"state"`
	// `AWS::Batch::JobQueue.Tags`.
	Tags interface{} `json:"tags"`
}

