package awsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DataSync::Agent`.
//
// The `AWS::DataSync::Agent` resource specifies an AWS DataSync agent to be deployed and activated on your host. The activation process associates your agent with your account. In the activation process, you specify information such as the AWS Region that you want to activate the agent in. You activate the agent in the AWS Region where your target locations (in Amazon S3, Amazon EFS, or Amazon FSx for Windows File Server) reside. Your tasks are created in this AWS Region .
//
// You can activate the agent in a virtual private cloud (VPC) or provide the agent access to a VPC endpoint so that you can run tasks without sending them over the public internet.
//
// You can specify an agent to be used for more than one location. If a task uses multiple agents, all of them must have a status of AVAILABLE for the task to run. If you use multiple agents for a source location, the status of all the agents must be AVAILABLE for the task to run.
//
// For more information, see [Activating an Agent](https://docs.aws.amazon.com/datasync/latest/userguide/activating-agent.html) in the *AWS DataSync User Guide* .
//
// Agents are automatically updated by AWS on a regular basis, using a mechanism that ensures minimal interruption to your tasks.
//
// TODO: EXAMPLE
//
type CfnAgent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ActivationKey() *string
	SetActivationKey(val *string)
	AgentName() *string
	SetAgentName(val *string)
	AttrAgentArn() *string
	AttrEndpointType() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SecurityGroupArns() *[]*string
	SetSecurityGroupArns(val *[]*string)
	Stack() awscdk.Stack
	SubnetArns() *[]*string
	SetSubnetArns(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VpcEndpointId() *string
	SetVpcEndpointId(val *string)
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

// The jsii proxy struct for CfnAgent
type jsiiProxy_CfnAgent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAgent) ActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrAgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAgentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) SecurityGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) SubnetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::Agent`.
func NewCfnAgent(scope constructs.Construct, id *string, props *CfnAgentProps) CfnAgent {
	_init_.Initialize()

	j := jsiiProxy_CfnAgent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnAgent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::Agent`.
func NewCfnAgent_Override(c CfnAgent, scope constructs.Construct, id *string, props *CfnAgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnAgent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAgent) SetActivationKey(val *string) {
	_jsii_.Set(
		j,
		"activationKey",
		val,
	)
}

func (j *jsiiProxy_CfnAgent) SetAgentName(val *string) {
	_jsii_.Set(
		j,
		"agentName",
		val,
	)
}

func (j *jsiiProxy_CfnAgent) SetSecurityGroupArns(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupArns",
		val,
	)
}

func (j *jsiiProxy_CfnAgent) SetSubnetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetArns",
		val,
	)
}

func (j *jsiiProxy_CfnAgent) SetVpcEndpointId(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAgent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnAgent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAgent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnAgent",
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
func CfnAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnAgent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAgent) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAgent) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAgent) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAgent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAgent) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAgent) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAgent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAgent) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAgent) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAgent) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAgent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAgent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAgent) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgent) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAgent`.
//
// TODO: EXAMPLE
//
type CfnAgentProps struct {
	// Your agent activation key.
	//
	// You can get the activation key either by sending an HTTP GET request with redirects that enable you to get the agent IP address (port 80). Alternatively, you can get it from the DataSync console.
	//
	// The redirect URL returned in the response provides you the activation key for your agent in the query string parameter `activationKey` . It might also include other activation-related parameters; however, these are merely defaults. The arguments you pass to this API call determine the actual configuration of your agent.
	//
	// For more information, see [Creating and activating an agent](https://docs.aws.amazon.com/datasync/latest/userguide/activating-agent.html) in the *AWS DataSync User Guide.*
	ActivationKey *string `json:"activationKey" yaml:"activationKey"`
	// The name you configured for your agent.
	//
	// This value is a text reference that is used to identify the agent in the console.
	AgentName *string `json:"agentName" yaml:"agentName"`
	// The Amazon Resource Names (ARNs) of the security groups used to protect your data transfer task subnets.
	//
	// See [SecurityGroupArns](https://docs.aws.amazon.com/datasync/latest/userguide/API_Ec2Config.html#DataSync-Type-Ec2Config-SecurityGroupArns) .
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	SecurityGroupArns *[]*string `json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Names (ARNs) of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	//
	// The agent that runs a task must be private. When you start a task that is associated with an agent created in a VPC, or one that has access to an IP address in a VPC, then the task is also private. In this case, DataSync creates four network interfaces for each task in your subnet. For a data transfer to work, the agent must be able to route to all these four network interfaces.
	SubnetArns *[]*string `json:"subnetArns" yaml:"subnetArns"`
	// The key-value pair that represents the tag that you want to associate with the agent.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your agents.
	//
	// > Valid characters for key and value are letters, spaces, and numbers representable in UTF-8 format, and the following special characters: + - = . _ : / @.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The ID of the virtual private cloud (VPC) endpoint that the agent has access to.
	//
	// This is the client-side VPC endpoint, powered by AWS PrivateLink . If you don't have an AWS PrivateLink VPC endpoint, see [AWS PrivateLink and VPC endpoints](https://docs.aws.amazon.com//vpc/latest/userguide/endpoint-services-overview.html) in the *Amazon VPC User Guide* .
	//
	// For more information about activating your agent in a private network based on a VPC, see [Using AWS DataSync in a Virtual Private Cloud](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-in-vpc.html) in the *AWS DataSync User Guide.*
	//
	// A VPC endpoint ID looks like this: `vpce-01234d5aff67890e1` .
	VpcEndpointId *string `json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

// A CloudFormation `AWS::DataSync::LocationEFS`.
//
// The `AWS::DataSync::LocationEFS` resource specifies an endpoint for an Amazon EFS location.
//
// TODO: EXAMPLE
//
type CfnLocationEFS interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Ec2Config() interface{}
	SetEc2Config(val interface{})
	EfsFilesystemArn() *string
	SetEfsFilesystemArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationEFS
type jsiiProxy_CfnLocationEFS struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationEFS) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Ec2Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) EfsFilesystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"efsFilesystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationEFS) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationEFS`.
func NewCfnLocationEFS(scope constructs.Construct, id *string, props *CfnLocationEFSProps) CfnLocationEFS {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationEFS{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationEFS`.
func NewCfnLocationEFS_Override(c CfnLocationEFS, scope constructs.Construct, id *string, props *CfnLocationEFSProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationEFS) SetEc2Config(val interface{}) {
	_jsii_.Set(
		j,
		"ec2Config",
		val,
	)
}

func (j *jsiiProxy_CfnLocationEFS) SetEfsFilesystemArn(val *string) {
	_jsii_.Set(
		j,
		"efsFilesystemArn",
		val,
	)
}

func (j *jsiiProxy_CfnLocationEFS) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationEFS_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationEFS_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
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
func CfnLocationEFS_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationEFS_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationEFS",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationEFS) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationEFS) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationEFS) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationEFS) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationEFS) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationEFS) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationEFS) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationEFS) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationEFS) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationEFS) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationEFS) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationEFS) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationEFS) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationEFS) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationEFS) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The subnet and the security group that DataSync uses to access the target EFS file system.
//
// The subnet must have at least one mount target for that file system. The security group that you provide must be able to communicate with the security group on the mount target in the subnet specified.
//
// TODO: EXAMPLE
//
type CfnLocationEFS_Ec2ConfigProperty struct {
	// The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	SecurityGroupArns *[]*string `json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Name (ARN) of the subnet that DataSync uses to access the target EFS file system.
	SubnetArn *string `json:"subnetArn" yaml:"subnetArn"`
}

// Properties for defining a `CfnLocationEFS`.
//
// TODO: EXAMPLE
//
type CfnLocationEFSProps struct {
	// The subnet and security group that the Amazon EFS file system uses.
	//
	// The security group that you provide needs to be able to communicate with the security group on the mount target in the subnet specified.
	//
	// The exact relationship between security group M (of the mount target) and security group S (which you provide for DataSync to use at this stage) is as follows:
	//
	// - Security group M (which you associate with the mount target) must allow inbound access for the Transmission Control Protocol (TCP) on the NFS port (2049) from security group S. You can enable inbound connections either by IP address (CIDR range) or security group.
	// - Security group S (provided to DataSync to access EFS) should have a rule that enables outbound connections to the NFS port on one of the file system’s mount targets. You can enable outbound connections either by IP address (CIDR range) or security group.
	//
	// For information about security groups and mount targets, see [Security Groups for Amazon EC2 Instances and Mount Targets](https://docs.aws.amazon.com/efs/latest/ug/security-considerations.html#network-access) in the *Amazon EFS User Guide.*
	Ec2Config interface{} `json:"ec2Config" yaml:"ec2Config"`
	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	EfsFilesystemArn *string `json:"efsFilesystemArn" yaml:"efsFilesystemArn"`
	// A subdirectory in the location’s path.
	//
	// This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination. By default, AWS DataSync uses the root directory.
	//
	// > `Subdirectory` must be specified with forward slashes. For example, `/path/to/folder` .
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationFSxLustre`.
//
// The `AWS::DataSync::LocationFSxLustre` resource specifies an endpoint for an Amazon FSx for Lustre file system.
//
// TODO: EXAMPLE
//
type CfnLocationFSxLustre interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	FsxFilesystemArn() *string
	SetFsxFilesystemArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SecurityGroupArns() *[]*string
	SetSecurityGroupArns(val *[]*string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationFSxLustre
type jsiiProxy_CfnLocationFSxLustre struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationFSxLustre) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) FsxFilesystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxFilesystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) SecurityGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxLustre) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationFSxLustre`.
func NewCfnLocationFSxLustre(scope constructs.Construct, id *string, props *CfnLocationFSxLustreProps) CfnLocationFSxLustre {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationFSxLustre{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationFSxLustre`.
func NewCfnLocationFSxLustre_Override(c CfnLocationFSxLustre, scope constructs.Construct, id *string, props *CfnLocationFSxLustreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationFSxLustre) SetFsxFilesystemArn(val *string) {
	_jsii_.Set(
		j,
		"fsxFilesystemArn",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxLustre) SetSecurityGroupArns(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupArns",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxLustre) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationFSxLustre_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationFSxLustre_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
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
func CfnLocationFSxLustre_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationFSxLustre_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxLustre",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationFSxLustre) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationFSxLustre) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationFSxLustre) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationFSxLustre) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationFSxLustre) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationFSxLustre) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationFSxLustre) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationFSxLustre) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationFSxLustre) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationFSxLustre) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationFSxLustre) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationFSxLustre) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationFSxLustre) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationFSxLustre) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationFSxLustre) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLocationFSxLustre`.
//
// TODO: EXAMPLE
//
type CfnLocationFSxLustreProps struct {
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn *string `json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// The ARNs of the security groups that are used to configure the FSx for Lustre file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	SecurityGroupArns *[]*string `json:"securityGroupArns" yaml:"securityGroupArns"`
	// A subdirectory in the location's path.
	//
	// This subdirectory in the FSx for Lustre file system is used to read data from the FSx for Lustre source location or write data to the FSx for Lustre destination.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationFSxWindows`.
//
// The `AWS::DataSync::LocationFSxWindows` resource specifies an endpoint for an Amazon FSx for Windows Server file system.
//
// TODO: EXAMPLE
//
type CfnLocationFSxWindows interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Domain() *string
	SetDomain(val *string)
	FsxFilesystemArn() *string
	SetFsxFilesystemArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	Ref() *string
	SecurityGroupArns() *[]*string
	SetSecurityGroupArns(val *[]*string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	User() *string
	SetUser(val *string)
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

// The jsii proxy struct for CfnLocationFSxWindows
type jsiiProxy_CfnLocationFSxWindows struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationFSxWindows) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) FsxFilesystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxFilesystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) SecurityGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxWindows) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationFSxWindows`.
func NewCfnLocationFSxWindows(scope constructs.Construct, id *string, props *CfnLocationFSxWindowsProps) CfnLocationFSxWindows {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationFSxWindows{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationFSxWindows`.
func NewCfnLocationFSxWindows_Override(c CfnLocationFSxWindows, scope constructs.Construct, id *string, props *CfnLocationFSxWindowsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetFsxFilesystemArn(val *string) {
	_jsii_.Set(
		j,
		"fsxFilesystemArn",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetSecurityGroupArns(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupArns",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

func (j *jsiiProxy_CfnLocationFSxWindows) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationFSxWindows_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationFSxWindows_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
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
func CfnLocationFSxWindows_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationFSxWindows_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationFSxWindows",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationFSxWindows) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationFSxWindows) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationFSxWindows) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationFSxWindows) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationFSxWindows) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationFSxWindows) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationFSxWindows) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationFSxWindows) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationFSxWindows) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationFSxWindows) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationFSxWindows) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationFSxWindows) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationFSxWindows) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationFSxWindows) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationFSxWindows) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLocationFSxWindows`.
//
// TODO: EXAMPLE
//
type CfnLocationFSxWindowsProps struct {
	// The Amazon Resource Name (ARN) for the FSx for Windows File Server file system.
	FsxFilesystemArn *string `json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows File Server file system.
	Password *string `json:"password" yaml:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are used to configure the FSx for Windows File Server file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	SecurityGroupArns *[]*string `json:"securityGroupArns" yaml:"securityGroupArns"`
	// The user who has the permissions to access files and folders in the FSx for Windows File Server file system.
	//
	// For information about choosing a user name that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#FSxWuser) .
	User *string `json:"user" yaml:"user"`
	// The name of the Windows domain that the FSx for Windows File Server belongs to.
	Domain *string `json:"domain" yaml:"domain"`
	// A subdirectory in the location's path.
	//
	// This subdirectory in the Amazon FSx for Windows File Server file system is used to read data from the Amazon FSx for Windows File Server source location or write data to the FSx for Windows File Server destination.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationHDFS`.
//
// The `AWS::DataSync::LocationHDFS` resource specifies an endpoint for a Hadoop Distributed File System (HDFS).
//
// TODO: EXAMPLE
//
type CfnLocationHDFS interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AgentArns() *[]*string
	SetAgentArns(val *[]*string)
	AttrLocationArn() *string
	AttrLocationUri() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	BlockSize() *float64
	SetBlockSize(val *float64)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	KerberosKeytab() *string
	SetKerberosKeytab(val *string)
	KerberosKrb5Conf() *string
	SetKerberosKrb5Conf(val *string)
	KerberosPrincipal() *string
	SetKerberosPrincipal(val *string)
	KmsKeyProviderUri() *string
	SetKmsKeyProviderUri(val *string)
	LogicalId() *string
	NameNodes() interface{}
	SetNameNodes(val interface{})
	Node() constructs.Node
	QopConfiguration() interface{}
	SetQopConfiguration(val interface{})
	Ref() *string
	ReplicationFactor() *float64
	SetReplicationFactor(val *float64)
	SimpleUser() *string
	SetSimpleUser(val *string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationHDFS
type jsiiProxy_CfnLocationHDFS struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationHDFS) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) BlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) KerberosKeytab() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) KerberosKrb5Conf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5Conf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) KerberosPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) KmsKeyProviderUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyProviderUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) NameNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) QopConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"qopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) SimpleUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simpleUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationHDFS) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationHDFS`.
func NewCfnLocationHDFS(scope constructs.Construct, id *string, props *CfnLocationHDFSProps) CfnLocationHDFS {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationHDFS{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationHDFS`.
func NewCfnLocationHDFS_Override(c CfnLocationHDFS, scope constructs.Construct, id *string, props *CfnLocationHDFSProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetAgentArns(val *[]*string) {
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetBlockSize(val *float64) {
	_jsii_.Set(
		j,
		"blockSize",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetKerberosKeytab(val *string) {
	_jsii_.Set(
		j,
		"kerberosKeytab",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetKerberosKrb5Conf(val *string) {
	_jsii_.Set(
		j,
		"kerberosKrb5Conf",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetKerberosPrincipal(val *string) {
	_jsii_.Set(
		j,
		"kerberosPrincipal",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetKmsKeyProviderUri(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyProviderUri",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetNameNodes(val interface{}) {
	_jsii_.Set(
		j,
		"nameNodes",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetQopConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"qopConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetReplicationFactor(val *float64) {
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetSimpleUser(val *string) {
	_jsii_.Set(
		j,
		"simpleUser",
		val,
	)
}

func (j *jsiiProxy_CfnLocationHDFS) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationHDFS_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationHDFS_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
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
func CfnLocationHDFS_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationHDFS_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationHDFS",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationHDFS) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationHDFS) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationHDFS) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationHDFS) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationHDFS) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationHDFS) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationHDFS) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationHDFS) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationHDFS) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationHDFS) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationHDFS) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationHDFS) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationHDFS) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationHDFS) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationHDFS) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The NameNode of the Hadoop Distributed File System (HDFS).
//
// The NameNode manages the file system's namespace and performs operations such as opening, closing, and renaming files and directories. The NameNode also contains the information to map blocks of data to the DataNodes.
//
// TODO: EXAMPLE
//
type CfnLocationHDFS_NameNodeProperty struct {
	// The hostname of the NameNode in the HDFS cluster.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the NameNode. An agent that's installed on-premises uses this hostname to communicate with the NameNode in the network.
	Hostname *string `json:"hostname" yaml:"hostname"`
	// The port that the NameNode uses to listen to client requests.
	Port *float64 `json:"port" yaml:"port"`
}

// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer privacy settings configured on the Hadoop Distributed File System (HDFS) cluster.
//
// TODO: EXAMPLE
//
type CfnLocationHDFS_QopConfigurationProperty struct {
	// The data transfer protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `dfs.data.transfer.protection` setting in the `hdfs-site.xml` file on your Hadoop cluster.
	DataTransferProtection *string `json:"dataTransferProtection" yaml:"dataTransferProtection"`
	// The Remote Procedure Call (RPC) protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `hadoop.rpc.protection` setting in your `core-site.xml` file on your Hadoop cluster.
	RpcProtection *string `json:"rpcProtection" yaml:"rpcProtection"`
}

// Properties for defining a `CfnLocationHDFS`.
//
// TODO: EXAMPLE
//
type CfnLocationHDFSProps struct {
	// The Amazon Resource Names (ARNs) of the agents that are used to connect to the HDFS cluster.
	AgentArns *[]*string `json:"agentArns" yaml:"agentArns"`
	// `AWS::DataSync::LocationHDFS.AuthenticationType`.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// The NameNode that manages the HDFS namespace.
	//
	// The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode.
	NameNodes interface{} `json:"nameNodes" yaml:"nameNodes"`
	// The size of data blocks to write into the HDFS cluster.
	//
	// The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
	BlockSize *float64 `json:"blockSize" yaml:"blockSize"`
	// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys.
	//
	// Provide the base64-encoded file text. If `KERBEROS` is specified for `AuthType` , this value is required.
	KerberosKeytab *string `json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The `krb5.conf` file that contains the Kerberos configuration information. You can load the `krb5.conf` by providing a string of the file's contents or an Amazon S3 presigned URL of the file. If `KERBEROS` is specified for `AuthType` , this value is required.
	KerberosKrb5Conf *string `json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// The Kerberos principal with access to the files and folders on the HDFS cluster.
	//
	// > If `KERBEROS` is specified for `AuthenticationType` , this parameter is required.
	KerberosPrincipal *string `json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string `json:"kmsKeyProviderUri" yaml:"kmsKeyProviderUri"`
	// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster.
	//
	// If `QopConfiguration` isn't specified, `RpcProtection` and `DataTransferProtection` default to `PRIVACY` . If you set `RpcProtection` or `DataTransferProtection` , the other parameter assumes the same value.
	QopConfiguration interface{} `json:"qopConfiguration" yaml:"qopConfiguration"`
	// The number of DataNodes to replicate the data to when writing to the HDFS cluster.
	//
	// By default, data is replicated to three DataNodes.
	ReplicationFactor *float64 `json:"replicationFactor" yaml:"replicationFactor"`
	// The user name used to identify the client on the host operating system.
	//
	// > If `SIMPLE` is specified for `AuthenticationType` , this parameter is required.
	SimpleUser *string `json:"simpleUser" yaml:"simpleUser"`
	// A subdirectory in the HDFS cluster.
	//
	// This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to `/` .
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationNFS`.
//
// The `AWS::DataSync::LocationNFS` resource specifies a file system on a Network File System (NFS) server that can be read from or written to.
//
// TODO: EXAMPLE
//
type CfnLocationNFS interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MountOptions() interface{}
	SetMountOptions(val interface{})
	Node() constructs.Node
	OnPremConfig() interface{}
	SetOnPremConfig(val interface{})
	Ref() *string
	ServerHostname() *string
	SetServerHostname(val *string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationNFS
type jsiiProxy_CfnLocationNFS struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationNFS) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) MountOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) OnPremConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) ServerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFS) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationNFS`.
func NewCfnLocationNFS(scope constructs.Construct, id *string, props *CfnLocationNFSProps) CfnLocationNFS {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationNFS{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationNFS`.
func NewCfnLocationNFS_Override(c CfnLocationNFS, scope constructs.Construct, id *string, props *CfnLocationNFSProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationNFS) SetMountOptions(val interface{}) {
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_CfnLocationNFS) SetOnPremConfig(val interface{}) {
	_jsii_.Set(
		j,
		"onPremConfig",
		val,
	)
}

func (j *jsiiProxy_CfnLocationNFS) SetServerHostname(val *string) {
	_jsii_.Set(
		j,
		"serverHostname",
		val,
	)
}

func (j *jsiiProxy_CfnLocationNFS) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationNFS_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationNFS_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
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
func CfnLocationNFS_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationNFS_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationNFS",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationNFS) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationNFS) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationNFS) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationNFS) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationNFS) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationNFS) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationNFS) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationNFS) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationNFS) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationNFS) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationNFS) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationNFS) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationNFS) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationNFS) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationNFS) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The NFS mount options that DataSync can use to mount your NFS share.
//
// TODO: EXAMPLE
//
type CfnLocationNFS_MountOptionsProperty struct {
	// The specific NFS version that you want DataSync to use to mount your NFS share.
	//
	// If the server refuses to use the version specified, the sync will fail. If you don't specify a version, DataSync defaults to `AUTOMATIC` . That is, DataSync automatically selects a version based on negotiation with the NFS server.
	//
	// You can specify the following NFS versions:
	//
	// - *[NFSv3](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc1813)* - stateless protocol version that allows for asynchronous writes on the server.
	// - *[NFSv4.0](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3530)* - stateful, firewall-friendly protocol version that supports delegations and pseudo file systems.
	// - *[NFSv4.1](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc5661)* - stateful protocol version that supports sessions, directory delegations, and parallel data processing. Version 4.1 also includes all features available in version 4.0.
	Version *string `json:"version" yaml:"version"`
}

// A list of Amazon Resource Names (ARNs) of agents to use for a Network File System (NFS) location.
//
// TODO: EXAMPLE
//
type CfnLocationNFS_OnPremConfigProperty struct {
	// ARNs of the agents to use for an NFS location.
	AgentArns *[]*string `json:"agentArns" yaml:"agentArns"`
}

// Properties for defining a `CfnLocationNFS`.
//
// TODO: EXAMPLE
//
type CfnLocationNFSProps struct {
	// Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect to an NFS server.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	OnPremConfig interface{} `json:"onPremConfig" yaml:"onPremConfig"`
	// The name of the NFS server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the NFS server. An agent that is installed on-premises uses this host name to mount the NFS server in a network.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	//
	// > This name must either be DNS-compliant or must be an IP version 4 (IPv4) address.
	ServerHostname *string `json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
	//
	// The NFS path should be a path that's exported by the NFS server, or a subdirectory of that path. The path should be such that it can be mounted by other NFS clients in your network.
	//
	// To see all the paths exported by your NFS server, run " `showmount -e nfs-server-name` " from an NFS client that has access to your server. You can specify any directory that appears in the results, and any subdirectory of that directory. Ensure that the NFS export is accessible without Kerberos authentication.
	//
	// To transfer all the data in the folder you specified, DataSync needs to have permissions to read all the data. To ensure this, either configure the NFS export with `no_root_squash,` or ensure that the permissions for all of the files that you want DataSync allow read access for all users. Doing either enables the agent to read the files. For the agent to access directories, you must additionally enable all execute access.
	//
	// If you are copying data to or from your AWS Snowcone device, see [NFS Server on AWS Snowcone](https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone) for more information.
	//
	// For information about NFS export configuration, see [18.7. The /etc/exports Configuration File](https://docs.aws.amazon.com/http://web.mit.edu/rhel-doc/5/RHEL-5-manual/Deployment_Guide-en-US/s1-nfs-server-config-exports.html) in the Red Hat Enterprise Linux documentation.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The NFS mount options that DataSync can use to mount your NFS share.
	MountOptions interface{} `json:"mountOptions" yaml:"mountOptions"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationObjectStorage`.
//
// The `AWS::DataSync::LocationObjectStorage` resource specifies an endpoint for a self-managed object storage bucket. For more information about self-managed object storage locations, see [Creating a Location for Object Storage](https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html) .
//
// TODO: EXAMPLE
//
type CfnLocationObjectStorage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessKey() *string
	SetAccessKey(val *string)
	AgentArns() *[]*string
	SetAgentArns(val *[]*string)
	AttrLocationArn() *string
	AttrLocationUri() *string
	BucketName() *string
	SetBucketName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SecretKey() *string
	SetSecretKey(val *string)
	ServerHostname() *string
	SetServerHostname(val *string)
	ServerPort() *float64
	SetServerPort(val *float64)
	ServerProtocol() *string
	SetServerProtocol(val *string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationObjectStorage
type jsiiProxy_CfnLocationObjectStorage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationObjectStorage) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) ServerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) ServerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) ServerProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStorage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationObjectStorage`.
func NewCfnLocationObjectStorage(scope constructs.Construct, id *string, props *CfnLocationObjectStorageProps) CfnLocationObjectStorage {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationObjectStorage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationObjectStorage`.
func NewCfnLocationObjectStorage_Override(c CfnLocationObjectStorage, scope constructs.Construct, id *string, props *CfnLocationObjectStorageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetAgentArns(val *[]*string) {
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetServerHostname(val *string) {
	_jsii_.Set(
		j,
		"serverHostname",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetServerPort(val *float64) {
	_jsii_.Set(
		j,
		"serverPort",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetServerProtocol(val *string) {
	_jsii_.Set(
		j,
		"serverProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnLocationObjectStorage) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationObjectStorage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationObjectStorage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
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
func CfnLocationObjectStorage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationObjectStorage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationObjectStorage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationObjectStorage) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationObjectStorage) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationObjectStorage) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationObjectStorage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationObjectStorage) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationObjectStorage) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationObjectStorage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationObjectStorage) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationObjectStorage) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationObjectStorage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationObjectStorage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationObjectStorage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationObjectStorage) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationObjectStorage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationObjectStorage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLocationObjectStorage`.
//
// TODO: EXAMPLE
//
type CfnLocationObjectStorageProps struct {
	// The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
	AgentArns *[]*string `json:"agentArns" yaml:"agentArns"`
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The name of the self-managed object storage server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname *string `json:"serverHostname" yaml:"serverHostname"`
	// Optional.
	//
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `AccessKey` and `SecretKey` to provide the user name and password, respectively.
	AccessKey *string `json:"accessKey" yaml:"accessKey"`
	// Optional.
	//
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `AccessKey` and `SecretKey` to provide the user name and password, respectively.
	SecretKey *string `json:"secretKey" yaml:"secretKey"`
	// The port that your self-managed object storage server accepts inbound network traffic on.
	//
	// The server port is set by default to TCP 80 (HTTP) or TCP 443 (HTTPS). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort *float64 `json:"serverPort" yaml:"serverPort"`
	// The protocol that the object storage server uses to communicate.
	//
	// Valid values are HTTP or HTTPS.
	ServerProtocol *string `json:"serverProtocol" yaml:"serverProtocol"`
	// The subdirectory in the self-managed object storage server that is used to read data from.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationS3`.
//
// The `AWS::DataSync::LocationS3` resource specifies an endpoint for an Amazon S3 bucket.
//
// For more information, see [Create an Amazon S3 location](https://docs.aws.amazon.com/datasync/latest/userguide/create-locations-cli.html#create-location-s3-cli) in the *AWS DataSync User Guide* .
//
// TODO: EXAMPLE
//
type CfnLocationS3 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	S3BucketArn() *string
	SetS3BucketArn(val *string)
	S3Config() interface{}
	SetS3Config(val interface{})
	S3StorageClass() *string
	SetS3StorageClass(val *string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
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

// The jsii proxy struct for CfnLocationS3
type jsiiProxy_CfnLocationS3 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationS3) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) S3BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) S3Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) S3StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationS3) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationS3`.
func NewCfnLocationS3(scope constructs.Construct, id *string, props *CfnLocationS3Props) CfnLocationS3 {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationS3{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationS3`.
func NewCfnLocationS3_Override(c CfnLocationS3, scope constructs.Construct, id *string, props *CfnLocationS3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationS3) SetS3BucketArn(val *string) {
	_jsii_.Set(
		j,
		"s3BucketArn",
		val,
	)
}

