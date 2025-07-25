package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
//
// To specify a VPN connection between a transit gateway and customer gateway, use the `TransitGatewayId` and `CustomerGatewayId` properties.
//
// To specify a VPN connection between a virtual private gateway and customer gateway, use the `VpnGatewayId` and `CustomerGatewayId` properties.
//
// For more information, see [AWS Site-to-Site VPN](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *AWS Site-to-Site VPN User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConnection := awscdk.Aws_ec2.NewCfnVPNConnection(this, jsii.String("MyCfnVPNConnection"), &CfnVPNConnectionProps{
//   	CustomerGatewayId: jsii.String("customerGatewayId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	EnableAcceleration: jsii.Boolean(false),
//   	LocalIpv4NetworkCidr: jsii.String("localIpv4NetworkCidr"),
//   	LocalIpv6NetworkCidr: jsii.String("localIpv6NetworkCidr"),
//   	OutsideIpAddressType: jsii.String("outsideIpAddressType"),
//   	RemoteIpv4NetworkCidr: jsii.String("remoteIpv4NetworkCidr"),
//   	RemoteIpv6NetworkCidr: jsii.String("remoteIpv6NetworkCidr"),
//   	StaticRoutesOnly: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	TransportTransitGatewayAttachmentId: jsii.String("transportTransitGatewayAttachmentId"),
//   	TunnelInsideIpVersion: jsii.String("tunnelInsideIpVersion"),
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   	VpnTunnelOptionsSpecifications: []interface{}{
//   		&VpnTunnelOptionsSpecificationProperty{
//   			DpdTimeoutAction: jsii.String("dpdTimeoutAction"),
//   			DpdTimeoutSeconds: jsii.Number(123),
//   			EnableTunnelLifecycleControl: jsii.Boolean(false),
//   			IkeVersions: []interface{}{
//   				map[string]*string{
//   					"value": jsii.String("value"),
//   				},
//   			},
//   			LogOptions: &VpnTunnelLogOptionsSpecificationProperty{
//   				CloudwatchLogOptions: &CloudwatchLogOptionsSpecificationProperty{
//   					LogEnabled: jsii.Boolean(false),
//   					LogGroupArn: jsii.String("logGroupArn"),
//   					LogOutputFormat: jsii.String("logOutputFormat"),
//   				},
//   			},
//   			Phase1DhGroupNumbers: []interface{}{
//   				&Phase1DHGroupNumbersRequestListValueProperty{
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			Phase1EncryptionAlgorithms: []interface{}{
//   				&Phase1EncryptionAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase1IntegrityAlgorithms: []interface{}{
//   				&Phase1IntegrityAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase1LifetimeSeconds: jsii.Number(123),
//   			Phase2DhGroupNumbers: []interface{}{
//   				&Phase2DHGroupNumbersRequestListValueProperty{
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			Phase2EncryptionAlgorithms: []interface{}{
//   				&Phase2EncryptionAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase2IntegrityAlgorithms: []interface{}{
//   				&Phase2IntegrityAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase2LifetimeSeconds: jsii.Number(123),
//   			PreSharedKey: jsii.String("preSharedKey"),
//   			RekeyFuzzPercentage: jsii.Number(123),
//   			RekeyMarginTimeSeconds: jsii.Number(123),
//   			ReplayWindowSize: jsii.Number(123),
//   			StartupAction: jsii.String("startupAction"),
//   			TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   			TunnelInsideIpv6Cidr: jsii.String("tunnelInsideIpv6Cidr"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html
//
type CfnVPNConnection interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The ID of the VPN connection.
	AttrVpnConnectionId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID of the customer gateway at your end of the VPN connection.
	CustomerGatewayId() *string
	SetCustomerGatewayId(val *string)
	// Indicate whether to enable acceleration for the VPN connection.
	EnableAcceleration() interface{}
	SetEnableAcceleration(val interface{})
	// The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.
	LocalIpv4NetworkCidr() *string
	SetLocalIpv4NetworkCidr(val *string)
	// The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.
	LocalIpv6NetworkCidr() *string
	SetLocalIpv6NetworkCidr(val *string)
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
	// The type of IP address assigned to the outside interface of the customer gateway device.
	OutsideIpAddressType() *string
	SetOutsideIpAddressType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IPv4 CIDR on the AWS side of the VPN connection.
	RemoteIpv4NetworkCidr() *string
	SetRemoteIpv4NetworkCidr(val *string)
	// The IPv6 CIDR on the AWS side of the VPN connection.
	RemoteIpv6NetworkCidr() *string
	SetRemoteIpv6NetworkCidr(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Indicates whether the VPN connection uses static routes only.
	StaticRoutesOnly() interface{}
	SetStaticRoutesOnly(val interface{})
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Any tags assigned to the VPN connection.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The ID of the transit gateway associated with the VPN connection.
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	// The transit gateway attachment ID to use for the VPN tunnel.
	TransportTransitGatewayAttachmentId() *string
	SetTransportTransitGatewayAttachmentId(val *string)
	// Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.
	TunnelInsideIpVersion() *string
	SetTunnelInsideIpVersion(val *string)
	// The type of VPN connection.
	Type() *string
	SetType(val *string)
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
	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	// The tunnel options for the VPN connection.
	VpnTunnelOptionsSpecifications() interface{}
	SetVpnTunnelOptionsSpecifications(val interface{})
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

// The jsii proxy struct for CfnVPNConnection
type jsiiProxy_CfnVPNConnection struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnVPNConnection) AttrVpnConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpnConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) EnableAcceleration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) LocalIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) LocalIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) OutsideIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) RemoteIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) RemoteIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) StaticRoutesOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) TransportTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) TunnelInsideIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnection) VpnTunnelOptionsSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpnTunnelOptionsSpecifications",
		&returns,
	)
	return returns
}


