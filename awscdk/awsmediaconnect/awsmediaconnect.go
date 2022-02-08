package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaConnect::Flow`.
//
// The AWS::MediaConnect::Flow resource defines a connection between one or more video sources and one or more outputs. For each flow, you specify the transport protocol to use, encryption information, and details for any outputs or entitlements that you want. AWS Elemental MediaConnect returns an ingest endpoint where you can send your live video as a single unicast stream. The service replicates and distributes the video to every output that you specify, whether inside or outside the AWS Cloud. You can also set up entitlements on a flow to allow other AWS accounts to access your content.
//
// TODO: EXAMPLE
//
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrFlowArn() *string
	AttrFlowAvailabilityZone() *string
	AttrSourceIngestIp() *string
	AttrSourceSourceArn() *string
	AttrSourceSourceIngestPort() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Source() interface{}
	SetSource(val interface{})
	SourceFailoverConfig() interface{}
	SetSourceFailoverConfig(val interface{})
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnFlow
type jsiiProxy_CfnFlow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlow) AttrFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrFlowAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceIngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceIngestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceSourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceFailoverConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFailoverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaConnect::Flow`.
func NewCfnFlow(scope constructs.Construct, id *string, props *CfnFlowProps) CfnFlow {
	_init_.Initialize()

	j := jsiiProxy_CfnFlow{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConnect::Flow`.
func NewCfnFlow_Override(c CfnFlow, scope constructs.Construct, id *string, props *CfnFlowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlow) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetSource(val interface{}) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnFlow) SetSourceFailoverConfig(val interface{}) {
	_jsii_.Set(
		j,
		"sourceFailoverConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
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
func CfnFlow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlow) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlow) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFlow) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFlow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the encryption of the flow.
//
// TODO: EXAMPLE
//
type CfnFlow_EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	Algorithm *string `json:"algorithm" yaml:"algorithm"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content.
	//
	// This parameter is not valid for static key encryption.
	ConstantInitializationVector *string `json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	DeviceId *string `json:"deviceId" yaml:"deviceId"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ).
	KeyType *string `json:"keyType" yaml:"keyType"`
	// The AWS Region that the API Gateway proxy endpoint was created in.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Region *string `json:"region" yaml:"region"`
	// An identifier for the content.
	//
	// The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The URL from the API Gateway proxy that you set up to talk to your key server.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Url *string `json:"url" yaml:"url"`
}

// The settings for source failover.
//
// TODO: EXAMPLE
//
type CfnFlow_FailoverConfigProperty struct {
	// The size of the buffer (delay) that the service maintains.
	//
	// A larger buffer means a longer delay in transmitting the stream, but more room for error correction. A smaller buffer means a shorter delay, but less room for error correction.
	RecoveryWindow *float64 `json:"recoveryWindow" yaml:"recoveryWindow"`
	// The state of source failover on the flow.
	//
	// If the state is disabled, the flow can have only one source. If the state is enabled, the flow can have one or two sources.
	State *string `json:"state" yaml:"state"`
}

// The details of the sources of the flow.
//
// If you are creating a flow with a VPC source, you must first create the flow with a temporary standard source by doing the following:
//
// - Use CloudFormation to create a flow with a standard source that uses the flow’s public IP address.
// - Use CloudFormation to create the VPC interface to add to this flow. This can also be done as part of the previous step.
// - After CloudFormation has created the flow and the VPC interface, update the source to point to the VPC interface that you created.
//
// TODO: EXAMPLE
//
type CfnFlow_SourceProperty struct {
	// The type of encryption that is used on the content ingested from the source.
	Decryption interface{} `json:"decryption" yaml:"decryption"`
	// A description of the source.
	//
	// This description is not visible outside of the current AWS account.
	Description *string `json:"description" yaml:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account.
	//
	// The entitlement is set by the content originator and the ARN is generated as part of the originator’s flow.
	EntitlementArn *string `json:"entitlementArn" yaml:"entitlementArn"`
	// The IP address that the flow listens on for incoming content.
	IngestIp *string `json:"ingestIp" yaml:"ingestIp"`
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	IngestPort *float64 `json:"ingestPort" yaml:"ingestPort"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *float64 `json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds for a RIST or Zixi-based source.
	MaxLatency *float64 `json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	MinLatency *float64 `json:"minLatency" yaml:"minLatency"`
	// The name of the source.
	Name *string `json:"name" yaml:"name"`
	// The protocol that is used by the source.
	//
	// For a full list of available protocols, see: [Source protocols](https://docs.aws.amazon.com/mediaconnect/latest/api/v1-flows-flowarn-source.html#v1-flows-flowarn-source-prop-setsourcerequest-protocol) in the *AWS Elemental MediaConnect API Reference* .
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The ARN of the source.
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
	// The port that the flow will be listening on for incoming content.
	SourceIngestPort *string `json:"sourceIngestPort" yaml:"sourceIngestPort"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `json:"streamId" yaml:"streamId"`
	// The name of the VPC interface that the source content comes from.
	VpcInterfaceName *string `json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that are allowed to contribute content to your source.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `json:"whitelistCidr" yaml:"whitelistCidr"`
}

