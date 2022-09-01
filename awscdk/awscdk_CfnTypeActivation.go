// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CloudFormation::TypeActivation`.
//
// Activates a public third-party extension, making it available for use in stack templates. For more information, see [Using public extensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html) in the *AWS CloudFormation User Guide* .
//
// Once you have activated a public third-party extension in your account and region, use [SetTypeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html) to specify configuration properties for the extension. For more information, see [Configuring extensions at the account level](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTypeActivation := cdk.NewCfnTypeActivation(this, jsii.String("MyCfnTypeActivation"), &cfnTypeActivationProps{
//   	autoUpdate: jsii.Boolean(false),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	loggingConfig: &loggingConfigProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		logRoleArn: jsii.String("logRoleArn"),
//   	},
//   	majorVersion: jsii.String("majorVersion"),
//   	publicTypeArn: jsii.String("publicTypeArn"),
//   	publisherId: jsii.String("publisherId"),
//   	type: jsii.String("type"),
//   	typeName: jsii.String("typeName"),
//   	typeNameAlias: jsii.String("typeNameAlias"),
//   	versionBump: jsii.String("versionBump"),
//   })
//
type CfnTypeActivation interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the activated extension, in this account and Region.
	AttrArn() *string
	// Whether to automatically update the extension in this account and region when a new *minor* version is published by the extension publisher.
	//
	// Major versions released by the publisher must be manually updated.
	//
	// The default is `true` .
	AutoUpdate() interface{}
	SetAutoUpdate(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the IAM execution role to use to activate the extension.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// Specifies logging configuration information for an extension.
	LoggingConfig() interface{}
	SetLoggingConfig(val interface{})
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
	// The major version of this extension you want to activate, if multiple major versions are available.
	//
	// The default is the latest major version. CloudFormation uses the latest available *minor* version of the major version selected.
	//
	// You can specify `MajorVersion` or `VersionBump` , but not both.
	MajorVersion() *string
	SetMajorVersion(val *string)
	// The tree node.
	Node() constructs.Node
	// The Amazon Resource Number (ARN) of the public extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublicTypeArn() *string
	SetPublicTypeArn(val *string)
	// The ID of the extension publisher.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublisherId() *string
	SetPublisherId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The extension type.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	Type() *string
	SetType(val *string)
	// The name of the extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	TypeName() *string
	SetTypeName(val *string)
	// An alias to assign to the public extension, in this account and region.
	//
	// If you specify an alias for the extension, CloudFormation treats the alias as the extension type name within this account and region. You must use the alias to refer to the extension in your templates, API calls, and CloudFormation console.
	//
	// An extension alias must be unique within a given account and region. You can activate the same public resource multiple times in the same account and region, using different type name aliases.
	TypeNameAlias() *string
	SetTypeNameAlias(val *string)
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
	// Manually updates a previously-activated type to a new major or minor version, if available.
	//
	// You can also use this parameter to update the value of `AutoUpdate` .
	//
	// - `MAJOR` : CloudFormation updates the extension to the newest major version, if one is available.
	// - `MINOR` : CloudFormation updates the extension to the newest minor version, if one is available.
	VersionBump() *string
	SetVersionBump(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target CfnResource)
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
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
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

// The jsii proxy struct for CfnTypeActivation
type jsiiProxy_CfnTypeActivation struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnTypeActivation) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) AutoUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) LoggingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) MajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) PublicTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) PublisherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) TypeNameAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTypeActivation) VersionBump() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionBump",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::TypeActivation`.
func NewCfnTypeActivation(scope constructs.Construct, id *string, props *CfnTypeActivationProps) CfnTypeActivation {
	_init_.Initialize()

	if err := validateNewCfnTypeActivationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTypeActivation{}

	_jsii_.Create(
		"aws-cdk-lib.CfnTypeActivation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::TypeActivation`.
func NewCfnTypeActivation_Override(c CfnTypeActivation, scope constructs.Construct, id *string, props *CfnTypeActivationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnTypeActivation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetAutoUpdate(val interface{}) {
	if err := j.validateSetAutoUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpdate",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetLoggingConfig(val interface{}) {
	if err := j.validateSetLoggingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetMajorVersion(val *string) {
	_jsii_.Set(
		j,
		"majorVersion",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetPublicTypeArn(val *string) {
	_jsii_.Set(
		j,
		"publicTypeArn",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetPublisherId(val *string) {
	_jsii_.Set(
		j,
		"publisherId",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetTypeNameAlias(val *string) {
	_jsii_.Set(
		j,
		"typeNameAlias",
		val,
	)
}

func (j *jsiiProxy_CfnTypeActivation)SetVersionBump(val *string) {
	_jsii_.Set(
		j,
		"versionBump",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTypeActivation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTypeActivation_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTypeActivation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTypeActivation_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
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
func CfnTypeActivation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTypeActivation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTypeActivation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTypeActivation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CfnTypeActivation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTypeActivation) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddDependsOn(target CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTypeActivation) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTypeActivation) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTypeActivation) GetAtt(attributeName *string) Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTypeActivation) Inspect(inspector TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTypeActivation) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTypeActivation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTypeActivation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTypeActivation) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

