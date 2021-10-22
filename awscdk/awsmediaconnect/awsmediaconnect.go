package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaConnect::Flow`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFlow_EncryptionProperty struct {
	// `CfnFlow.EncryptionProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnFlow.EncryptionProperty.Algorithm`.
	Algorithm *string `json:"algorithm"`
	// `CfnFlow.EncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnFlow.EncryptionProperty.DeviceId`.
	DeviceId *string `json:"deviceId"`
	// `CfnFlow.EncryptionProperty.KeyType`.
	KeyType *string `json:"keyType"`
	// `CfnFlow.EncryptionProperty.Region`.
	Region *string `json:"region"`
	// `CfnFlow.EncryptionProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `CfnFlow.EncryptionProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnFlow.EncryptionProperty.Url`.
	Url *string `json:"url"`
}

type CfnFlow_FailoverConfigProperty struct {
	// `CfnFlow.FailoverConfigProperty.RecoveryWindow`.
	RecoveryWindow *float64 `json:"recoveryWindow"`
	// `CfnFlow.FailoverConfigProperty.State`.
	State *string `json:"state"`
}

type CfnFlow_SourceProperty struct {
	// `CfnFlow.SourceProperty.Decryption`.
	Decryption interface{} `json:"decryption"`
	// `CfnFlow.SourceProperty.Description`.
	Description *string `json:"description"`
	// `CfnFlow.SourceProperty.EntitlementArn`.
	EntitlementArn *string `json:"entitlementArn"`
	// `CfnFlow.SourceProperty.IngestIp`.
	IngestIp *string `json:"ingestIp"`
	// `CfnFlow.SourceProperty.IngestPort`.
	IngestPort *float64 `json:"ingestPort"`
	// `CfnFlow.SourceProperty.MaxBitrate`.
	MaxBitrate *float64 `json:"maxBitrate"`
	// `CfnFlow.SourceProperty.MaxLatency`.
	MaxLatency *float64 `json:"maxLatency"`
	// `CfnFlow.SourceProperty.MinLatency`.
	MinLatency *float64 `json:"minLatency"`
	// `CfnFlow.SourceProperty.Name`.
	Name *string `json:"name"`
	// `CfnFlow.SourceProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnFlow.SourceProperty.SourceArn`.
	SourceArn *string `json:"sourceArn"`
	// `CfnFlow.SourceProperty.SourceIngestPort`.
	SourceIngestPort *string `json:"sourceIngestPort"`
	// `CfnFlow.SourceProperty.StreamId`.
	StreamId *string `json:"streamId"`
	// `CfnFlow.SourceProperty.VpcInterfaceName`.
	VpcInterfaceName *string `json:"vpcInterfaceName"`
	// `CfnFlow.SourceProperty.WhitelistCidr`.
	WhitelistCidr *string `json:"whitelistCidr"`
}

// A CloudFormation `AWS::MediaConnect::FlowEntitlement`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFlowEntitlement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFlowEntitlement) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFlowEntitlement_EncryptionProperty struct {
	// `CfnFlowEntitlement.EncryptionProperty.Algorithm`.
	Algorithm *string `json:"algorithm"`
	// `CfnFlowEntitlement.EncryptionProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnFlowEntitlement.EncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnFlowEntitlement.EncryptionProperty.DeviceId`.
	DeviceId *string `json:"deviceId"`
	// `CfnFlowEntitlement.EncryptionProperty.KeyType`.
	KeyType *string `json:"keyType"`
	// `CfnFlowEntitlement.EncryptionProperty.Region`.
	Region *string `json:"region"`
	// `CfnFlowEntitlement.EncryptionProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `CfnFlowEntitlement.EncryptionProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnFlowEntitlement.EncryptionProperty.Url`.
	Url *string `json:"url"`
}

// Properties for defining a `AWS::MediaConnect::FlowEntitlement`.
type CfnFlowEntitlementProps struct {
	// `AWS::MediaConnect::FlowEntitlement.Description`.
	Description *string `json:"description"`
	// `AWS::MediaConnect::FlowEntitlement.FlowArn`.
	FlowArn *string `json:"flowArn"`
	// `AWS::MediaConnect::FlowEntitlement.Name`.
	Name *string `json:"name"`
	// `AWS::MediaConnect::FlowEntitlement.Subscribers`.
	Subscribers *[]*string `json:"subscribers"`
	// `AWS::MediaConnect::FlowEntitlement.DataTransferSubscriberFeePercent`.
	DataTransferSubscriberFeePercent *float64 `json:"dataTransferSubscriberFeePercent"`
	// `AWS::MediaConnect::FlowEntitlement.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `AWS::MediaConnect::FlowEntitlement.EntitlementStatus`.
	EntitlementStatus *string `json:"entitlementStatus"`
}

// A CloudFormation `AWS::MediaConnect::FlowOutput`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFlowOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFlowOutput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFlowOutput_EncryptionProperty struct {
	// `CfnFlowOutput.EncryptionProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnFlowOutput.EncryptionProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnFlowOutput.EncryptionProperty.Algorithm`.
	Algorithm *string `json:"algorithm"`
	// `CfnFlowOutput.EncryptionProperty.KeyType`.
	KeyType *string `json:"keyType"`
}

type CfnFlowOutput_VpcInterfaceAttachmentProperty struct {
	// `CfnFlowOutput.VpcInterfaceAttachmentProperty.VpcInterfaceName`.
	VpcInterfaceName *string `json:"vpcInterfaceName"`
}

