package awsmanagedblockchain

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ManagedBlockchain::Member`.
//
// Creates a member within a Managed Blockchain network.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMember := awscdk.Aws_managedblockchain.NewCfnMember(this, jsii.String("MyCfnMember"), &cfnMemberProps{
//   	memberConfiguration: &memberConfigurationProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		memberFrameworkConfiguration: &memberFrameworkConfigurationProperty{
//   			memberFabricConfiguration: &memberFabricConfigurationProperty{
//   				adminPassword: jsii.String("adminPassword"),
//   				adminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	invitationId: jsii.String("invitationId"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		framework: jsii.String("framework"),
//   		frameworkVersion: jsii.String("frameworkVersion"),
//   		name: jsii.String("name"),
//   		votingPolicy: &votingPolicyProperty{
//   			approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   				proposalDurationInHours: jsii.Number(123),
//   				thresholdComparator: jsii.String("thresholdComparator"),
//   				thresholdPercentage: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		networkFrameworkConfiguration: &networkFrameworkConfigurationProperty{
//   			networkFabricConfiguration: &networkFabricConfigurationProperty{
//   				edition: jsii.String("edition"),
//   			},
//   		},
//   	},
//   	networkId: jsii.String("networkId"),
//   })
//
type CfnMember interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier of the member.
	AttrMemberId() *string
	// The unique identifier of the network to which the member belongs.
	AttrNetworkId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The unique identifier of the invitation to join the network sent to the account that creates the member.
	InvitationId() *string
	SetInvitationId(val *string)
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
	// Configuration properties of the member.
	MemberConfiguration() interface{}
	SetMemberConfiguration(val interface{})
	// Configuration properties of the network to which the member belongs.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The unique identifier of the network to which the member belongs.
	NetworkId() *string
	SetNetworkId(val *string)
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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

// The jsii proxy struct for CfnMember
type jsiiProxy_CfnMember struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMember) AttrMemberId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMemberId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) AttrNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) InvitationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invitationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) MemberConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMember) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ManagedBlockchain::Member`.
func NewCfnMember(scope constructs.Construct, id *string, props *CfnMemberProps) CfnMember {
	_init_.Initialize()

	j := jsiiProxy_CfnMember{}

	_jsii_.Create(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ManagedBlockchain::Member`.
func NewCfnMember_Override(c CfnMember, scope constructs.Construct, id *string, props *CfnMemberProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMember) SetInvitationId(val *string) {
	_jsii_.Set(
		j,
		"invitationId",
		val,
	)
}

func (j *jsiiProxy_CfnMember) SetMemberConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"memberConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnMember) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnMember) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMember_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMember_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnMember_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMember_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_managedblockchain.CfnMember",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMember) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMember) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMember) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMember) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMember) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMember) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMember) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMember) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMember) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMember) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMember) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMember) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMember) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMember) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMember) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A policy type that defines the voting rules for the network.
//
// The rules decide if a proposal is approved. Approval may be based on criteria such as the percentage of `YES` votes and the duration of the proposal. The policy applies to all proposals and is specified when the network is created.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvalThresholdPolicyProperty := &approvalThresholdPolicyProperty{
//   	proposalDurationInHours: jsii.Number(123),
//   	thresholdComparator: jsii.String("thresholdComparator"),
//   	thresholdPercentage: jsii.Number(123),
//   }
//
type CfnMember_ApprovalThresholdPolicyProperty struct {
	// The duration from the time that a proposal is created until it expires.
	//
	// If members cast neither the required number of `YES` votes to approve the proposal nor the number of `NO` votes required to reject it before the duration expires, the proposal is `EXPIRED` and `ProposalActions` are not carried out.
	ProposalDurationInHours *float64 `field:"optional" json:"proposalDurationInHours" yaml:"proposalDurationInHours"`
	// Determines whether the vote percentage must be greater than the `ThresholdPercentage` or must be greater than or equal to the `ThreholdPercentage` to be approved.
	ThresholdComparator *string `field:"optional" json:"thresholdComparator" yaml:"thresholdComparator"`
	// The percentage of votes among all members that must be `YES` for a proposal to be approved.
	//
	// For example, a `ThresholdPercentage` value of `50` indicates 50%. The `ThresholdComparator` determines the precise comparison. If a `ThresholdPercentage` value of `50` is specified on a network with 10 members, along with a `ThresholdComparator` value of `GREATER_THAN` , this indicates that 6 `YES` votes are required for the proposal to be approved.
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

// Configuration properties of the member.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberConfigurationProperty := &memberConfigurationProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	memberFrameworkConfiguration: &memberFrameworkConfigurationProperty{
//   		memberFabricConfiguration: &memberFabricConfigurationProperty{
//   			adminPassword: jsii.String("adminPassword"),
//   			adminUsername: jsii.String("adminUsername"),
//   		},
//   	},
//   }
//
type CfnMember_MemberConfigurationProperty struct {
	// The name of the member.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of the member.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties of the blockchain framework relevant to the member.
	MemberFrameworkConfiguration interface{} `field:"optional" json:"memberFrameworkConfiguration" yaml:"memberFrameworkConfiguration"`
}

// Configuration properties for Hyperledger Fabric for a member in a Managed Blockchain network using the Hyperledger Fabric framework.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberFabricConfigurationProperty := &memberFabricConfigurationProperty{
//   	adminPassword: jsii.String("adminPassword"),
//   	adminUsername: jsii.String("adminUsername"),
//   }
//
type CfnMember_MemberFabricConfigurationProperty struct {
	// The password for the member's initial administrative user.
	//
	// The `AdminPassword` must be at least eight characters long and no more than 32 characters. It must contain at least one uppercase letter, one lowercase letter, and one digit. It cannot have a single quotation mark (‘), a double quotation marks (“), a forward slash(/), a backward slash(\), @, or a space.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// The user name for the member's initial administrative user.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
}

// Configuration properties relevant to a member for the blockchain framework that the Managed Blockchain network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberFrameworkConfigurationProperty := &memberFrameworkConfigurationProperty{
//   	memberFabricConfiguration: &memberFabricConfigurationProperty{
//   		adminPassword: jsii.String("adminPassword"),
//   		adminUsername: jsii.String("adminUsername"),
//   	},
//   }
//
type CfnMember_MemberFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric.
	MemberFabricConfiguration interface{} `field:"optional" json:"memberFabricConfiguration" yaml:"memberFabricConfiguration"`
}

// Configuration properties of the network to which the member belongs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	framework: jsii.String("framework"),
//   	frameworkVersion: jsii.String("frameworkVersion"),
//   	name: jsii.String("name"),
//   	votingPolicy: &votingPolicyProperty{
//   		approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   			proposalDurationInHours: jsii.Number(123),
//   			thresholdComparator: jsii.String("thresholdComparator"),
//   			thresholdPercentage: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	networkFrameworkConfiguration: &networkFrameworkConfigurationProperty{
//   		networkFabricConfiguration: &networkFabricConfigurationProperty{
//   			edition: jsii.String("edition"),
//   		},
//   	},
//   }
//
type CfnMember_NetworkConfigurationProperty struct {
	// The blockchain framework that the network uses.
	Framework *string `field:"required" json:"framework" yaml:"framework"`
	// The version of the blockchain framework that the network uses.
	FrameworkVersion *string `field:"required" json:"frameworkVersion" yaml:"frameworkVersion"`
	// The name of the network.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The voting rules for the network to decide if a proposal is accepted.
	VotingPolicy interface{} `field:"required" json:"votingPolicy" yaml:"votingPolicy"`
	// Attributes of the blockchain framework for the network.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties relevant to the network for the blockchain framework that the network uses.
	NetworkFrameworkConfiguration interface{} `field:"optional" json:"networkFrameworkConfiguration" yaml:"networkFrameworkConfiguration"`
}

// Hyperledger Fabric configuration properties for the network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkFabricConfigurationProperty := &networkFabricConfigurationProperty{
//   	edition: jsii.String("edition"),
//   }
//
type CfnMember_NetworkFabricConfigurationProperty struct {
	// The edition of Amazon Managed Blockchain that the network uses.
	//
	// Valid values are `standard` and `starter` . For more information, see
	Edition *string `field:"required" json:"edition" yaml:"edition"`
}

// Configuration properties relevant to the network for the blockchain framework that the network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkFrameworkConfigurationProperty := &networkFrameworkConfigurationProperty{
//   	networkFabricConfiguration: &networkFabricConfigurationProperty{
//   		edition: jsii.String("edition"),
//   	},
//   }
//
type CfnMember_NetworkFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric for a member in a Managed Blockchain network using the Hyperledger Fabric framework.
	NetworkFabricConfiguration interface{} `field:"optional" json:"networkFabricConfiguration" yaml:"networkFabricConfiguration"`
}

// The voting rules for the network to decide if a proposal is accepted.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   votingPolicyProperty := &votingPolicyProperty{
//   	approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   		proposalDurationInHours: jsii.Number(123),
//   		thresholdComparator: jsii.String("thresholdComparator"),
//   		thresholdPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnMember_VotingPolicyProperty struct {
	// Defines the rules for the network for voting on proposals, such as the percentage of `YES` votes required for the proposal to be approved and the duration of the proposal.
	//
	// The policy applies to all proposals and is specified when the network is created.
	ApprovalThresholdPolicy interface{} `field:"optional" json:"approvalThresholdPolicy" yaml:"approvalThresholdPolicy"`
}

// Properties for defining a `CfnMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberProps := &cfnMemberProps{
//   	memberConfiguration: &memberConfigurationProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		memberFrameworkConfiguration: &memberFrameworkConfigurationProperty{
//   			memberFabricConfiguration: &memberFabricConfigurationProperty{
//   				adminPassword: jsii.String("adminPassword"),
//   				adminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	invitationId: jsii.String("invitationId"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		framework: jsii.String("framework"),
//   		frameworkVersion: jsii.String("frameworkVersion"),
//   		name: jsii.String("name"),
//   		votingPolicy: &votingPolicyProperty{
//   			approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   				proposalDurationInHours: jsii.Number(123),
//   				thresholdComparator: jsii.String("thresholdComparator"),
//   				thresholdPercentage: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		networkFrameworkConfiguration: &networkFrameworkConfigurationProperty{
//   			networkFabricConfiguration: &networkFabricConfigurationProperty{
//   				edition: jsii.String("edition"),
//   			},
//   		},
//   	},
//   	networkId: jsii.String("networkId"),
//   }
//
type CfnMemberProps struct {
	// Configuration properties of the member.
	MemberConfiguration interface{} `field:"required" json:"memberConfiguration" yaml:"memberConfiguration"`
	// The unique identifier of the invitation to join the network sent to the account that creates the member.
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Configuration properties of the network to which the member belongs.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The unique identifier of the network to which the member belongs.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}

// A CloudFormation `AWS::ManagedBlockchain::Node`.
//
// Creates a node on the specified blockchain network.
//
// Applies to Hyperledger Fabric and Ethereum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNode := awscdk.Aws_managedblockchain.NewCfnNode(this, jsii.String("MyCfnNode"), &cfnNodeProps{
//   	networkId: jsii.String("networkId"),
//   	nodeConfiguration: &nodeConfigurationProperty{
//   		availabilityZone: jsii.String("availabilityZone"),
//   		instanceType: jsii.String("instanceType"),
//   	},
//
//   	// the properties below are optional
//   	memberId: jsii.String("memberId"),
//   })
//
type CfnNode interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the node.
	AttrArn() *string
	// The unique identifier of the member in which the node is created.
	//
	// Applies only to Hyperledger Fabric.
	AttrMemberId() *string
	// The unique identifier of the network that the node is in.
	AttrNetworkId() *string
	// The unique identifier of the node.
	AttrNodeId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
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
	// The unique identifier of the member to which the node belongs.
	//
	// Applies only to Hyperledger Fabric.
	MemberId() *string
	SetMemberId(val *string)
	// The unique identifier of the network for the node.
	//
	// Ethereum public networks have the following `NetworkId` s:
	//
	// - `n-ethereum-mainnet`
	// - `n-ethereum-rinkeby`
	// - `n-ethereum-ropsten`.
	NetworkId() *string
	SetNetworkId(val *string)
	// The tree node.
	Node() constructs.Node
	// Configuration properties of a peer node.
	NodeConfiguration() interface{}
	SetNodeConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

// The jsii proxy struct for CfnNode
type jsiiProxy_CfnNode struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnNode) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) AttrMemberId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMemberId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) AttrNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) AttrNodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) MemberId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) NodeConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNode) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ManagedBlockchain::Node`.
func NewCfnNode(scope constructs.Construct, id *string, props *CfnNodeProps) CfnNode {
	_init_.Initialize()

	j := jsiiProxy_CfnNode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ManagedBlockchain::Node`.
