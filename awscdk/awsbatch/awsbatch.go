package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
	"github.com/aws/constructs-go/constructs/v3"
)

// Properties for how to prepare compute resources that are provisioned for a compute environment.
// Experimental.
type AllocationStrategy string

const (
	AllocationStrategy_BEST_FIT AllocationStrategy = "BEST_FIT"
	AllocationStrategy_BEST_FIT_PROGRESSIVE AllocationStrategy = "BEST_FIT_PROGRESSIVE"
	AllocationStrategy_SPOT_CAPACITY_OPTIMIZED AllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnComputeEnvironment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnComputeEnvironment(scope awscdk.Construct, id *string, props *CfnComputeEnvironmentProps) CfnComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_CfnComputeEnvironment{}

	_jsii_.Create(
		"monocdk.aws_batch.CfnComputeEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::ComputeEnvironment`.
func NewCfnComputeEnvironment_Override(c CfnComputeEnvironment, scope awscdk.Construct, id *string, props *CfnComputeEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.CfnComputeEnvironment",
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
		"monocdk.aws_batch.CfnComputeEnvironment",
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
		"monocdk.aws_batch.CfnComputeEnvironment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnComputeEnvironment",
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
		"monocdk.aws_batch.CfnComputeEnvironment",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) OnPrepare() {
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
func (c *jsiiProxy_CfnComputeEnvironment) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnComputeEnvironment) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnComputeEnvironment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnComputeEnvironment) Validate() *[]*string {
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnJobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnJobDefinition(scope awscdk.Construct, id *string, props *CfnJobDefinitionProps) CfnJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnJobDefinition{}

	_jsii_.Create(
		"monocdk.aws_batch.CfnJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::JobDefinition`.
func NewCfnJobDefinition_Override(c CfnJobDefinition, scope awscdk.Construct, id *string, props *CfnJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.CfnJobDefinition",
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
		"monocdk.aws_batch.CfnJobDefinition",
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
		"monocdk.aws_batch.CfnJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnJobDefinition",
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
		"monocdk.aws_batch.CfnJobDefinition",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) OnPrepare() {
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
func (c *jsiiProxy_CfnJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnJobDefinition) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnJobDefinition) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnJobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnJobDefinition) Validate() *[]*string {
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnJobQueue) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnJobQueue(scope awscdk.Construct, id *string, props *CfnJobQueueProps) CfnJobQueue {
	_init_.Initialize()

	j := jsiiProxy_CfnJobQueue{}

	_jsii_.Create(
		"monocdk.aws_batch.CfnJobQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::JobQueue`.
func NewCfnJobQueue_Override(c CfnJobQueue, scope awscdk.Construct, id *string, props *CfnJobQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.CfnJobQueue",
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
		"monocdk.aws_batch.CfnJobQueue",
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
		"monocdk.aws_batch.CfnJobQueue",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnJobQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnJobQueue",
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
		"monocdk.aws_batch.CfnJobQueue",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) OnPrepare() {
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
func (c *jsiiProxy_CfnJobQueue) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnJobQueue) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnJobQueue) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnJobQueue) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnJobQueue) Validate() *[]*string {
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

// Batch Compute Environment.
//
// Defines a batch compute environment to run batch jobs on.
// Experimental.
type ComputeEnvironment interface {
	awscdk.Resource
	IComputeEnvironment
	ComputeEnvironmentArn() *string
	ComputeEnvironmentName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
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

// The jsii proxy struct for ComputeEnvironment
type jsiiProxy_ComputeEnvironment struct {
	internal.Type__awscdkResource
	jsiiProxy_IComputeEnvironment
}

func (j *jsiiProxy_ComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeEnvironment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeEnvironment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeEnvironment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewComputeEnvironment(scope constructs.Construct, id *string, props *ComputeEnvironmentProps) ComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_ComputeEnvironment{}

	_jsii_.Create(
		"monocdk.aws_batch.ComputeEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewComputeEnvironment_Override(c ComputeEnvironment, scope constructs.Construct, id *string, props *ComputeEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.ComputeEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

// Fetches an existing batch compute environment by its amazon resource name.
// Experimental.
func ComputeEnvironment_FromComputeEnvironmentArn(scope constructs.Construct, id *string, computeEnvironmentArn *string) IComputeEnvironment {
	_init_.Initialize()

	var returns IComputeEnvironment

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.ComputeEnvironment",
		"fromComputeEnvironmentArn",
		[]interface{}{scope, id, computeEnvironmentArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.ComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ComputeEnvironment_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.ComputeEnvironment",
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
func (c *jsiiProxy_ComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_ComputeEnvironment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_ComputeEnvironment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_ComputeEnvironment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_ComputeEnvironment) OnPrepare() {
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
func (c *jsiiProxy_ComputeEnvironment) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_ComputeEnvironment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_ComputeEnvironment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_ComputeEnvironment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_ComputeEnvironment) ToString() *string {
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
func (c *jsiiProxy_ComputeEnvironment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a new Compute Environment.
// Experimental.
type ComputeEnvironmentProps struct {
	// A name for the compute environment.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	ComputeEnvironmentName *string `json:"computeEnvironmentName"`
	// The details of the required compute resources for the managed compute environment.
	//
	// If specified, and this is an unmanaged compute environment, will throw an error.
	//
	// By default, AWS Batch managed compute environments use a recent, approved version of the
	// Amazon ECS-optimized AMI for compute resources.
	// Experimental.
	ComputeResources *ComputeResources `json:"computeResources"`
	// The state of the compute environment.
	//
	// If the state is set to true, then the compute
	// environment accepts jobs from a queue and can scale out automatically based on queues.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// Determines if AWS should manage the allocation of compute resources for processing jobs.
	//
	// If set to false, then you are in charge of providing the compute resource details.
	// Experimental.
	Managed *bool `json:"managed"`
	// The IAM role used by Batch to make calls to other AWS services on your behalf for managing the resources that you use with the service.
	//
	// By default, this role is created for you using
	// the AWS managed service policy for Batch.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
}

// Property to specify if the compute environment uses On-Demand, SpotFleet, Fargate, or Fargate Spot compute resources.
// Experimental.
type ComputeResourceType string

const (
	ComputeResourceType_ON_DEMAND ComputeResourceType = "ON_DEMAND"
	ComputeResourceType_SPOT ComputeResourceType = "SPOT"
	ComputeResourceType_FARGATE ComputeResourceType = "FARGATE"
	ComputeResourceType_FARGATE_SPOT ComputeResourceType = "FARGATE_SPOT"
)

// Properties for defining the structure of the batch compute cluster.
// Experimental.
type ComputeResources struct {
	// The VPC network that all compute resources will be connected to.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated.
	//
	// This could be due to availability of the instance type in
	// the region or Amazon EC2 service limits. If this is not specified, the default for the EC2
	// ComputeResourceType is BEST_FIT, which will use only the best fitting instance type, waiting for
	// additional capacity if it's not available. This allocation strategy keeps costs lower but can limit
	// scaling. If you are using Spot Fleets with BEST_FIT then the Spot Fleet IAM Role must be specified.
	// BEST_FIT_PROGRESSIVE will select an additional instance type that is large enough to meet the
	// requirements of the jobs in the queue, with a preference for an instance type with a lower cost.
	// The default value for the SPOT instance type is SPOT_CAPACITY_OPTIMIZED, which is only available for
	// for this type of compute resources and will select an additional instance type that is large enough
	// to meet the requirements of the jobs in the queue, with a preference for an instance type that is
	// less likely to be interrupted.
	// Experimental.
	AllocationStrategy AllocationStrategy `json:"allocationStrategy"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for
	// that instance type before instances are launched. For example, if your maximum percentage is 20%,
	// then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. You always
	// pay the lowest (market) price and never more than your maximum percentage. If you leave this field empty,
	// the default value is 100% of the On-Demand price.
	// Experimental.
	BidPercentage *float64 `json:"bidPercentage"`
	// Key-value pair tags to be applied to resources that are launched in the compute environment.
	//
	// For AWS Batch, these take the form of "String1": "String2", where String1 is the tag key and
	// String2 is the tag value—for example, { "Name": "AWS Batch Instance - C4OnDemand" }.
	// Experimental.
	ComputeResourcesTags *map[string]*string `json:"computeResourcesTags"`
	// The desired number of EC2 vCPUS in the compute environment.
	// Experimental.
	DesiredvCpus *float64 `json:"desiredvCpus"`
	// The EC2 key pair that is used for instances launched in the compute environment.
	//
	// If no key is defined, then SSH access is not allowed to provisioned compute resources.
	// Experimental.
	Ec2KeyPair *string `json:"ec2KeyPair"`
	// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
	// Experimental.
	Image awsec2.IMachineImage `json:"image"`
	// The Amazon ECS instance profile applied to Amazon EC2 instances in a compute environment.
	//
	// You can specify
	// the short name or full Amazon Resource Name (ARN) of an instance profile. For example, ecsInstanceRole or
	// arn:aws:iam::<aws_account_id>:instance-profile/ecsInstanceRole . For more information, see Amazon ECS
	// Instance Role in the AWS Batch User Guide.
	// Experimental.
	InstanceRole *string `json:"instanceRole"`
	// The types of EC2 instances that may be launched in the compute environment.
	//
	// You can specify instance
	// families to launch any instance type within those families (for example, c4 or p3), or you can specify
	// specific sizes within a family (such as c4.8xlarge). You can also choose optimal to pick instance types
	// (from the C, M, and R instance families) on the fly that match the demand of your job queues.
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `json:"instanceTypes"`
	// An optional launch template to associate with your compute resources.
	//
	// For more information, see README file.
	// Experimental.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate"`
	// The maximum number of EC2 vCPUs that an environment can reach.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU.
	// Experimental.
	MaxvCpus *float64 `json:"maxvCpus"`
	// The minimum number of EC2 vCPUs that an environment should maintain (even if the compute environment state is DISABLED).
	//
	// Each vCPU is equivalent to 1,024 CPU shares. By keeping this set to 0 you will not have instance time wasted when
	// there is no work to be run. If you set this above zero you will maintain that number of vCPUs at all times.
	// Experimental.
	MinvCpus *float64 `json:"minvCpus"`
	// The Amazon EC2 placement group to associate with your compute resources.
	// Experimental.
	PlacementGroup *string `json:"placementGroup"`
	// The EC2 security group(s) associated with instances launched in the compute environment.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment.
	// For more information, see Amazon EC2 Spot Fleet Role in the AWS Batch User Guide.
	// Experimental.
	SpotFleetRole awsiam.IRole `json:"spotFleetRole"`
	// The type of compute environment: ON_DEMAND, SPOT, FARGATE, or FARGATE_SPOT.
	// Experimental.
	Type ComputeResourceType `json:"type"`
	// The VPC subnets into which the compute resources are launched.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Exposed secret for log configuration.
// Experimental.
type ExposedSecret interface {
	OptionName() *string
	SetOptionName(val *string)
	SecretArn() *string
	SetSecretArn(val *string)
}

// The jsii proxy struct for ExposedSecret
type jsiiProxy_ExposedSecret struct {
	_ byte // padding
}

func (j *jsiiProxy_ExposedSecret) OptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExposedSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewExposedSecret(optionName *string, secretArn *string) ExposedSecret {
	_init_.Initialize()

	j := jsiiProxy_ExposedSecret{}

	_jsii_.Create(
		"monocdk.aws_batch.ExposedSecret",
		[]interface{}{optionName, secretArn},
		&j,
	)

	return &j
}

// Experimental.
func NewExposedSecret_Override(e ExposedSecret, optionName *string, secretArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.ExposedSecret",
		[]interface{}{optionName, secretArn},
		e,
	)
}

func (j *jsiiProxy_ExposedSecret) SetOptionName(val *string) {
	_jsii_.Set(
		j,
		"optionName",
		val,
	)
}

func (j *jsiiProxy_ExposedSecret) SetSecretArn(val *string) {
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

// User Parameters Store Parameter.
// Experimental.
func ExposedSecret_FromParametersStore(optionName *string, parameter awsssm.IParameter) ExposedSecret {
	_init_.Initialize()

	var returns ExposedSecret

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.ExposedSecret",
		"fromParametersStore",
		[]interface{}{optionName, parameter},
		&returns,
	)

	return returns
}

// Use Secrets Manager Secret.
// Experimental.
func ExposedSecret_FromSecretsManager(optionName *string, secret awssecretsmanager.ISecret) ExposedSecret {
	_init_.Initialize()

	var returns ExposedSecret

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.ExposedSecret",
		"fromSecretsManager",
		[]interface{}{optionName, secret},
		&returns,
	)

	return returns
}

// Properties of a compute environment.
// Experimental.
type IComputeEnvironment interface {
	awscdk.IResource
	// The ARN of this compute environment.
	// Experimental.
	ComputeEnvironmentArn() *string
	// The name of this compute environment.
	// Experimental.
	ComputeEnvironmentName() *string
}

// The jsii proxy for IComputeEnvironment
type jsiiProxy_IComputeEnvironment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

// An interface representing a job definition - either a new one, created with the CDK, *using the {@link JobDefinition} class, or existing ones, referenced using the {@link JobDefinition.fromJobDefinitionArn} method.
// Experimental.
type IJobDefinition interface {
	awscdk.IResource
	// The ARN of this batch job definition.
	// Experimental.
	JobDefinitionArn() *string
	// The name of the batch job definition.
	// Experimental.
	JobDefinitionName() *string
}

// The jsii proxy for IJobDefinition
type jsiiProxy_IJobDefinition struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

// Properties of a Job Queue.
// Experimental.
type IJobQueue interface {
	awscdk.IResource
	// The ARN of this batch job queue.
	// Experimental.
	JobQueueArn() *string
	// A name for the job queue.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobQueueName() *string
}

// The jsii proxy for IJobQueue
type jsiiProxy_IJobQueue struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IJobQueue) JobQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) JobQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueName",
		&returns,
	)
	return returns
}

// Properties for specifying multi-node properties for compute resources.
// Experimental.
type IMultiNodeProps interface {
	// The number of nodes associated with a multi-node parallel job.
	// Experimental.
	Count() *float64
	// The number of nodes associated with a multi-node parallel job.
	// Experimental.
	SetCount(c *float64)
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	// Experimental.
	MainNode() *float64
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	// Experimental.
	SetMainNode(m *float64)
	// A list of node ranges and their properties associated with a multi-node parallel job.
	// Experimental.
	RangeProps() *[]INodeRangeProps
	// A list of node ranges and their properties associated with a multi-node parallel job.
	// Experimental.
	SetRangeProps(r *[]INodeRangeProps)
}

// The jsii proxy for IMultiNodeProps
type jsiiProxy_IMultiNodeProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IMultiNodeProps) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IMultiNodeProps) MainNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mainNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps) SetMainNode(val *float64) {
	_jsii_.Set(
		j,
		"mainNode",
		val,
	)
}

func (j *jsiiProxy_IMultiNodeProps) RangeProps() *[]INodeRangeProps {
	var returns *[]INodeRangeProps
	_jsii_.Get(
		j,
		"rangeProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiNodeProps) SetRangeProps(val *[]INodeRangeProps) {
	_jsii_.Set(
		j,
		"rangeProps",
		val,
	)
}

// Properties for a multi-node batch job.
// Experimental.
type INodeRangeProps interface {
	// The container details for the node range.
	// Experimental.
	Container() *JobDefinitionContainer
	// The container details for the node range.
	// Experimental.
	SetContainer(c *JobDefinitionContainer)
	// The minimum node index value to apply this container definition against.
	//
	// You may nest node ranges, for example 0:10 and 4:5, in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	FromNodeIndex() *float64
	// The minimum node index value to apply this container definition against.
	//
	// You may nest node ranges, for example 0:10 and 4:5, in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	SetFromNodeIndex(f *float64)
	// The maximum node index value to apply this container definition against. If omitted, the highest value is used relative.
	//
	// to the number of nodes associated with the job. You may nest node ranges, for example 0:10 and 4:5,
	// in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	ToNodeIndex() *float64
	// The maximum node index value to apply this container definition against. If omitted, the highest value is used relative.
	//
	// to the number of nodes associated with the job. You may nest node ranges, for example 0:10 and 4:5,
	// in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	SetToNodeIndex(t *float64)
}

// The jsii proxy for INodeRangeProps
type jsiiProxy_INodeRangeProps struct {
	_ byte // padding
}

func (j *jsiiProxy_INodeRangeProps) Container() *JobDefinitionContainer {
	var returns *JobDefinitionContainer
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetContainer(val *JobDefinitionContainer) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_INodeRangeProps) FromNodeIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromNodeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetFromNodeIndex(val *float64) {
	_jsii_.Set(
		j,
		"fromNodeIndex",
		val,
	)
}

func (j *jsiiProxy_INodeRangeProps) ToNodeIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toNodeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRangeProps) SetToNodeIndex(val *float64) {
	_jsii_.Set(
		j,
		"toNodeIndex",
		val,
	)
}

// Batch Job Definition.
//
// Defines a batch job definition to execute a specific batch job.
// Experimental.
type JobDefinition interface {
	awscdk.Resource
	IJobDefinition
	Env() *awscdk.ResourceEnvironment
	JobDefinitionArn() *string
	JobDefinitionName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
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

// The jsii proxy struct for JobDefinition
type jsiiProxy_JobDefinition struct {
	internal.Type__awscdkResource
	jsiiProxy_IJobDefinition
}

func (j *jsiiProxy_JobDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewJobDefinition(scope constructs.Construct, id *string, props *JobDefinitionProps) JobDefinition {
	_init_.Initialize()

	j := jsiiProxy_JobDefinition{}

	_jsii_.Create(
		"monocdk.aws_batch.JobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJobDefinition_Override(j JobDefinition, scope constructs.Construct, id *string, props *JobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.JobDefinition",
		[]interface{}{scope, id, props},
		j,
	)
}

// Imports an existing batch job definition by its amazon resource name.
// Experimental.
func JobDefinition_FromJobDefinitionArn(scope constructs.Construct, id *string, jobDefinitionArn *string) IJobDefinition {
	_init_.Initialize()

	var returns IJobDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"fromJobDefinitionArn",
		[]interface{}{scope, id, jobDefinitionArn},
		&returns,
	)

	return returns
}

// Imports an existing batch job definition by its name.
//
// If name is specified without a revision then the latest active revision is used.
// Experimental.
func JobDefinition_FromJobDefinitionName(scope constructs.Construct, id *string, jobDefinitionName *string) IJobDefinition {
	_init_.Initialize()

	var returns IJobDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"fromJobDefinitionName",
		[]interface{}{scope, id, jobDefinitionName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func JobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func JobDefinition_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
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
func (j *jsiiProxy_JobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (j *jsiiProxy_JobDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (j *jsiiProxy_JobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
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
func (j *jsiiProxy_JobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (j *jsiiProxy_JobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (j *jsiiProxy_JobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a job definition container.
// Experimental.
type JobDefinitionContainer struct {
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// Whether or not to assign a public IP to the job.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The IAM role that AWS Batch can assume.
	//
	// Required when using Fargate.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all
	// containers in a job should not exceed the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `json:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// Currently all node groups in a
	// multi-node parallel job must use the same instance type. This parameter is not valid
	// for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The IAM role that the container can assume for AWS permissions.
	// Experimental.
	JobRole awsiam.IRole `json:"jobRole"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	//
	// For now, only the `devices` property is supported.
	// Experimental.
	LinuxParams awsecs.LinuxParameters `json:"linuxParams"`
	// The log configuration specification for the container.
	// Experimental.
	LogConfiguration *LogConfiguration `json:"logConfiguration"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed
	// the memory specified here, the container is killed. You must specify at least 4 MiB of memory for EC2 and 512 MiB for Fargate.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The mount points for data volumes in your container.
	// Experimental.
	MountPoints *[]*awsecs.MountPoint `json:"mountPoints"`
	// Fargate platform version.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	// Experimental.
	Privileged *bool `json:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	// Experimental.
	ReadOnly *bool `json:"readOnly"`
	// A list of ulimits to set in the container.
	// Experimental.
	Ulimits *[]*awsecs.Ulimit `json:"ulimits"`
	// The user name to use inside the container.
	// Experimental.
	User *string `json:"user"`
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU for EC2 and 0.25 for Fargate.
	// Experimental.
	Vcpus *float64 `json:"vcpus"`
	// A list of data volumes used in a job.
	// Experimental.
	Volumes *[]*awsecs.Volume `json:"volumes"`
}

// Construction properties of the {@link JobDefinition} construct.
// Experimental.
type JobDefinitionProps struct {
	// An object with various properties specific to container-based jobs.
	// Experimental.
	Container *JobDefinitionContainer `json:"container"`
	// The name of the job definition.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobDefinitionName *string `json:"jobDefinitionName"`
	// An object with various properties specific to multi-node parallel jobs.
	// Experimental.
	NodeProps IMultiNodeProps `json:"nodeProps"`
	// When you submit a job, you can specify parameters that should replace the placeholders or override the default job definition parameters.
	//
	// Parameters
	// in job submission requests take precedence over the defaults in a job definition.
	// This allows you to use the same job definition for multiple jobs that use the same
	// format, and programmatically change values in the command at submission time.
	// Experimental.
	Parameters *map[string]*string `json:"parameters"`
	// The platform capabilities required by the job definition.
	// Experimental.
	PlatformCapabilities *[]PlatformCapabilities `json:"platformCapabilities"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and
	// 10 attempts. If the value of attempts is greater than one, the job is retried on failure
	// the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// The timeout configuration for jobs that are submitted with this job definition.
	//
	// You can specify
	// a timeout duration after which AWS Batch terminates your jobs if they have not finished.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
}

// Batch Job Queue.
//
// Defines a batch job queue to define how submitted batch jobs
// should be ran based on specified batch compute environments.
// Experimental.
type JobQueue interface {
	awscdk.Resource
	IJobQueue
	Env() *awscdk.ResourceEnvironment
	JobQueueArn() *string
	JobQueueName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
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

// The jsii proxy struct for JobQueue
type jsiiProxy_JobQueue struct {
	internal.Type__awscdkResource
	jsiiProxy_IJobQueue
}

func (j *jsiiProxy_JobQueue) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobQueue) JobQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobQueue) JobQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobQueue) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobQueue) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobQueue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewJobQueue(scope constructs.Construct, id *string, props *JobQueueProps) JobQueue {
	_init_.Initialize()

	j := jsiiProxy_JobQueue{}

	_jsii_.Create(
		"monocdk.aws_batch.JobQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJobQueue_Override(j JobQueue, scope constructs.Construct, id *string, props *JobQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.JobQueue",
		[]interface{}{scope, id, props},
		j,
	)
}

// Fetches an existing batch job queue by its amazon resource name.
// Experimental.
func JobQueue_FromJobQueueArn(scope constructs.Construct, id *string, jobQueueArn *string) IJobQueue {
	_init_.Initialize()

	var returns IJobQueue

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobQueue",
		"fromJobQueueArn",
		[]interface{}{scope, id, jobQueueArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func JobQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func JobQueue_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobQueue",
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
func (j *jsiiProxy_JobQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (j *jsiiProxy_JobQueue) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobQueue) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobQueue) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobQueue) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (j *jsiiProxy_JobQueue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
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
func (j *jsiiProxy_JobQueue) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobQueue) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (j *jsiiProxy_JobQueue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (j *jsiiProxy_JobQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JobQueue) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for mapping a compute environment to a job queue.
// Experimental.
type JobQueueComputeEnvironment struct {
	// The batch compute environment to use for processing submitted jobs to this queue.
	// Experimental.
	ComputeEnvironment IComputeEnvironment `json:"computeEnvironment"`
	// The order in which this compute environment will be selected for dynamic allocation of resources to process submitted jobs.
	// Experimental.
	Order *float64 `json:"order"`
}

// Properties of a batch job queue.
// Experimental.
type JobQueueProps struct {
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to
	// determine which compute environment should execute a given job. Compute environments must be in the VALID state before you can associate them
	// with a job queue. You can associate up to three compute environments with a job queue.
	// Experimental.
	ComputeEnvironments *[]*JobQueueComputeEnvironment `json:"computeEnvironments"`
	// The state of the job queue.
	//
	// If set to true, it is able to accept jobs.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// A name for the job queue.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobQueueName *string `json:"jobQueueName"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority (or a higher integer value for the priority parameter) are evaluated first
	// when associated with the same compute environment. Priority is determined in descending order, for example, a job queue with a priority value
	// of 10 is given scheduling preference over a job queue with a priority value of 1.
	// Experimental.
	Priority *float64 `json:"priority"`
}

// Launch template property specification.
// Experimental.
type LaunchTemplateSpecification struct {
	// The Launch template name.
	// Experimental.
	LaunchTemplateName *string `json:"launchTemplateName"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `json:"version"`
}

// Log configuration options to send to a custom log driver for the container.
// Experimental.
type LogConfiguration struct {
	// The log driver to use for the container.
	// Experimental.
	LogDriver LogDriver `json:"logDriver"`
	// The configuration options to send to the log driver.
	// Experimental.
	Options interface{} `json:"options"`
	// The secrets to pass to the log configuration as options.
	//
	// For more information, see https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data-secrets.html#secrets-logconfig
	// Experimental.
	SecretOptions *[]ExposedSecret `json:"secretOptions"`
}

// The log driver to use for the container.
// Experimental.
type LogDriver string

const (
	LogDriver_AWSLOGS LogDriver = "AWSLOGS"
	LogDriver_FLUENTD LogDriver = "FLUENTD"
	LogDriver_GELF LogDriver = "GELF"
	LogDriver_JOURNALD LogDriver = "JOURNALD"
	LogDriver_LOGENTRIES LogDriver = "LOGENTRIES"
	LogDriver_JSON_FILE LogDriver = "JSON_FILE"
	LogDriver_SPLUNK LogDriver = "SPLUNK"
	LogDriver_SYSLOG LogDriver = "SYSLOG"
)

// Platform capabilities.
// Experimental.
type PlatformCapabilities string

const (
	PlatformCapabilities_EC2 PlatformCapabilities = "EC2"
	PlatformCapabilities_FARGATE PlatformCapabilities = "FARGATE"
)

