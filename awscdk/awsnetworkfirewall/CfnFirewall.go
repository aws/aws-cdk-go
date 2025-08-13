package awsnetworkfirewall

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the firewall to provide stateful, managed, network firewall and intrusion detection and prevention filtering for your VPCs in Amazon VPC .
//
// The firewall defines the configuration settings for an AWS Network Firewall firewall. The settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall AWS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewall := awscdk.Aws_networkfirewall.NewCfnFirewall(this, jsii.String("MyCfnFirewall"), &CfnFirewallProps{
//   	FirewallName: jsii.String("firewallName"),
//   	FirewallPolicyArn: jsii.String("firewallPolicyArn"),
//
//   	// the properties below are optional
//   	AvailabilityZoneChangeProtection: jsii.Boolean(false),
//   	AvailabilityZoneMappings: []interface{}{
//   		&AvailabilityZoneMappingProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   		},
//   	},
//   	DeleteProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnabledAnalysisTypes: []*string{
//   		jsii.String("enabledAnalysisTypes"),
//   	},
//   	FirewallPolicyChangeProtection: jsii.Boolean(false),
//   	SubnetChangeProtection: jsii.Boolean(false),
//   	SubnetMappings: []interface{}{
//   		&SubnetMappingProperty{
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			IpAddressType: jsii.String("ipAddressType"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html
//
type CfnFirewall interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The unique IDs of the firewall endpoints for all of the subnets that you attached to the firewall.
	//
	// The subnets are not listed in any particular order. For example: `["us-west-2c:vpce-111122223333", "us-west-2a:vpce-987654321098", "us-west-2b:vpce-012345678901"]` .
	AttrEndpointIds() *[]*string
	// The Amazon Resource Name (ARN) of the firewall.
	AttrFirewallArn() *string
	// The name of the firewallresource.
	AttrFirewallId() *string
	// A setting indicating whether the firewall is protected against changes to its Availability Zone configuration.
	AvailabilityZoneChangeProtection() interface{}
	SetAvailabilityZoneChangeProtection(val interface{})
	// The Availability Zones where the firewall endpoints are created for a transit gateway-attached firewall.
	AvailabilityZoneMappings() interface{}
	SetAvailabilityZoneMappings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A flag indicating whether it is possible to delete the firewall.
	DeleteProtection() interface{}
	SetDeleteProtection(val interface{})
	// A description of the firewall.
	Description() *string
	SetDescription(val *string)
	// An optional setting indicating the specific traffic analysis types to enable on the firewall.
	EnabledAnalysisTypes() *[]*string
	SetEnabledAnalysisTypes(val *[]*string)
	// The descriptive name of the firewall.
	FirewallName() *string
	SetFirewallName(val *string)
	// The Amazon Resource Name (ARN) of the firewall policy.
	FirewallPolicyArn() *string
	SetFirewallPolicyArn(val *string)
	// A setting indicating whether the firewall is protected against a change to the firewall policy association.
	FirewallPolicyChangeProtection() interface{}
	SetFirewallPolicyChangeProtection(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A setting indicating whether the firewall is protected against changes to the subnet associations.
	SubnetChangeProtection() interface{}
	SetSubnetChangeProtection(val interface{})
	// The primary public subnets that Network Firewall is using for the firewall.
	SubnetMappings() interface{}
	SetSubnetMappings(val interface{})
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An array of key-value pairs to apply to this resource.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The unique identifier of the transit gateway associated with this firewall.
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
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
	// The unique identifier of the VPC where the firewall is in use.
	VpcId() *string
	SetVpcId(val *string)
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

// The jsii proxy struct for CfnFirewall
type jsiiProxy_CfnFirewall struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnFirewall) AttrEndpointIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEndpointIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) AttrFirewallArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFirewallArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) AttrFirewallId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFirewallId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) AvailabilityZoneChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) AvailabilityZoneMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) DeleteProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) EnabledAnalysisTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledAnalysisTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) FirewallName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) FirewallPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) FirewallPolicyChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallPolicyChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) SubnetChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) SubnetMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewall) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


func NewCfnFirewall(scope constructs.Construct, id *string, props *CfnFirewallProps) CfnFirewall {
	_init_.Initialize()

	if err := validateNewCfnFirewallParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewall{}

	_jsii_.Create(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnFirewall_Override(c CfnFirewall, scope constructs.Construct, id *string, props *CfnFirewallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFirewall)SetAvailabilityZoneChangeProtection(val interface{}) {
	if err := j.validateSetAvailabilityZoneChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneChangeProtection",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetAvailabilityZoneMappings(val interface{}) {
	if err := j.validateSetAvailabilityZoneMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneMappings",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetDeleteProtection(val interface{}) {
	if err := j.validateSetDeleteProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteProtection",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetEnabledAnalysisTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"enabledAnalysisTypes",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetFirewallName(val *string) {
	if err := j.validateSetFirewallNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallName",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetFirewallPolicyArn(val *string) {
	if err := j.validateSetFirewallPolicyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyArn",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetFirewallPolicyChangeProtection(val interface{}) {
	if err := j.validateSetFirewallPolicyChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyChangeProtection",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetSubnetChangeProtection(val interface{}) {
	if err := j.validateSetSubnetChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetChangeProtection",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetSubnetMappings(val interface{}) {
	if err := j.validateSetSubnetMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetMappings",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetTransitGatewayId(val *string) {
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_CfnFirewall)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFirewall_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewall_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnFirewall_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewall_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
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
func CfnFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewall_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_networkfirewall.CfnFirewall",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewall) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFirewall) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewall) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewall) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFirewall) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFirewall) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFirewall) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFirewall) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFirewall) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFirewall) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFirewall) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFirewall) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewall) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewall) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewall) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewall) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFirewall) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFirewall) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewall) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

