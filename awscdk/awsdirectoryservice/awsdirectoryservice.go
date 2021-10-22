package awsdirectoryservice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdirectoryservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DirectoryService::MicrosoftAD`.
type CfnMicrosoftAD interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAlias() *string
	AttrDnsIpAddresses() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreateAlias() interface{}
	SetCreateAlias(val interface{})
	CreationStack() *[]*string
	Edition() *string
	SetEdition(val *string)
	EnableSso() interface{}
	SetEnableSso(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	Ref() *string
	ShortName() *string
	SetShortName(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	VpcSettings() interface{}
	SetVpcSettings(val interface{})
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

// The jsii proxy struct for CfnMicrosoftAD
type jsiiProxy_CfnMicrosoftAD struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMicrosoftAD) AttrAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) AttrDnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) CreateAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) EnableSso() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSso",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftAD) VpcSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcSettings",
		&returns,
	)
	return returns
}


// Create a new `AWS::DirectoryService::MicrosoftAD`.
func NewCfnMicrosoftAD(scope constructs.Construct, id *string, props *CfnMicrosoftADProps) CfnMicrosoftAD {
	_init_.Initialize()

	j := jsiiProxy_CfnMicrosoftAD{}

	_jsii_.Create(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DirectoryService::MicrosoftAD`.
func NewCfnMicrosoftAD_Override(c CfnMicrosoftAD, scope constructs.Construct, id *string, props *CfnMicrosoftADProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetCreateAlias(val interface{}) {
	_jsii_.Set(
		j,
		"createAlias",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetEnableSso(val interface{}) {
	_jsii_.Set(
		j,
		"enableSso",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetShortName(val *string) {
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftAD) SetVpcSettings(val interface{}) {
	_jsii_.Set(
		j,
		"vpcSettings",
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
func CfnMicrosoftAD_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMicrosoftAD_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
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
func CfnMicrosoftAD_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMicrosoftAD_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_directoryservice.CfnMicrosoftAD",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnMicrosoftAD) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnMicrosoftAD) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnMicrosoftAD) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnMicrosoftAD) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnMicrosoftAD) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnMicrosoftAD) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnMicrosoftAD) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnMicrosoftAD) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnMicrosoftAD) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnMicrosoftAD) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMicrosoftAD) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMicrosoftAD) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnMicrosoftAD) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnMicrosoftAD) ToString() *string {
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
func (c *jsiiProxy_CfnMicrosoftAD) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnMicrosoftAD_VpcSettingsProperty struct {
	// `CfnMicrosoftAD.VpcSettingsProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `CfnMicrosoftAD.VpcSettingsProperty.VpcId`.
	VpcId *string `json:"vpcId"`
}

// Properties for defining a `AWS::DirectoryService::MicrosoftAD`.
type CfnMicrosoftADProps struct {
	// `AWS::DirectoryService::MicrosoftAD.Name`.
	Name *string `json:"name"`
	// `AWS::DirectoryService::MicrosoftAD.Password`.
	Password *string `json:"password"`
	// `AWS::DirectoryService::MicrosoftAD.VpcSettings`.
	VpcSettings interface{} `json:"vpcSettings"`
	// `AWS::DirectoryService::MicrosoftAD.CreateAlias`.
	CreateAlias interface{} `json:"createAlias"`
	// `AWS::DirectoryService::MicrosoftAD.Edition`.
	Edition *string `json:"edition"`
	// `AWS::DirectoryService::MicrosoftAD.EnableSso`.
	EnableSso interface{} `json:"enableSso"`
	// `AWS::DirectoryService::MicrosoftAD.ShortName`.
	ShortName *string `json:"shortName"`
}

// A CloudFormation `AWS::DirectoryService::SimpleAD`.
type CfnSimpleAD interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAlias() *string
	AttrDnsIpAddresses() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreateAlias() interface{}
	SetCreateAlias(val interface{})
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EnableSso() interface{}
	SetEnableSso(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	Ref() *string
	ShortName() *string
	SetShortName(val *string)
	Size() *string
	SetSize(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	VpcSettings() interface{}
	SetVpcSettings(val interface{})
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

// The jsii proxy struct for CfnSimpleAD
type jsiiProxy_CfnSimpleAD struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSimpleAD) AttrAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) AttrDnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) CreateAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) EnableSso() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSso",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleAD) VpcSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcSettings",
		&returns,
	)
	return returns
}


// Create a new `AWS::DirectoryService::SimpleAD`.
func NewCfnSimpleAD(scope constructs.Construct, id *string, props *CfnSimpleADProps) CfnSimpleAD {
	_init_.Initialize()

	j := jsiiProxy_CfnSimpleAD{}

	_jsii_.Create(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DirectoryService::SimpleAD`.
func NewCfnSimpleAD_Override(c CfnSimpleAD, scope constructs.Construct, id *string, props *CfnSimpleADProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetCreateAlias(val interface{}) {
	_jsii_.Set(
		j,
		"createAlias",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetEnableSso(val interface{}) {
	_jsii_.Set(
		j,
		"enableSso",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetShortName(val *string) {
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetSize(val *string) {
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleAD) SetVpcSettings(val interface{}) {
	_jsii_.Set(
		j,
		"vpcSettings",
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
func CfnSimpleAD_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSimpleAD_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
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
func CfnSimpleAD_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimpleAD_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_directoryservice.CfnSimpleAD",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnSimpleAD) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSimpleAD) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSimpleAD) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSimpleAD) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnSimpleAD) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnSimpleAD) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSimpleAD) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSimpleAD) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSimpleAD) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSimpleAD) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnSimpleAD) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSimpleAD) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSimpleAD) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSimpleAD) ToString() *string {
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
func (c *jsiiProxy_CfnSimpleAD) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnSimpleAD_VpcSettingsProperty struct {
	// `CfnSimpleAD.VpcSettingsProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `CfnSimpleAD.VpcSettingsProperty.VpcId`.
	VpcId *string `json:"vpcId"`
}

// Properties for defining a `AWS::DirectoryService::SimpleAD`.
type CfnSimpleADProps struct {
	// `AWS::DirectoryService::SimpleAD.Name`.
	Name *string `json:"name"`
	// `AWS::DirectoryService::SimpleAD.Password`.
	Password *string `json:"password"`
	// `AWS::DirectoryService::SimpleAD.Size`.
	Size *string `json:"size"`
	// `AWS::DirectoryService::SimpleAD.VpcSettings`.
	VpcSettings interface{} `json:"vpcSettings"`
	// `AWS::DirectoryService::SimpleAD.CreateAlias`.
	CreateAlias interface{} `json:"createAlias"`
	// `AWS::DirectoryService::SimpleAD.Description`.
	Description *string `json:"description"`
	// `AWS::DirectoryService::SimpleAD.EnableSso`.
	EnableSso interface{} `json:"enableSso"`
	// `AWS::DirectoryService::SimpleAD.ShortName`.
	ShortName *string `json:"shortName"`
}

