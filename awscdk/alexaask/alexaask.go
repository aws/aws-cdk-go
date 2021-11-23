package alexaask

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/alexaask/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `Alexa::ASK::Skill`.
//
// TODO: EXAMPLE
//
type CfnSkill interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AuthenticationConfiguration() interface{}
	SetAuthenticationConfiguration(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SkillPackage() interface{}
	SetSkillPackage(val interface{})
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	VendorId() *string
	SetVendorId(val *string)
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

// The jsii proxy struct for CfnSkill
type jsiiProxy_CfnSkill struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSkill) AuthenticationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) SkillPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skillPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkill) VendorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorId",
		&returns,
	)
	return returns
}


// Create a new `Alexa::ASK::Skill`.
func NewCfnSkill(scope constructs.Construct, id *string, props *CfnSkillProps) CfnSkill {
	_init_.Initialize()

	j := jsiiProxy_CfnSkill{}

	_jsii_.Create(
		"aws-cdk-lib.alexa_ask.CfnSkill",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `Alexa::ASK::Skill`.
func NewCfnSkill_Override(c CfnSkill, scope constructs.Construct, id *string, props *CfnSkillProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.alexa_ask.CfnSkill",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSkill) SetAuthenticationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"authenticationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnSkill) SetSkillPackage(val interface{}) {
	_jsii_.Set(
		j,
		"skillPackage",
		val,
	)
}

func (j *jsiiProxy_CfnSkill) SetVendorId(val *string) {
	_jsii_.Set(
		j,
		"vendorId",
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
func CfnSkill_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.alexa_ask.CfnSkill",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSkill_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.alexa_ask.CfnSkill",
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
func CfnSkill_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.alexa_ask.CfnSkill",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSkill_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.alexa_ask.CfnSkill",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnSkill) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSkill) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSkill) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSkill) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnSkill) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnSkill) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSkill) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSkill) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSkill) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSkill) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnSkill) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSkill) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSkill) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSkill) ToString() *string {
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
func (c *jsiiProxy_CfnSkill) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnSkill_AuthenticationConfigurationProperty struct {
	// `CfnSkill.AuthenticationConfigurationProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnSkill.AuthenticationConfigurationProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret"`
	// `CfnSkill.AuthenticationConfigurationProperty.RefreshToken`.
	RefreshToken *string `json:"refreshToken"`
}

// TODO: EXAMPLE
//
type CfnSkill_OverridesProperty struct {
	// `CfnSkill.OverridesProperty.Manifest`.
	Manifest interface{} `json:"manifest"`
}

// TODO: EXAMPLE
//
type CfnSkill_SkillPackageProperty struct {
	// `CfnSkill.SkillPackageProperty.Overrides`.
	Overrides interface{} `json:"overrides"`
	// `CfnSkill.SkillPackageProperty.S3Bucket`.
	S3Bucket *string `json:"s3Bucket"`
	// `CfnSkill.SkillPackageProperty.S3BucketRole`.
	S3BucketRole *string `json:"s3BucketRole"`
	// `CfnSkill.SkillPackageProperty.S3Key`.
	S3Key *string `json:"s3Key"`
	// `CfnSkill.SkillPackageProperty.S3ObjectVersion`.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
}

// Properties for defining a `Alexa::ASK::Skill`.
//
// TODO: EXAMPLE
//
type CfnSkillProps struct {
	// `Alexa::ASK::Skill.AuthenticationConfiguration`.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration"`
	// `Alexa::ASK::Skill.SkillPackage`.
	SkillPackage interface{} `json:"skillPackage"`
	// `Alexa::ASK::Skill.VendorId`.
	VendorId *string `json:"vendorId"`
}