// A CloudFormation `AWS::MediaConnect::FlowEntitlement`.
//
// The AWS::MediaConnect::FlowEntitlement resource defines the permission that an AWS account grants to another AWS account to allow access to the content in a specific AWS Elemental MediaConnect flow. The content originator grants an entitlement to a specific AWS account (the subscriber). When an entitlement is granted, the subscriber can create a flow using the originator's flow as the source. Each flow can have up to 50 entitlements.
//
// TODO: EXAMPLE
//
type CfnFlowEntitlement interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrEntitlementArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataTransferSubscriberFeePercent() *float64
	SetDataTransferSubscriberFeePercent(val *float64)
	Description() *string
	SetDescription(val *string)
	Encryption() interface{}
	SetEncryption(val interface{})
	EntitlementStatus() *string
	SetEntitlementStatus(val *string)
	FlowArn() *string
	SetFlowArn(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Subscribers() *[]*string
	SetSubscribers(val *[]*string)
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

// The jsii proxy struct for CfnFlowEntitlement
type jsiiProxy_CfnFlowEntitlement struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlowEntitlement) AttrEntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEntitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) DataTransferSubscriberFeePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataTransferSubscriberFeePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Encryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) EntitlementStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) Subscribers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subscribers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowEntitlement) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaConnect::FlowEntitlement`.
func NewCfnFlowEntitlement(scope constructs.Construct, id *string, props *CfnFlowEntitlementProps) CfnFlowEntitlement {
	_init_.Initialize()

	j := jsiiProxy_CfnFlowEntitlement{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConnect::FlowEntitlement`.
