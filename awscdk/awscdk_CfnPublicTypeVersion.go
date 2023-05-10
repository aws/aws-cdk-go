// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CloudFormation::PublicTypeVersion`.
//
// Tests and publishes a registered extension as a public, third-party extension.
//
// CloudFormation first tests the extension to make sure it meets all necessary requirements for being published in the CloudFormation registry. If it does, CloudFormation then publishes it to the registry as a public third-party extension in this Region. Public extensions are available for use by all CloudFormation users.
//
// - For resource types, testing includes passing all contracts tests defined for the type.
// - For modules, testing includes determining if the module's model meets all necessary requirements.
//
// For more information, see [Testing your public extension prior to publishing](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-testing) in the *CloudFormation CLI User Guide* .
//
// If you don't specify a version, CloudFormation uses the default version of the extension in your account and Region for testing.
//
// To perform testing, CloudFormation assumes the execution role specified when the type was registered.
//
// An extension must have a test status of `PASSED` before it can be published. For more information, see [Publishing extensions to make them available for public use](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-publish.html) in the *CloudFormation CLI User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublicTypeVersion := monocdk.NewCfnPublicTypeVersion(this, jsii.String("MyCfnPublicTypeVersion"), &CfnPublicTypeVersionProps{
//   	Arn: jsii.String("arn"),
//   	LogDeliveryBucket: jsii.String("logDeliveryBucket"),
//   	PublicVersionNumber: jsii.String("publicVersionNumber"),
//   	Type: jsii.String("type"),
//   	TypeName: jsii.String("typeName"),
//   })
//
type CfnPublicTypeVersion interface {
	CfnResource
	IInspectable
	// The Amazon Resource Number (ARN) of the extension.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Arn() *string
	SetArn(val *string)
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication.
	AttrPublicTypeArn() *string
	// The publisher ID of the extension publisher.
	AttrPublisherId() *string
	// The Amazon Resource Number (ARN) assigned to this version of the extension.
	AttrTypeVersionArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The S3 bucket to which CloudFormation delivers the contract test execution logs.
	//
	// CloudFormation delivers the logs by the time contract testing has completed and the extension has been assigned a test type status of `PASSED` or `FAILED` .
	//
	// The user initiating the stack operation must be able to access items in the specified S3 bucket. Specifically, the user needs the following permissions:
	//
	// - GetObject
	// - PutObject
	//
	// For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	LogDeliveryBucket() *string
	SetLogDeliveryBucket(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
	// The version number to assign to this version of the extension.
	//
	// Use the following format, and adhere to semantic versioning when assigning a version number to your extension:
	//
	// `MAJOR.MINOR.PATCH`
	//
	// For more information, see [Semantic Versioning 2.0.0](https://docs.aws.amazon.com/https://semver.org/) .
	//
	// If you don't specify a version number, CloudFormation increments the version number by one minor version release.
	//
	// You cannot specify a version number the first time you publish a type. AWS CloudFormation automatically sets the first version number to be `1.0.0` .
	PublicVersionNumber() *string
	SetPublicVersionNumber(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() Stack
	// The type of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	Type() *string
	SetType(val *string)
	// The name of the extension to test.
	//
	// Conditional: You must specify `Arn` , or `TypeName` and `Type` .
	TypeName() *string
	SetTypeName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPublicTypeVersion
type jsiiProxy_CfnPublicTypeVersion struct {
	jsiiProxy_CfnResource
	jsiiProxy_IInspectable
}

func (j *jsiiProxy_CfnPublicTypeVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrPublicTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrPublisherId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublisherId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) AttrTypeVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnOptions() ICfnResourceOptions {
	var returns ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) LogDeliveryBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDeliveryBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) PublicVersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicTypeVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::PublicTypeVersion`.
func NewCfnPublicTypeVersion(scope Construct, id *string, props *CfnPublicTypeVersionProps) CfnPublicTypeVersion {
	_init_.Initialize()

	if err := validateNewCfnPublicTypeVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPublicTypeVersion{}

	_jsii_.Create(
		"monocdk.CfnPublicTypeVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::PublicTypeVersion`.
func NewCfnPublicTypeVersion_Override(c CfnPublicTypeVersion, scope Construct, id *string, props *CfnPublicTypeVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.CfnPublicTypeVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion)SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion)SetLogDeliveryBucket(val *string) {
	_jsii_.Set(
		j,
		"logDeliveryBucket",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion)SetPublicVersionNumber(val *string) {
	_jsii_.Set(
		j,
		"publicVersionNumber",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion)SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnPublicTypeVersion)SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
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
func CfnPublicTypeVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPublicTypeVersion_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnPublicTypeVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPublicTypeVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnPublicTypeVersion_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnPublicTypeVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPublicTypeVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPublicTypeVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CfnPublicTypeVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicTypeVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.CfnPublicTypeVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddDependsOn(target CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) ApplyRemovalPolicy(policy RemovalPolicy, options *RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) GetAtt(attributeName *string) Reference {
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

func (c *jsiiProxy_CfnPublicTypeVersion) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnPublicTypeVersion) Inspect(inspector TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnPublicTypeVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPublicTypeVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicTypeVersion) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

