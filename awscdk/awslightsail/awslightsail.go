package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lightsail::Disk`.
type CfnDisk interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AddOns() interface{}
	SetAddOns(val interface{})
	AttrAttachedTo() *string
	AttrAttachmentState() *string
	AttrDiskArn() *string
	AttrIops() *float64
	AttrIsAttached() awscdk.IResolvable
	AttrPath() *string
	AttrResourceType() *string
	AttrState() *string
	AttrSupportCode() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DiskName() *string
	SetDiskName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SizeInGb() *float64
	SetSizeInGb(val *float64)
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

// The jsii proxy struct for CfnDisk
type jsiiProxy_CfnDisk struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDisk) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrDiskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDiskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIsAttached() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) SizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk(scope constructs.Construct, id *string, props *CfnDiskProps) CfnDisk {
	_init_.Initialize()

	j := jsiiProxy_CfnDisk{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk_Override(c CfnDisk, scope constructs.Construct, id *string, props *CfnDiskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDisk) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetDiskName(val *string) {
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetSizeInGb(val *float64) {
	_jsii_.Set(
		j,
		"sizeInGb",
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
func CfnDisk_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDisk_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
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
func CfnDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDisk_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDisk) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDisk) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDisk) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDisk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDisk) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDisk) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDisk) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDisk) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDisk) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDisk) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDisk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDisk) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDisk) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDisk) ToString() *string {
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
func (c *jsiiProxy_CfnDisk) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDisk_AddOnProperty struct {
	// `CfnDisk.AddOnProperty.AddOnType`.
	AddOnType *string `json:"addOnType"`
	// `CfnDisk.AddOnProperty.AutoSnapshotAddOnRequest`.
	AutoSnapshotAddOnRequest interface{} `json:"autoSnapshotAddOnRequest"`
	// `CfnDisk.AddOnProperty.Status`.
	Status *string `json:"status"`
}

type CfnDisk_AutoSnapshotAddOnProperty struct {
	// `CfnDisk.AutoSnapshotAddOnProperty.SnapshotTimeOfDay`.
	SnapshotTimeOfDay *string `json:"snapshotTimeOfDay"`
}

// Properties for defining a `AWS::Lightsail::Disk`.
type CfnDiskProps struct {
	// `AWS::Lightsail::Disk.DiskName`.
	DiskName *string `json:"diskName"`
	// `AWS::Lightsail::Disk.SizeInGb`.
	SizeInGb *float64 `json:"sizeInGb"`
	// `AWS::Lightsail::Disk.AddOns`.
	AddOns interface{} `json:"addOns"`
	// `AWS::Lightsail::Disk.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::Lightsail::Disk.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Lightsail::Instance`.
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AddOns() interface{}
	SetAddOns(val interface{})
	AttrHardwareCpuCount() *float64
	AttrHardwareRamSizeInGb() *float64
	AttrInstanceArn() *string
	AttrIsStaticIp() awscdk.IResolvable
	AttrKeyPairName() *string
	AttrLocationAvailabilityZone() *string
	AttrLocationRegionName() *string
	AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string
	AttrPrivateIpAddress() *string
	AttrPublicIpAddress() *string
	AttrResourceType() *string
	AttrSshKeyName() *string
	AttrStateCode() *float64
	AttrStateName() *string
	AttrSupportCode() *string
	AttrUserData() *string
	AttrUserName() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	BlueprintId() *string
	SetBlueprintId(val *string)
	BundleId() *string
	SetBundleId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Hardware() interface{}
	SetHardware(val interface{})
	InstanceName() *string
	SetInstanceName(val *string)
	Location() interface{}
	SetLocation(val interface{})
	LogicalId() *string
	Networking() interface{}
	SetNetworking(val interface{})
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	State() interface{}
	SetState(val interface{})
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

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareRamSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareRamSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrIsStaticIp() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrKeyPairName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKeyPairName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkingMonthlyTransferGbPerMonthAllocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrStateCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrUserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Hardware() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Networking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) State() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBlueprintId(val *string) {
	_jsii_.Set(
		j,
		"blueprintId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetHardware(val interface{}) {
	_jsii_.Set(
		j,
		"hardware",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceName(val *string) {
	_jsii_.Set(
		j,
		"instanceName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetLocation(val interface{}) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetNetworking(val interface{}) {
	_jsii_.Set(
		j,
		"networking",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetState(val interface{}) {
	_jsii_.Set(
		j,
		"state",
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
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
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
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnInstance) ToString() *string {
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
func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInstance_AddOnProperty struct {
	// `CfnInstance.AddOnProperty.AddOnType`.
	AddOnType *string `json:"addOnType"`
	// `CfnInstance.AddOnProperty.AutoSnapshotAddOnRequest`.
	AutoSnapshotAddOnRequest interface{} `json:"autoSnapshotAddOnRequest"`
	// `CfnInstance.AddOnProperty.Status`.
	Status *string `json:"status"`
}

type CfnInstance_AutoSnapshotAddOnProperty struct {
	// `CfnInstance.AutoSnapshotAddOnProperty.SnapshotTimeOfDay`.
	SnapshotTimeOfDay *string `json:"snapshotTimeOfDay"`
}

type CfnInstance_DiskProperty struct {
	// `CfnInstance.DiskProperty.DiskName`.
	DiskName *string `json:"diskName"`
	// `CfnInstance.DiskProperty.Path`.
	Path *string `json:"path"`
	// `CfnInstance.DiskProperty.AttachedTo`.
	AttachedTo *string `json:"attachedTo"`
	// `CfnInstance.DiskProperty.AttachmentState`.
	AttachmentState *string `json:"attachmentState"`
	// `CfnInstance.DiskProperty.IOPS`.
	Iops *float64 `json:"iops"`
	// `CfnInstance.DiskProperty.IsSystemDisk`.
	IsSystemDisk interface{} `json:"isSystemDisk"`
	// `CfnInstance.DiskProperty.SizeInGb`.
	SizeInGb *string `json:"sizeInGb"`
}

type CfnInstance_HardwareProperty struct {
	// `CfnInstance.HardwareProperty.CpuCount`.
	CpuCount *float64 `json:"cpuCount"`
	// `CfnInstance.HardwareProperty.Disks`.
	Disks interface{} `json:"disks"`
	// `CfnInstance.HardwareProperty.RamSizeInGb`.
	RamSizeInGb *float64 `json:"ramSizeInGb"`
}

type CfnInstance_LocationProperty struct {
	// `CfnInstance.LocationProperty.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `CfnInstance.LocationProperty.RegionName`.
	RegionName *string `json:"regionName"`
}

