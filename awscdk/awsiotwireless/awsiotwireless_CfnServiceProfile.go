package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTWireless::ServiceProfile`.
//
// Creates a new service profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProfile := awscdk.Aws_iotwireless.NewCfnServiceProfile(this, jsii.String("MyCfnServiceProfile"), &cfnServiceProfileProps{
//   	loRaWan: &loRaWANServiceProfileProperty{
//   		addGwMetadata: jsii.Boolean(false),
//   		channelMask: jsii.String("channelMask"),
//   		devStatusReqFreq: jsii.Number(123),
//   		dlBucketSize: jsii.Number(123),
//   		dlRate: jsii.Number(123),
//   		dlRatePolicy: jsii.String("dlRatePolicy"),
//   		drMax: jsii.Number(123),
//   		drMin: jsii.Number(123),
//   		hrAllowed: jsii.Boolean(false),
//   		minGwDiversity: jsii.Number(123),
//   		nwkGeoLoc: jsii.Boolean(false),
//   		prAllowed: jsii.Boolean(false),
//   		raAllowed: jsii.Boolean(false),
//   		reportDevStatusBattery: jsii.Boolean(false),
//   		reportDevStatusMargin: jsii.Boolean(false),
//   		targetPer: jsii.Number(123),
//   		ulBucketSize: jsii.Number(123),
//   		ulRate: jsii.Number(123),
//   		ulRatePolicy: jsii.String("ulRatePolicy"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnServiceProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the service profile created.
	AttrArn() *string
	// The ID of the service profile created.
	AttrId() *string
	// The ChannelMask value.
	AttrLoRaWanChannelMask() *string
	// The DevStatusReqFreq value.
	AttrLoRaWanDevStatusReqFreq() *float64
	// The DLBucketSize value.
	AttrLoRaWanDlBucketSize() *float64
	// The DLRate value.
	AttrLoRaWanDlRate() *float64
	// The DLRatePolicy value.
	AttrLoRaWanDlRatePolicy() *string
	// The DRMax value.
	AttrLoRaWanDrMax() *float64
	// The DRMin value.
	AttrLoRaWanDrMin() *float64
	// The HRAllowed value that describes whether handover roaming is allowed.
	AttrLoRaWanHrAllowed() awscdk.IResolvable
	// The MinGwDiversity value.
	AttrLoRaWanMinGwDiversity() *float64
	// The NwkGeoLoc value.
	AttrLoRaWanNwkGeoLoc() awscdk.IResolvable
	// The PRAllowed value that describes whether passive roaming is allowed.
	AttrLoRaWanPrAllowed() awscdk.IResolvable
	// The RAAllowed value that describes whether roaming activation is allowed.
	AttrLoRaWanRaAllowed() awscdk.IResolvable
	// The ReportDevStatusBattery value.
	AttrLoRaWanReportDevStatusBattery() awscdk.IResolvable
	// The ReportDevStatusMargin value.
	AttrLoRaWanReportDevStatusMargin() awscdk.IResolvable
	AttrLoRaWanResponse() awscdk.IResolvable
	// The TargetPer value.
	AttrLoRaWanTargetPer() *float64
	// The UlBucketSize value.
	AttrLoRaWanUlBucketSize() *float64
	// The ULRate value.
	AttrLoRaWanUlRate() *float64
	// The ULRatePolicy value.
	AttrLoRaWanUlRatePolicy() *string
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
	// LoRaWAN service profile object.
	LoRaWan() interface{}
	SetLoRaWan(val interface{})
	// The name of the new resource.
	Name() *string
	SetName(val *string)
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
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnServiceProfile
type jsiiProxy_CfnServiceProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnServiceProfile) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanChannelMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoRaWanChannelMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDevStatusReqFreq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanDevStatusReqFreq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanDlBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanDlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoRaWanDlRatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDrMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanDrMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanDrMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanDrMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanHrAllowed() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanHrAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanMinGwDiversity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanMinGwDiversity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanNwkGeoLoc() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanNwkGeoLoc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanPrAllowed() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanPrAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanRaAllowed() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanRaAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanReportDevStatusBattery() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanReportDevStatusBattery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanReportDevStatusMargin() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanReportDevStatusMargin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanResponse() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLoRaWanResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanTargetPer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanTargetPer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanUlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanUlBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanUlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLoRaWanUlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) AttrLoRaWanUlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoRaWanUlRatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) LoRaWan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loRaWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceProfile) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTWireless::ServiceProfile`.
func NewCfnServiceProfile(scope constructs.Construct, id *string, props *CfnServiceProfileProps) CfnServiceProfile {
	_init_.Initialize()

	if err := validateNewCfnServiceProfileParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTWireless::ServiceProfile`.
func NewCfnServiceProfile_Override(c CfnServiceProfile, scope constructs.Construct, id *string, props *CfnServiceProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServiceProfile)SetLoRaWan(val interface{}) {
	if err := j.validateSetLoRaWanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loRaWan",
		val,
	)
}

func (j *jsiiProxy_CfnServiceProfile)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnServiceProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceProfile_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnServiceProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnServiceProfile_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
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
func CfnServiceProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotwireless.CfnServiceProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceProfile) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServiceProfile) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceProfile) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServiceProfile) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServiceProfile) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServiceProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServiceProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServiceProfile) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnServiceProfile) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnServiceProfile) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServiceProfile) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServiceProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnServiceProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceProfile) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

