package awsram

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::RAM::ResourceShare`.
//
// Specifies a resource share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceShare := awscdk.Aws_ram.NewCfnResourceShare(this, jsii.String("MyCfnResourceShare"), &cfnResourceShareProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	allowExternalPrincipals: jsii.Boolean(false),
//   	permissionArns: []*string{
//   		jsii.String("permissionArns"),
//   	},
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnResourceShare interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true` lets you share with individual AWS accounts that are *not* in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true` .
	AllowExternalPrincipals() interface{}
	SetAllowExternalPrincipals(val interface{})
	// The Amazon Resource Name (ARN) of the resource share.
	AttrArn() *string
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
	// Specifies the name of the resource share.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	PermissionArns() *[]*string
	SetPermissionArns(val *[]*string)
	// Specifies a list of one or more principals to associate with the resource share.
	//
	// You can include the following values:
	//
	// - An AWS account ID, for example: `123456789012`
	// - An [Amazon Resoure Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of an organization in AWS Organizations , for example: `arn:aws:organizations::123456789012:organization/o-exampleorgid`
	// - An ARN of an organizational unit (OU) in AWS Organizations , for example: `arn:aws:organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123`
	// - An ARN of an IAM role, for example: `arn:aws:iam::123456789012:role/rolename`
	// - An ARN of an IAM user, for example: `arn:aws:iam::123456789012user/username`
	//
	// > Not all resource types can be shared with IAM roles and users. For more information, see [Sharing with IAM roles and users](https://docs.aws.amazon.com//ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types) in the *AWS Resource Access Manager User Guide* .
	Principals() *[]*string
	SetPrincipals(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies a list of one or more ARNs of the resources to associate with the resource share.
	ResourceArns() *[]*string
	SetResourceArns(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies one or more tags to attach to the resource share itself.
	//
	// It doesn't attach the tags to the resources associated with the resource share.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnResourceShare
type jsiiProxy_CfnResourceShare struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceShare) AllowExternalPrincipals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) PermissionArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Principals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) ResourceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceShare) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RAM::ResourceShare`.
func NewCfnResourceShare(scope constructs.Construct, id *string, props *CfnResourceShareProps) CfnResourceShare {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceShare{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RAM::ResourceShare`.
func NewCfnResourceShare_Override(c CfnResourceShare, scope constructs.Construct, id *string, props *CfnResourceShareProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceShare) SetAllowExternalPrincipals(val interface{}) {
	_jsii_.Set(
		j,
		"allowExternalPrincipals",
		val,
	)
}

func (j *jsiiProxy_CfnResourceShare) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnResourceShare) SetPermissionArns(val *[]*string) {
	_jsii_.Set(
		j,
		"permissionArns",
		val,
	)
}

func (j *jsiiProxy_CfnResourceShare) SetPrincipals(val *[]*string) {
	_jsii_.Set(
		j,
		"principals",
		val,
	)
}

func (j *jsiiProxy_CfnResourceShare) SetResourceArns(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceArns",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceShare_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceShare_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
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
func CfnResourceShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceShare_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ram.CfnResourceShare",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceShare) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceShare) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceShare) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceShare) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceShare) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceShare) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceShare) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceShare) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceShare) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceShare) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceShare) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceShare) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourceShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceShareProps := &cfnResourceShareProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	allowExternalPrincipals: jsii.Boolean(false),
//   	permissionArns: []*string{
//   		jsii.String("permissionArns"),
//   	},
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResourceShareProps struct {
	// Specifies the name of the resource share.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true` lets you share with individual AWS accounts that are *not* in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true` .
	AllowExternalPrincipals interface{} `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	// Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	PermissionArns *[]*string `field:"optional" json:"permissionArns" yaml:"permissionArns"`
	// Specifies a list of one or more principals to associate with the resource share.
	//
	// You can include the following values:
	//
	// - An AWS account ID, for example: `123456789012`
	// - An [Amazon Resoure Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of an organization in AWS Organizations , for example: `arn:aws:organizations::123456789012:organization/o-exampleorgid`
	// - An ARN of an organizational unit (OU) in AWS Organizations , for example: `arn:aws:organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123`
	// - An ARN of an IAM role, for example: `arn:aws:iam::123456789012:role/rolename`
	// - An ARN of an IAM user, for example: `arn:aws:iam::123456789012user/username`
	//
	// > Not all resource types can be shared with IAM roles and users. For more information, see [Sharing with IAM roles and users](https://docs.aws.amazon.com//ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types) in the *AWS Resource Access Manager User Guide* .
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Specifies a list of one or more ARNs of the resources to associate with the resource share.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Specifies one or more tags to attach to the resource share itself.
	//
	// It doesn't attach the tags to the resources associated with the resource share.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