type CfnInstance_MonthlyTransferProperty struct {
	// `CfnInstance.MonthlyTransferProperty.GbPerMonthAllocated`.
	GbPerMonthAllocated *string `json:"gbPerMonthAllocated"`
}

type CfnInstance_NetworkingProperty struct {
	// `CfnInstance.NetworkingProperty.Ports`.
	Ports interface{} `json:"ports"`
	// `CfnInstance.NetworkingProperty.MonthlyTransfer`.
	MonthlyTransfer *float64 `json:"monthlyTransfer"`
}

type CfnInstance_PortProperty struct {
	// `CfnInstance.PortProperty.AccessDirection`.
	AccessDirection *string `json:"accessDirection"`
	// `CfnInstance.PortProperty.AccessFrom`.
	AccessFrom *string `json:"accessFrom"`
	// `CfnInstance.PortProperty.AccessType`.
	AccessType *string `json:"accessType"`
	// `CfnInstance.PortProperty.CidrListAliases`.
	CidrListAliases *[]*string `json:"cidrListAliases"`
	// `CfnInstance.PortProperty.Cidrs`.
	Cidrs *[]*string `json:"cidrs"`
	// `CfnInstance.PortProperty.CommonName`.
	CommonName *string `json:"commonName"`
	// `CfnInstance.PortProperty.FromPort`.
	FromPort *float64 `json:"fromPort"`
	// `CfnInstance.PortProperty.Ipv6Cidrs`.
	Ipv6Cidrs *[]*string `json:"ipv6Cidrs"`
	// `CfnInstance.PortProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnInstance.PortProperty.ToPort`.
	ToPort *float64 `json:"toPort"`
}

type CfnInstance_StateProperty struct {
	// `CfnInstance.StateProperty.Code`.
	Code *float64 `json:"code"`
	// `CfnInstance.StateProperty.Name`.
	Name *string `json:"name"`
}

// Properties for defining a `AWS::Lightsail::Instance`.
type CfnInstanceProps struct {
	// `AWS::Lightsail::Instance.BlueprintId`.
	BlueprintId *string `json:"blueprintId"`
	// `AWS::Lightsail::Instance.BundleId`.
	BundleId *string `json:"bundleId"`
	// `AWS::Lightsail::Instance.InstanceName`.
	InstanceName *string `json:"instanceName"`
	// `AWS::Lightsail::Instance.AddOns`.
	AddOns interface{} `json:"addOns"`
	// `AWS::Lightsail::Instance.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
	// `AWS::Lightsail::Instance.Hardware`.
	Hardware interface{} `json:"hardware"`
	// `AWS::Lightsail::Instance.Location`.
	Location interface{} `json:"location"`
	// `AWS::Lightsail::Instance.Networking`.
	Networking interface{} `json:"networking"`
	// `AWS::Lightsail::Instance.State`.
	State interface{} `json:"state"`
	// `AWS::Lightsail::Instance.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