func (j *jsiiProxy_CfnLocationS3) SetS3Config(val interface{}) {
	_jsii_.Set(
		j,
		"s3Config",
		val,
	)
}

func (j *jsiiProxy_CfnLocationS3) SetS3StorageClass(val *string) {
	_jsii_.Set(
		j,
		"s3StorageClass",
		val,
	)
}

func (j *jsiiProxy_CfnLocationS3) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationS3_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationS3_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
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
func CfnLocationS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationS3_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationS3",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationS3) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationS3) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationS3) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationS3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationS3) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationS3) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationS3) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationS3) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationS3) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationS3) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationS3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationS3) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationS3) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationS3) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role used to access an Amazon S3 bucket.
//
// For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
//
// TODO: EXAMPLE
//
type CfnLocationS3_S3ConfigProperty struct {
	// The ARN of the IAM role for accessing the S3 bucket.
	BucketAccessRoleArn *string `json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
}

// Properties for defining a `CfnLocationS3`.
//
// TODO: EXAMPLE
//
type CfnLocationS3Props struct {
	// The ARN of the Amazon S3 bucket.
	S3BucketArn *string `json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that is used to access an Amazon S3 bucket.
	//
	// For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
	S3Config interface{} `json:"s3Config" yaml:"s3Config"`
	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination.
	//
	// For buckets in AWS Regions , the storage class defaults to S3 Standard.
	//
	// For more information about S3 storage classes, see [Amazon S3 Storage Classes](https://docs.aws.amazon.com/s3/storage-classes/) . Some storage classes have behaviors that can affect your S3 storage costs. For detailed information, see [Considerations When Working with Amazon S3 Storage Classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) .
	S3StorageClass *string `json:"s3StorageClass" yaml:"s3StorageClass"`
	// A subdirectory in the Amazon S3 bucket.
	//
	// This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::LocationSMB`.
//
// The `AWS::DataSync::LocationSMB` resource specifies a Server Message Block (SMB) location.
//
// TODO: EXAMPLE
//
type CfnLocationSMB interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AgentArns() *[]*string
	SetAgentArns(val *[]*string)
	AttrLocationArn() *string
	AttrLocationUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Domain() *string
	SetDomain(val *string)
	LogicalId() *string
	MountOptions() interface{}
	SetMountOptions(val interface{})
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	Ref() *string
	ServerHostname() *string
	SetServerHostname(val *string)
	Stack() awscdk.Stack
	Subdirectory() *string
	SetSubdirectory(val *string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	User() *string
	SetUser(val *string)
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

// The jsii proxy struct for CfnLocationSMB
type jsiiProxy_CfnLocationSMB struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLocationSMB) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) AttrLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) AttrLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) MountOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) ServerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationSMB) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::LocationSMB`.
func NewCfnLocationSMB(scope constructs.Construct, id *string, props *CfnLocationSMBProps) CfnLocationSMB {
	_init_.Initialize()

	j := jsiiProxy_CfnLocationSMB{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::LocationSMB`.