func NewCfnNode_Override(c CfnNode, scope constructs.Construct, id *string, props *CfnNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNode) SetMemberId(val *string) {
	_jsii_.Set(
		j,
		"memberId",
		val,
	)
}

func (j *jsiiProxy_CfnNode) SetNetworkId(val *string) {
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_CfnNode) SetNodeConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"nodeConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnNode_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnNode_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNode_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_managedblockchain.CfnNode",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNode) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNode) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNode) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNode) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNode) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNode) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnNode) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNode) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNode) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNode) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnNode) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNode) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNode) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration properties of a peer node within a membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeConfigurationProperty := &nodeConfigurationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	instanceType: jsii.String("instanceType"),
//   }
//
type CfnNode_NodeConfigurationProperty struct {
	// The Availability Zone in which the node exists.
	//
	// Required for Ethereum nodes.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The Amazon Managed Blockchain instance type for the node.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

// Properties for defining a `CfnNode`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNodeProps := &cfnNodeProps{
//   	networkId: jsii.String("networkId"),
//   	nodeConfiguration: &nodeConfigurationProperty{
//   		availabilityZone: jsii.String("availabilityZone"),
//   		instanceType: jsii.String("instanceType"),
//   	},
//
//   	// the properties below are optional
//   	memberId: jsii.String("memberId"),
//   }
//
type CfnNodeProps struct {
	// The unique identifier of the network for the node.
	//
	// Ethereum public networks have the following `NetworkId` s:
	//
	// - `n-ethereum-mainnet`
	// - `n-ethereum-rinkeby`
	// - `n-ethereum-ropsten`.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Configuration properties of a peer node.
	NodeConfiguration interface{} `field:"required" json:"nodeConfiguration" yaml:"nodeConfiguration"`
	// The unique identifier of the member to which the node belongs.
	//
	// Applies only to Hyperledger Fabric.
	MemberId *string `field:"optional" json:"memberId" yaml:"memberId"`
}

