package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::MediaConnect::FlowSource resource is used to add additional sources to an existing flow.
//
// Adding an additional source requires Failover to be enabled. When you enable Failover, the additional source must use the same protocol as the existing source. A source is the external video content that includes configuration information (encryption and source type) and a network address. Each flow has at least one source. A standard source comes from a source other than another AWS Elemental MediaConnect flow, such as an on-premises encoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowSource := awscdk.Aws_mediaconnect.NewCfnFlowSource(this, jsii.String("MyCfnFlowSource"), &CfnFlowSourceProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Decryption: &EncryptionProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		Algorithm: jsii.String("algorithm"),
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		DeviceId: jsii.String("deviceId"),
//   		KeyType: jsii.String("keyType"),
//   		Region: jsii.String("region"),
//   		ResourceId: jsii.String("resourceId"),
//   		SecretArn: jsii.String("secretArn"),
//   		Url: jsii.String("url"),
//   	},
//   	EntitlementArn: jsii.String("entitlementArn"),
//   	FlowArn: jsii.String("flowArn"),
//   	GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   		BridgeArn: jsii.String("bridgeArn"),
//
//   		// the properties below are optional
//   		VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	IngestPort: jsii.Number(123),
//   	MaxBitrate: jsii.Number(123),
//   	MaxLatency: jsii.Number(123),
//   	MinLatency: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SenderControlPort: jsii.Number(123),
//   	SenderIpAddress: jsii.String("senderIpAddress"),
//   	SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   	SourceListenerPort: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	WhitelistCidr: jsii.String("whitelistCidr"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html
//
type CfnFlowSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The IP address that the flow listens on for incoming content.
	AttrIngestIp() *string
	// The ARN of the source.
	AttrSourceArn() *string
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	AttrSourceIngestPort() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The type of encryption that is used on the content ingested from the source.
	Decryption() interface{}
	SetDecryption(val interface{})
	// A description of the source.
	Description() *string
	SetDescription(val *string)
	// The ARN of the entitlement that allows you to subscribe to the flow.
	EntitlementArn() *string
	SetEntitlementArn(val *string)
	// The Amazon Resource Name (ARN) of the flow this source is connected to.
	FlowArn() *string
	SetFlowArn(val *string)
	// The source configuration for cloud flows receiving a stream from a bridge.
	GatewayBridgeSource() interface{}
	SetGatewayBridgeSource(val interface{})
	// The port that the flow listens on for incoming content.
	IngestPort() *float64
	SetIngestPort(val *float64)
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
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	// The maximum latency in milliseconds.
	MaxLatency() *float64
	SetMaxLatency(val *float64)
	// The minimum latency in milliseconds for SRT-based streams.
	MinLatency() *float64
	SetMinLatency(val *float64)
	// The name of the source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The protocol that the source uses to deliver the content to MediaConnect.
	Protocol() *string
	SetProtocol(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The port that the flow uses to send outbound requests to initiate connection with the sender.
	SenderControlPort() *float64
	SetSenderControlPort(val *float64)
	// The IP address that the flow communicates with to initiate connection with the sender.
	SenderIpAddress() *string
	SetSenderIpAddress(val *string)
	// Source IP or domain name for SRT-caller protocol.
	SourceListenerAddress() *string
	SetSourceListenerAddress(val *string)
	// Source port for SRT-caller protocol.
	SourceListenerPort() *float64
	SetSourceListenerPort(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The stream ID that you want to use for this transport.
	StreamId() *string
	SetStreamId(val *string)
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
	// The name of the VPC interface that you want to send your output to.
	VpcInterfaceName() *string
	SetVpcInterfaceName(val *string)
	// The range of IP addresses that are allowed to contribute content to your source.
	WhitelistCidr() *string
	SetWhitelistCidr(val *string)
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

// The jsii proxy struct for CfnFlowSource
type jsiiProxy_CfnFlowSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlowSource) AttrIngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIngestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) AttrSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) AttrSourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Decryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) EntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) GatewayBridgeSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayBridgeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) IngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) SenderControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) SenderIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) SourceListenerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) SourceListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) VpcInterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSource) WhitelistCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidr",
		&returns,
	)
	return returns
}


func NewCfnFlowSource(scope constructs.Construct, id *string, props *CfnFlowSourceProps) CfnFlowSource {
	_init_.Initialize()

	if err := validateNewCfnFlowSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnFlowSource_Override(c CfnFlowSource, scope constructs.Construct, id *string, props *CfnFlowSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetDecryption(val interface{}) {
	if err := j.validateSetDecryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decryption",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetEntitlementArn(val *string) {
	_jsii_.Set(
		j,
		"entitlementArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetFlowArn(val *string) {
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetGatewayBridgeSource(val interface{}) {
	if err := j.validateSetGatewayBridgeSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayBridgeSource",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetIngestPort(val *float64) {
	_jsii_.Set(
		j,
		"ingestPort",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetMaxBitrate(val *float64) {
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetMaxLatency(val *float64) {
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetMinLatency(val *float64) {
	_jsii_.Set(
		j,
		"minLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetSenderControlPort(val *float64) {
	_jsii_.Set(
		j,
		"senderControlPort",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetSenderIpAddress(val *string) {
	_jsii_.Set(
		j,
		"senderIpAddress",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetSourceListenerAddress(val *string) {
	_jsii_.Set(
		j,
		"sourceListenerAddress",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetSourceListenerPort(val *float64) {
	_jsii_.Set(
		j,
		"sourceListenerPort",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetStreamId(val *string) {
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetVpcInterfaceName(val *string) {
	_jsii_.Set(
		j,
		"vpcInterfaceName",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource)SetWhitelistCidr(val *string) {
	_jsii_.Set(
		j,
		"whitelistCidr",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlowSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowSource_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnFlowSource_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowSource_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
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
func CfnFlowSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowSource) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlowSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlowSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFlowSource) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFlowSource) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFlowSource) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlowSource) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowSource) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowSource) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowSource) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlowSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFlowSource) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFlowSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowSource) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

