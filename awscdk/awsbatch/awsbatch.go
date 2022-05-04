package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   		updateToLatestImageVersion: jsii.Boolean(false),
//   	},
//   	replaceComputeEnvironment: jsii.Boolean(false),
//   	serviceRole: jsii.String("serviceRole"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	unmanagedvCpus: jsii.Number(123),
//   	updatePolicy: &updatePolicyProperty{
//   		jobExecutionTimeoutMinutes: jsii.Number(123),
//   		terminateJobsOnUpdate: jsii.Boolean(false),
//   	},
//   })
//
type CfnComputeEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrComputeEnvironmentArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Batch::ComputeEnvironment.ReplaceComputeEnvironment`.
	ReplaceComputeEnvironment() interface{}
	SetReplaceComputeEnvironment(val interface{})
	// `AWS::Batch::ComputeEnvironment.ServiceRole`.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
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
	UpdatedProperites() *map[string]interface{}
	// `AWS::Batch::ComputeEnvironment.UpdatePolicy`.
	UpdatePolicy() interface{}
	SetUpdatePolicy(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnComputeEnvironment) ReplaceComputeEnvironment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceComputeEnvironment",
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

func (j *jsiiProxy_CfnComputeEnvironment) UpdatePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatePolicy",
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

func (j *jsiiProxy_CfnComputeEnvironment) SetReplaceComputeEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"replaceComputeEnvironment",
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

func (j *jsiiProxy_CfnComputeEnvironment) SetUpdatePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"updatePolicy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
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
// Deprecated: use `x instanceof Construct` instead.
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
//   	updateToLatestImageVersion: jsii.Boolean(false),
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
	// `CfnComputeEnvironment.ComputeResourcesProperty.UpdateToLatestImageVersion`.
	UpdateToLatestImageVersion interface{} `json:"updateToLatestImageVersion" yaml:"updateToLatestImageVersion"`
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

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import batch "github.com/aws/aws-cdk-go/awscdk/aws_batch"
//   updatePolicyProperty := &updatePolicyProperty{
//   	jobExecutionTimeoutMinutes: jsii.Number(123),
//   	terminateJobsOnUpdate: jsii.Boolean(false),
//   }
//
type CfnComputeEnvironment_UpdatePolicyProperty struct {
	// `CfnComputeEnvironment.UpdatePolicyProperty.JobExecutionTimeoutMinutes`.
	JobExecutionTimeoutMinutes *float64 `json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// `CfnComputeEnvironment.UpdatePolicyProperty.TerminateJobsOnUpdate`.
	TerminateJobsOnUpdate interface{} `json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
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
//   		updateToLatestImageVersion: jsii.Boolean(false),
//   	},
//   	replaceComputeEnvironment: jsii.Boolean(false),
//   	serviceRole: jsii.String("serviceRole"),
//   	state: jsii.String("state"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	unmanagedvCpus: jsii.Number(123),
//   	updatePolicy: &updatePolicyProperty{
//   		jobExecutionTimeoutMinutes: jsii.Number(123),
//   		terminateJobsOnUpdate: jsii.Boolean(false),
//   	},
//   }
//
type CfnComputeEnvironmentProps struct {
	// `AWS::Batch::ComputeEnvironment.Type`.
	Type *string `json:"type" yaml:"type"`
	// `AWS::Batch::ComputeEnvironment.ComputeEnvironmentName`.
	ComputeEnvironmentName *string `json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// `AWS::Batch::ComputeEnvironment.ComputeResources`.
	ComputeResources interface{} `json:"computeResources" yaml:"computeResources"`
	// `AWS::Batch::ComputeEnvironment.ReplaceComputeEnvironment`.
	ReplaceComputeEnvironment interface{} `json:"replaceComputeEnvironment" yaml:"replaceComputeEnvironment"`
	// `AWS::Batch::ComputeEnvironment.ServiceRole`.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// `AWS::Batch::ComputeEnvironment.State`.
	State *string `json:"state" yaml:"state"`
	// `AWS::Batch::ComputeEnvironment.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Batch::ComputeEnvironment.UnmanagedvCpus`.
	UnmanagedvCpus *float64 `json:"unmanagedvCpus" yaml:"unmanagedvCpus"`
	// `AWS::Batch::ComputeEnvironment.UpdatePolicy`.
	UpdatePolicy interface{} `json:"updatePolicy" yaml:"updatePolicy"`
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
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Batch::JobDefinition.ContainerProperties`.
	ContainerProperties() interface{}
	SetContainerProperties(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
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
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Batch::JobQueue.ComputeEnvironmentOrder`.
	ComputeEnvironmentOrder() interface{}
	SetComputeEnvironmentOrder(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// `AWS::Batch::JobQueue.Priority`.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Batch::JobQueue.SchedulingPolicyArn`.
	SchedulingPolicyArn() *string
	SetSchedulingPolicyArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
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
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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
// Deprecated: use `x instanceof Construct` instead.
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
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// `AWS::Batch::SchedulingPolicy.Name`.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Batch::SchedulingPolicy.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnSchedulingPolicy) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnSchedulingPolicy(scope constructs.Construct, id *string, props *CfnSchedulingPolicyProps) CfnSchedulingPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnSchedulingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Batch::SchedulingPolicy`.
func NewCfnSchedulingPolicy_Override(c CfnSchedulingPolicy, scope constructs.Construct, id *string, props *CfnSchedulingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
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
func CfnSchedulingPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchedulingPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnSchedulingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
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
		"aws-cdk-lib.aws_batch.CfnSchedulingPolicy",
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

func (c *jsiiProxy_CfnSchedulingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