// Properties for defining a `AWS::MediaConnect::FlowOutput`.
type CfnFlowOutputProps struct {
	// `AWS::MediaConnect::FlowOutput.FlowArn`.
	FlowArn *string `json:"flowArn"`
	// `AWS::MediaConnect::FlowOutput.Protocol`.
	Protocol *string `json:"protocol"`
	// `AWS::MediaConnect::FlowOutput.CidrAllowList`.
	CidrAllowList *[]*string `json:"cidrAllowList"`
	// `AWS::MediaConnect::FlowOutput.Description`.
	Description *string `json:"description"`
	// `AWS::MediaConnect::FlowOutput.Destination`.
	Destination *string `json:"destination"`
	// `AWS::MediaConnect::FlowOutput.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `AWS::MediaConnect::FlowOutput.MaxLatency`.
	MaxLatency *float64 `json:"maxLatency"`
	// `AWS::MediaConnect::FlowOutput.MinLatency`.
	MinLatency *float64 `json:"minLatency"`
	// `AWS::MediaConnect::FlowOutput.Name`.
	Name *string `json:"name"`
	// `AWS::MediaConnect::FlowOutput.Port`.
	Port *float64 `json:"port"`
	// `AWS::MediaConnect::FlowOutput.RemoteId`.
	RemoteId *string `json:"remoteId"`
	// `AWS::MediaConnect::FlowOutput.SmoothingLatency`.
	SmoothingLatency *float64 `json:"smoothingLatency"`
	// `AWS::MediaConnect::FlowOutput.StreamId`.
	StreamId *string `json:"streamId"`
	// `AWS::MediaConnect::FlowOutput.VpcInterfaceAttachment`.
	VpcInterfaceAttachment interface{} `json:"vpcInterfaceAttachment"`
}

// Properties for defining a `AWS::MediaConnect::Flow`.
type CfnFlowProps struct {
	// `AWS::MediaConnect::Flow.Name`.
	Name *string `json:"name"`
	// `AWS::MediaConnect::Flow.Source`.
	Source interface{} `json:"source"`
	// `AWS::MediaConnect::Flow.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::MediaConnect::Flow.SourceFailoverConfig`.
	SourceFailoverConfig interface{} `json:"sourceFailoverConfig"`
}

// A CloudFormation `AWS::MediaConnect::FlowSource`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFlowSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFlowSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFlowSource_EncryptionProperty struct {
	// `CfnFlowSource.EncryptionProperty.Algorithm`.
	Algorithm *string `json:"algorithm"`
	// `CfnFlowSource.EncryptionProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnFlowSource.EncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnFlowSource.EncryptionProperty.DeviceId`.
	DeviceId *string `json:"deviceId"`
	// `CfnFlowSource.EncryptionProperty.KeyType`.
	KeyType *string `json:"keyType"`
	// `CfnFlowSource.EncryptionProperty.Region`.
	Region *string `json:"region"`
	// `CfnFlowSource.EncryptionProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `CfnFlowSource.EncryptionProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnFlowSource.EncryptionProperty.Url`.
	Url *string `json:"url"`
}

// Properties for defining a `AWS::MediaConnect::FlowSource`.
type CfnFlowSourceProps struct {
	// `AWS::MediaConnect::FlowSource.Description`.
	Description *string `json:"description"`
	// `AWS::MediaConnect::FlowSource.Name`.
	Name *string `json:"name"`
	// `AWS::MediaConnect::FlowSource.Decryption`.
	Decryption interface{} `json:"decryption"`
	// `AWS::MediaConnect::FlowSource.EntitlementArn`.
	EntitlementArn *string `json:"entitlementArn"`
	// `AWS::MediaConnect::FlowSource.FlowArn`.
	FlowArn *string `json:"flowArn"`
	// `AWS::MediaConnect::FlowSource.IngestPort`.
	IngestPort *float64 `json:"ingestPort"`
	// `AWS::MediaConnect::FlowSource.MaxBitrate`.
	MaxBitrate *float64 `json:"maxBitrate"`
	// `AWS::MediaConnect::FlowSource.MaxLatency`.
	MaxLatency *float64 `json:"maxLatency"`
	// `AWS::MediaConnect::FlowSource.Protocol`.
	Protocol *string `json:"protocol"`
	// `AWS::MediaConnect::FlowSource.StreamId`.
	StreamId *string `json:"streamId"`
	// `AWS::MediaConnect::FlowSource.VpcInterfaceName`.
	VpcInterfaceName *string `json:"vpcInterfaceName"`
	// `AWS::MediaConnect::FlowSource.WhitelistCidr`.
	WhitelistCidr *string `json:"whitelistCidr"`
}

// A CloudFormation `AWS::MediaConnect::FlowVpcInterface`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFlowVpcInterface) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFlowVpcInterface) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::MediaConnect::FlowVpcInterface`.
type CfnFlowVpcInterfaceProps struct {
	// `AWS::MediaConnect::FlowVpcInterface.FlowArn`.
	FlowArn *string `json:"flowArn"`
	// `AWS::MediaConnect::FlowVpcInterface.Name`.
	Name *string `json:"name"`
	// `AWS::MediaConnect::FlowVpcInterface.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::MediaConnect::FlowVpcInterface.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `AWS::MediaConnect::FlowVpcInterface.SubnetId`.
	SubnetId *string `json:"subnetId"`
}

