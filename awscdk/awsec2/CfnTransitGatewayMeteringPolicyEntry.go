package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an entry in a transit gateway metering policy to define traffic measurement rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMeteringPolicyEntry := awscdk.Aws_ec2.NewCfnTransitGatewayMeteringPolicyEntry(this, jsii.String("MyCfnTransitGatewayMeteringPolicyEntry"), &CfnTransitGatewayMeteringPolicyEntryProps{
//   	MeteredAccount: jsii.String("meteredAccount"),
//   	PolicyRuleNumber: jsii.Number(123),
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//
//   	// the properties below are optional
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: jsii.String("destinationPortRange"),
//   	DestinationTransitGatewayAttachmentId: jsii.String("destinationTransitGatewayAttachmentId"),
//   	DestinationTransitGatewayAttachmentType: jsii.String("destinationTransitGatewayAttachmentType"),
//   	Protocol: jsii.String("protocol"),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: jsii.String("sourcePortRange"),
//   	SourceTransitGatewayAttachmentId: jsii.String("sourceTransitGatewayAttachmentId"),
//   	SourceTransitGatewayAttachmentType: jsii.String("sourceTransitGatewayAttachmentType"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html
//
type CfnTransitGatewayMeteringPolicyEntry interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsec2.ITransitGatewayMeteringPolicyEntryRef
	// The state of the metering policy entry.
	AttrState() *string
	// The date and time when the metering policy entry update becomes effective.
	AttrUpdateEffectiveAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Describes an IPv4 CIDR block.
	DestinationCidrBlock() *string
	SetDestinationCidrBlock(val *string)
	// Describes a range of ports.
	DestinationPortRange() *string
	SetDestinationPortRange(val *string)
	// The ID of the source attachment through which traffic leaves a transit gateway.
	DestinationTransitGatewayAttachmentId() *string
	SetDestinationTransitGatewayAttachmentId(val *string)
	DestinationTransitGatewayAttachmentType() *string
	SetDestinationTransitGatewayAttachmentType(val *string)
	Env() *interfaces.ResourceEnvironment
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
	// The AWS account ID to which the metered traffic is attributed.
	MeteredAccount() *string
	SetMeteredAccount(val *string)
	// The tree node.
	Node() constructs.Node
	// The rule number of the metering policy entry.
	PolicyRuleNumber() *float64
	SetPolicyRuleNumber(val *float64)
	// The protocol of the traffic.
	Protocol() *string
	SetProtocol(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Describes an IPv4 CIDR block.
	SourceCidrBlock() *string
	SetSourceCidrBlock(val *string)
	// Describes a range of ports.
	SourcePortRange() *string
	SetSourcePortRange(val *string)
	// The ID of the source attachment through which traffic enters a transit gateway.
	SourceTransitGatewayAttachmentId() *string
	SetSourceTransitGatewayAttachmentId(val *string)
	SourceTransitGatewayAttachmentType() *string
	SetSourceTransitGatewayAttachmentType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A reference to a TransitGatewayMeteringPolicyEntry resource.
	TransitGatewayMeteringPolicyEntryRef() *interfacesawsec2.TransitGatewayMeteringPolicyEntryReference
	// The ID of the transit gateway metering policy for which the entry is being created.
	TransitGatewayMeteringPolicyId() *string
	SetTransitGatewayMeteringPolicyId(val *string)
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

// The jsii proxy struct for CfnTransitGatewayMeteringPolicyEntry
type jsiiProxy_CfnTransitGatewayMeteringPolicyEntry struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsec2ITransitGatewayMeteringPolicyEntryRef
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AttrUpdateEffectiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateEffectiveAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) MeteredAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) PolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) SourceCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) TransitGatewayMeteringPolicyEntryRef() *interfacesawsec2.TransitGatewayMeteringPolicyEntryReference {
	var returns *interfacesawsec2.TransitGatewayMeteringPolicyEntryReference
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) TransitGatewayMeteringPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::TransitGatewayMeteringPolicyEntry`.
func NewCfnTransitGatewayMeteringPolicyEntry(scope constructs.Construct, id *string, props *CfnTransitGatewayMeteringPolicyEntryProps) CfnTransitGatewayMeteringPolicyEntry {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayMeteringPolicyEntryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayMeteringPolicyEntry{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::TransitGatewayMeteringPolicyEntry`.
func NewCfnTransitGatewayMeteringPolicyEntry_Override(c CfnTransitGatewayMeteringPolicyEntry, scope constructs.Construct, id *string, props *CfnTransitGatewayMeteringPolicyEntryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetDestinationCidrBlock(val *string) {
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetDestinationPortRange(val *string) {
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentId(val *string) {
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentType(val *string) {
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetMeteredAccount(val *string) {
	if err := j.validateSetMeteredAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteredAccount",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetPolicyRuleNumber(val *float64) {
	if err := j.validateSetPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetSourceCidrBlock(val *string) {
	_jsii_.Set(
		j,
		"sourceCidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetSourcePortRange(val *string) {
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetSourceTransitGatewayAttachmentId(val *string) {
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetSourceTransitGatewayAttachmentType(val *string) {
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry)SetTransitGatewayMeteringPolicyId(val *string) {
	if err := j.validateSetTransitGatewayMeteringPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayMeteringPolicyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTransitGatewayMeteringPolicyEntry_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMeteringPolicyEntry_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTransitGatewayMeteringPolicyEntry_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMeteringPolicyEntry_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnTransitGatewayMeteringPolicyEntry.
func CfnTransitGatewayMeteringPolicyEntry_IsCfnTransitGatewayMeteringPolicyEntry(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMeteringPolicyEntry_IsCfnTransitGatewayMeteringPolicyEntryParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		"isCfnTransitGatewayMeteringPolicyEntry",
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
func CfnTransitGatewayMeteringPolicyEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMeteringPolicyEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayMeteringPolicyEntry_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMeteringPolicyEntry",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntry) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

