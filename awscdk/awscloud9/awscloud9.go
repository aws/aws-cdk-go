package awscloud9

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloud9/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Cloud9::EnvironmentEC2`.
//
// The `AWS::Cloud9::EnvironmentEC2` resource creates an Amazon EC2 development environment in AWS Cloud9 . For more information, see [Creating an Environment](https://docs.aws.amazon.com/cloud9/latest/user-guide/create-environment.html) in the *AWS Cloud9 User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud9 "github.com/aws/aws-cdk-go/awscdk/aws_cloud9"
//   cfnEnvironmentEC2 := cloud9.NewCfnEnvironmentEC2(this, jsii.String("MyCfnEnvironmentEC2"), &cfnEnvironmentEC2Props{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	automaticStopTimeMinutes: jsii.Number(123),
//   	connectionType: jsii.String("connectionType"),
//   	description: jsii.String("description"),
//   	imageId: jsii.String("imageId"),
//   	name: jsii.String("name"),
//   	ownerArn: jsii.String("ownerArn"),
//   	repositories: []interface{}{
//   		&repositoryProperty{
//   			pathComponent: jsii.String("pathComponent"),
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnEnvironmentEC2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the development environment, such as `arn:aws:cloud9:us-east-2:123456789012:environment:2bc3642873c342e485f7e0c561234567` .
	AttrArn() *string
	// The name of the environment.
	AttrName() *string
	// The number of minutes until the running instance is shut down after the environment was last used.
	AutomaticStopTimeMinutes() *float64
	SetAutomaticStopTimeMinutes(val *float64)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are `CONNECT_SSH` (default) and `CONNECT_SSM` (connected through AWS Systems Manager ).
	ConnectionType() *string
	SetConnectionType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The description of the environment to create.
	Description() *string
	SetDescription(val *string)
	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
	//
	// To choose an AMI for the instance, you must specify a valid AMI alias or a valid AWS Systems Manager path.
	//
	// The default AMI is used if the parameter isn't explicitly assigned a value in the request.
	//
	// *AMI aliases*
	//
	// - *Amazon Linux (default): `amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `ubuntu-18.04-x86_64`
	//
	// *SSM paths*
	//
	// - *Amazon Linux (default): `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
	ImageId() *string
	SetImageId(val *string)
	// The type of instance to connect to the environment (for example, `t2.micro` ).
	InstanceType() *string
	SetInstanceType(val *string)
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
	// The name of the environment.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Amazon Resource Name (ARN) of the environment owner.
	//
	// This ARN can be the ARN of any AWS Identity and Access Management principal. If this value is not specified, the ARN defaults to this environment's creator.
	OwnerArn() *string
	SetOwnerArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Any AWS CodeCommit source code repositories to be cloned into the development environment.
	Repositories() interface{}
	SetRepositories(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The ID of the subnet in Amazon Virtual Private Cloud (Amazon VPC) that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	SubnetId() *string
	SetSubnetId(val *string)
	// An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.
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

// The jsii proxy struct for CfnEnvironmentEC2
type jsiiProxy_CfnEnvironmentEC2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEnvironmentEC2) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) AutomaticStopTimeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticStopTimeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) OwnerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Repositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentEC2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2(scope awscdk.Construct, id *string, props *CfnEnvironmentEC2Props) CfnEnvironmentEC2 {
	_init_.Initialize()

	j := jsiiProxy_CfnEnvironmentEC2{}

	_jsii_.Create(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cloud9::EnvironmentEC2`.
func NewCfnEnvironmentEC2_Override(c CfnEnvironmentEC2, scope awscdk.Construct, id *string, props *CfnEnvironmentEC2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetAutomaticStopTimeMinutes(val *float64) {
	_jsii_.Set(
		j,
		"automaticStopTimeMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetOwnerArn(val *string) {
	_jsii_.Set(
		j,
		"ownerArn",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetRepositories(val interface{}) {
	_jsii_.Set(
		j,
		"repositories",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironmentEC2) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
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
func CfnEnvironmentEC2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEnvironmentEC2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEnvironmentEC2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentEC2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cloud9.CfnEnvironmentEC2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEnvironmentEC2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironmentEC2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Repository` property type specifies an AWS CodeCommit source code repository to be cloned into an AWS Cloud9 development environment.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud9 "github.com/aws/aws-cdk-go/awscdk/aws_cloud9"
//   repositoryProperty := &repositoryProperty{
//   	pathComponent: jsii.String("pathComponent"),
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   }
//
type CfnEnvironmentEC2_RepositoryProperty struct {
	// The path within the development environment's default file system location to clone the AWS CodeCommit repository into.
	//
	// For example, `/REPOSITORY_NAME` would clone the repository into the `/home/USER_NAME/environment/REPOSITORY_NAME` directory in the environment.
	PathComponent *string `json:"pathComponent" yaml:"pathComponent"`
	// The clone URL of the AWS CodeCommit repository to be cloned.
	//
	// For example, for an AWS CodeCommit repository this might be `https://git-codecommit.us-east-2.amazonaws.com/v1/repos/REPOSITORY_NAME` .
	RepositoryUrl *string `json:"repositoryUrl" yaml:"repositoryUrl"`
}

// Properties for defining a `CfnEnvironmentEC2`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud9 "github.com/aws/aws-cdk-go/awscdk/aws_cloud9"
//   cfnEnvironmentEC2Props := &cfnEnvironmentEC2Props{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	automaticStopTimeMinutes: jsii.Number(123),
//   	connectionType: jsii.String("connectionType"),
//   	description: jsii.String("description"),
//   	imageId: jsii.String("imageId"),
//   	name: jsii.String("name"),
//   	ownerArn: jsii.String("ownerArn"),
//   	repositories: []interface{}{
//   		&repositoryProperty{
//   			pathComponent: jsii.String("pathComponent"),
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentEC2Props struct {
	// The type of instance to connect to the environment (for example, `t2.micro` ).
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// The number of minutes until the running instance is shut down after the environment was last used.
	AutomaticStopTimeMinutes *float64 `json:"automaticStopTimeMinutes" yaml:"automaticStopTimeMinutes"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are `CONNECT_SSH` (default) and `CONNECT_SSM` (connected through AWS Systems Manager ).
	ConnectionType *string `json:"connectionType" yaml:"connectionType"`
	// The description of the environment to create.
	Description *string `json:"description" yaml:"description"`
	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
	//
	// To choose an AMI for the instance, you must specify a valid AMI alias or a valid AWS Systems Manager path.
	//
	// The default AMI is used if the parameter isn't explicitly assigned a value in the request.
	//
	// *AMI aliases*
	//
	// - *Amazon Linux (default): `amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `ubuntu-18.04-x86_64`
	//
	// *SSM paths*
	//
	// - *Amazon Linux (default): `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
	ImageId *string `json:"imageId" yaml:"imageId"`
	// The name of the environment.
	Name *string `json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the environment owner.
	//
	// This ARN can be the ARN of any AWS Identity and Access Management principal. If this value is not specified, the ARN defaults to this environment's creator.
	OwnerArn *string `json:"ownerArn" yaml:"ownerArn"`
	// Any AWS CodeCommit source code repositories to be cloned into the development environment.
	Repositories interface{} `json:"repositories" yaml:"repositories"`
	// The ID of the subnet in Amazon Virtual Private Cloud (Amazon VPC) that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// The class for different repository providers.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   })
//
// Experimental.
type CloneRepository interface {
	// Experimental.
	PathComponent() *string
	// Experimental.
	RepositoryUrl() *string
}

// The jsii proxy struct for CloneRepository
type jsiiProxy_CloneRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CloneRepository) PathComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloneRepository) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}


// import repository to cloud9 environment from AWS CodeCommit.
// Experimental.
func CloneRepository_FromCodeCommit(repository awscodecommit.IRepository, path *string) CloneRepository {
	_init_.Initialize()

	var returns CloneRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.CloneRepository",
		"fromCodeCommit",
		[]interface{}{repository, path},
		&returns,
	)

	return returns
}

// A Cloud9 Environment with Amazon EC2.
//
// Example:
//   // create a cloud9 ec2 environment in a new VPC
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(3),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   })
//
//   // or create the cloud9 environment in the default VPC with specific instanceType
//   defaultVpc := ec2.vpc.fromLookup(this, jsii.String("DefaultVPC"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &ec2EnvironmentProps{
//   	vpc: defaultVpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   })
//
//   // or specify in a different subnetSelection
//   c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	subnetSelection: &subnetSelection{
//   		subnetType: ec2.subnetType_PRIVATE,
//   	},
//   })
//
//   // print the Cloud9 IDE URL in the output
//   // print the Cloud9 IDE URL in the output
//   NewCfnOutput(this, jsii.String("URL"), &cfnOutputProps{
//   	value: c9env.ideUrl,
//   })
//
// Experimental.
type Ec2Environment interface {
	awscdk.Resource
	IEc2Environment
	// The environment ARN of this Cloud9 environment.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The environment name of this Cloud9 environment.
	// Experimental.
	Ec2EnvironmentName() *string
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
	// The environment ID of this Cloud9 environment.
	// Experimental.
	EnvironmentId() *string
	// The complete IDE URL of this Cloud9 environment.
	// Experimental.
	IdeUrl() *string
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
	// VPC ID.
	// Experimental.
	Vpc() awsec2.IVpc
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

// The jsii proxy struct for Ec2Environment
type jsiiProxy_Ec2Environment struct {
	internal.Type__awscdkResource
	jsiiProxy_IEc2Environment
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) IdeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ideUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2Environment(scope constructs.Construct, id *string, props *Ec2EnvironmentProps) Ec2Environment {
	_init_.Initialize()

	j := jsiiProxy_Ec2Environment{}

	_jsii_.Create(
		"monocdk.aws_cloud9.Ec2Environment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2Environment_Override(e Ec2Environment, scope constructs.Construct, id *string, props *Ec2EnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloud9.Ec2Environment",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from EnvironmentEc2Name.
// Experimental.
func Ec2Environment_FromEc2EnvironmentName(scope constructs.Construct, id *string, ec2EnvironmentName *string) IEc2Environment {
	_init_.Initialize()

	var returns IEc2Environment

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.Ec2Environment",
		"fromEc2EnvironmentName",
		[]interface{}{scope, id, ec2EnvironmentName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Ec2Environment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.Ec2Environment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Ec2Environment_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloud9.Ec2Environment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_Ec2Environment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Environment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_Ec2Environment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Environment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_Ec2Environment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for Ec2Environment.
//
// Example:
//   // create a cloud9 ec2 environment in a new VPC
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(3),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   })
//
//   // or create the cloud9 environment in the default VPC with specific instanceType
//   defaultVpc := ec2.vpc.fromLookup(this, jsii.String("DefaultVPC"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &ec2EnvironmentProps{
//   	vpc: defaultVpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   })
//
//   // or specify in a different subnetSelection
//   c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	subnetSelection: &subnetSelection{
//   		subnetType: ec2.subnetType_PRIVATE,
//   	},
//   })
//
//   // print the Cloud9 IDE URL in the output
//   // print the Cloud9 IDE URL in the output
//   NewCfnOutput(this, jsii.String("URL"), &cfnOutputProps{
//   	value: c9env.ideUrl,
//   })
//
// Experimental.
type Ec2EnvironmentProps struct {
	// The VPC that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The AWS CodeCommit repository to be cloned.
	// Experimental.
	ClonedRepositories *[]CloneRepository `json:"clonedRepositories" yaml:"clonedRepositories"`
	// Description of the environment.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Name of the environment.
	// Experimental.
	Ec2EnvironmentName *string `json:"ec2EnvironmentName" yaml:"ec2EnvironmentName"`
	// The type of instance to connect to the environment.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The subnetSelection of the VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
}

// A Cloud9 Environment.
// Experimental.
type IEc2Environment interface {
	awscdk.IResource
	// The arn of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The name of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentName() *string
}

// The jsii proxy for IEc2Environment
type jsiiProxy_IEc2Environment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