func NewCfnFlowEntitlement_Override(c CfnFlowEntitlement, scope constructs.Construct, id *string, props *CfnFlowEntitlementProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetDataTransferSubscriberFeePercent(val *float64) {
	_jsii_.Set(
		j,
		"dataTransferSubscriberFeePercent",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetEntitlementStatus(val *string) {
	_jsii_.Set(
		j,
		"entitlementStatus",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetFlowArn(val *string) {
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlowEntitlement) SetSubscribers(val *[]*string) {
	_jsii_.Set(
		j,
		"subscribers",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlowEntitlement_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlowEntitlement_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
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
func CfnFlowEntitlement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowEntitlement_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowEntitlement",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFlowEntitlement) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlowEntitlement) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlowEntitlement) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowEntitlement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFlowEntitlement) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFlowEntitlement) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowEntitlement) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlowEntitlement) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlowEntitlement) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlowEntitlement) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFlowEntitlement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowEntitlement) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFlowEntitlement) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFlowEntitlement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowEntitlement) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the encryption of the flow.
//
// TODO: EXAMPLE
//
type CfnFlowEntitlement_EncryptionProperty struct {
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	Algorithm *string `json:"algorithm" yaml:"algorithm"`
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content.
	//
	// This parameter is not valid for static key encryption.
	ConstantInitializationVector *string `json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	DeviceId *string `json:"deviceId" yaml:"deviceId"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ).
	KeyType *string `json:"keyType" yaml:"keyType"`
	// The AWS Region that the API Gateway proxy endpoint was created in.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Region *string `json:"region" yaml:"region"`
	// An identifier for the content.
	//
	// The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The URL from the API Gateway proxy that you set up to talk to your key server.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Url *string `json:"url" yaml:"url"`
}

// Properties for defining a `CfnFlowEntitlement`.
//
// TODO: EXAMPLE
//
type CfnFlowEntitlementProps struct {
	// A description of the entitlement.
	//
	// This description appears only on the MediaConnect console and is not visible outside of the current AWS account.
	Description *string `json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `json:"flowArn" yaml:"flowArn"`
	// The name of the entitlement.
	//
	// This value must be unique within the current flow.
	Name *string `json:"name" yaml:"name"`
	// The AWS account IDs that you want to share your content with.
	//
	// The receiving accounts (subscribers) will be allowed to create their own flows using your content as the source.
	Subscribers *[]*string `json:"subscribers" yaml:"subscribers"`
	// The percentage of the entitlement data transfer fee that you want the subscriber to be responsible for.
	DataTransferSubscriberFeePercent *float64 `json:"dataTransferSubscriberFeePercent" yaml:"dataTransferSubscriberFeePercent"`
	// The type of encryption that MediaConnect will use on the output that is associated with the entitlement.
	Encryption interface{} `json:"encryption" yaml:"encryption"`
	// An indication of whether the new entitlement should be enabled or disabled as soon as it is created.
	//
	// If you don’t specify the entitlementStatus field in your request, MediaConnect sets it to ENABLED.
	EntitlementStatus *string `json:"entitlementStatus" yaml:"entitlementStatus"`
}

// A CloudFormation `AWS::MediaConnect::FlowOutput`.
//
// The AWS::MediaConnect::FlowOutput resource defines the destination address, protocol, and port that AWS Elemental MediaConnect sends the ingested video to. Each flow can have up to 50 outputs. An output can have the same protocol or a different protocol from the source.
//
// TODO: EXAMPLE
//
type CfnFlowOutput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrOutputArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CidrAllowList() *[]*string
	SetCidrAllowList(val *[]*string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	Destination() *string
	SetDestination(val *string)
	Encryption() interface{}
	SetEncryption(val interface{})
	FlowArn() *string
	SetFlowArn(val *string)
	LogicalId() *string
	MaxLatency() *float64
	SetMaxLatency(val *float64)
	MinLatency() *float64
	SetMinLatency(val *float64)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	Protocol() *string
	SetProtocol(val *string)
	Ref() *string
	RemoteId() *string
	SetRemoteId(val *string)
	SmoothingLatency() *float64
	SetSmoothingLatency(val *float64)
	Stack() awscdk.Stack
	StreamId() *string
	SetStreamId(val *string)
	UpdatedProperites() *map[string]interface{}
	VpcInterfaceAttachment() interface{}
	SetVpcInterfaceAttachment(val interface{})
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

// The jsii proxy struct for CfnFlowOutput
type jsiiProxy_CfnFlowOutput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlowOutput) AttrOutputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOutputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) CidrAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Encryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) RemoteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) SmoothingLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"smoothingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutput) VpcInterfaceAttachment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInterfaceAttachment",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaConnect::FlowOutput`.
func NewCfnFlowOutput(scope constructs.Construct, id *string, props *CfnFlowOutputProps) CfnFlowOutput {
	_init_.Initialize()

	j := jsiiProxy_CfnFlowOutput{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConnect::FlowOutput`.
func NewCfnFlowOutput_Override(c CfnFlowOutput, scope constructs.Construct, id *string, props *CfnFlowOutputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetCidrAllowList(val *[]*string) {
	_jsii_.Set(
		j,
		"cidrAllowList",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetFlowArn(val *string) {
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetMaxLatency(val *float64) {
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetMinLatency(val *float64) {
	_jsii_.Set(
		j,
		"minLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetRemoteId(val *string) {
	_jsii_.Set(
		j,
		"remoteId",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetSmoothingLatency(val *float64) {
	_jsii_.Set(
		j,
		"smoothingLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetStreamId(val *string) {
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_CfnFlowOutput) SetVpcInterfaceAttachment(val interface{}) {
	_jsii_.Set(
		j,
		"vpcInterfaceAttachment",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlowOutput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlowOutput_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
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
func CfnFlowOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowOutput_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowOutput",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFlowOutput) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlowOutput) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlowOutput) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFlowOutput) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFlowOutput) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlowOutput) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlowOutput) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlowOutput) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFlowOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowOutput) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFlowOutput) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFlowOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowOutput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the encryption of the flow.
//
// TODO: EXAMPLE
//
type CfnFlowOutput_EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	Algorithm *string `json:"algorithm" yaml:"algorithm"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ).
	KeyType *string `json:"keyType" yaml:"keyType"`
}

// The VPC interface that you want to send your output to.
//
// TODO: EXAMPLE
//
type CfnFlowOutput_VpcInterfaceAttachmentProperty struct {
	// The name of the VPC interface that you want to send your output to.
	VpcInterfaceName *string `json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

// Properties for defining a `CfnFlowOutput`.
//
// TODO: EXAMPLE
//
type CfnFlowOutputProps struct {
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `json:"flowArn" yaml:"flowArn"`
	// The protocol to use for the output.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The range of IP addresses that are allowed to initiate output requests to this flow.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList *[]*string `json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// This description is not visible outside of the current AWS account even if the account grants entitlements to other accounts.
	Description *string `json:"description" yaml:"description"`
	// The IP address where you want to send the output.
	Destination *string `json:"destination" yaml:"destination"`
	// The encryption credentials that you want to use for the output.
	Encryption interface{} `json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *float64 `json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	MinLatency *float64 `json:"minLatency" yaml:"minLatency"`
	// The name of the VPC interface.
	Name *string `json:"name" yaml:"name"`
	// The port to use when MediaConnect distributes content to the output.
	Port *float64 `json:"port" yaml:"port"`
	// The identifier that is assigned to the Zixi receiver.
	//
	// This parameter applies only to outputs that use Zixi pull.
	RemoteId *string `json:"remoteId" yaml:"remoteId"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *float64 `json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `json:"streamId" yaml:"streamId"`
	// The VPC interface that you want to send your output to.
	VpcInterfaceAttachment interface{} `json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

// Properties for defining a `CfnFlow`.
//
// TODO: EXAMPLE
//
type CfnFlowProps struct {
	// The name of the flow.
	Name *string `json:"name" yaml:"name"`
	// The settings for the source that you want to use for the new flow.
	Source interface{} `json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS Region.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The settings for source failover.
	SourceFailoverConfig interface{} `json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

// A CloudFormation `AWS::MediaConnect::FlowSource`.
//
// The AWS::MediaConnect::FlowSource resource is the external video content that includes configuration information (encryption and source type) and a network address. Each flow has at least one source. A standard source comes from a source other than another AWS Elemental MediaConnect flow, such as an on-premises encoder. An entitled source comes from a MediaConnect flow that is owned by another AWS account and has granted an entitlement to your account.
//
// Note: MediaConnect does not currently support using CloudFormation to add sources that use the SRT-listener protocol.
//
// TODO: EXAMPLE
//
type CfnFlowSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrIngestIp() *string
	AttrSourceArn() *string
	AttrSourceIngestPort() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Decryption() interface{}
	SetDecryption(val interface{})
	Description() *string
	SetDescription(val *string)
	EntitlementArn() *string
	SetEntitlementArn(val *string)
	FlowArn() *string
	SetFlowArn(val *string)
	IngestPort() *float64
	SetIngestPort(val *float64)
	LogicalId() *string
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	MaxLatency() *float64
	SetMaxLatency(val *float64)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Protocol() *string
	SetProtocol(val *string)
	Ref() *string
	Stack() awscdk.Stack
	StreamId() *string
	SetStreamId(val *string)
	UpdatedProperites() *map[string]interface{}
	VpcInterfaceName() *string
	SetVpcInterfaceName(val *string)
	WhitelistCidr() *string
	SetWhitelistCidr(val *string)
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


// Create a new `AWS::MediaConnect::FlowSource`.
func NewCfnFlowSource(scope constructs.Construct, id *string, props *CfnFlowSourceProps) CfnFlowSource {
	_init_.Initialize()

	j := jsiiProxy_CfnFlowSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConnect::FlowSource`.
func NewCfnFlowSource_Override(c CfnFlowSource, scope constructs.Construct, id *string, props *CfnFlowSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetDecryption(val interface{}) {
	_jsii_.Set(
		j,
		"decryption",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetEntitlementArn(val *string) {
	_jsii_.Set(
		j,
		"entitlementArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetFlowArn(val *string) {
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetIngestPort(val *float64) {
	_jsii_.Set(
		j,
		"ingestPort",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetMaxBitrate(val *float64) {
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetMaxLatency(val *float64) {
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetStreamId(val *string) {
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetVpcInterfaceName(val *string) {
	_jsii_.Set(
		j,
		"vpcInterfaceName",
		val,
	)
}

func (j *jsiiProxy_CfnFlowSource) SetWhitelistCidr(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlowSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowSource",
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
func CfnFlowSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFlowSource) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlowSource) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlowSource) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFlowSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFlowSource) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlowSource) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlowSource) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlowSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFlowSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the encryption of the flow.
//
// TODO: EXAMPLE
//
type CfnFlowSource_EncryptionProperty struct {
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	Algorithm *string `json:"algorithm" yaml:"algorithm"`
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content.
	//
	// This parameter is not valid for static key encryption.
	ConstantInitializationVector *string `json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	DeviceId *string `json:"deviceId" yaml:"deviceId"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ).
	KeyType *string `json:"keyType" yaml:"keyType"`
	// The AWS Region that the API Gateway proxy endpoint was created in.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Region *string `json:"region" yaml:"region"`
	// An identifier for the content.
	//
	// The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The URL from the API Gateway proxy that you set up to talk to your key server.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	Url *string `json:"url" yaml:"url"`
}

// Properties for defining a `CfnFlowSource`.
//
// TODO: EXAMPLE
//
type CfnFlowSourceProps struct {
	// A description of the source.
	//
	// This description is not visible outside of the current AWS account.
	Description *string `json:"description" yaml:"description"`
	// The name of the source.
	Name *string `json:"name" yaml:"name"`
	// The type of encryption that is used on the content ingested from the source.
	Decryption interface{} `json:"decryption" yaml:"decryption"`
	// The ARN of the entitlement that allows you to subscribe to the flow.
	//
	// The entitlement is set by the content originator, and the ARN is generated as part of the originator's flow.
	EntitlementArn *string `json:"entitlementArn" yaml:"entitlementArn"`
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `json:"flowArn" yaml:"flowArn"`
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	IngestPort *float64 `json:"ingestPort" yaml:"ingestPort"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *float64 `json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based, Zixi-based, and Fujitsu-based streams.
	MaxLatency *float64 `json:"maxLatency" yaml:"maxLatency"`
	// The protocol that the source uses to deliver the content to MediaConnect.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `json:"streamId" yaml:"streamId"`
	// The name of the VPC interface that you want to send your output to.
	VpcInterfaceName *string `json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that are allowed to contribute content to your source.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `json:"whitelistCidr" yaml:"whitelistCidr"`
}

// A CloudFormation `AWS::MediaConnect::FlowVpcInterface`.
//
// The AWS::MediaConnect::FlowVpcInterface resource is a connection between your AWS Elemental MediaConnect flow and a virtual private cloud (VPC) that you created using the Amazon Virtual Private Cloud service.
//
// To avoid streaming your content over the public internet, you can add up to two VPC interfaces to your flow and use those connections to transfer content between your VPC and MediaConnect.
//
// You can update an existing flow to add a VPC interface. If you haven’t created the flow yet, you must create the flow with a temporary standard source by doing the following:
//
// - Use CloudFormation to create a flow with a standard source that uses to the flow’s public IP address.
// - Use CloudFormation to create a VPC interface to add to this flow. This can also be done as part of the previous step.
// - After CloudFormation has created the flow and the VPC interface, update the source to point to the VPC interface that you created.
//
// TODO: EXAMPLE
//
type CfnFlowVpcInterface interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrNetworkInterfaceIds() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	FlowArn() *string
	SetFlowArn(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	Stack() awscdk.Stack
	SubnetId() *string
	SetSubnetId(val *string)
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

// The jsii proxy struct for CfnFlowVpcInterface
type jsiiProxy_CfnFlowVpcInterface struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlowVpcInterface) AttrNetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNetworkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowVpcInterface) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaConnect::FlowVpcInterface`.
func NewCfnFlowVpcInterface(scope constructs.Construct, id *string, props *CfnFlowVpcInterfaceProps) CfnFlowVpcInterface {
	_init_.Initialize()

	j := jsiiProxy_CfnFlowVpcInterface{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConnect::FlowVpcInterface`.
func NewCfnFlowVpcInterface_Override(c CfnFlowVpcInterface, scope constructs.Construct, id *string, props *CfnFlowVpcInterfaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowVpcInterface) SetFlowArn(val *string) {
	_jsii_.Set(
		j,
		"flowArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowVpcInterface) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlowVpcInterface) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowVpcInterface) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnFlowVpcInterface) SetSubnetId(val *string) {
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
func CfnFlowVpcInterface_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlowVpcInterface_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
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
func CfnFlowVpcInterface_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowVpcInterface_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlowVpcInterface",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFlowVpcInterface) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFlowVpcInterface) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFlowVpcInterface) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowVpcInterface) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFlowVpcInterface) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFlowVpcInterface) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFlowVpcInterface) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFlowVpcInterface) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFlowVpcInterface) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFlowVpcInterface) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFlowVpcInterface) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowVpcInterface) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFlowVpcInterface) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFlowVpcInterface) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowVpcInterface) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnFlowVpcInterface`.
//
// TODO: EXAMPLE
//
type CfnFlowVpcInterfaceProps struct {
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `json:"flowArn" yaml:"flowArn"`
	// The name of the VPC Interface.
	//
	// This value must be unique within the current flow.
	Name *string `json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that you created when you set up MediaConnect as a trusted service.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The VPC security groups that you want MediaConnect to use for your VPC configuration.
	//
	// You must include at least one security group in the request.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet IDs that you want to use for your VPC interface.
	//
	// A range of IP addresses in your VPC. When you create your VPC, you specify a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. This is the primary CIDR block for your VPC. When you create a subnet for your VPC, you specify the CIDR block for the subnet, which is a subset of the VPC CIDR block.
	//
	// The subnets that you use across all VPC interfaces on the flow must be in the same Availability Zone as the flow.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
}