func NewCfnLocationSMB_Override(c CfnLocationSMB, scope constructs.Construct, id *string, props *CfnLocationSMBProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetAgentArns(val *[]*string) {
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetMountOptions(val interface{}) {
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetServerHostname(val *string) {
	_jsii_.Set(
		j,
		"serverHostname",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetSubdirectory(val *string) {
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

func (j *jsiiProxy_CfnLocationSMB) SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLocationSMB_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLocationSMB_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
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
func CfnLocationSMB_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationSMB_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnLocationSMB",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLocationSMB) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLocationSMB) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLocationSMB) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationSMB) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLocationSMB) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLocationSMB) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLocationSMB) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLocationSMB) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLocationSMB) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLocationSMB) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLocationSMB) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLocationSMB) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLocationSMB) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLocationSMB) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLocationSMB) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The mount options used by DataSync to access the SMB server.
//
// TODO: EXAMPLE
//
type CfnLocationSMB_MountOptionsProperty struct {
	// The specific SMB version that you want DataSync to use to mount your SMB share.
	//
	// If you don't specify a version, DataSync defaults to `AUTOMATIC` . That is, DataSync automatically selects a version based on negotiation with the SMB server.
	Version *string `json:"version" yaml:"version"`
}

// Properties for defining a `CfnLocationSMB`.
//
// TODO: EXAMPLE
//
type CfnLocationSMBProps struct {
	// The Amazon Resource Names (ARNs) of agents to use for a Server Message Block (SMB) location.
	AgentArns *[]*string `json:"agentArns" yaml:"agentArns"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	Password *string `json:"password" yaml:"password"`
	// The name of the SMB server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the SMB server. An agent that is installed on-premises uses this hostname to mount the SMB server in a network.
	//
	// > This name must either be DNS-compliant or must be an IP version 4 (IPv4) address.
	ServerHostname *string `json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination.
	//
	// The SMB path should be a path that's exported by the SMB server, or a subdirectory of that path. The path should be such that it can be mounted by other SMB clients in your network.
	//
	// > `Subdirectory` must be specified with forward slashes. For example, `/path/to/folder` .
	//
	// To transfer all the data in the folder you specified, DataSync must have permissions to mount the SMB share, as well as to access all the data in that share. To ensure this, either make sure that the user name and password specified belongs to the user who can mount the share, and who has the appropriate permissions for all of the files and directories that you want DataSync to access, or use credentials of a member of the Backup Operators group to mount the share. Doing either one enables the agent to access the data. For the agent to access directories, you must additionally enable all execute access.
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// The user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// For information about choosing a user name that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#SMBuser) .
	User *string `json:"user" yaml:"user"`
	// The name of the Windows domain that the SMB server belongs to.
	Domain *string `json:"domain" yaml:"domain"`
	// The mount options used by DataSync to access the SMB server.
	MountOptions interface{} `json:"mountOptions" yaml:"mountOptions"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataSync::Task`.
//
// The `AWS::DataSync::Task` resource specifies a task. A task is a set of two locations (source and destination) and a set of `Options` that you use to control the behavior of a task. If you don't specify `Options` when you create a task, AWS DataSync populates them with service defaults.
//
// TODO: EXAMPLE
//
type CfnTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDestinationNetworkInterfaceArns() *[]*string
	AttrErrorCode() *string
	AttrErrorDetail() *string
	AttrSourceNetworkInterfaceArns() *[]*string
	AttrStatus() *string
	AttrTaskArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CloudWatchLogGroupArn() *string
	SetCloudWatchLogGroupArn(val *string)
	CreationStack() *[]*string
	DestinationLocationArn() *string
	SetDestinationLocationArn(val *string)
	Excludes() interface{}
	SetExcludes(val interface{})
	Includes() interface{}
	SetIncludes(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Options() interface{}
	SetOptions(val interface{})
	Ref() *string
	Schedule() interface{}
	SetSchedule(val interface{})
	SourceLocationArn() *string
	SetSourceLocationArn(val *string)
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

// The jsii proxy struct for CfnTask
type jsiiProxy_CfnTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTask) AttrDestinationNetworkInterfaceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDestinationNetworkInterfaceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrErrorCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrErrorDetail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrErrorDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrSourceNetworkInterfaceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSourceNetworkInterfaceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrTaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTaskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) DestinationLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Excludes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Includes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Options() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) SourceLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::Task`.
func NewCfnTask(scope constructs.Construct, id *string, props *CfnTaskProps) CfnTask {
	_init_.Initialize()

	j := jsiiProxy_CfnTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::Task`.
func NewCfnTask_Override(c CfnTask, scope constructs.Construct, id *string, props *CfnTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datasync.CfnTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTask) SetCloudWatchLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetDestinationLocationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationLocationArn",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetExcludes(val interface{}) {
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetIncludes(val interface{}) {
	_jsii_.Set(
		j,
		"includes",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetOptions(val interface{}) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnTask) SetSourceLocationArn(val *string) {
	_jsii_.Set(
		j,
		"sourceLocationArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnTask",
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
func CfnTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datasync.CfnTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datasync.CfnTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTask) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTask) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTask) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTask) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTask) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTask) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTask) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTask) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTask) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies which files, folders, and objects to include or exclude when transferring files from source to destination.
//
// TODO: EXAMPLE
//
type CfnTask_FilterRuleProperty struct {
	// The type of filter rule to apply.
	//
	// AWS DataSync only supports the SIMPLE_PATTERN rule type.
	FilterType *string `json:"filterType" yaml:"filterType"`
	// A single filter string that consists of the patterns to include or exclude.
	//
	// The patterns are delimited by "|" (that is, a pipe), for example: `/folder1|/folder2`
	Value *string `json:"value" yaml:"value"`
}