func NewCfnVPNConnection(scope constructs.Construct, id *string, props *CfnVPNConnectionProps) CfnVPNConnection {
	_init_.Initialize()

	if err := validateNewCfnVPNConnectionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPNConnection{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnVPNConnection_Override(c CfnVPNConnection, scope constructs.Construct, id *string, props *CfnVPNConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetCustomerGatewayId(val *string) {
	if err := j.validateSetCustomerGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGatewayId",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetEnableAcceleration(val interface{}) {
	if err := j.validateSetEnableAccelerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceleration",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetLocalIpv4NetworkCidr(val *string) {
	_jsii_.Set(
		j,
		"localIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetLocalIpv6NetworkCidr(val *string) {
	_jsii_.Set(
		j,
		"localIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetOutsideIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"outsideIpAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetRemoteIpv4NetworkCidr(val *string) {
	_jsii_.Set(
		j,
		"remoteIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetRemoteIpv6NetworkCidr(val *string) {
	_jsii_.Set(
		j,
		"remoteIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetStaticRoutesOnly(val interface{}) {
	if err := j.validateSetStaticRoutesOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticRoutesOnly",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetTransitGatewayId(val *string) {
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetTransportTransitGatewayAttachmentId(val *string) {
	_jsii_.Set(
		j,
		"transportTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetTunnelInsideIpVersion(val *string) {
	_jsii_.Set(
		j,
		"tunnelInsideIpVersion",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetVpnGatewayId(val *string) {
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

func (j *jsiiProxy_CfnVPNConnection)SetVpnTunnelOptionsSpecifications(val interface{}) {
	if err := j.validateSetVpnTunnelOptionsSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnTunnelOptionsSpecifications",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVPNConnection_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPNConnection_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnVPNConnection_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPNConnection_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
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
func CfnVPNConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPNConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPNConnection_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPNConnection) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVPNConnection) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVPNConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVPNConnection) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnVPNConnection) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnVPNConnection) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVPNConnection) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnection) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnection) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVPNConnection) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPNConnection) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnVPNConnection) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnVPNConnection) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnection) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

