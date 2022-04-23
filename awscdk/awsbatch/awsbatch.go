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
	// Batch will use the best fitting instance type will be used when assigning a batch job in this compute environment.
	// Experimental.
	AllocationStrategy_BEST_FIT AllocationStrategy = "BEST_FIT"
	// Batch will select additional instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types with a lower cost per unit vCPU.
	// Experimental.
	AllocationStrategy_BEST_FIT_PROGRESSIVE AllocationStrategy = "BEST_FIT_PROGRESSIVE"
	// This is only available for Spot Instance compute resources and will select additional instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types that are less likely to be interrupted.
	// Experimental.
	AllocationStrategy_SPOT_CAPACITY_OPTIMIZED AllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
)

// A CloudFormation `AWS::Batch::ComputeEnvironment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnComputeEnvironment := batch.NewCfnComputeEnvironment(this, jsii.String("MyCfnComputeEnvironment"), &cfnComputeEnvironmentProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	computeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	computeResources: &computeResourcesProperty{
//   		maxvCpus: jsii.Number(123),
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		bidPercentage: jsii.Number(123),
//   		desiredvCpus: jsii.Number(123),
//   		ec2Configuration: []interface{}{
//   			&ec2ConfigurationObjectProperty{
//   				imageType: jsii.String("imageType"),
//
//   				// the properties below are optional
//   				imageIdOverride: jsii.String("imageIdOverride"),
//   			},
//   		},
//   		ec2KeyPair: jsii.String("ec2KeyPair"),
//   		imageId: jsii.String("imageId"),
//   		instanceRole: jsii.String("instanceRole"),
//   		instanceTypes: []*string{
//   			jsii.String("instanceTypes"),
//   		},
//   		launchTemplate: &launchTemplateSpecificationProperty{
//   			launchTemplateId: jsii.String("launchTemplateId"),
//   			launchTemplateName: jsii.String("launchTemplateName"),
//   			version: jsii.String("version"),
//   		},
//   		minvCpus: jsii.Number(123),
//   		placementGroup: jsii.String("placementGroup"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		spotIamFleetRole: jsii.String("spotIamFleetRole"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	unmanagedvCpus: jsii.Number(123),
//   })
//
type CfnComputeEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrComputeEnvironmentArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Batch::ComputeEnvironment.ComputeEnvironmentName`.
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	// `AWS::Batch::ComputeEnvironment.ComputeResources`.
	ComputeResources() interface{}
	SetComputeResources(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Batch::ComputeEnvironment.ServiceRole`.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Batch::ComputeEnvironment.State`.
	State() *string
	SetState(val *string)
	// `AWS::Batch::ComputeEnvironment.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Batch::ComputeEnvironment.Type`.
	Type() *string
	SetType(val *string)
	// `AWS::Batch::ComputeEnvironment.UnmanagedvCpus`.
	UnmanagedvCpus() *float64
	SetUnmanagedvCpus(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComputeEnvironment
type jsiiProxy_CfnComputeEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComputeEnvironment) AttrComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComputeEnvironmentArn",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnComputeEnvironment) UnmanagedvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unmanagedvCpus",
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

func (j *jsiiProxy_CfnComputeEnvironment) SetUnmanagedvCpus(val *float64) {
	_jsii_.Set(
		j,
		"unmanagedvCpus",
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

func (c *jsiiProxy_CfnComputeEnvironment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnComputeEnvironment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnComputeEnvironment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnComputeEnvironment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   computeResourcesProperty := &computeResourcesProperty{
//   	maxvCpus: jsii.Number(123),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	bidPercentage: jsii.Number(123),
//   	desiredvCpus: jsii.Number(123),
//   	ec2Configuration: []interface{}{
//   		&ec2ConfigurationObjectProperty{
//   			imageType: jsii.String("imageType"),
//
//   			// the properties below are optional
//   			imageIdOverride: jsii.String("imageIdOverride"),
//   		},
//   	},
//   	ec2KeyPair: jsii.String("ec2KeyPair"),
//   	imageId: jsii.String("imageId"),
//   	instanceRole: jsii.String("instanceRole"),
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   		version: jsii.String("version"),
//   	},
//   	minvCpus: jsii.Number(123),
//   	placementGroup: jsii.String("placementGroup"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	spotIamFleetRole: jsii.String("spotIamFleetRole"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnComputeEnvironment_ComputeResourcesProperty struct {
	// `CfnComputeEnvironment.ComputeResourcesProperty.MaxvCpus`.
	MaxvCpus *float64 `json:"maxvCpus" yaml:"maxvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Subnets`.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy" yaml:"allocationStrategy"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.BidPercentage`.
	BidPercentage *float64 `json:"bidPercentage" yaml:"bidPercentage"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.DesiredvCpus`.
	DesiredvCpus *float64 `json:"desiredvCpus" yaml:"desiredvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Ec2Configuration`.
	Ec2Configuration interface{} `json:"ec2Configuration" yaml:"ec2Configuration"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Ec2KeyPair`.
	Ec2KeyPair *string `json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.ImageId`.
	ImageId *string `json:"imageId" yaml:"imageId"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.InstanceRole`.
	InstanceRole *string `json:"instanceRole" yaml:"instanceRole"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.InstanceTypes`.
	InstanceTypes *[]*string `json:"instanceTypes" yaml:"instanceTypes"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.LaunchTemplate`.
	LaunchTemplate interface{} `json:"launchTemplate" yaml:"launchTemplate"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.MinvCpus`.
	MinvCpus *float64 `json:"minvCpus" yaml:"minvCpus"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.PlacementGroup`.
	PlacementGroup *string `json:"placementGroup" yaml:"placementGroup"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.SpotIamFleetRole`.
	SpotIamFleetRole *string `json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// `CfnComputeEnvironment.ComputeResourcesProperty.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   ec2ConfigurationObjectProperty := &ec2ConfigurationObjectProperty{
//   	imageType: jsii.String("imageType"),
//
//   	// the properties below are optional
//   	imageIdOverride: jsii.String("imageIdOverride"),
//   }
//
type CfnComputeEnvironment_Ec2ConfigurationObjectProperty struct {
	// `CfnComputeEnvironment.Ec2ConfigurationObjectProperty.ImageType`.
	ImageType *string `json:"imageType" yaml:"imageType"`
	// `CfnComputeEnvironment.Ec2ConfigurationObjectProperty.ImageIdOverride`.
	ImageIdOverride *string `json:"imageIdOverride" yaml:"imageIdOverride"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   launchTemplateSpecificationProperty := &launchTemplateSpecificationProperty{
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   	version: jsii.String("version"),
//   }
//
type CfnComputeEnvironment_LaunchTemplateSpecificationProperty struct {
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.LaunchTemplateId`.
	LaunchTemplateId *string `json:"launchTemplateId" yaml:"launchTemplateId"`
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.LaunchTemplateName`.
	LaunchTemplateName *string `json:"launchTemplateName" yaml:"launchTemplateName"`
	// `CfnComputeEnvironment.LaunchTemplateSpecificationProperty.Version`.
	Version *string `json:"version" yaml:"version"`
}

// Properties for defining a `CfnComputeEnvironment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnComputeEnvironmentProps := &cfnComputeEnvironmentProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	computeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	computeResources: &computeResourcesProperty{
//   		maxvCpus: jsii.Number(123),
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		bidPercentage: jsii.Number(123),
//   		desiredvCpus: jsii.Number(123),
//   		ec2Configuration: []interface{}{
//   			&ec2ConfigurationObjectProperty{
//   				imageType: jsii.String("imageType"),
//
//   				// the properties below are optional
//   				imageIdOverride: jsii.String("imageIdOverride"),
//   			},
//   		},
//   		ec2KeyPair: jsii.String("ec2KeyPair"),
//   		imageId: jsii.String("imageId"),
//   		instanceRole: jsii.String("instanceRole"),
//   		instanceTypes: []*string{
//   			jsii.String("instanceTypes"),
//   		},
//   		launchTemplate: &launchTemplateSpecificationProperty{
//   			launchTemplateId: jsii.String("launchTemplateId"),
//   			launchTemplateName: jsii.String("launchTemplateName"),
//   			version: jsii.String("version"),
//   		},
//   		minvCpus: jsii.Number(123),
//   		placementGroup: jsii.String("placementGroup"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		spotIamFleetRole: jsii.String("spotIamFleetRole"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	unmanagedvCpus: jsii.Number(123),
//   }
//
type CfnComputeEnvironmentProps struct {
	// `AWS::Batch::ComputeEnvironment.Type`.
	Type *string `json:"type" yaml:"type"`
	// `AWS::Batch::ComputeEnvironment.ComputeEnvironmentName`.
	ComputeEnvironmentName *string `json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// `AWS::Batch::ComputeEnvironment.ComputeResources`.
	ComputeResources interface{} `json:"computeResources" yaml:"computeResources"`
	// `AWS::Batch::ComputeEnvironment.ServiceRole`.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// `AWS::Batch::ComputeEnvironment.State`.
	State *string `json:"state" yaml:"state"`
	// `AWS::Batch::ComputeEnvironment.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Batch::ComputeEnvironment.UnmanagedvCpus`.
	UnmanagedvCpus *float64 `json:"unmanagedvCpus" yaml:"unmanagedvCpus"`
}

// A CloudFormation `AWS::Batch::JobDefinition`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   var parameters interface{}
//   var tags interface{}
//   cfnJobDefinition := batch.NewCfnJobDefinition(this, jsii.String("MyCfnJobDefinition"), &cfnJobDefinitionProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	containerProperties: &containerPropertiesProperty{
//   		image: jsii.String("image"),
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&environmentProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   			platformVersion: jsii.String("platformVersion"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		jobRoleArn: jsii.String("jobRoleArn"),
//   		linuxParameters: &linuxParametersProperty{
//   			devices: []interface{}{
//   				&deviceProperty{
//   					containerPath: jsii.String("containerPath"),
//   					hostPath: jsii.String("hostPath"),
//   					permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			initProcessEnabled: jsii.Boolean(false),
//   			maxSwap: jsii.Number(123),
//   			sharedMemorySize: jsii.Number(123),
//   			swappiness: jsii.Number(123),
//   			tmpfs: []interface{}{
//   				&tmpfsProperty{
//   					containerPath: jsii.String("containerPath"),
//   					size: jsii.Number(123),
//
//   					// the properties below are optional
//   					mountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			options: options,
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		memory: jsii.Number(123),
//   		mountPoints: []interface{}{
//   			&mountPointsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				readOnly: jsii.Boolean(false),
//   				sourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		networkConfiguration: &networkConfigurationProperty{
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		privileged: jsii.Boolean(false),
//   		readonlyRootFilesystem: jsii.Boolean(false),
//   		resourceRequirements: []interface{}{
//   			&resourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		secrets: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		ulimits: []interface{}{
//   			&ulimitProperty{
//   				hardLimit: jsii.Number(123),
//   				name: jsii.String("name"),
//   				softLimit: jsii.Number(123),
//   			},
//   		},
//   		user: jsii.String("user"),
//   		vcpus: jsii.Number(123),
//   		volumes: []interface{}{
//   			&volumesProperty{
//   				efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   					fileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					authorizationConfig: &authorizationConfigProperty{
//   						accessPointId: jsii.String("accessPointId"),
//   						iam: jsii.String("iam"),
//   					},
//   					rootDirectory: jsii.String("rootDirectory"),
//   					transitEncryption: jsii.String("transitEncryption"),
//   					transitEncryptionPort: jsii.Number(123),
//   				},
//   				host: &volumesHostProperty{
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	nodeProperties: &nodePropertiesProperty{
//   		mainNode: jsii.Number(123),
//   		nodeRangeProperties: []interface{}{
//   			&nodeRangePropertyProperty{
//   				targetNodes: jsii.String("targetNodes"),
//
//   				// the properties below are optional
//   				container: &containerPropertiesProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					environment: []interface{}{
//   						&environmentProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					executionRoleArn: jsii.String("executionRoleArn"),
//   					fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   						platformVersion: jsii.String("platformVersion"),
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					jobRoleArn: jsii.String("jobRoleArn"),
//   					linuxParameters: &linuxParametersProperty{
//   						devices: []interface{}{
//   							&deviceProperty{
//   								containerPath: jsii.String("containerPath"),
//   								hostPath: jsii.String("hostPath"),
//   								permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						initProcessEnabled: jsii.Boolean(false),
//   						maxSwap: jsii.Number(123),
//   						sharedMemorySize: jsii.Number(123),
//   						swappiness: jsii.Number(123),
//   						tmpfs: []interface{}{
//   							&tmpfsProperty{
//   								containerPath: jsii.String("containerPath"),
//   								size: jsii.Number(123),
//
//   								// the properties below are optional
//   								mountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					logConfiguration: &logConfigurationProperty{
//   						logDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						options: options,
//   						secretOptions: []interface{}{
//   							&secretProperty{
//   								name: jsii.String("name"),
//   								valueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					memory: jsii.Number(123),
//   					mountPoints: []interface{}{
//   						&mountPointsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							readOnly: jsii.Boolean(false),
//   							sourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					networkConfiguration: &networkConfigurationProperty{
//   						assignPublicIp: jsii.String("assignPublicIp"),
//   					},
//   					privileged: jsii.Boolean(false),
//   					readonlyRootFilesystem: jsii.Boolean(false),
//   					resourceRequirements: []interface{}{
//   						&resourceRequirementProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					secrets: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					ulimits: []interface{}{
//   						&ulimitProperty{
//   							hardLimit: jsii.Number(123),
//   							name: jsii.String("name"),
//   							softLimit: jsii.Number(123),
//   						},
//   					},
//   					user: jsii.String("user"),
//   					vcpus: jsii.Number(123),
//   					volumes: []interface{}{
//   						&volumesProperty{
//   							efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   								fileSystemId: jsii.String("fileSystemId"),
//
//   								// the properties below are optional
//   								authorizationConfig: &authorizationConfigProperty{
//   									accessPointId: jsii.String("accessPointId"),
//   									iam: jsii.String("iam"),
//   								},
//   								rootDirectory: jsii.String("rootDirectory"),
//   								transitEncryption: jsii.String("transitEncryption"),
//   								transitEncryptionPort: jsii.Number(123),
//   							},
//   							host: &volumesHostProperty{
//   								sourcePath: jsii.String("sourcePath"),
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		numNodes: jsii.Number(123),
//   	},
//   	parameters: parameters,
//   	platformCapabilities: []*string{
//   		jsii.String("platformCapabilities"),
//   	},
//   	propagateTags: jsii.Boolean(false),
//   	retryStrategy: &retryStrategyProperty{
//   		attempts: jsii.Number(123),
//   		evaluateOnExit: []interface{}{
//   			&evaluateOnExitProperty{
//   				action: jsii.String("action"),
//
//   				// the properties below are optional
//   				onExitCode: jsii.String("onExitCode"),
//   				onReason: jsii.String("onReason"),
//   				onStatusReason: jsii.String("onStatusReason"),
//   			},
//   		},
//   	},
//   	schedulingPriority: jsii.Number(123),
//   	tags: tags,
//   	timeout: &timeoutProperty{
//   		attemptDurationSeconds: jsii.Number(123),
//   	},
//   })
//
type CfnJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Batch::JobDefinition.ContainerProperties`.
	ContainerProperties() interface{}
	SetContainerProperties(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Batch::JobDefinition.JobDefinitionName`.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::Batch::JobDefinition.NodeProperties`.
	NodeProperties() interface{}
	SetNodeProperties(val interface{})
	// `AWS::Batch::JobDefinition.Parameters`.
	Parameters() interface{}
	SetParameters(val interface{})
	// `AWS::Batch::JobDefinition.PlatformCapabilities`.
	PlatformCapabilities() *[]*string
	SetPlatformCapabilities(val *[]*string)
	// `AWS::Batch::JobDefinition.PropagateTags`.
	PropagateTags() interface{}
	SetPropagateTags(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Batch::JobDefinition.RetryStrategy`.
	RetryStrategy() interface{}
	SetRetryStrategy(val interface{})
	// `AWS::Batch::JobDefinition.SchedulingPriority`.
	SchedulingPriority() *float64
	SetSchedulingPriority(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Batch::JobDefinition.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Batch::JobDefinition.Timeout`.
	Timeout() interface{}
	SetTimeout(val interface{})
	// `AWS::Batch::JobDefinition.Type`.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnJobDefinition) SchedulingPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingPriority",
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

func (j *jsiiProxy_CfnJobDefinition) SetSchedulingPriority(val *float64) {
	_jsii_.Set(
		j,
		"schedulingPriority",
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

func (c *jsiiProxy_CfnJobDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnJobDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   authorizationConfigProperty := &authorizationConfigProperty{
//   	accessPointId: jsii.String("accessPointId"),
//   	iam: jsii.String("iam"),
//   }
//
type CfnJobDefinition_AuthorizationConfigProperty struct {
	// `CfnJobDefinition.AuthorizationConfigProperty.AccessPointId`.
	AccessPointId *string `json:"accessPointId" yaml:"accessPointId"`
	// `CfnJobDefinition.AuthorizationConfigProperty.Iam`.
	Iam *string `json:"iam" yaml:"iam"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   containerPropertiesProperty := &containerPropertiesProperty{
//   	image: jsii.String("image"),
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	environment: []interface{}{
//   		&environmentProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   		platformVersion: jsii.String("platformVersion"),
//   	},
//   	instanceType: jsii.String("instanceType"),
//   	jobRoleArn: jsii.String("jobRoleArn"),
//   	linuxParameters: &linuxParametersProperty{
//   		devices: []interface{}{
//   			&deviceProperty{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//   				permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   			},
//   		},
//   		initProcessEnabled: jsii.Boolean(false),
//   		maxSwap: jsii.Number(123),
//   		sharedMemorySize: jsii.Number(123),
//   		swappiness: jsii.Number(123),
//   		tmpfs: []interface{}{
//   			&tmpfsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				size: jsii.Number(123),
//
//   				// the properties below are optional
//   				mountOptions: []*string{
//   					jsii.String("mountOptions"),
//   				},
//   			},
//   		},
//   	},
//   	logConfiguration: &logConfigurationProperty{
//   		logDriver: jsii.String("logDriver"),
//
//   		// the properties below are optional
//   		options: options,
//   		secretOptions: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	memory: jsii.Number(123),
//   	mountPoints: []interface{}{
//   		&mountPointsProperty{
//   			containerPath: jsii.String("containerPath"),
//   			readOnly: jsii.Boolean(false),
//   			sourceVolume: jsii.String("sourceVolume"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		assignPublicIp: jsii.String("assignPublicIp"),
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	resourceRequirements: []interface{}{
//   		&resourceRequirementProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	secrets: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   	ulimits: []interface{}{
//   		&ulimitProperty{
//   			hardLimit: jsii.Number(123),
//   			name: jsii.String("name"),
//   			softLimit: jsii.Number(123),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	vcpus: jsii.Number(123),
//   	volumes: []interface{}{
//   		&volumesProperty{
//   			efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   				fileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfigProperty{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &volumesHostProperty{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnJobDefinition_ContainerPropertiesProperty struct {
	// `CfnJobDefinition.ContainerPropertiesProperty.Image`.
	Image *string `json:"image" yaml:"image"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Command`.
	Command *[]*string `json:"command" yaml:"command"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Environment`.
	Environment interface{} `json:"environment" yaml:"environment"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// `CfnJobDefinition.ContainerPropertiesProperty.FargatePlatformConfiguration`.
	FargatePlatformConfiguration interface{} `json:"fargatePlatformConfiguration" yaml:"fargatePlatformConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.InstanceType`.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// `CfnJobDefinition.ContainerPropertiesProperty.JobRoleArn`.
	JobRoleArn *string `json:"jobRoleArn" yaml:"jobRoleArn"`
	// `CfnJobDefinition.ContainerPropertiesProperty.LinuxParameters`.
	LinuxParameters interface{} `json:"linuxParameters" yaml:"linuxParameters"`
	// `CfnJobDefinition.ContainerPropertiesProperty.LogConfiguration`.
	LogConfiguration interface{} `json:"logConfiguration" yaml:"logConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Memory`.
	Memory *float64 `json:"memory" yaml:"memory"`
	// `CfnJobDefinition.ContainerPropertiesProperty.MountPoints`.
	MountPoints interface{} `json:"mountPoints" yaml:"mountPoints"`
	// `CfnJobDefinition.ContainerPropertiesProperty.NetworkConfiguration`.
	NetworkConfiguration interface{} `json:"networkConfiguration" yaml:"networkConfiguration"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Privileged`.
	Privileged interface{} `json:"privileged" yaml:"privileged"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ReadonlyRootFilesystem`.
	ReadonlyRootFilesystem interface{} `json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// `CfnJobDefinition.ContainerPropertiesProperty.ResourceRequirements`.
	ResourceRequirements interface{} `json:"resourceRequirements" yaml:"resourceRequirements"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Secrets`.
	Secrets interface{} `json:"secrets" yaml:"secrets"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Ulimits`.
	Ulimits interface{} `json:"ulimits" yaml:"ulimits"`
	// `CfnJobDefinition.ContainerPropertiesProperty.User`.
	User *string `json:"user" yaml:"user"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Vcpus`.
	Vcpus *float64 `json:"vcpus" yaml:"vcpus"`
	// `CfnJobDefinition.ContainerPropertiesProperty.Volumes`.
	Volumes interface{} `json:"volumes" yaml:"volumes"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   deviceProperty := &deviceProperty{
//   	containerPath: jsii.String("containerPath"),
//   	hostPath: jsii.String("hostPath"),
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   }
//
type CfnJobDefinition_DeviceProperty struct {
	// `CfnJobDefinition.DeviceProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath" yaml:"containerPath"`
	// `CfnJobDefinition.DeviceProperty.HostPath`.
	HostPath *string `json:"hostPath" yaml:"hostPath"`
	// `CfnJobDefinition.DeviceProperty.Permissions`.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   efsVolumeConfigurationProperty := &efsVolumeConfigurationProperty{
//   	fileSystemId: jsii.String("fileSystemId"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfigProperty{
//   		accessPointId: jsii.String("accessPointId"),
//   		iam: jsii.String("iam"),
//   	},
//   	rootDirectory: jsii.String("rootDirectory"),
//   	transitEncryption: jsii.String("transitEncryption"),
//   	transitEncryptionPort: jsii.Number(123),
//   }
//
type CfnJobDefinition_EfsVolumeConfigurationProperty struct {
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.FileSystemId`.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.AuthorizationConfig`.
	AuthorizationConfig interface{} `json:"authorizationConfig" yaml:"authorizationConfig"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.RootDirectory`.
	RootDirectory *string `json:"rootDirectory" yaml:"rootDirectory"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.TransitEncryption`.
	TransitEncryption *string `json:"transitEncryption" yaml:"transitEncryption"`
	// `CfnJobDefinition.EfsVolumeConfigurationProperty.TransitEncryptionPort`.
	TransitEncryptionPort *float64 `json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   environmentProperty := &environmentProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnJobDefinition_EnvironmentProperty struct {
	// `CfnJobDefinition.EnvironmentProperty.Name`.
	Name *string `json:"name" yaml:"name"`
	// `CfnJobDefinition.EnvironmentProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   evaluateOnExitProperty := &evaluateOnExitProperty{
//   	action: jsii.String("action"),
//
//   	// the properties below are optional
//   	onExitCode: jsii.String("onExitCode"),
//   	onReason: jsii.String("onReason"),
//   	onStatusReason: jsii.String("onStatusReason"),
//   }
//
type CfnJobDefinition_EvaluateOnExitProperty struct {
	// `CfnJobDefinition.EvaluateOnExitProperty.Action`.
	Action *string `json:"action" yaml:"action"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnExitCode`.
	OnExitCode *string `json:"onExitCode" yaml:"onExitCode"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnReason`.
	OnReason *string `json:"onReason" yaml:"onReason"`
	// `CfnJobDefinition.EvaluateOnExitProperty.OnStatusReason`.
	OnStatusReason *string `json:"onStatusReason" yaml:"onStatusReason"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   fargatePlatformConfigurationProperty := &fargatePlatformConfigurationProperty{
//   	platformVersion: jsii.String("platformVersion"),
//   }
//
type CfnJobDefinition_FargatePlatformConfigurationProperty struct {
	// `CfnJobDefinition.FargatePlatformConfigurationProperty.PlatformVersion`.
	PlatformVersion *string `json:"platformVersion" yaml:"platformVersion"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   linuxParametersProperty := &linuxParametersProperty{
//   	devices: []interface{}{
//   		&deviceProperty{
//   			containerPath: jsii.String("containerPath"),
//   			hostPath: jsii.String("hostPath"),
//   			permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   		},
//   	},
//   	initProcessEnabled: jsii.Boolean(false),
//   	maxSwap: jsii.Number(123),
//   	sharedMemorySize: jsii.Number(123),
//   	swappiness: jsii.Number(123),
//   	tmpfs: []interface{}{
//   		&tmpfsProperty{
//   			containerPath: jsii.String("containerPath"),
//   			size: jsii.Number(123),
//
//   			// the properties below are optional
//   			mountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_LinuxParametersProperty struct {
	// `CfnJobDefinition.LinuxParametersProperty.Devices`.
	Devices interface{} `json:"devices" yaml:"devices"`
	// `CfnJobDefinition.LinuxParametersProperty.InitProcessEnabled`.
	InitProcessEnabled interface{} `json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// `CfnJobDefinition.LinuxParametersProperty.MaxSwap`.
	MaxSwap *float64 `json:"maxSwap" yaml:"maxSwap"`
	// `CfnJobDefinition.LinuxParametersProperty.SharedMemorySize`.
	SharedMemorySize *float64 `json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// `CfnJobDefinition.LinuxParametersProperty.Swappiness`.
	Swappiness *float64 `json:"swappiness" yaml:"swappiness"`
	// `CfnJobDefinition.LinuxParametersProperty.Tmpfs`.
	Tmpfs interface{} `json:"tmpfs" yaml:"tmpfs"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   logConfigurationProperty := &logConfigurationProperty{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: options,
//   	secretOptions: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
type CfnJobDefinition_LogConfigurationProperty struct {
	// `CfnJobDefinition.LogConfigurationProperty.LogDriver`.
	LogDriver *string `json:"logDriver" yaml:"logDriver"`
	// `CfnJobDefinition.LogConfigurationProperty.Options`.
	Options interface{} `json:"options" yaml:"options"`
	// `CfnJobDefinition.LogConfigurationProperty.SecretOptions`.
	SecretOptions interface{} `json:"secretOptions" yaml:"secretOptions"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   mountPointsProperty := &mountPointsProperty{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type CfnJobDefinition_MountPointsProperty struct {
	// `CfnJobDefinition.MountPointsProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath" yaml:"containerPath"`
	// `CfnJobDefinition.MountPointsProperty.ReadOnly`.
	ReadOnly interface{} `json:"readOnly" yaml:"readOnly"`
	// `CfnJobDefinition.MountPointsProperty.SourceVolume`.
	SourceVolume *string `json:"sourceVolume" yaml:"sourceVolume"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   }
//
type CfnJobDefinition_NetworkConfigurationProperty struct {
	// `CfnJobDefinition.NetworkConfigurationProperty.AssignPublicIp`.
	AssignPublicIp *string `json:"assignPublicIp" yaml:"assignPublicIp"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   nodePropertiesProperty := &nodePropertiesProperty{
//   	mainNode: jsii.Number(123),
//   	nodeRangeProperties: []interface{}{
//   		&nodeRangePropertyProperty{
//   			targetNodes: jsii.String("targetNodes"),
//
//   			// the properties below are optional
//   			container: &containerPropertiesProperty{
//   				image: jsii.String("image"),
//
//   				// the properties below are optional
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				environment: []interface{}{
//   					&environmentProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   					platformVersion: jsii.String("platformVersion"),
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				jobRoleArn: jsii.String("jobRoleArn"),
//   				linuxParameters: &linuxParametersProperty{
//   					devices: []interface{}{
//   						&deviceProperty{
//   							containerPath: jsii.String("containerPath"),
//   							hostPath: jsii.String("hostPath"),
//   							permissions: []*string{
//   								jsii.String("permissions"),
//   							},
//   						},
//   					},
//   					initProcessEnabled: jsii.Boolean(false),
//   					maxSwap: jsii.Number(123),
//   					sharedMemorySize: jsii.Number(123),
//   					swappiness: jsii.Number(123),
//   					tmpfs: []interface{}{
//   						&tmpfsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							size: jsii.Number(123),
//
//   							// the properties below are optional
//   							mountOptions: []*string{
//   								jsii.String("mountOptions"),
//   							},
//   						},
//   					},
//   				},
//   				logConfiguration: &logConfigurationProperty{
//   					logDriver: jsii.String("logDriver"),
//
//   					// the properties below are optional
//   					options: options,
//   					secretOptions: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   				},
//   				memory: jsii.Number(123),
//   				mountPoints: []interface{}{
//   					&mountPointsProperty{
//   						containerPath: jsii.String("containerPath"),
//   						readOnly: jsii.Boolean(false),
//   						sourceVolume: jsii.String("sourceVolume"),
//   					},
//   				},
//   				networkConfiguration: &networkConfigurationProperty{
//   					assignPublicIp: jsii.String("assignPublicIp"),
//   				},
//   				privileged: jsii.Boolean(false),
//   				readonlyRootFilesystem: jsii.Boolean(false),
//   				resourceRequirements: []interface{}{
//   					&resourceRequirementProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				secrets: []interface{}{
//   					&secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   				ulimits: []interface{}{
//   					&ulimitProperty{
//   						hardLimit: jsii.Number(123),
//   						name: jsii.String("name"),
//   						softLimit: jsii.Number(123),
//   					},
//   				},
//   				user: jsii.String("user"),
//   				vcpus: jsii.Number(123),
//   				volumes: []interface{}{
//   					&volumesProperty{
//   						efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   							fileSystemId: jsii.String("fileSystemId"),
//
//   							// the properties below are optional
//   							authorizationConfig: &authorizationConfigProperty{
//   								accessPointId: jsii.String("accessPointId"),
//   								iam: jsii.String("iam"),
//   							},
//   							rootDirectory: jsii.String("rootDirectory"),
//   							transitEncryption: jsii.String("transitEncryption"),
//   							transitEncryptionPort: jsii.Number(123),
//   						},
//   						host: &volumesHostProperty{
//   							sourcePath: jsii.String("sourcePath"),
//   						},
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	numNodes: jsii.Number(123),
//   }
//
type CfnJobDefinition_NodePropertiesProperty struct {
	// `CfnJobDefinition.NodePropertiesProperty.MainNode`.
	MainNode *float64 `json:"mainNode" yaml:"mainNode"`
	// `CfnJobDefinition.NodePropertiesProperty.NodeRangeProperties`.
	NodeRangeProperties interface{} `json:"nodeRangeProperties" yaml:"nodeRangeProperties"`
	// `CfnJobDefinition.NodePropertiesProperty.NumNodes`.
	NumNodes *float64 `json:"numNodes" yaml:"numNodes"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   nodeRangePropertyProperty := &nodeRangePropertyProperty{
//   	targetNodes: jsii.String("targetNodes"),
//
//   	// the properties below are optional
//   	container: &containerPropertiesProperty{
//   		image: jsii.String("image"),
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&environmentProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   			platformVersion: jsii.String("platformVersion"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		jobRoleArn: jsii.String("jobRoleArn"),
//   		linuxParameters: &linuxParametersProperty{
//   			devices: []interface{}{
//   				&deviceProperty{
//   					containerPath: jsii.String("containerPath"),
//   					hostPath: jsii.String("hostPath"),
//   					permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			initProcessEnabled: jsii.Boolean(false),
//   			maxSwap: jsii.Number(123),
//   			sharedMemorySize: jsii.Number(123),
//   			swappiness: jsii.Number(123),
//   			tmpfs: []interface{}{
//   				&tmpfsProperty{
//   					containerPath: jsii.String("containerPath"),
//   					size: jsii.Number(123),
//
//   					// the properties below are optional
//   					mountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			options: options,
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		memory: jsii.Number(123),
//   		mountPoints: []interface{}{
//   			&mountPointsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				readOnly: jsii.Boolean(false),
//   				sourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		networkConfiguration: &networkConfigurationProperty{
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		privileged: jsii.Boolean(false),
//   		readonlyRootFilesystem: jsii.Boolean(false),
//   		resourceRequirements: []interface{}{
//   			&resourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		secrets: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		ulimits: []interface{}{
//   			&ulimitProperty{
//   				hardLimit: jsii.Number(123),
//   				name: jsii.String("name"),
//   				softLimit: jsii.Number(123),
//   			},
//   		},
//   		user: jsii.String("user"),
//   		vcpus: jsii.Number(123),
//   		volumes: []interface{}{
//   			&volumesProperty{
//   				efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   					fileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					authorizationConfig: &authorizationConfigProperty{
//   						accessPointId: jsii.String("accessPointId"),
//   						iam: jsii.String("iam"),
//   					},
//   					rootDirectory: jsii.String("rootDirectory"),
//   					transitEncryption: jsii.String("transitEncryption"),
//   					transitEncryptionPort: jsii.Number(123),
//   				},
//   				host: &volumesHostProperty{
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_NodeRangePropertyProperty struct {
	// `CfnJobDefinition.NodeRangePropertyProperty.TargetNodes`.
	TargetNodes *string `json:"targetNodes" yaml:"targetNodes"`
	// `CfnJobDefinition.NodeRangePropertyProperty.Container`.
	Container interface{} `json:"container" yaml:"container"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   resourceRequirementProperty := &resourceRequirementProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnJobDefinition_ResourceRequirementProperty struct {
	// `CfnJobDefinition.ResourceRequirementProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnJobDefinition.ResourceRequirementProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   retryStrategyProperty := &retryStrategyProperty{
//   	attempts: jsii.Number(123),
//   	evaluateOnExit: []interface{}{
//   		&evaluateOnExitProperty{
//   			action: jsii.String("action"),
//
//   			// the properties below are optional
//   			onExitCode: jsii.String("onExitCode"),
//   			onReason: jsii.String("onReason"),
//   			onStatusReason: jsii.String("onStatusReason"),
//   		},
//   	},
//   }
//
type CfnJobDefinition_RetryStrategyProperty struct {
	// `CfnJobDefinition.RetryStrategyProperty.Attempts`.
	Attempts *float64 `json:"attempts" yaml:"attempts"`
	// `CfnJobDefinition.RetryStrategyProperty.EvaluateOnExit`.
	EvaluateOnExit interface{} `json:"evaluateOnExit" yaml:"evaluateOnExit"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   secretProperty := &secretProperty{
//   	name: jsii.String("name"),
//   	valueFrom: jsii.String("valueFrom"),
//   }
//
type CfnJobDefinition_SecretProperty struct {
	// `CfnJobDefinition.SecretProperty.Name`.
	Name *string `json:"name" yaml:"name"`
	// `CfnJobDefinition.SecretProperty.ValueFrom`.
	ValueFrom *string `json:"valueFrom" yaml:"valueFrom"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   timeoutProperty := &timeoutProperty{
//   	attemptDurationSeconds: jsii.Number(123),
//   }
//
type CfnJobDefinition_TimeoutProperty struct {
	// `CfnJobDefinition.TimeoutProperty.AttemptDurationSeconds`.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   tmpfsProperty := &tmpfsProperty{
//   	containerPath: jsii.String("containerPath"),
//   	size: jsii.Number(123),
//
//   	// the properties below are optional
//   	mountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   }
//
type CfnJobDefinition_TmpfsProperty struct {
	// `CfnJobDefinition.TmpfsProperty.ContainerPath`.
	ContainerPath *string `json:"containerPath" yaml:"containerPath"`
	// `CfnJobDefinition.TmpfsProperty.Size`.
	Size *float64 `json:"size" yaml:"size"`
	// `CfnJobDefinition.TmpfsProperty.MountOptions`.
	MountOptions *[]*string `json:"mountOptions" yaml:"mountOptions"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   ulimitProperty := &ulimitProperty{
//   	hardLimit: jsii.Number(123),
//   	name: jsii.String("name"),
//   	softLimit: jsii.Number(123),
//   }
//
type CfnJobDefinition_UlimitProperty struct {
	// `CfnJobDefinition.UlimitProperty.HardLimit`.
	HardLimit *float64 `json:"hardLimit" yaml:"hardLimit"`
	// `CfnJobDefinition.UlimitProperty.Name`.
	Name *string `json:"name" yaml:"name"`
	// `CfnJobDefinition.UlimitProperty.SoftLimit`.
	SoftLimit *float64 `json:"softLimit" yaml:"softLimit"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   volumesHostProperty := &volumesHostProperty{
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type CfnJobDefinition_VolumesHostProperty struct {
	// `CfnJobDefinition.VolumesHostProperty.SourcePath`.
	SourcePath *string `json:"sourcePath" yaml:"sourcePath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   volumesProperty := &volumesProperty{
//   	efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   		fileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			accessPointId: jsii.String("accessPointId"),
//   			iam: jsii.String("iam"),
//   		},
//   		rootDirectory: jsii.String("rootDirectory"),
//   		transitEncryption: jsii.String("transitEncryption"),
//   		transitEncryptionPort: jsii.Number(123),
//   	},
//   	host: &volumesHostProperty{
//   		sourcePath: jsii.String("sourcePath"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnJobDefinition_VolumesProperty struct {
	// `CfnJobDefinition.VolumesProperty.EfsVolumeConfiguration`.
	EfsVolumeConfiguration interface{} `json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// `CfnJobDefinition.VolumesProperty.Host`.
	Host interface{} `json:"host" yaml:"host"`
	// `CfnJobDefinition.VolumesProperty.Name`.
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnJobDefinition`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var options interface{}
//   var parameters interface{}
//   var tags interface{}
//   cfnJobDefinitionProps := &cfnJobDefinitionProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	containerProperties: &containerPropertiesProperty{
//   		image: jsii.String("image"),
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&environmentProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   			platformVersion: jsii.String("platformVersion"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		jobRoleArn: jsii.String("jobRoleArn"),
//   		linuxParameters: &linuxParametersProperty{
//   			devices: []interface{}{
//   				&deviceProperty{
//   					containerPath: jsii.String("containerPath"),
//   					hostPath: jsii.String("hostPath"),
//   					permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			initProcessEnabled: jsii.Boolean(false),
//   			maxSwap: jsii.Number(123),
//   			sharedMemorySize: jsii.Number(123),
//   			swappiness: jsii.Number(123),
//   			tmpfs: []interface{}{
//   				&tmpfsProperty{
//   					containerPath: jsii.String("containerPath"),
//   					size: jsii.Number(123),
//
//   					// the properties below are optional
//   					mountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			options: options,
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		memory: jsii.Number(123),
//   		mountPoints: []interface{}{
//   			&mountPointsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				readOnly: jsii.Boolean(false),
//   				sourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		networkConfiguration: &networkConfigurationProperty{
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		privileged: jsii.Boolean(false),
//   		readonlyRootFilesystem: jsii.Boolean(false),
//   		resourceRequirements: []interface{}{
//   			&resourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		secrets: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		ulimits: []interface{}{
//   			&ulimitProperty{
//   				hardLimit: jsii.Number(123),
//   				name: jsii.String("name"),
//   				softLimit: jsii.Number(123),
//   			},
//   		},
//   		user: jsii.String("user"),
//   		vcpus: jsii.Number(123),
//   		volumes: []interface{}{
//   			&volumesProperty{
//   				efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   					fileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					authorizationConfig: &authorizationConfigProperty{
//   						accessPointId: jsii.String("accessPointId"),
//   						iam: jsii.String("iam"),
//   					},
//   					rootDirectory: jsii.String("rootDirectory"),
//   					transitEncryption: jsii.String("transitEncryption"),
//   					transitEncryptionPort: jsii.Number(123),
//   				},
//   				host: &volumesHostProperty{
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	nodeProperties: &nodePropertiesProperty{
//   		mainNode: jsii.Number(123),
//   		nodeRangeProperties: []interface{}{
//   			&nodeRangePropertyProperty{
//   				targetNodes: jsii.String("targetNodes"),
//
//   				// the properties below are optional
//   				container: &containerPropertiesProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					environment: []interface{}{
//   						&environmentProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					executionRoleArn: jsii.String("executionRoleArn"),
//   					fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   						platformVersion: jsii.String("platformVersion"),
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					jobRoleArn: jsii.String("jobRoleArn"),
//   					linuxParameters: &linuxParametersProperty{
//   						devices: []interface{}{
//   							&deviceProperty{
//   								containerPath: jsii.String("containerPath"),
//   								hostPath: jsii.String("hostPath"),
//   								permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						initProcessEnabled: jsii.Boolean(false),
//   						maxSwap: jsii.Number(123),
//   						sharedMemorySize: jsii.Number(123),
//   						swappiness: jsii.Number(123),
//   						tmpfs: []interface{}{
//   							&tmpfsProperty{
//   								containerPath: jsii.String("containerPath"),
//   								size: jsii.Number(123),
//
//   								// the properties below are optional
//   								mountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					logConfiguration: &logConfigurationProperty{
//   						logDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						options: options,
//   						secretOptions: []interface{}{
//   							&secretProperty{
//   								name: jsii.String("name"),
//   								valueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					memory: jsii.Number(123),
//   					mountPoints: []interface{}{
//   						&mountPointsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							readOnly: jsii.Boolean(false),
//   							sourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					networkConfiguration: &networkConfigurationProperty{
//   						assignPublicIp: jsii.String("assignPublicIp"),
//   					},
//   					privileged: jsii.Boolean(false),
//   					readonlyRootFilesystem: jsii.Boolean(false),
//   					resourceRequirements: []interface{}{
//   						&resourceRequirementProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					secrets: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					ulimits: []interface{}{
//   						&ulimitProperty{
//   							hardLimit: jsii.Number(123),
//   							name: jsii.String("name"),
//   							softLimit: jsii.Number(123),
//   						},
//   					},
//   					user: jsii.String("user"),
//   					vcpus: jsii.Number(123),
//   					volumes: []interface{}{
//   						&volumesProperty{
//   							efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   								fileSystemId: jsii.String("fileSystemId"),
//
//   								// the properties below are optional
//   								authorizationConfig: &authorizationConfigProperty{
//   									accessPointId: jsii.String("accessPointId"),
//   									iam: jsii.String("iam"),
//   								},
//   								rootDirectory: jsii.String("rootDirectory"),
//   								transitEncryption: jsii.String("transitEncryption"),
//   								transitEncryptionPort: jsii.Number(123),
//   							},
//   							host: &volumesHostProperty{
//   								sourcePath: jsii.String("sourcePath"),
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		numNodes: jsii.Number(123),
//   	},
//   	parameters: parameters,
//   	platformCapabilities: []*string{
//   		jsii.String("platformCapabilities"),
//   	},
//   	propagateTags: jsii.Boolean(false),
//   	retryStrategy: &retryStrategyProperty{
//   		attempts: jsii.Number(123),
//   		evaluateOnExit: []interface{}{
//   			&evaluateOnExitProperty{
//   				action: jsii.String("action"),
//
//   				// the properties below are optional
//   				onExitCode: jsii.String("onExitCode"),
//   				onReason: jsii.String("onReason"),
//   				onStatusReason: jsii.String("onStatusReason"),
//   			},
//   		},
//   	},
//   	schedulingPriority: jsii.Number(123),
//   	tags: tags,
//   	timeout: &timeoutProperty{
//   		attemptDurationSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnJobDefinitionProps struct {
	// `AWS::Batch::JobDefinition.Type`.
	Type *string `json:"type" yaml:"type"`
	// `AWS::Batch::JobDefinition.ContainerProperties`.
	ContainerProperties interface{} `json:"containerProperties" yaml:"containerProperties"`
	// `AWS::Batch::JobDefinition.JobDefinitionName`.
	JobDefinitionName *string `json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// `AWS::Batch::JobDefinition.NodeProperties`.
	NodeProperties interface{} `json:"nodeProperties" yaml:"nodeProperties"`
	// `AWS::Batch::JobDefinition.Parameters`.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// `AWS::Batch::JobDefinition.PlatformCapabilities`.
	PlatformCapabilities *[]*string `json:"platformCapabilities" yaml:"platformCapabilities"`
	// `AWS::Batch::JobDefinition.PropagateTags`.
	PropagateTags interface{} `json:"propagateTags" yaml:"propagateTags"`
	// `AWS::Batch::JobDefinition.RetryStrategy`.
	RetryStrategy interface{} `json:"retryStrategy" yaml:"retryStrategy"`
	// `AWS::Batch::JobDefinition.SchedulingPriority`.
	SchedulingPriority *float64 `json:"schedulingPriority" yaml:"schedulingPriority"`
	// `AWS::Batch::JobDefinition.Tags`.
	Tags interface{} `json:"tags" yaml:"tags"`
	// `AWS::Batch::JobDefinition.Timeout`.
	Timeout interface{} `json:"timeout" yaml:"timeout"`
}

// A CloudFormation `AWS::Batch::JobQueue`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnJobQueue := batch.NewCfnJobQueue(this, jsii.String("MyCfnJobQueue"), &cfnJobQueueProps{
//   	computeEnvironmentOrder: []interface{}{
//   		&computeEnvironmentOrderProperty{
//   			computeEnvironment: jsii.String("computeEnvironment"),
//   			order: jsii.Number(123),
//   		},
//   	},
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	jobQueueName: jsii.String("jobQueueName"),
//   	schedulingPolicyArn: jsii.String("schedulingPolicyArn"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnJobQueue interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrJobQueueArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::Batch::JobQueue.ComputeEnvironmentOrder`.
	ComputeEnvironmentOrder() interface{}
	SetComputeEnvironmentOrder(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Batch::JobQueue.JobQueueName`.
	JobQueueName() *string
	SetJobQueueName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::Batch::JobQueue.Priority`.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::Batch::JobQueue.SchedulingPolicyArn`.
	SchedulingPolicyArn() *string
	SetSchedulingPolicyArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Batch::JobQueue.State`.
	State() *string
	SetState(val *string)
	// `AWS::Batch::JobQueue.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnJobQueue
type jsiiProxy_CfnJobQueue struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobQueue) AttrJobQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrJobQueueArn",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnJobQueue) SchedulingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyArn",
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

func (j *jsiiProxy_CfnJobQueue) SetSchedulingPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"schedulingPolicyArn",
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

func (c *jsiiProxy_CfnJobQueue) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJobQueue) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJobQueue) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJobQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJobQueue) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJobQueue) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJobQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnJobQueue) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJobQueue) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnJobQueue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnJobQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnJobQueue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnJobQueue) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   computeEnvironmentOrderProperty := &computeEnvironmentOrderProperty{
//   	computeEnvironment: jsii.String("computeEnvironment"),
//   	order: jsii.Number(123),
//   }
//
type CfnJobQueue_ComputeEnvironmentOrderProperty struct {
	// `CfnJobQueue.ComputeEnvironmentOrderProperty.ComputeEnvironment`.
	ComputeEnvironment *string `json:"computeEnvironment" yaml:"computeEnvironment"`
	// `CfnJobQueue.ComputeEnvironmentOrderProperty.Order`.
	Order *float64 `json:"order" yaml:"order"`
}

// Properties for defining a `CfnJobQueue`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnJobQueueProps := &cfnJobQueueProps{
//   	computeEnvironmentOrder: []interface{}{
//   		&computeEnvironmentOrderProperty{
//   			computeEnvironment: jsii.String("computeEnvironment"),
//   			order: jsii.Number(123),
//   		},
//   	},
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	jobQueueName: jsii.String("jobQueueName"),
//   	schedulingPolicyArn: jsii.String("schedulingPolicyArn"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnJobQueueProps struct {
	// `AWS::Batch::JobQueue.ComputeEnvironmentOrder`.
	ComputeEnvironmentOrder interface{} `json:"computeEnvironmentOrder" yaml:"computeEnvironmentOrder"`
	// `AWS::Batch::JobQueue.Priority`.
	Priority *float64 `json:"priority" yaml:"priority"`
	// `AWS::Batch::JobQueue.JobQueueName`.
	JobQueueName *string `json:"jobQueueName" yaml:"jobQueueName"`
	// `AWS::Batch::JobQueue.SchedulingPolicyArn`.
	SchedulingPolicyArn *string `json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
	// `AWS::Batch::JobQueue.State`.
	State *string `json:"state" yaml:"state"`
	// `AWS::Batch::JobQueue.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Batch::SchedulingPolicy`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnSchedulingPolicy := batch.NewCfnSchedulingPolicy(this, jsii.String("MyCfnSchedulingPolicy"), &cfnSchedulingPolicyProps{
//   	fairsharePolicy: &fairsharePolicyProperty{
//   		computeReservation: jsii.Number(123),
//   		shareDecaySeconds: jsii.Number(123),
//   		shareDistribution: []interface{}{
//   			&shareAttributesProperty{
//   				shareIdentifier: jsii.String("shareIdentifier"),
//   				weightFactor: jsii.Number(123),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnSchedulingPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Batch::SchedulingPolicy.FairsharePolicy`.
	FairsharePolicy() interface{}
	SetFairsharePolicy(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// `AWS::Batch::SchedulingPolicy.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::Batch::SchedulingPolicy.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSchedulingPolicy
type jsiiProxy_CfnSchedulingPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchedulingPolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) FairsharePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fairsharePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Batch::SchedulingPolicy`.
func NewCfnSchedulingPolicy(scope awscdk.Construct, id *string, props *CfnSchedulingPolicyProps) CfnSchedulingPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnSchedulingPolicy{}

	_jsii_.Create(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::SchedulingPolicy`.
func NewCfnSchedulingPolicy_Override(c CfnSchedulingPolicy, scope awscdk.Construct, id *string, props *CfnSchedulingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchedulingPolicy) SetFairsharePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"fairsharePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnSchedulingPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
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
func CfnSchedulingPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSchedulingPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSchedulingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchedulingPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_batch.CfnSchedulingPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSchedulingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   fairsharePolicyProperty := &fairsharePolicyProperty{
//   	computeReservation: jsii.Number(123),
//   	shareDecaySeconds: jsii.Number(123),
//   	shareDistribution: []interface{}{
//   		&shareAttributesProperty{
//   			shareIdentifier: jsii.String("shareIdentifier"),
//   			weightFactor: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnSchedulingPolicy_FairsharePolicyProperty struct {
	// `CfnSchedulingPolicy.FairsharePolicyProperty.ComputeReservation`.
	ComputeReservation *float64 `json:"computeReservation" yaml:"computeReservation"`
	// `CfnSchedulingPolicy.FairsharePolicyProperty.ShareDecaySeconds`.
	ShareDecaySeconds *float64 `json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// `CfnSchedulingPolicy.FairsharePolicyProperty.ShareDistribution`.
	ShareDistribution interface{} `json:"shareDistribution" yaml:"shareDistribution"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   shareAttributesProperty := &shareAttributesProperty{
//   	shareIdentifier: jsii.String("shareIdentifier"),
//   	weightFactor: jsii.Number(123),
//   }
//
type CfnSchedulingPolicy_ShareAttributesProperty struct {
	// `CfnSchedulingPolicy.ShareAttributesProperty.ShareIdentifier`.
	ShareIdentifier *string `json:"shareIdentifier" yaml:"shareIdentifier"`
	// `CfnSchedulingPolicy.ShareAttributesProperty.WeightFactor`.
	WeightFactor *float64 `json:"weightFactor" yaml:"weightFactor"`
}

// Properties for defining a `CfnSchedulingPolicy`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   cfnSchedulingPolicyProps := &cfnSchedulingPolicyProps{
//   	fairsharePolicy: &fairsharePolicyProperty{
//   		computeReservation: jsii.Number(123),
//   		shareDecaySeconds: jsii.Number(123),
//   		shareDistribution: []interface{}{
//   			&shareAttributesProperty{
//   				shareIdentifier: jsii.String("shareIdentifier"),
//   				weightFactor: jsii.Number(123),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSchedulingPolicyProps struct {
	// `AWS::Batch::SchedulingPolicy.FairsharePolicy`.
	FairsharePolicy interface{} `json:"fairsharePolicy" yaml:"fairsharePolicy"`
	// `AWS::Batch::SchedulingPolicy.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::Batch::SchedulingPolicy.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Batch Compute Environment.
//
// Defines a batch compute environment to run batch jobs on.
//
// Example:
//   var vpc vpc
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
//   			generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type ComputeEnvironment interface {
	awscdk.Resource
	IComputeEnvironment
	// The ARN of this compute environment.
	// Experimental.
	ComputeEnvironmentArn() *string
	// The name of this compute environment.
	// Experimental.
	ComputeEnvironmentName() *string
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
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

func (c *jsiiProxy_ComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (c *jsiiProxy_ComputeEnvironment) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeEnvironment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_ComputeEnvironment) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeEnvironment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   var vpc vpc
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
//   			generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type ComputeEnvironmentProps struct {
	// A name for the compute environment.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	ComputeEnvironmentName *string `json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// The details of the required compute resources for the managed compute environment.
	//
	// If specified, and this is an unmanaged compute environment, will throw an error.
	//
	// By default, AWS Batch managed compute environments use a recent, approved version of the
	// Amazon ECS-optimized AMI for compute resources.
	// Experimental.
	ComputeResources *ComputeResources `json:"computeResources" yaml:"computeResources"`
	// The state of the compute environment.
	//
	// If the state is set to true, then the compute
	// environment accepts jobs from a queue and can scale out automatically based on queues.
	// Experimental.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// Determines if AWS should manage the allocation of compute resources for processing jobs.
	//
	// If set to false, then you are in charge of providing the compute resource details.
	// Experimental.
	Managed *bool `json:"managed" yaml:"managed"`
	// The IAM role used by Batch to make calls to other AWS services on your behalf for managing the resources that you use with the service.
	//
	// By default, this role is created for you using
	// the AWS managed service policy for Batch.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole" yaml:"serviceRole"`
}

// Property to specify if the compute environment uses On-Demand, SpotFleet, Fargate, or Fargate Spot compute resources.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//
//   spotEnvironment := batch.NewComputeEnvironment(this, jsii.String("MySpotEnvironment"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		type: batch.computeResourceType_SPOT,
//   		bidPercentage: jsii.Number(75),
//   		 // Bids for resources at 75% of the on-demand price
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type ComputeResourceType string

const (
	// Resources will be EC2 On-Demand resources.
	// Experimental.
	ComputeResourceType_ON_DEMAND ComputeResourceType = "ON_DEMAND"
	// Resources will be EC2 SpotFleet resources.
	// Experimental.
	ComputeResourceType_SPOT ComputeResourceType = "SPOT"
	// Resources will be Fargate resources.
	// Experimental.
	ComputeResourceType_FARGATE ComputeResourceType = "FARGATE"
	// Resources will be Fargate Spot resources.
	//
	// Fargate Spot uses spare capacity in the AWS cloud to run your fault-tolerant,
	// time-flexible jobs at up to a 70% discount. If AWS needs the resources back,
	// jobs running on Fargate Spot will be interrupted with two minutes of notification.
	// Experimental.
	ComputeResourceType_FARGATE_SPOT ComputeResourceType = "FARGATE_SPOT"
)

// Properties for defining the structure of the batch compute cluster.
//
// Example:
//   var vpc vpc
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
//   			generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type ComputeResources struct {
	// The VPC network that all compute resources will be connected to.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
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
	AllocationStrategy AllocationStrategy `json:"allocationStrategy" yaml:"allocationStrategy"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for
	// that instance type before instances are launched. For example, if your maximum percentage is 20%,
	// then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. You always
	// pay the lowest (market) price and never more than your maximum percentage. If you leave this field empty,
	// the default value is 100% of the On-Demand price.
	// Experimental.
	BidPercentage *float64 `json:"bidPercentage" yaml:"bidPercentage"`
	// Key-value pair tags to be applied to resources that are launched in the compute environment.
	//
	// For AWS Batch, these take the form of "String1": "String2", where String1 is the tag key and
	// String2 is the tag value—for example, { "Name": "AWS Batch Instance - C4OnDemand" }.
	// Experimental.
	ComputeResourcesTags *map[string]*string `json:"computeResourcesTags" yaml:"computeResourcesTags"`
	// The desired number of EC2 vCPUS in the compute environment.
	// Experimental.
	DesiredvCpus *float64 `json:"desiredvCpus" yaml:"desiredvCpus"`
	// The EC2 key pair that is used for instances launched in the compute environment.
	//
	// If no key is defined, then SSH access is not allowed to provisioned compute resources.
	// Experimental.
	Ec2KeyPair *string `json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
	// Experimental.
	Image awsec2.IMachineImage `json:"image" yaml:"image"`
	// The Amazon ECS instance profile applied to Amazon EC2 instances in a compute environment.
	//
	// You can specify
	// the short name or full Amazon Resource Name (ARN) of an instance profile. For example, ecsInstanceRole or
	// arn:aws:iam::<aws_account_id>:instance-profile/ecsInstanceRole . For more information, see Amazon ECS
	// Instance Role in the AWS Batch User Guide.
	// Experimental.
	InstanceRole *string `json:"instanceRole" yaml:"instanceRole"`
	// The types of EC2 instances that may be launched in the compute environment.
	//
	// You can specify instance
	// families to launch any instance type within those families (for example, c4 or p3), or you can specify
	// specific sizes within a family (such as c4.8xlarge). You can also choose optimal to pick instance types
	// (from the C, M, and R instance families) on the fly that match the demand of your job queues.
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `json:"instanceTypes" yaml:"instanceTypes"`
	// An optional launch template to associate with your compute resources.
	//
	// For more information, see README file.
	// Experimental.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of EC2 vCPUs that an environment can reach.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU.
	// Experimental.
	MaxvCpus *float64 `json:"maxvCpus" yaml:"maxvCpus"`
	// The minimum number of EC2 vCPUs that an environment should maintain (even if the compute environment state is DISABLED).
	//
	// Each vCPU is equivalent to 1,024 CPU shares. By keeping this set to 0 you will not have instance time wasted when
	// there is no work to be run. If you set this above zero you will maintain that number of vCPUs at all times.
	// Experimental.
	MinvCpus *float64 `json:"minvCpus" yaml:"minvCpus"`
	// The Amazon EC2 placement group to associate with your compute resources.
	// Experimental.
	PlacementGroup *string `json:"placementGroup" yaml:"placementGroup"`
	// The EC2 security group(s) associated with instances launched in the compute environment.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// This property will be ignored if you set the environment type to ON_DEMAND.
	//
	// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment.
	// For more information, see Amazon EC2 Spot Fleet Role in the AWS Batch User Guide.
	// Experimental.
	SpotFleetRole awsiam.IRole `json:"spotFleetRole" yaml:"spotFleetRole"`
	// The type of compute environment: ON_DEMAND, SPOT, FARGATE, or FARGATE_SPOT.
	// Experimental.
	Type ComputeResourceType `json:"type" yaml:"type"`
	// The VPC subnets into which the compute resources are launched.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Exposed secret for log configuration.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ExposedSecret interface {
	// Name of the option.
	// Experimental.
	OptionName() *string
	// Experimental.
	SetOptionName(val *string)
	// ARN of the secret option.
	// Experimental.
	SecretArn() *string
	// Experimental.
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
	// Experimental.
	SetCount(c *float64)
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	// Experimental.
	MainNode() *float64
	// Experimental.
	SetMainNode(m *float64)
	// A list of node ranges and their properties associated with a multi-node parallel job.
	// Experimental.
	RangeProps() *[]INodeRangeProps
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
	// Experimental.
	SetContainer(c *JobDefinitionContainer)
	// The minimum node index value to apply this container definition against.
	//
	// You may nest node ranges, for example 0:10 and 4:5, in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	FromNodeIndex() *float64
	// Experimental.
	SetFromNodeIndex(f *float64)
	// The maximum node index value to apply this container definition against. If omitted, the highest value is used relative.
	//
	// to the number of nodes associated with the job. You may nest node ranges, for example 0:10 and 4:5,
	// in which case the 4:5 range properties override the 0:10 properties.
	// Experimental.
	ToNodeIndex() *float64
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
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   repo := ecr.repository.fromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))
//
//   batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.NewEcrImage(repo, jsii.String("latest")),
//   	},
//   })
//
// Experimental.
type JobDefinition interface {
	awscdk.Resource
	IJobDefinition
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
	// The ARN of this batch job definition.
	// Experimental.
	JobDefinitionArn() *string
	// The name of the batch job definition.
	// Experimental.
	JobDefinitionName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
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

func (j *jsiiProxy_JobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (j *jsiiProxy_JobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (j *jsiiProxy_JobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type JobDefinitionContainer struct {
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `json:"image" yaml:"image"`
	// Whether or not to assign a public IP to the job.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp" yaml:"assignPublicIp"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command" yaml:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The IAM role that AWS Batch can assume.
	//
	// Required when using Fargate.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole" yaml:"executionRole"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all
	// containers in a job should not exceed the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `json:"gpuCount" yaml:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// Currently all node groups in a
	// multi-node parallel job must use the same instance type. This parameter is not valid
	// for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The IAM role that the container can assume for AWS permissions.
	// Experimental.
	JobRole awsiam.IRole `json:"jobRole" yaml:"jobRole"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	//
	// For now, only the `devices` property is supported.
	// Experimental.
	LinuxParams awsecs.LinuxParameters `json:"linuxParams" yaml:"linuxParams"`
	// The log configuration specification for the container.
	// Experimental.
	LogConfiguration *LogConfiguration `json:"logConfiguration" yaml:"logConfiguration"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed
	// the memory specified here, the container is killed. You must specify at least 4 MiB of memory for EC2 and 512 MiB for Fargate.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The mount points for data volumes in your container.
	// Experimental.
	MountPoints *[]*awsecs.MountPoint `json:"mountPoints" yaml:"mountPoints"`
	// Fargate platform version.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion" yaml:"platformVersion"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	// Experimental.
	Privileged *bool `json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	// Experimental.
	ReadOnly *bool `json:"readOnly" yaml:"readOnly"`
	// A list of ulimits to set in the container.
	// Experimental.
	Ulimits *[]*awsecs.Ulimit `json:"ulimits" yaml:"ulimits"`
	// The user name to use inside the container.
	// Experimental.
	User *string `json:"user" yaml:"user"`
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU for EC2 and 0.25 for Fargate.
	// Experimental.
	Vcpus *float64 `json:"vcpus" yaml:"vcpus"`
	// A list of data volumes used in a job.
	// Experimental.
	Volumes *[]*awsecs.Volume `json:"volumes" yaml:"volumes"`
}

// Construction properties of the {@link JobDefinition} construct.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type JobDefinitionProps struct {
	// An object with various properties specific to container-based jobs.
	// Experimental.
	Container *JobDefinitionContainer `json:"container" yaml:"container"`
	// The name of the job definition.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobDefinitionName *string `json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// An object with various properties specific to multi-node parallel jobs.
	// Experimental.
	NodeProps IMultiNodeProps `json:"nodeProps" yaml:"nodeProps"`
	// When you submit a job, you can specify parameters that should replace the placeholders or override the default job definition parameters.
	//
	// Parameters
	// in job submission requests take precedence over the defaults in a job definition.
	// This allows you to use the same job definition for multiple jobs that use the same
	// format, and programmatically change values in the command at submission time.
	// Experimental.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The platform capabilities required by the job definition.
	// Experimental.
	PlatformCapabilities *[]PlatformCapabilities `json:"platformCapabilities" yaml:"platformCapabilities"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and
	// 10 attempts. If the value of attempts is greater than one, the job is retried on failure
	// the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout configuration for jobs that are submitted with this job definition.
	//
	// You can specify
	// a timeout duration after which AWS Batch terminates your jobs if they have not finished.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
}

// Batch Job Queue.
//
// Defines a batch job queue to define how submitted batch jobs
// should be ran based on specified batch compute environments.
//
// Example:
//   var sharedComputeEnvs computeEnvironment
//   highPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
//   	computeEnvironments: []jobQueueComputeEnvironment{
//   		&jobQueueComputeEnvironment{
//   			computeEnvironment: sharedComputeEnvs,
//   			order: jsii.Number(1),
//   		},
//   	},
//   	priority: jsii.Number(2),
//   })
//
//   lowPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
//   	computeEnvironments: []*jobQueueComputeEnvironment{
//   		&jobQueueComputeEnvironment{
//   			computeEnvironment: sharedComputeEnvs,
//   			order: jsii.Number(1),
//   		},
//   	},
//   	priority: jsii.Number(1),
//   })
//
// Experimental.
type JobQueue interface {
	awscdk.Resource
	IJobQueue
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
	// The ARN of this batch job queue.
	// Experimental.
	JobQueueArn() *string
	// A name for the job queue.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobQueueName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
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

func (j *jsiiProxy_JobQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (j *jsiiProxy_JobQueue) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobQueue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (j *jsiiProxy_JobQueue) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobQueue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//
//   var computeEnvironment computeEnvironment
//   jobQueueComputeEnvironment := &jobQueueComputeEnvironment{
//   	computeEnvironment: computeEnvironment,
//   	order: jsii.Number(123),
//   }
//
// Experimental.
type JobQueueComputeEnvironment struct {
	// The batch compute environment to use for processing submitted jobs to this queue.
	// Experimental.
	ComputeEnvironment IComputeEnvironment `json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order in which this compute environment will be selected for dynamic allocation of resources to process submitted jobs.
	// Experimental.
	Order *float64 `json:"order" yaml:"order"`
}

// Properties of a batch job queue.
//
// Example:
//   var sharedComputeEnvs computeEnvironment
//   highPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
//   	computeEnvironments: []jobQueueComputeEnvironment{
//   		&jobQueueComputeEnvironment{
//   			computeEnvironment: sharedComputeEnvs,
//   			order: jsii.Number(1),
//   		},
//   	},
//   	priority: jsii.Number(2),
//   })
//
//   lowPrioQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &jobQueueProps{
//   	computeEnvironments: []*jobQueueComputeEnvironment{
//   		&jobQueueComputeEnvironment{
//   			computeEnvironment: sharedComputeEnvs,
//   			order: jsii.Number(1),
//   		},
//   	},
//   	priority: jsii.Number(1),
//   })
//
// Experimental.
type JobQueueProps struct {
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to
	// determine which compute environment should execute a given job. Compute environments must be in the VALID state before you can associate them
	// with a job queue. You can associate up to three compute environments with a job queue.
	// Experimental.
	ComputeEnvironments *[]*JobQueueComputeEnvironment `json:"computeEnvironments" yaml:"computeEnvironments"`
	// The state of the job queue.
	//
	// If set to true, it is able to accept jobs.
	// Experimental.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// A name for the job queue.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobQueueName *string `json:"jobQueueName" yaml:"jobQueueName"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority (or a higher integer value for the priority parameter) are evaluated first
	// when associated with the same compute environment. Priority is determined in descending order, for example, a job queue with a priority value
	// of 10 is given scheduling preference over a job queue with a priority value of 1.
	// Experimental.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Launch template property specification.
//
// Example:
//   var vpc vpc
//   var myLaunchTemplate cfnLaunchTemplate
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		launchTemplate: &launchTemplateSpecification{
//   			launchTemplateName: string(myLaunchTemplate.launchTemplateName),
//   		},
//   		vpc: vpc,
//   	},
//   	computeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
//   })
//
// Experimental.
type LaunchTemplateSpecification struct {
	// The Launch template name.
	// Experimental.
	LaunchTemplateName *string `json:"launchTemplateName" yaml:"launchTemplateName"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Log configuration options to send to a custom log driver for the container.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type LogConfiguration struct {
	// The log driver to use for the container.
	// Experimental.
	LogDriver LogDriver `json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	// Experimental.
	Options interface{} `json:"options" yaml:"options"`
	// The secrets to pass to the log configuration as options.
	//
	// For more information, see https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data-secrets.html#secrets-logconfig
	// Experimental.
	SecretOptions *[]ExposedSecret `json:"secretOptions" yaml:"secretOptions"`
}

// The log driver to use for the container.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type LogDriver string

const (
	// Specifies the Amazon CloudWatch Logs logging driver.
	// Experimental.
	LogDriver_AWSLOGS LogDriver = "AWSLOGS"
	// Specifies the Fluentd logging driver.
	// Experimental.
	LogDriver_FLUENTD LogDriver = "FLUENTD"
	// Specifies the Graylog Extended Format (GELF) logging driver.
	// Experimental.
	LogDriver_GELF LogDriver = "GELF"
	// Specifies the journald logging driver.
	// Experimental.
	LogDriver_JOURNALD LogDriver = "JOURNALD"
	// Specifies the logentries logging driver.
	// Experimental.
	LogDriver_LOGENTRIES LogDriver = "LOGENTRIES"
	// Specifies the JSON file logging driver.
	// Experimental.
	LogDriver_JSON_FILE LogDriver = "JSON_FILE"
	// Specifies the Splunk logging driver.
	// Experimental.
	LogDriver_SPLUNK LogDriver = "SPLUNK"
	// Specifies the syslog logging driver.
	// Experimental.
	LogDriver_SYSLOG LogDriver = "SYSLOG"
)

// Platform capabilities.
// Experimental.
type PlatformCapabilities string

const (
	// Specifies EC2 environment.
	// Experimental.
	PlatformCapabilities_EC2 PlatformCapabilities = "EC2"
	// Specifies Fargate environment.
	// Experimental.
	PlatformCapabilities_FARGATE PlatformCapabilities = "FARGATE"
)

