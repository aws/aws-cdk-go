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
// The `Alexa::ASK::Skill` resource creates an Alexa skill that enables customers to access new abilities. For more information about developing a skill, see the  .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   cfnSkill := awscdk.Alexa_ask.NewCfnSkill(this, jsii.String("MyCfnSkill"), &cfnSkillProps{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	skillPackage: &skillPackageProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//
//   		// the properties below are optional
//   		overrides: &overridesProperty{
//   			manifest: manifest,
//   		},
//   		s3BucketRole: jsii.String("s3BucketRole"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	vendorId: jsii.String("vendorId"),
//   })
//
type CfnSkill interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
	//
	// Only Login with Amazon clients created through the  are supported. The client ID, client secret, and refresh token are required.
	AuthenticationConfiguration() interface{}
	SetAuthenticationConfiguration(val interface{})
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Configuration for the skill package that contains the components of the Alexa skill.
	//
	// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. For more information about the skill package format, see the  .
	SkillPackage() interface{}
	SetSkillPackage(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The vendor ID associated with the Amazon developer account that will host the skill.
	//
	// Details for retrieving the vendor ID are in  . The provided LWA credentials must be linked to the developer account associated with this vendor ID.
	VendorId() *string
	SetVendorId(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnSkill) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSkill) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSkill) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSkill) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSkill) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSkill) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSkill) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnSkill) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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

func (c *jsiiProxy_CfnSkill) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AuthenticationConfiguration` property type specifies the Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
//
// Only Login with Amazon security profiles created through the  are supported for authentication. A client ID, client secret, and refresh token are required. You can generate a client ID and client secret by creating a new  on the Amazon Developer Portal or you can retrieve them from an existing profile. You can then retrieve the refresh token using the Alexa Skills Kit CLI. For instructions, see  in the  .
//
// `AuthenticationConfiguration` is a property of the `Alexa::ASK::Skill` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationProperty := &authenticationConfigurationProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnSkill_AuthenticationConfigurationProperty struct {
	// Client ID from Login with Amazon (LWA).
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Client secret from Login with Amazon (LWA).
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Refresh token from Login with Amazon (LWA).
	//
	// This token is secret.
	RefreshToken *string `field:"required" json:"refreshToken" yaml:"refreshToken"`
}

// The `Overrides` property type provides overrides to the skill package to apply when creating or updating the skill.
//
// Values provided here do not modify the contents of the original skill package. Currently, only overriding values inside of the skill manifest component of the package is supported.
//
// `Overrides` is a property of the `Alexa::ASK::Skill SkillPackage` property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   overridesProperty := &overridesProperty{
//   	manifest: manifest,
//   }
//
type CfnSkill_OverridesProperty struct {
	// Overrides to apply to the skill manifest inside of the skill package.
	//
	// The skill manifest contains metadata about the skill. For more information, see  .
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

// The `SkillPackage` property type contains configuration details for the skill package that contains the components of the Alexa skill.
//
// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. More details about the skill package format are located in the  .
//
// `SkillPackage` is a property of the `Alexa::ASK::Skill` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   skillPackageProperty := &skillPackageProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//
//   	// the properties below are optional
//   	overrides: &overridesProperty{
//   		manifest: manifest,
//   	},
//   	s3BucketRole: jsii.String("s3BucketRole"),
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
type CfnSkill_SkillPackageProperty struct {
	// The name of the Amazon S3 bucket where the .zip file that contains the skill package is stored.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The location and name of the skill package .zip file.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// Overrides to the skill package to apply when creating or updating the skill.
	//
	// Values provided here do not modify the contents of the original skill package. Currently, only overriding values inside of the skill manifest component of the package is supported.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// ARN of the IAM role that grants the Alexa service ( `alexa-appkit.amazon.com` ) permission to access the bucket and retrieve the skill package. This property is optional. If you do not provide it, the bucket must be publicly accessible or configured with a policy that allows this access. Otherwise, AWS CloudFormation cannot create the skill.
	S3BucketRole *string `field:"optional" json:"s3BucketRole" yaml:"s3BucketRole"`
	// If you have S3 versioning enabled, the version ID of the skill package.zip file.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

// Properties for defining a `CfnSkill`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifest interface{}
//
//   cfnSkillProps := &cfnSkillProps{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	skillPackage: &skillPackageProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//
//   		// the properties below are optional
//   		overrides: &overridesProperty{
//   			manifest: manifest,
//   		},
//   		s3BucketRole: jsii.String("s3BucketRole"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	vendorId: jsii.String("vendorId"),
//   }
//
type CfnSkillProps struct {
	// Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
	//
	// Only Login with Amazon clients created through the  are supported. The client ID, client secret, and refresh token are required.
	AuthenticationConfiguration interface{} `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Configuration for the skill package that contains the components of the Alexa skill.
	//
	// Skill packages are retrieved from an Amazon S3 bucket and key and used to create and update the skill. For more information about the skill package format, see the  .
	SkillPackage interface{} `field:"required" json:"skillPackage" yaml:"skillPackage"`
	// The vendor ID associated with the Amazon developer account that will host the skill.
	//
	// Details for retrieving the vendor ID are in  . The provided LWA credentials must be linked to the developer account associated with this vendor ID.
	VendorId *string `field:"required" json:"vendorId" yaml:"vendorId"`
}

