package awsnetworkmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a VPC attachment on an edge location of a core network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcAttachment := awscdk.Aws_networkmanager.NewCfnVpcAttachment(this, jsii.String("MyCfnVpcAttachment"), &CfnVpcAttachmentProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	SubnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	VpcArn: jsii.String("vpcArn"),
//
//   	// the properties below are optional
//   	Options: &VpcOptionsProperty{
//   		ApplianceModeSupport: jsii.Boolean(false),
//   		DnsSupport: jsii.Boolean(false),
//   		Ipv6Support: jsii.Boolean(false),
//   		SecurityGroupReferencingSupport: jsii.Boolean(false),
//   	},
//   	ProposedNetworkFunctionGroupChange: &ProposedNetworkFunctionGroupChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		NetworkFunctionGroupName: jsii.String("networkFunctionGroupName"),
//   		Tags: []cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ProposedSegmentChange: &ProposedSegmentChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		Tags: []*cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Tags: []*cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-vpcattachment.html
//
type CfnVpcAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The ID of the VPC attachment.
	AttrAttachmentId() *string
	// The policy rule number associated with the attachment.
	AttrAttachmentPolicyRuleNumber() *float64
	// The type of attachment.
	//
	// This will be `VPC` .
	AttrAttachmentType() *string
	// The ARN of the core network.
	AttrCoreNetworkArn() *string
	// The timestamp when the VPC attachment was created.
	AttrCreatedAt() *string
	// The Region where the core network edge is located.
	AttrEdgeLocation() *string
	// The name of the network function group.
	AttrNetworkFunctionGroupName() *string
	// The ID of the VPC attachment owner.
	AttrOwnerAccountId() *string
	// The resource ARN for the VPC attachment.
	AttrResourceArn() *string
	// The name of the attachment's segment.
	AttrSegmentName() *string
	// The state of the attachment.
	//
	// This can be: `REJECTED` | `PENDING_ATTACHMENT_ACCEPTANCE` | `CREATING` | `FAILED` | `AVAILABLE` | `UPDATING` | `PENDING_NETWORK_UPDATE` | `PENDING_TAG_ACCEPTANCE` | `DELETING` .
	AttrState() *string
	// The timestamp when the VPC attachment was last updated.
	AttrUpdatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The core network ID.
	CoreNetworkId() *string
	SetCoreNetworkId(val *string)
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
	// Options for creating the VPC attachment.
	Options() interface{}
	SetOptions(val interface{})
	// Describes proposed changes to a network function group.
	ProposedNetworkFunctionGroupChange() interface{}
	SetProposedNetworkFunctionGroupChange(val interface{})
	// Describes a proposed segment change.
	ProposedSegmentChange() interface{}
	SetProposedSegmentChange(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The subnet ARNs.
	SubnetArns() *[]*string
	SetSubnetArns(val *[]*string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags associated with the VPC attachment.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	// The ARN of the VPC attachment.
	VpcArn() *string
	SetVpcArn(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnVpcAttachment
type jsiiProxy_CfnVpcAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnVpcAttachment) AttrAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrAttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAttachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrCoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCoreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrEdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEdgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrNetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrSegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSegmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) CoreNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) Options() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) ProposedNetworkFunctionGroupChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proposedNetworkFunctionGroupChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) ProposedSegmentChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proposedSegmentChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) SubnetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcAttachment) VpcArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcArn",
		&returns,
	)
	return returns
}


func NewCfnVpcAttachment(scope constructs.Construct, id *string, props *CfnVpcAttachmentProps) CfnVpcAttachment {
	_init_.Initialize()

	if err := validateNewCfnVpcAttachmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVpcAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnVpcAttachment_Override(c CfnVpcAttachment, scope constructs.Construct, id *string, props *CfnVpcAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetCoreNetworkId(val *string) {
	if err := j.validateSetCoreNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkId",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetOptions(val interface{}) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetProposedNetworkFunctionGroupChange(val interface{}) {
	if err := j.validateSetProposedNetworkFunctionGroupChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proposedNetworkFunctionGroupChange",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetProposedSegmentChange(val interface{}) {
	if err := j.validateSetProposedSegmentChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proposedSegmentChange",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetSubnetArns(val *[]*string) {
	if err := j.validateSetSubnetArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetArns",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnVpcAttachment)SetVpcArn(val *string) {
	if err := j.validateSetVpcArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVpcAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVpcAttachment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnVpcAttachment_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVpcAttachment_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		"isCfnResource",
		[]interface{}{x},
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
func CfnVpcAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVpcAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVpcAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_networkmanager.CfnVpcAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnVpcAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcAttachment) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