// Represents the options that are available to control the behavior of a [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) operation. This behavior includes preserving metadata, such as user ID (UID), group ID (GID), and file permissions; overwriting files in the destination; data integrity verification; and so on.
//
// A task has a set of default options associated with it. If you don't specify an option in [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) , the default value is used. You can override the default options on each task execution by specifying an overriding `Options` value to [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) .
//
// TODO: EXAMPLE
//
type CfnTask_OptionsProperty struct {
	// A file metadata value that shows the last time that a file was accessed (that is, when the file was read or written to).
	//
	// If you set `Atime` to `BEST_EFFORT` , AWS DataSync attempts to preserve the original `Atime` attribute on all source files (that is, the version before the PREPARING phase). However, `Atime` 's behavior is not fully standard across platforms, so AWS DataSync can only do this on a best-effort basis.
	//
	// Default value: `BEST_EFFORT`
	//
	// `BEST_EFFORT` : Attempt to preserve the per-file `Atime` value (recommended).
	//
	// `NONE` : Ignore `Atime` .
	//
	// > If `Atime` is set to `BEST_EFFORT` , `Mtime` must be set to `PRESERVE` .
	// >
	// > If `Atime` is set to `NONE` , `Mtime` must also be `NONE` .
	Atime *string `json:"atime" yaml:"atime"`
	// A value that limits the bandwidth used by AWS DataSync .
	//
	// For example, if you want AWS DataSync to use a maximum of 1 MB, set this value to `1048576` (=1024*1024).
	BytesPerSecond *float64 `json:"bytesPerSecond" yaml:"bytesPerSecond"`
	// The group ID (GID) of the file's owners.
	//
	// Default value: `INT_VALUE`
	//
	// `INT_VALUE` : Preserve the integer value of the user ID (UID) and group ID (GID) (recommended).
	//
	// `NAME` : Currently not supported.
	//
	// `NONE` : Ignore the UID and GID.
	Gid *string `json:"gid" yaml:"gid"`
	// A value that determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.
	//
	// For more information about providing a log group for DataSync, see [CloudWatchLogGroupArn](https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateTask.html#DataSync-CreateTask-request-CloudWatchLogGroupArn) . If set to `OFF` , no logs are published. `BASIC` publishes logs on errors for individual files transferred, and `TRANSFER` publishes logs for every file or object that is transferred and integrity checked.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
	// A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.
	//
	// This option is required for cases when you need to run the same task more than one time.
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Preserve original `Mtime` (recommended)
	//
	// `NONE` : Ignore `Mtime` .
	//
	// > If `Mtime` is set to `PRESERVE` , `Atime` must be set to `BEST_EFFORT` .
	// >
	// > If `Mtime` is set to `NONE` , `Atime` must also be set to `NONE` .
	Mtime *string `json:"mtime" yaml:"mtime"`
	// A value that determines whether files at the destination should be overwritten or preserved when copying files.
	//
	// If set to `NEVER` a destination file will not be replaced by a source file, even if the destination file differs from the source file. If you modify files in the destination and you sync the files, you can use this value to protect against overwriting those changes.
	//
	// Some storage classes have specific behaviors that can affect your S3 storage cost. For detailed information, see [Considerations when working with Amazon S3 storage classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) in the *AWS DataSync User Guide* .
	OverwriteMode *string `json:"overwriteMode" yaml:"overwriteMode"`
	// A value that determines which users or groups can access a file for a specific purpose, such as reading, writing, or execution of the file.
	//
	// This option should be set only for Network File System (NFS), Amazon EFS, and Amazon S3 locations. For more information about what metadata is copied by DataSync, see [Metadata Copied by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html#metadata-copied) .
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Preserve POSIX-style permissions (recommended).
	//
	// `NONE` : Ignore permissions.
	//
	// > AWS DataSync can preserve extant permissions of a source location.
	PosixPermissions *string `json:"posixPermissions" yaml:"posixPermissions"`
	// A value that specifies whether files in the destination that don't exist in the source file system are preserved.
	//
	// This option can affect your storage costs. If your task deletes objects, you might incur minimum storage duration charges for certain storage classes. For detailed information, see [Considerations when working with Amazon S3 storage classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) in the *AWS DataSync User Guide* .
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Ignore destination files that aren't present in the source (recommended).
	//
	// `REMOVE` : Delete destination files that aren't present in the source.
	PreserveDeletedFiles *string `json:"preserveDeletedFiles" yaml:"preserveDeletedFiles"`
	// A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and re-create the files with that device name and metadata on the destination.
	//
	// DataSync does not copy the contents of such devices, only the name and metadata.
	//
	// > AWS DataSync can't sync the actual contents of such devices, because they are nonterminal and don't return an end-of-file (EOF) marker.
	//
	// Default value: `NONE`
	//
	// `NONE` : Ignore special devices (recommended).
	//
	// `PRESERVE` : Preserve character and block device metadata. This option isn't currently supported for Amazon EFS.
	PreserveDevices *string `json:"preserveDevices" yaml:"preserveDevices"`
	// A value that determines which components of the SMB security descriptor are copied from source to destination objects.
	//
	// This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. For more information about how DataSync handles metadata, see [How DataSync Handles Metadata and Special Files](https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html) .
	//
	// Default value: `OWNER_DACL`
	//
	// `OWNER_DACL` : For each copied object, DataSync copies the following metadata:
	//
	// - Object owner.
	// - NTFS discretionary access control lists (DACLs), which determine whether to grant access to an object.
	//
	// When you use option, DataSync does NOT copy the NTFS system access control lists (SACLs), which are used by administrators to log attempts to access a secured object.
	//
	// `OWNER_DACL_SACL` : For each copied object, DataSync copies the following metadata:
	//
	// - Object owner.
	// - NTFS discretionary access control lists (DACLs), which determine whether to grant access to an object.
	// - NTFS system access control lists (SACLs), which are used by administrators to log attempts to access a secured object.
	//
	// Copying SACLs requires granting additional permissions to the Windows user that DataSync uses to access your SMB location. For information about choosing a user that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#SMBuser) .
	//
	// `NONE` : None of the SMB security descriptor components are copied. Destination objects are owned by the user that was provided for accessing the destination location. DACLs and SACLs are set based on the destination server’s configuration.
	SecurityDescriptorCopyFlags *string `json:"securityDescriptorCopyFlags" yaml:"securityDescriptorCopyFlags"`
	// A value that determines whether tasks should be queued before executing the tasks.
	//
	// If set to `ENABLED` , the tasks will be queued. The default is `ENABLED` .
	//
	// If you use the same agent to run multiple tasks, you can enable the tasks to run in series. For more information, see [Queueing task executions](https://docs.aws.amazon.com/datasync/latest/userguide/run-task.html#queue-task-execution) .
	TaskQueueing *string `json:"taskQueueing" yaml:"taskQueueing"`
	// A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing it to the destination location.
	//
	// `CHANGED` : DataSync copies only data or metadata that is new or different from the source location to the destination location.
	//
	// `ALL` : DataSync copies all source location content to the destination, without comparing it to existing content on the destination.
	TransferMode *string `json:"transferMode" yaml:"transferMode"`
	// The user ID (UID) of the file's owner.
	//
	// Default value: `INT_VALUE`
	//
	// `INT_VALUE` : Preserve the integer value of the UID and group ID (GID) (recommended).
	//
	// `NAME` : Currently not supported
	//
	// `NONE` : Ignore the UID and GID.
	Uid *string `json:"uid" yaml:"uid"`
	// A value that determines whether a data integrity verification is performed at the end of a task execution after all data and metadata have been transferred.
	//
	// For more information, see [Configure task settings](https://docs.aws.amazon.com/datasync/latest/userguide/create-task.html) .
	//
	// Default value: `POINT_IN_TIME_CONSISTENT`
	//
	// `ONLY_FILES_TRANSFERRED` (recommended): Perform verification only on files that were transferred.
	//
	// `POINT_IN_TIME_CONSISTENT` : Scan the entire source and entire destination at the end of the transfer to verify that the source and destination are fully synchronized. This option isn't supported when transferring to S3 Glacier or S3 Glacier Deep Archive storage classes.
	//
	// `NONE` : No additional verification is done at the end of the transfer, but all data transmissions are integrity-checked with checksum verification during the transfer.
	VerifyMode *string `json:"verifyMode" yaml:"verifyMode"`
}

// Specifies the schedule you want your task to use for repeated executions.
//
// For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) .
//
// TODO: EXAMPLE
//
type CfnTask_TaskScheduleProperty struct {
	// A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
}

// Properties for defining a `CfnTask`.
//
// TODO: EXAMPLE
//
type CfnTaskProps struct {
	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	DestinationLocationArn *string `json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// The Amazon Resource Name (ARN) of the source location for the task.
	SourceLocationArn *string `json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	//
	// For more information about how to use CloudWatch Logs with DataSync, see [Monitoring Your Task](https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html#cloudwatchlogs) in the *AWS DataSync User Guide.*
	//
	// For more information about these groups, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) in the *Amazon CloudWatch Logs User Guide* .
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// A list of filter rules that determines which files to exclude from a task.
	//
	// The list should contain a single filter string that consists of the patterns to exclude. The patterns are delimited by "|" (that is, a pipe), for example, `"/folder1|/folder2"` .
	Excludes interface{} `json:"excludes" yaml:"excludes"`
	// A list of filter rules that determines which files to include when running a task.
	//
	// The pattern contains a single filter string that consists of the patterns to include. The patterns are delimited by "|" (that is, a pipe), for example, `"/folder1|/folder2"` .
	Includes interface{} `json:"includes" yaml:"includes"`
	// The name of a task.
	//
	// This value is a text reference that is used to identify the task in the console.
	Name *string `json:"name" yaml:"name"`
	// The set of configuration options that control the behavior of a single execution of the task that occurs when you call `StartTaskExecution` .
	//
	// You can configure these options to preserve metadata such as user ID (UID) and group ID (GID), file permissions, data integrity verification, and so on.
	//
	// For each individual task execution, you can override these options by specifying the `OverrideOptions` before starting the task execution. For more information, see the [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) operation.
	Options interface{} `json:"options" yaml:"options"`
	// Specifies a schedule used to periodically transfer files from a source to a destination location.
	//
	// The schedule should be specified in UTC time. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	Schedule interface{} `json:"schedule" yaml:"schedule"`
	// The key-value pair that represents the tag that you want to add to the resource.
	//
	// The value can be an empty string.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

