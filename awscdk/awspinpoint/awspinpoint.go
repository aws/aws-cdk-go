package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Pinpoint::ADMChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the ADM channel to send push notifications through the Amazon Device Messaging (ADM) service to apps that run on Amazon devices, such as Kindle Fire tablets. Before you can use Amazon Pinpoint to send messages to Amazon devices, you have to enable the ADM channel for an Amazon Pinpoint application.
//
// The ADMChannel resource represents the status and authentication settings for the ADM channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnADMChannel := awscdk.Aws_pinpoint.NewCfnADMChannel(this, jsii.String("MyCfnADMChannel"), &cfnADMChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   })
//
type CfnADMChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the ADM channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The Client ID that you received from Amazon to send messages by using ADM.
	ClientId() *string
	SetClientId(val *string)
	// The Client Secret that you received from Amazon to send messages by using ADM.
	ClientSecret() *string
	SetClientSecret(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the ADM channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnADMChannel
type jsiiProxy_CfnADMChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnADMChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnADMChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::ADMChannel`.
func NewCfnADMChannel(scope constructs.Construct, id *string, props *CfnADMChannelProps) CfnADMChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnADMChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::ADMChannel`.
func NewCfnADMChannel_Override(c CfnADMChannel, scope constructs.Construct, id *string, props *CfnADMChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnADMChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnADMChannel) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_CfnADMChannel) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_CfnADMChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnADMChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnADMChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
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
func CfnADMChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnADMChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnADMChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnADMChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnADMChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnADMChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnADMChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnADMChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnADMChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnADMChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnADMChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnADMChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnADMChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnADMChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnADMChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnADMChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnADMChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnADMChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnADMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnADMChannelProps := &cfnADMChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnADMChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the ADM channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Client ID that you received from Amazon to send messages by using ADM.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret that you received from Amazon to send messages by using ADM.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Specifies whether to enable the ADM channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// A CloudFormation `AWS::Pinpoint::APNSChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the APNs channel to send push notification messages to the Apple Push Notification service (APNs). Before you can use Amazon Pinpoint to send notifications to APNs, you have to enable the APNs channel for an Amazon Pinpoint application.
//
// The APNSChannel resource represents the status and authentication settings for the APNs channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSChannel := awscdk.Aws_pinpoint.NewCfnAPNSChannel(this, jsii.String("MyCfnAPNSChannel"), &cfnAPNSChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   })
//
type CfnAPNSChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the APNs channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId() *string
	SetBundleId(val *string)
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate() *string
	SetCertificate(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	// Specifies whether to enable the APNs channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey() *string
	SetPrivateKey(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId() *string
	SetTeamId(val *string)
	// The authentication key to use for APNs tokens.
	TokenKey() *string
	SetTokenKey(val *string)
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId() *string
	SetTokenKeyId(val *string)
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

// The jsii proxy struct for CfnAPNSChannel
type jsiiProxy_CfnAPNSChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAPNSChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::APNSChannel`.
func NewCfnAPNSChannel(scope constructs.Construct, id *string, props *CfnAPNSChannelProps) CfnAPNSChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnAPNSChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::APNSChannel`.
func NewCfnAPNSChannel_Override(c CfnAPNSChannel, scope constructs.Construct, id *string, props *CfnAPNSChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAPNSChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAPNSChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
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
func CfnAPNSChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPNSChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAPNSChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAPNSChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSChannelProps := &cfnAPNSChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the APNs channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the APNs channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

// A CloudFormation `AWS::Pinpoint::APNSSandboxChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the APNs sandbox channel to send push notification messages to the sandbox environment of the Apple Push Notification service (APNs). Before you can use Amazon Pinpoint to send notifications to the APNs sandbox environment, you have to enable the APNs sandbox channel for an Amazon Pinpoint application.
//
// The APNSSandboxChannel resource represents the status and authentication settings of the APNs sandbox channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSSandboxChannel := awscdk.Aws_pinpoint.NewCfnAPNSSandboxChannel(this, jsii.String("MyCfnAPNSSandboxChannel"), &cfnAPNSSandboxChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   })
//
type CfnAPNSSandboxChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the APNs sandbox channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId() *string
	SetBundleId(val *string)
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate() *string
	SetCertificate(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	// Specifies whether to enable the APNs Sandbox channel for the Amazon Pinpoint application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey() *string
	SetPrivateKey(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId() *string
	SetTeamId(val *string)
	// The authentication key to use for APNs tokens.
	TokenKey() *string
	SetTokenKey(val *string)
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId() *string
	SetTokenKeyId(val *string)
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

// The jsii proxy struct for CfnAPNSSandboxChannel
type jsiiProxy_CfnAPNSSandboxChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::APNSSandboxChannel`.
func NewCfnAPNSSandboxChannel(scope constructs.Construct, id *string, props *CfnAPNSSandboxChannelProps) CfnAPNSSandboxChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnAPNSSandboxChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::APNSSandboxChannel`.
func NewCfnAPNSSandboxChannel_Override(c CfnAPNSSandboxChannel, scope constructs.Construct, id *string, props *CfnAPNSSandboxChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSSandboxChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAPNSSandboxChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAPNSSandboxChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
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
func CfnAPNSSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPNSSandboxChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSSandboxChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSSandboxChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAPNSSandboxChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSSandboxChannelProps := &cfnAPNSSandboxChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSSandboxChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the APNs sandbox channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the APNs Sandbox channel for the Amazon Pinpoint application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

// A CloudFormation `AWS::Pinpoint::APNSVoipChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the APNs VoIP channel to send VoIP notification messages to the Apple Push Notification service (APNs). Before you can use Amazon Pinpoint to send VoIP notifications to APNs, you have to enable the APNs VoIP channel for an Amazon Pinpoint application.
//
// The APNSVoipChannel resource represents the status and authentication settings of the APNs VoIP channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipChannel := awscdk.Aws_pinpoint.NewCfnAPNSVoipChannel(this, jsii.String("MyCfnAPNSVoipChannel"), &cfnAPNSVoipChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   })
//
type CfnAPNSVoipChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the APNs VoIP channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId() *string
	SetBundleId(val *string)
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate() *string
	SetCertificate(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	// Specifies whether to enable the APNs VoIP channel for the Amazon Pinpoint application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey() *string
	SetPrivateKey(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId() *string
	SetTeamId(val *string)
	// The authentication key to use for APNs tokens.
	TokenKey() *string
	SetTokenKey(val *string)
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId() *string
	SetTokenKeyId(val *string)
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

// The jsii proxy struct for CfnAPNSVoipChannel
type jsiiProxy_CfnAPNSVoipChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAPNSVoipChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::APNSVoipChannel`.
func NewCfnAPNSVoipChannel(scope constructs.Construct, id *string, props *CfnAPNSVoipChannelProps) CfnAPNSVoipChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnAPNSVoipChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::APNSVoipChannel`.
func NewCfnAPNSVoipChannel_Override(c CfnAPNSVoipChannel, scope constructs.Construct, id *string, props *CfnAPNSVoipChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAPNSVoipChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAPNSVoipChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
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
func CfnAPNSVoipChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPNSVoipChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAPNSVoipChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAPNSVoipChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipChannelProps := &cfnAPNSVoipChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSVoipChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the APNs VoIP channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the APNs VoIP channel for the Amazon Pinpoint application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

// A CloudFormation `AWS::Pinpoint::APNSVoipSandboxChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the APNs VoIP sandbox channel to send VoIP notification messages to the sandbox environment of the Apple Push Notification service (APNs). Before you can use Amazon Pinpoint to send VoIP notifications to the APNs sandbox environment, you have to enable the APNs VoIP sandbox channel for an Amazon Pinpoint application.
//
// The APNSVoipSandboxChannel resource represents the status and authentication settings of the APNs VoIP sandbox channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipSandboxChannel := awscdk.Aws_pinpoint.NewCfnAPNSVoipSandboxChannel(this, jsii.String("MyCfnAPNSVoipSandboxChannel"), &cfnAPNSVoipSandboxChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   })
//
type CfnAPNSVoipSandboxChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the application that the APNs VoIP sandbox channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId() *string
	SetBundleId(val *string)
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using an APNs certificate.
	Certificate() *string
	SetCertificate(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod() *string
	SetDefaultAuthenticationMethod(val *string)
	// Specifies whether the APNs VoIP sandbox channel is enabled for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with the APNs sandbox environment.
	PrivateKey() *string
	SetPrivateKey(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The identifier that's assigned to your Apple developer account team.
	//
	// This identifier is used for APNs tokens.
	TeamId() *string
	SetTeamId(val *string)
	// The authentication key to use for APNs tokens.
	TokenKey() *string
	SetTokenKey(val *string)
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using APNs tokens.
	TokenKeyId() *string
	SetTokenKeyId(val *string)
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

// The jsii proxy struct for CfnAPNSVoipSandboxChannel
type jsiiProxy_CfnAPNSVoipSandboxChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) DefaultAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) TokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) TokenKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::APNSVoipSandboxChannel`.
func NewCfnAPNSVoipSandboxChannel(scope constructs.Construct, id *string, props *CfnAPNSVoipSandboxChannelProps) CfnAPNSVoipSandboxChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnAPNSVoipSandboxChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::APNSVoipSandboxChannel`.
func NewCfnAPNSVoipSandboxChannel_Override(c CfnAPNSVoipSandboxChannel, scope constructs.Construct, id *string, props *CfnAPNSVoipSandboxChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetDefaultAuthenticationMethod(val *string) {
	_jsii_.Set(
		j,
		"defaultAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetTeamId(val *string) {
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetTokenKey(val *string) {
	_jsii_.Set(
		j,
		"tokenKey",
		val,
	)
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannel) SetTokenKeyId(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAPNSVoipSandboxChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAPNSVoipSandboxChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
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
func CfnAPNSVoipSandboxChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPNSVoipSandboxChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnAPNSVoipSandboxChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAPNSVoipSandboxChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipSandboxChannelProps := &cfnAPNSVoipSandboxChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSVoipSandboxChannelProps struct {
	// The unique identifier for the application that the APNs VoIP sandbox channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether the APNs VoIP sandbox channel is enabled for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with the APNs sandbox environment.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple developer account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

// A CloudFormation `AWS::Pinpoint::App`.
//
// An *app* is an Amazon Pinpoint application, also referred to as a *project* . An application is a collection of related settings, customer information, segments, campaigns, and other types of Amazon Pinpoint resources.
//
// The App resource represents an Amazon Pinpoint application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnApp := awscdk.Aws_pinpoint.NewCfnApp(this, jsii.String("MyCfnApp"), &cfnAppProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: tags,
//   })
//
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the application.
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
	// The display name of the application.
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
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::App`.
func NewCfnApp(scope constructs.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::App`.
func NewCfnApp_Override(c CfnApp, scope constructs.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp) SetName(val *string) {
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
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
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
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnAppProps := &cfnAppProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: tags,
//   }
//
type CfnAppProps struct {
	// The display name of the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Pinpoint::ApplicationSettings`.
//
// Specifies the settings for an Amazon Pinpoint application. In Amazon Pinpoint, an *application* (also referred to as an *app* or *project* ) is a collection of related settings, customer information, segments, and campaigns, and other types of Amazon Pinpoint resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationSettings := awscdk.Aws_pinpoint.NewCfnApplicationSettings(this, jsii.String("MyCfnApplicationSettings"), &cfnApplicationSettingsProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	quietTime: &quietTimeProperty{
//   		end: jsii.String("end"),
//   		start: jsii.String("start"),
//   	},
//   })
//
type CfnApplicationSettings interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The settings for the Lambda function to use by default as a code hook for campaigns in the application.
	//
	// To override these settings for a specific campaign, use the Campaign resource to define custom Lambda function settings for the campaign.
	CampaignHook() interface{}
	SetCampaignHook(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies whether to enable application-related alarms in Amazon CloudWatch.
	CloudWatchMetricsEnabled() interface{}
	SetCloudWatchMetricsEnabled(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default sending limits for campaigns in the application.
	//
	// To override these limits for a specific campaign, use the Campaign resource to define custom limits for the campaign.
	Limits() interface{}
	SetLimits(val interface{})
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
	// The default quiet time for campaigns in the application.
	//
	// Quiet time is a specific time range when campaigns don't send messages to endpoints, if all the following conditions are met:
	//
	// - The `EndpointDemographic.Timezone` property of the endpoint is set to a valid value.
	//
	// - The current time in the endpoint's time zone is later than or equal to the time specified by the `QuietTime.Start` property for the application (or a campaign that has custom quiet time settings).
	//
	// - The current time in the endpoint's time zone is earlier than or equal to the time specified by the `QuietTime.End` property for the application (or a campaign that has custom quiet time settings).
	//
	// If any of the preceding conditions isn't met, the endpoint will receive messages from a campaign, even if quiet time is enabled.
	//
	// To override the default quiet time settings for a specific campaign, use the Campaign resource to define a custom quiet time for the campaign.
	QuietTime() interface{}
	SetQuietTime(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnApplicationSettings
type jsiiProxy_CfnApplicationSettings struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationSettings) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CampaignHook() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"campaignHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CloudWatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) Limits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) QuietTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quietTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettings) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::ApplicationSettings`.
func NewCfnApplicationSettings(scope constructs.Construct, id *string, props *CfnApplicationSettingsProps) CfnApplicationSettings {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::ApplicationSettings`.
func NewCfnApplicationSettings_Override(c CfnApplicationSettings, scope constructs.Construct, id *string, props *CfnApplicationSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationSettings) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationSettings) SetCampaignHook(val interface{}) {
	_jsii_.Set(
		j,
		"campaignHook",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationSettings) SetCloudWatchMetricsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cloudWatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationSettings) SetLimits(val interface{}) {
	_jsii_.Set(
		j,
		"limits",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationSettings) SetQuietTime(val interface{}) {
	_jsii_.Set(
		j,
		"quietTime",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApplicationSettings_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApplicationSettings_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
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
func CfnApplicationSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationSettings_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnApplicationSettings",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplicationSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationSettings) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the Lambda function to use by default as a code hook for campaigns in the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignHookProperty := &campaignHookProperty{
//   	lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   	mode: jsii.String("mode"),
//   	webUrl: jsii.String("webUrl"),
//   }
//
type CfnApplicationSettings_CampaignHookProperty struct {
	// The name or Amazon Resource Name (ARN) of the Lambda function that Amazon Pinpoint invokes to send messages for campaigns in the application.
	LambdaFunctionName *string `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// The mode that Amazon Pinpoint uses to invoke the Lambda function. Possible values are:.
	//
	// - `FILTER` - Invoke the function to customize the segment that's used by a campaign.
	// - `DELIVERY` - (Deprecated) Previously, invoked the function to send a campaign through a custom channel. This functionality is not supported anymore. To send a campaign through a custom channel, use the `CustomDeliveryConfiguration` and `CampaignCustomMessage` objects of the campaign.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The web URL that Amazon Pinpoint calls to invoke the Lambda function over HTTPS.
	WebUrl *string `field:"optional" json:"webUrl" yaml:"webUrl"`
}

// Specifies the default sending limits for campaigns in the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsProperty := &limitsProperty{
//   	daily: jsii.Number(123),
//   	maximumDuration: jsii.Number(123),
//   	messagesPerSecond: jsii.Number(123),
//   	total: jsii.Number(123),
//   }
//
type CfnApplicationSettings_LimitsProperty struct {
	// The maximum number of messages that a campaign can send to a single endpoint during a 24-hour period.
	//
	// The maximum value is 100.
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// The maximum amount of time, in seconds, that a campaign can attempt to deliver a message after the scheduled start time for the campaign.
	//
	// The minimum value is 60 seconds.
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// The maximum number of messages that a campaign can send each second.
	//
	// The minimum value is 50. The maximum value is 20,000.
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// The maximum number of messages that a campaign can send to a single endpoint during the course of the campaign.
	//
	// The maximum value is 100.
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

// Specifies the start and end times that define a time range when messages aren't sent to endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quietTimeProperty := &quietTimeProperty{
//   	end: jsii.String("end"),
//   	start: jsii.String("start"),
//   }
//
type CfnApplicationSettings_QuietTimeProperty struct {
	// The specific time when quiet time ends.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	End *string `field:"required" json:"end" yaml:"end"`
	// The specific time when quiet time begins.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	Start *string `field:"required" json:"start" yaml:"start"`
}

// Properties for defining a `CfnApplicationSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationSettingsProps := &cfnApplicationSettingsProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	quietTime: &quietTimeProperty{
//   		end: jsii.String("end"),
//   		start: jsii.String("start"),
//   	},
//   }
//
type CfnApplicationSettingsProps struct {
	// The unique identifier for the Amazon Pinpoint application.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The settings for the Lambda function to use by default as a code hook for campaigns in the application.
	//
	// To override these settings for a specific campaign, use the Campaign resource to define custom Lambda function settings for the campaign.
	CampaignHook interface{} `field:"optional" json:"campaignHook" yaml:"campaignHook"`
	// Specifies whether to enable application-related alarms in Amazon CloudWatch.
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// The default sending limits for campaigns in the application.
	//
	// To override these limits for a specific campaign, use the Campaign resource to define custom limits for the campaign.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// The default quiet time for campaigns in the application.
	//
	// Quiet time is a specific time range when campaigns don't send messages to endpoints, if all the following conditions are met:
	//
	// - The `EndpointDemographic.Timezone` property of the endpoint is set to a valid value.
	//
	// - The current time in the endpoint's time zone is later than or equal to the time specified by the `QuietTime.Start` property for the application (or a campaign that has custom quiet time settings).
	//
	// - The current time in the endpoint's time zone is earlier than or equal to the time specified by the `QuietTime.End` property for the application (or a campaign that has custom quiet time settings).
	//
	// If any of the preceding conditions isn't met, the endpoint will receive messages from a campaign, even if quiet time is enabled.
	//
	// To override the default quiet time settings for a specific campaign, use the Campaign resource to define a custom quiet time for the campaign.
	QuietTime interface{} `field:"optional" json:"quietTime" yaml:"quietTime"`
}

// A CloudFormation `AWS::Pinpoint::BaiduChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the Baidu channel to send notifications to the Baidu Cloud Push notification service. Before you can use Amazon Pinpoint to send notifications to the Baidu Cloud Push service, you have to enable the Baidu channel for an Amazon Pinpoint application.
//
// The BaiduChannel resource represents the status and authentication settings of the Baidu channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBaiduChannel := awscdk.Aws_pinpoint.NewCfnBaiduChannel(this, jsii.String("MyCfnBaiduChannel"), &cfnBaiduChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//   	secretKey: jsii.String("secretKey"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   })
//
type CfnBaiduChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The API key that you received from the Baidu Cloud Push service to communicate with the service.
	ApiKey() *string
	SetApiKey(val *string)
	// The unique identifier for the Amazon Pinpoint application that you're configuring the Baidu channel for.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the Baidu channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The secret key that you received from the Baidu Cloud Push service to communicate with the service.
	SecretKey() *string
	SetSecretKey(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnBaiduChannel
type jsiiProxy_CfnBaiduChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBaiduChannel) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBaiduChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::BaiduChannel`.
func NewCfnBaiduChannel(scope constructs.Construct, id *string, props *CfnBaiduChannelProps) CfnBaiduChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnBaiduChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::BaiduChannel`.
func NewCfnBaiduChannel_Override(c CfnBaiduChannel, scope constructs.Construct, id *string, props *CfnBaiduChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBaiduChannel) SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_CfnBaiduChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnBaiduChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnBaiduChannel) SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBaiduChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBaiduChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
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
func CfnBaiduChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBaiduChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnBaiduChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBaiduChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBaiduChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnBaiduChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBaiduChannelProps := &cfnBaiduChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//   	secretKey: jsii.String("secretKey"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBaiduChannelProps struct {
	// The API key that you received from the Baidu Cloud Push service to communicate with the service.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that you're configuring the Baidu channel for.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The secret key that you received from the Baidu Cloud Push service to communicate with the service.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Specifies whether to enable the Baidu channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// A CloudFormation `AWS::Pinpoint::Campaign`.
//
// Specifies the settings for a campaign. A *campaign* is a messaging initiative that engages a specific segment of users for an Amazon Pinpoint application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//   var tags interface{}
//
//   cfnCampaign := awscdk.Aws_pinpoint.NewCfnCampaign(this, jsii.String("MyCfnCampaign"), &cfnCampaignProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//   	schedule: &scheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		eventFilter: &campaignEventFilterProperty{
//   			dimensions: &eventDimensionsProperty{
//   				attributes: attributes,
//   				eventType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				metrics: metrics,
//   			},
//   			filterType: jsii.String("filterType"),
//   		},
//   		frequency: jsii.String("frequency"),
//   		isLocalTime: jsii.Boolean(false),
//   		quietTime: &quietTimeProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   		startTime: jsii.String("startTime"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	segmentId: jsii.String("segmentId"),
//
//   	// the properties below are optional
//   	additionalTreatments: []interface{}{
//   		&writeTreatmentResourceProperty{
//   			customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   				deliveryUri: jsii.String("deliveryUri"),
//   				endpointTypes: []*string{
//   					jsii.String("endpointTypes"),
//   				},
//   			},
//   			messageConfiguration: &messageConfigurationProperty{
//   				admMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				apnsMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				baiduMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				customMessage: &campaignCustomMessageProperty{
//   					data: jsii.String("data"),
//   				},
//   				defaultMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				emailMessage: &campaignEmailMessageProperty{
//   					body: jsii.String("body"),
//   					fromAddress: jsii.String("fromAddress"),
//   					htmlBody: jsii.String("htmlBody"),
//   					title: jsii.String("title"),
//   				},
//   				gcmMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				inAppMessage: &campaignInAppMessageProperty{
//   					content: []interface{}{
//   						&inAppMessageContentProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							bodyConfig: &inAppMessageBodyConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								body: jsii.String("body"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							headerConfig: &inAppMessageHeaderConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								header: jsii.String("header"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							primaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   							secondaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   						},
//   					},
//   					customConfig: customConfig,
//   					layout: jsii.String("layout"),
//   				},
//   				smsMessage: &campaignSmsMessageProperty{
//   					body: jsii.String("body"),
//   					entityId: jsii.String("entityId"),
//   					messageType: jsii.String("messageType"),
//   					originationNumber: jsii.String("originationNumber"),
//   					senderId: jsii.String("senderId"),
//   					templateId: jsii.String("templateId"),
//   				},
//   			},
//   			schedule: &scheduleProperty{
//   				endTime: jsii.String("endTime"),
//   				eventFilter: &campaignEventFilterProperty{
//   					dimensions: &eventDimensionsProperty{
//   						attributes: attributes,
//   						eventType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						metrics: metrics,
//   					},
//   					filterType: jsii.String("filterType"),
//   				},
//   				frequency: jsii.String("frequency"),
//   				isLocalTime: jsii.Boolean(false),
//   				quietTime: &quietTimeProperty{
//   					end: jsii.String("end"),
//   					start: jsii.String("start"),
//   				},
//   				startTime: jsii.String("startTime"),
//   				timeZone: jsii.String("timeZone"),
//   			},
//   			sizePercent: jsii.Number(123),
//   			templateConfiguration: &templateConfigurationProperty{
//   				emailTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				pushTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				smsTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				voiceTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   			},
//   			treatmentDescription: jsii.String("treatmentDescription"),
//   			treatmentName: jsii.String("treatmentName"),
//   		},
//   	},
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	holdoutPercent: jsii.Number(123),
//   	isPaused: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		session: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	messageConfiguration: &messageConfigurationProperty{
//   		admMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		apnsMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		baiduMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		customMessage: &campaignCustomMessageProperty{
//   			data: jsii.String("data"),
//   		},
//   		defaultMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		emailMessage: &campaignEmailMessageProperty{
//   			body: jsii.String("body"),
//   			fromAddress: jsii.String("fromAddress"),
//   			htmlBody: jsii.String("htmlBody"),
//   			title: jsii.String("title"),
//   		},
//   		gcmMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		inAppMessage: &campaignInAppMessageProperty{
//   			content: []interface{}{
//   				&inAppMessageContentProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					bodyConfig: &inAppMessageBodyConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						body: jsii.String("body"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					headerConfig: &inAppMessageHeaderConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						header: jsii.String("header"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					primaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   					secondaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			customConfig: customConfig,
//   			layout: jsii.String("layout"),
//   		},
//   		smsMessage: &campaignSmsMessageProperty{
//   			body: jsii.String("body"),
//   			entityId: jsii.String("entityId"),
//   			messageType: jsii.String("messageType"),
//   			originationNumber: jsii.String("originationNumber"),
//   			senderId: jsii.String("senderId"),
//   			templateId: jsii.String("templateId"),
//   		},
//   	},
//   	priority: jsii.Number(123),
//   	segmentVersion: jsii.Number(123),
//   	tags: tags,
//   	templateConfiguration: &templateConfigurationProperty{
//   		emailTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		pushTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		smsTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		voiceTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	treatmentDescription: jsii.String("treatmentDescription"),
//   	treatmentName: jsii.String("treatmentName"),
//   })
//
type CfnCampaign interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An array of requests that defines additional treatments for the campaign, in addition to the default treatment for the campaign.
	AdditionalTreatments() interface{}
	SetAdditionalTreatments(val interface{})
	// The unique identifier for the Amazon Pinpoint application that the campaign is associated with.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the campaign.
	AttrArn() *string
	// The unique identifier for the campaign.
	AttrCampaignId() *string
	// Specifies the Lambda function to use as a code hook for a campaign.
	CampaignHook() interface{}
	SetCampaignHook(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration() interface{}
	SetCustomDeliveryConfiguration(val interface{})
	// A custom description of the campaign.
	Description() *string
	SetDescription(val *string)
	// The allocated percentage of users (segment members) who shouldn't receive messages from the campaign.
	HoldoutPercent() *float64
	SetHoldoutPercent(val *float64)
	// Specifies whether to pause the campaign.
	//
	// A paused campaign doesn't run unless you resume it by changing this value to `false` . If you restart a campaign, the campaign restarts from the beginning and not at the point you paused it.
	IsPaused() interface{}
	SetIsPaused(val interface{})
	// The messaging limits for the campaign.
	Limits() interface{}
	SetLimits(val interface{})
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
	// The message configuration settings for the campaign.
	MessageConfiguration() interface{}
	SetMessageConfiguration(val interface{})
	// The name of the campaign.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// An integer between 1 and 5, inclusive, that represents the priority of the in-app message campaign, where 1 is the highest priority and 5 is the lowest.
	//
	// If there are multiple messages scheduled to be displayed at the same time, the priority determines the order in which those messages are displayed.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schedule settings for the campaign.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The unique identifier for the segment to associate with the campaign.
	SegmentId() *string
	SetSegmentId(val *string)
	// The version of the segment to associate with the campaign.
	SegmentVersion() *float64
	SetSegmentVersion(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The message template to use for the treatment.
	TemplateConfiguration() interface{}
	SetTemplateConfiguration(val interface{})
	// A custom description of the default treatment for the campaign.
	TreatmentDescription() *string
	SetTreatmentDescription(val *string)
	// A custom name of the default treatment for the campaign, if the campaign has multiple treatments.
	//
	// A *treatment* is a variation of a campaign that's used for A/B testing.
	TreatmentName() *string
	SetTreatmentName(val *string)
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

// The jsii proxy struct for CfnCampaign
type jsiiProxy_CfnCampaign struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCampaign) AdditionalTreatments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalTreatments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrCampaignId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCampaignId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CampaignHook() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"campaignHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CustomDeliveryConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) HoldoutPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdoutPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) IsPaused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPaused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Limits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) MessageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SegmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SegmentVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TemplateConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TreatmentDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatmentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TreatmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::Campaign`.
func NewCfnCampaign(scope constructs.Construct, id *string, props *CfnCampaignProps) CfnCampaign {
	_init_.Initialize()

	j := jsiiProxy_CfnCampaign{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::Campaign`.
func NewCfnCampaign_Override(c CfnCampaign, scope constructs.Construct, id *string, props *CfnCampaignProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCampaign) SetAdditionalTreatments(val interface{}) {
	_jsii_.Set(
		j,
		"additionalTreatments",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetCampaignHook(val interface{}) {
	_jsii_.Set(
		j,
		"campaignHook",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetCustomDeliveryConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"customDeliveryConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetHoldoutPercent(val *float64) {
	_jsii_.Set(
		j,
		"holdoutPercent",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetIsPaused(val interface{}) {
	_jsii_.Set(
		j,
		"isPaused",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetLimits(val interface{}) {
	_jsii_.Set(
		j,
		"limits",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetMessageConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"messageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetSegmentId(val *string) {
	_jsii_.Set(
		j,
		"segmentId",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetSegmentVersion(val *float64) {
	_jsii_.Set(
		j,
		"segmentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetTemplateConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"templateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetTreatmentDescription(val *string) {
	_jsii_.Set(
		j,
		"treatmentDescription",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign) SetTreatmentName(val *string) {
	_jsii_.Set(
		j,
		"treatmentName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCampaign_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCampaign_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
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
func CfnCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaign_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaign) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCampaign) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCampaign) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCampaign) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCampaign) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCampaign) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCampaign) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies attribute-based criteria for including or excluding endpoints from a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDimensionProperty := &attributeDimensionProperty{
//   	attributeType: jsii.String("attributeType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnCampaign_AttributeDimensionProperty struct {
	// The type of segment dimension to use. Valid values are:.
	//
	// - `INCLUSIVE` – endpoints that have attributes matching the values are included in the segment.
	// - `EXCLUSIVE` – endpoints that have attributes matching the values are excluded from the segment.
	// - `CONTAINS` – endpoints that have attributes' substrings match the values are included in the segment.
	// - `BEFORE` – endpoints with attributes read as ISO_INSTANT datetimes before the value are included in the segment.
	// - `AFTER` – endpoints with attributes read as ISO_INSTANT datetimes after the value are included in the segment.
	// - `BETWEEN` – endpoints with attributes read as ISO_INSTANT datetimes between the values are included in the segment.
	// - `ON` – endpoints with attributes read as ISO_INSTANT dates on the value are included in the segment. Time is ignored in this comparison.
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `AttributeType` property, endpoints are included or excluded from the segment if their attribute values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Specifies the contents of a message that's sent through a custom channel to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignCustomMessageProperty := &campaignCustomMessageProperty{
//   	data: jsii.String("data"),
//   }
//
type CfnCampaign_CampaignCustomMessageProperty struct {
	// The raw, JSON-formatted string to use as the payload for the message.
	//
	// The maximum size is 5 KB.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

// Specifies the content and "From" address for an email message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignEmailMessageProperty := &campaignEmailMessageProperty{
//   	body: jsii.String("body"),
//   	fromAddress: jsii.String("fromAddress"),
//   	htmlBody: jsii.String("htmlBody"),
//   	title: jsii.String("title"),
//   }
//
type CfnCampaign_CampaignEmailMessageProperty struct {
	// The body of the email for recipients whose email clients don't render HTML content.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The verified email address to send the email from.
	//
	// The default address is the `FromAddress` specified for the email channel for the application.
	FromAddress *string `field:"optional" json:"fromAddress" yaml:"fromAddress"`
	// The body of the email, in HTML format, for recipients whose email clients render HTML content.
	HtmlBody *string `field:"optional" json:"htmlBody" yaml:"htmlBody"`
	// The subject line, or title, of the email.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

// Specifies the settings for events that cause a campaign to be sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   campaignEventFilterProperty := &campaignEventFilterProperty{
//   	dimensions: &eventDimensionsProperty{
//   		attributes: attributes,
//   		eventType: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		metrics: metrics,
//   	},
//   	filterType: jsii.String("filterType"),
//   }
//
type CfnCampaign_CampaignEventFilterProperty struct {
	// The dimension settings of the event filter for the campaign.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The type of event that causes the campaign to be sent.
	//
	// Valid values are: `SYSTEM` , sends the campaign when a system event occurs; and, `ENDPOINT` , sends the campaign when an endpoint event (Events resource) occurs.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

// Specifies settings for invoking an Lambda function that customizes a segment for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignHookProperty := &campaignHookProperty{
//   	lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   	mode: jsii.String("mode"),
//   	webUrl: jsii.String("webUrl"),
//   }
//
type CfnCampaign_CampaignHookProperty struct {
	// The name or Amazon Resource Name (ARN) of the Lambda function that Amazon Pinpoint invokes to customize a segment for a campaign.
	LambdaFunctionName *string `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// The mode that Amazon Pinpoint uses to invoke the Lambda function. Possible values are:.
	//
	// - `FILTER` - Invoke the function to customize the segment that's used by a campaign.
	// - `DELIVERY` - (Deprecated) Previously, invoked the function to send a campaign through a custom channel. This functionality is not supported anymore. To send a campaign through a custom channel, use the `CustomDeliveryConfiguration` and `CampaignCustomMessage` objects of the campaign.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The web URL that Amazon Pinpoint calls to invoke the Lambda function over HTTPS.
	WebUrl *string `field:"optional" json:"webUrl" yaml:"webUrl"`
}

// Specifies the appearance of an in-app message, including the message type, the title and body text, text and background colors, and the configurations of buttons that appear in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//
//   campaignInAppMessageProperty := &campaignInAppMessageProperty{
//   	content: []interface{}{
//   		&inAppMessageContentProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			bodyConfig: &inAppMessageBodyConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				body: jsii.String("body"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			headerConfig: &inAppMessageHeaderConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				header: jsii.String("header"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			primaryBtn: &inAppMessageButtonProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   			secondaryBtn: &inAppMessageButtonProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	customConfig: customConfig,
//   	layout: jsii.String("layout"),
//   }
//
type CfnCampaign_CampaignInAppMessageProperty struct {
	// An array that contains configurtion information about the in-app message for the campaign, including title and body text, text colors, background colors, image URLs, and button configurations.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that describes how the in-app message will appear. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
}

// Specifies the content and settings for an SMS message that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignSmsMessageProperty := &campaignSmsMessageProperty{
//   	body: jsii.String("body"),
//   	entityId: jsii.String("entityId"),
//   	messageType: jsii.String("messageType"),
//   	originationNumber: jsii.String("originationNumber"),
//   	senderId: jsii.String("senderId"),
//   	templateId: jsii.String("templateId"),
//   }
//
type CfnCampaign_CampaignSmsMessageProperty struct {
	// The body of the SMS message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The entity ID or Principal Entity (PE) id received from the regulatory body for sending SMS in your country.
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The SMS message type.
	//
	// Valid values are `TRANSACTIONAL` (for messages that are critical or time-sensitive, such as a one-time passwords) and `PROMOTIONAL` (for messsages that aren't critical or time-sensitive, such as marketing messages).
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
	// The long code to send the SMS message from.
	//
	// This value should be one of the dedicated long codes that's assigned to your AWS account. Although it isn't required, we recommend that you specify the long code using an E.164 format to ensure prompt and accurate delivery of the message. For example, +12065550100.
	OriginationNumber *string `field:"optional" json:"originationNumber" yaml:"originationNumber"`
	// The alphabetic Sender ID to display as the sender of the message on a recipient's device.
	//
	// Support for sender IDs varies by country or region. To specify a phone number as the sender, omit this parameter and use `OriginationNumber` instead. For more information about support for Sender ID by country, see the [Amazon Pinpoint User Guide](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) .
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The template ID received from the regulatory body for sending SMS in your country.
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

// Specifies the delivery configuration settings for sending a campaign or campaign treatment through a custom channel.
//
// This object is required if you use the `CampaignCustomMessage` object to define the message to send for the campaign or campaign treatment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDeliveryConfigurationProperty := &customDeliveryConfigurationProperty{
//   	deliveryUri: jsii.String("deliveryUri"),
//   	endpointTypes: []*string{
//   		jsii.String("endpointTypes"),
//   	},
//   }
//
type CfnCampaign_CustomDeliveryConfigurationProperty struct {
	// The destination to send the campaign or treatment to. This value can be one of the following:.
	//
	// - The name or Amazon Resource Name (ARN) of an AWS Lambda function to invoke to handle delivery of the campaign or treatment.
	// - The URL for a web application or service that supports HTTPS and can receive the message. The URL has to be a full URL, including the HTTPS protocol.
	DeliveryUri *string `field:"optional" json:"deliveryUri" yaml:"deliveryUri"`
	// The types of endpoints to send the campaign or treatment to.
	//
	// Each valid value maps to a type of channel that you can associate with an endpoint by using the `ChannelType` property of an endpoint.
	EndpointTypes *[]*string `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
}

// Specifies the default behavior for a button that appears in an in-app message.
//
// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultButtonConfigurationProperty := &defaultButtonConfigurationProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	borderRadius: jsii.Number(123),
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   	text: jsii.String("text"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnCampaign_DefaultButtonConfigurationProperty struct {
	// The background color of a button, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The border radius of a button.
	BorderRadius *float64 `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// The text that appears on a button in an in-app message.
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The color of the body text in a button, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the dimensions for an event filter that determines when a campaign is sent or a journey activity is performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   eventDimensionsProperty := &eventDimensionsProperty{
//   	attributes: attributes,
//   	eventType: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	metrics: metrics,
//   }
//
type CfnCampaign_EventDimensionsProperty struct {
	// One or more custom attributes that your application reports to Amazon Pinpoint.
	//
	// You can use these attributes as selection criteria when you create an event filter.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The name of the event that causes the campaign to be sent or the journey activity to be performed.
	//
	// This can be a standard event that Amazon Pinpoint generates, such as `_email.delivered` . For campaigns, this can also be a custom event that's specific to your application. For information about standard events, see [Streaming Amazon Pinpoint Events](https://docs.aws.amazon.com/pinpoint/latest/developerguide/event-streams.html) in the *Amazon Pinpoint Developer Guide* .
	EventType interface{} `field:"optional" json:"eventType" yaml:"eventType"`
	// One or more custom metrics that your application reports to Amazon Pinpoint .
	//
	// You can use these metrics as selection criteria when you create an event filter.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
}

// Specifies the configuration of main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageBodyConfigProperty := &inAppMessageBodyConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	body: jsii.String("body"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnCampaign_InAppMessageBodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the configuration of a button that appears in an in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageButtonProperty := &inAppMessageButtonProperty{
//   	android: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	defaultConfig: &defaultButtonConfigurationProperty{
//   		backgroundColor: jsii.String("backgroundColor"),
//   		borderRadius: jsii.Number(123),
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   		text: jsii.String("text"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	ios: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	web: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   }
//
type CfnCampaign_InAppMessageButtonProperty struct {
	// An object that defines the default behavior for a button in in-app messages sent to Android.
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// An object that defines the default behavior for a button in an in-app message.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// An object that defines the default behavior for a button in in-app messages sent to iOS devices.
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// An object that defines the default behavior for a button in in-app messages for web applications.
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

// Specifies the configuration and contents of an in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageContentProperty := &inAppMessageContentProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	bodyConfig: &inAppMessageBodyConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		body: jsii.String("body"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	headerConfig: &inAppMessageHeaderConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		header: jsii.String("header"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	imageUrl: jsii.String("imageUrl"),
//   	primaryBtn: &inAppMessageButtonProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   	secondaryBtn: &inAppMessageButtonProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   }
//
type CfnCampaign_InAppMessageContentProperty struct {
	// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specifies the configuration of main body text in an in-app message template.
	BodyConfig interface{} `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// Specifies the configuration and content of the header or title text of the in-app message.
	HeaderConfig interface{} `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// The URL of the image that appears on an in-app message banner.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// An object that contains configuration information about the primary button in an in-app message.
	PrimaryBtn interface{} `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// An object that contains configuration information about the secondary button in an in-app message.
	SecondaryBtn interface{} `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}

// Specifies the configuration and content of the header or title text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageHeaderConfigProperty := &inAppMessageHeaderConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	header: jsii.String("header"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnCampaign_InAppMessageHeaderConfigProperty struct {
	// The text alignment of the title of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The header or title text of the in-app message.
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the limits on the messages that a campaign can send.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsProperty := &limitsProperty{
//   	daily: jsii.Number(123),
//   	maximumDuration: jsii.Number(123),
//   	messagesPerSecond: jsii.Number(123),
//   	session: jsii.Number(123),
//   	total: jsii.Number(123),
//   }
//
type CfnCampaign_LimitsProperty struct {
	// The maximum number of messages that a campaign can send to a single endpoint during a 24-hour period.
	//
	// The maximum value is 100.
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// The maximum amount of time, in seconds, that a campaign can attempt to deliver a message after the scheduled start time for the campaign.
	//
	// The minimum value is 60 seconds.
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// The maximum number of messages that a campaign can send each second.
	//
	// The minimum value is 50. The maximum value is 20,000.
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// `CfnCampaign.LimitsProperty.Session`.
	Session *float64 `field:"optional" json:"session" yaml:"session"`
	// The maximum number of messages that a campaign can send to a single endpoint during the course of the campaign.
	//
	// The maximum value is 100.
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

// Specifies the message configuration settings for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//
//   messageConfigurationProperty := &messageConfigurationProperty{
//   	admMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	apnsMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	baiduMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	customMessage: &campaignCustomMessageProperty{
//   		data: jsii.String("data"),
//   	},
//   	defaultMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	emailMessage: &campaignEmailMessageProperty{
//   		body: jsii.String("body"),
//   		fromAddress: jsii.String("fromAddress"),
//   		htmlBody: jsii.String("htmlBody"),
//   		title: jsii.String("title"),
//   	},
//   	gcmMessage: &messageProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		jsonBody: jsii.String("jsonBody"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		rawContent: jsii.String("rawContent"),
//   		silentPush: jsii.Boolean(false),
//   		timeToLive: jsii.Number(123),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	inAppMessage: &campaignInAppMessageProperty{
//   		content: []interface{}{
//   			&inAppMessageContentProperty{
//   				backgroundColor: jsii.String("backgroundColor"),
//   				bodyConfig: &inAppMessageBodyConfigProperty{
//   					alignment: jsii.String("alignment"),
//   					body: jsii.String("body"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				headerConfig: &inAppMessageHeaderConfigProperty{
//   					alignment: jsii.String("alignment"),
//   					header: jsii.String("header"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				imageUrl: jsii.String("imageUrl"),
//   				primaryBtn: &inAppMessageButtonProperty{
//   					android: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					defaultConfig: &defaultButtonConfigurationProperty{
//   						backgroundColor: jsii.String("backgroundColor"),
//   						borderRadius: jsii.Number(123),
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   						text: jsii.String("text"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					ios: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					web: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   				},
//   				secondaryBtn: &inAppMessageButtonProperty{
//   					android: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					defaultConfig: &defaultButtonConfigurationProperty{
//   						backgroundColor: jsii.String("backgroundColor"),
//   						borderRadius: jsii.Number(123),
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   						text: jsii.String("text"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					ios: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   					web: &overrideButtonConfigurationProperty{
//   						buttonAction: jsii.String("buttonAction"),
//   						link: jsii.String("link"),
//   					},
//   				},
//   			},
//   		},
//   		customConfig: customConfig,
//   		layout: jsii.String("layout"),
//   	},
//   	smsMessage: &campaignSmsMessageProperty{
//   		body: jsii.String("body"),
//   		entityId: jsii.String("entityId"),
//   		messageType: jsii.String("messageType"),
//   		originationNumber: jsii.String("originationNumber"),
//   		senderId: jsii.String("senderId"),
//   		templateId: jsii.String("templateId"),
//   	},
//   }
//
type CfnCampaign_MessageConfigurationProperty struct {
	// The message that the campaign sends through the ADM (Amazon Device Messaging) channel.
	//
	// If specified, this message overrides the default message.
	AdmMessage interface{} `field:"optional" json:"admMessage" yaml:"admMessage"`
	// The message that the campaign sends through the APNs (Apple Push Notification service) channel.
	//
	// If specified, this message overrides the default message.
	ApnsMessage interface{} `field:"optional" json:"apnsMessage" yaml:"apnsMessage"`
	// The message that the campaign sends through the Baidu (Baidu Cloud Push) channel.
	//
	// If specified, this message overrides the default message.
	BaiduMessage interface{} `field:"optional" json:"baiduMessage" yaml:"baiduMessage"`
	// The message that the campaign sends through a custom channel, as specified by the delivery configuration ( `CustomDeliveryConfiguration` ) settings for the campaign.
	//
	// If specified, this message overrides the default message.
	CustomMessage interface{} `field:"optional" json:"customMessage" yaml:"customMessage"`
	// The default message that the campaign sends through all the channels that are configured for the campaign.
	DefaultMessage interface{} `field:"optional" json:"defaultMessage" yaml:"defaultMessage"`
	// The message that the campaign sends through the email channel.
	//
	// If specified, this message overrides the default message.
	EmailMessage interface{} `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// The message that the campaign sends through the GCM channel, which enables Amazon Pinpoint to send push notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// If specified, this message overrides the default message.
	GcmMessage interface{} `field:"optional" json:"gcmMessage" yaml:"gcmMessage"`
	// The default message for the in-app messaging channel.
	//
	// This message overrides the default message ( `DefaultMessage` ).
	InAppMessage interface{} `field:"optional" json:"inAppMessage" yaml:"inAppMessage"`
	// The message that the campaign sends through the SMS channel.
	//
	// If specified, this message overrides the default message.
	SmsMessage interface{} `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

// Specifies the content and settings for a push notification that's sent to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &messageProperty{
//   	action: jsii.String("action"),
//   	body: jsii.String("body"),
//   	imageIconUrl: jsii.String("imageIconUrl"),
//   	imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   	imageUrl: jsii.String("imageUrl"),
//   	jsonBody: jsii.String("jsonBody"),
//   	mediaUrl: jsii.String("mediaUrl"),
//   	rawContent: jsii.String("rawContent"),
//   	silentPush: jsii.Boolean(false),
//   	timeToLive: jsii.Number(123),
//   	title: jsii.String("title"),
//   	url: jsii.String("url"),
//   }
//
type CfnCampaign_MessageProperty struct {
	// The action to occur if a recipient taps the push notification. Valid values are:.
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This setting uses the deep-linking features of iOS and Android.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The body of the notification message.
	//
	// The maximum number of characters is 200.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of the image to display as the push notification icon, such as the icon for the app.
	ImageIconUrl *string `field:"optional" json:"imageIconUrl" yaml:"imageIconUrl"`
	// The URL of the image to display as the small, push notification icon, such as a small version of the icon for the app.
	ImageSmallIconUrl *string `field:"optional" json:"imageSmallIconUrl" yaml:"imageSmallIconUrl"`
	// The URL of an image to display in the push notification.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The JSON payload to use for a silent push notification.
	JsonBody *string `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// The URL of the image or video to display in the push notification.
	MediaUrl *string `field:"optional" json:"mediaUrl" yaml:"mediaUrl"`
	// The raw, JSON-formatted string to use as the payload for the notification message.
	//
	// If specified, this value overrides all other content for the message.
	RawContent *string `field:"optional" json:"rawContent" yaml:"rawContent"`
	// Specifies whether the notification is a silent push notification, which is a push notification that doesn't display on a recipient's device.
	//
	// Silent push notifications can be used for cases such as updating an app's configuration, displaying messages in an in-app message center, or supporting phone home functionality.
	SilentPush interface{} `field:"optional" json:"silentPush" yaml:"silentPush"`
	// The number of seconds that the push notification service should keep the message, if the service is unable to deliver the notification the first time.
	//
	// This value is converted to an expiration value when it's sent to a push notification service. If this value is `0` , the service treats the notification as if it expires immediately and the service doesn't store or try to deliver the notification again.
	//
	// This value doesn't apply to messages that are sent through the Amazon Device Messaging (ADM) service.
	TimeToLive *float64 `field:"optional" json:"timeToLive" yaml:"timeToLive"`
	// The title to display above the notification message on a recipient's device.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in a recipient's default mobile browser, if a recipient taps the push notification and the value of the `Action` property is `URL` .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Specifies metric-based criteria for including or excluding endpoints from a segment.
//
// These criteria derive from custom metrics that you define for endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	value: jsii.Number(123),
//   }
//
type CfnCampaign_MetricDimensionProperty struct {
	// The operator to use when comparing metric values.
	//
	// Valid values are: `GREATER_THAN` , `LESS_THAN` , `GREATER_THAN_OR_EQUAL` , `LESS_THAN_OR_EQUAL` , and `EQUAL` .
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value to compare.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// Specifies the configuration of a button with settings that are specific to a certain device type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideButtonConfigurationProperty := &overrideButtonConfigurationProperty{
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   }
//
type CfnCampaign_OverrideButtonConfigurationProperty struct {
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	Link *string `field:"optional" json:"link" yaml:"link"`
}

// Specifies the start and end times that define a time range when messages aren't sent to endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quietTimeProperty := &quietTimeProperty{
//   	end: jsii.String("end"),
//   	start: jsii.String("start"),
//   }
//
type CfnCampaign_QuietTimeProperty struct {
	// The specific time when quiet time ends.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	End *string `field:"required" json:"end" yaml:"end"`
	// The specific time when quiet time begins.
	//
	// This value has to use 24-hour notation and be in HH:MM format, where HH is the hour (with a leading zero, if applicable) and MM is the minutes. For example, use `02:30` to represent 2:30 AM, or `14:30` to represent 2:30 PM.
	Start *string `field:"required" json:"start" yaml:"start"`
}

// Specifies the schedule settings for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   scheduleProperty := &scheduleProperty{
//   	endTime: jsii.String("endTime"),
//   	eventFilter: &campaignEventFilterProperty{
//   		dimensions: &eventDimensionsProperty{
//   			attributes: attributes,
//   			eventType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			metrics: metrics,
//   		},
//   		filterType: jsii.String("filterType"),
//   	},
//   	frequency: jsii.String("frequency"),
//   	isLocalTime: jsii.Boolean(false),
//   	quietTime: &quietTimeProperty{
//   		end: jsii.String("end"),
//   		start: jsii.String("start"),
//   	},
//   	startTime: jsii.String("startTime"),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnCampaign_ScheduleProperty struct {
	// The scheduled time, in ISO 8601 format, when the campaign ended or will end.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The type of event that causes the campaign to be sent, if the value of the `Frequency` property is `EVENT` .
	EventFilter interface{} `field:"optional" json:"eventFilter" yaml:"eventFilter"`
	// Specifies how often the campaign is sent or whether the campaign is sent in response to a specific event.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Specifies whether the start and end times for the campaign schedule use each recipient's local time.
	//
	// To base the schedule on each recipient's local time, set this value to `true` .
	IsLocalTime interface{} `field:"optional" json:"isLocalTime" yaml:"isLocalTime"`
	// The default quiet time for the campaign.
	//
	// Quiet time is a specific time range when a campaign doesn't send messages to endpoints, if all the following conditions are met:
	//
	// - The `EndpointDemographic.Timezone` property of the endpoint is set to a valid value.
	// - The current time in the endpoint's time zone is later than or equal to the time specified by the `QuietTime.Start` property for the campaign.
	// - The current time in the endpoint's time zone is earlier than or equal to the time specified by the `QuietTime.End` property for the campaign.
	//
	// If any of the preceding conditions isn't met, the endpoint will receive messages from the campaign, even if quiet time is enabled.
	QuietTime interface{} `field:"optional" json:"quietTime" yaml:"quietTime"`
	// The scheduled time when the campaign began or will begin.
	//
	// Valid values are: `IMMEDIATE` , to start the campaign immediately; or, a specific time in ISO 8601 format.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The starting UTC offset for the campaign schedule, if the value of the `IsLocalTime` property is `true` .
	//
	// Valid values are: `UTC, UTC+01, UTC+02, UTC+03, UTC+03:30, UTC+04, UTC+04:30, UTC+05, UTC+05:30, UTC+05:45, UTC+06, UTC+06:30, UTC+07, UTC+08, UTC+09, UTC+09:30, UTC+10, UTC+10:30, UTC+11, UTC+12, UTC+13, UTC-02, UTC-03, UTC-04, UTC-05, UTC-06, UTC-07, UTC-08, UTC-09, UTC-10,` and `UTC-11` .
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

// Specifies the dimension type and values for a segment dimension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setDimensionProperty := &setDimensionProperty{
//   	dimensionType: jsii.String("dimensionType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnCampaign_SetDimensionProperty struct {
	// The type of segment dimension to use.
	//
	// Valid values are: `INCLUSIVE` , endpoints that match the criteria are included in the segment; and, `EXCLUSIVE` , endpoints that match the criteria are excluded from the segment.
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `DimensionType` property, endpoints are included or excluded from the segment if their values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Specifies the message template to use for the message, for each type of channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateConfigurationProperty := &templateConfigurationProperty{
//   	emailTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	pushTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	smsTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	voiceTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnCampaign_TemplateConfigurationProperty struct {
	// The email template to use for the message.
	EmailTemplate interface{} `field:"optional" json:"emailTemplate" yaml:"emailTemplate"`
	// The push notification template to use for the message.
	PushTemplate interface{} `field:"optional" json:"pushTemplate" yaml:"pushTemplate"`
	// The SMS template to use for the message.
	SmsTemplate interface{} `field:"optional" json:"smsTemplate" yaml:"smsTemplate"`
	// The voice template to use for the message.
	//
	// This object isn't supported for campaigns.
	VoiceTemplate interface{} `field:"optional" json:"voiceTemplate" yaml:"voiceTemplate"`
}

// Specifies the name and version of the message template to use for the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProperty := &templateProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnCampaign_TemplateProperty struct {
	// The name of the message template to use for the message.
	//
	// If specified, this value must match the name of an existing message template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier for the version of the message template to use for the message.
	//
	// If specified, this value must match the identifier for an existing template version. To retrieve a list of versions and version identifiers for a template, use the Template Versions resource.
	//
	// If you don't specify a value for this property, Amazon Pinpoint uses the *active version* of the template. The *active version* is typically the version of a template that's been most recently reviewed and approved for use, depending on your workflow. It isn't necessarily the latest version of a template.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Specifies the settings for a campaign treatment.
//
// A *treatment* is a variation of a campaign that's used for A/B testing of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//
//   writeTreatmentResourceProperty := &writeTreatmentResourceProperty{
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	messageConfiguration: &messageConfigurationProperty{
//   		admMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		apnsMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		baiduMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		customMessage: &campaignCustomMessageProperty{
//   			data: jsii.String("data"),
//   		},
//   		defaultMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		emailMessage: &campaignEmailMessageProperty{
//   			body: jsii.String("body"),
//   			fromAddress: jsii.String("fromAddress"),
//   			htmlBody: jsii.String("htmlBody"),
//   			title: jsii.String("title"),
//   		},
//   		gcmMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		inAppMessage: &campaignInAppMessageProperty{
//   			content: []interface{}{
//   				&inAppMessageContentProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					bodyConfig: &inAppMessageBodyConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						body: jsii.String("body"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					headerConfig: &inAppMessageHeaderConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						header: jsii.String("header"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					primaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   					secondaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			customConfig: customConfig,
//   			layout: jsii.String("layout"),
//   		},
//   		smsMessage: &campaignSmsMessageProperty{
//   			body: jsii.String("body"),
//   			entityId: jsii.String("entityId"),
//   			messageType: jsii.String("messageType"),
//   			originationNumber: jsii.String("originationNumber"),
//   			senderId: jsii.String("senderId"),
//   			templateId: jsii.String("templateId"),
//   		},
//   	},
//   	schedule: &scheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		eventFilter: &campaignEventFilterProperty{
//   			dimensions: &eventDimensionsProperty{
//   				attributes: attributes,
//   				eventType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				metrics: metrics,
//   			},
//   			filterType: jsii.String("filterType"),
//   		},
//   		frequency: jsii.String("frequency"),
//   		isLocalTime: jsii.Boolean(false),
//   		quietTime: &quietTimeProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   		startTime: jsii.String("startTime"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	sizePercent: jsii.Number(123),
//   	templateConfiguration: &templateConfigurationProperty{
//   		emailTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		pushTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		smsTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		voiceTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	treatmentDescription: jsii.String("treatmentDescription"),
//   	treatmentName: jsii.String("treatmentName"),
//   }
//
type CfnCampaign_WriteTreatmentResourceProperty struct {
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// The message configuration settings for the treatment.
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// The schedule settings for the treatment.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The allocated percentage of users (segment members) to send the treatment to.
	SizePercent *float64 `field:"optional" json:"sizePercent" yaml:"sizePercent"`
	// The message template to use for the treatment.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the treatment.
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name for the treatment.
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//   var tags interface{}
//
//   cfnCampaignProps := &cfnCampaignProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//   	schedule: &scheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		eventFilter: &campaignEventFilterProperty{
//   			dimensions: &eventDimensionsProperty{
//   				attributes: attributes,
//   				eventType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				metrics: metrics,
//   			},
//   			filterType: jsii.String("filterType"),
//   		},
//   		frequency: jsii.String("frequency"),
//   		isLocalTime: jsii.Boolean(false),
//   		quietTime: &quietTimeProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   		startTime: jsii.String("startTime"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	segmentId: jsii.String("segmentId"),
//
//   	// the properties below are optional
//   	additionalTreatments: []interface{}{
//   		&writeTreatmentResourceProperty{
//   			customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   				deliveryUri: jsii.String("deliveryUri"),
//   				endpointTypes: []*string{
//   					jsii.String("endpointTypes"),
//   				},
//   			},
//   			messageConfiguration: &messageConfigurationProperty{
//   				admMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				apnsMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				baiduMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				customMessage: &campaignCustomMessageProperty{
//   					data: jsii.String("data"),
//   				},
//   				defaultMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				emailMessage: &campaignEmailMessageProperty{
//   					body: jsii.String("body"),
//   					fromAddress: jsii.String("fromAddress"),
//   					htmlBody: jsii.String("htmlBody"),
//   					title: jsii.String("title"),
//   				},
//   				gcmMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				inAppMessage: &campaignInAppMessageProperty{
//   					content: []interface{}{
//   						&inAppMessageContentProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							bodyConfig: &inAppMessageBodyConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								body: jsii.String("body"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							headerConfig: &inAppMessageHeaderConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								header: jsii.String("header"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							primaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   							secondaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   						},
//   					},
//   					customConfig: customConfig,
//   					layout: jsii.String("layout"),
//   				},
//   				smsMessage: &campaignSmsMessageProperty{
//   					body: jsii.String("body"),
//   					entityId: jsii.String("entityId"),
//   					messageType: jsii.String("messageType"),
//   					originationNumber: jsii.String("originationNumber"),
//   					senderId: jsii.String("senderId"),
//   					templateId: jsii.String("templateId"),
//   				},
//   			},
//   			schedule: &scheduleProperty{
//   				endTime: jsii.String("endTime"),
//   				eventFilter: &campaignEventFilterProperty{
//   					dimensions: &eventDimensionsProperty{
//   						attributes: attributes,
//   						eventType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						metrics: metrics,
//   					},
//   					filterType: jsii.String("filterType"),
//   				},
//   				frequency: jsii.String("frequency"),
//   				isLocalTime: jsii.Boolean(false),
//   				quietTime: &quietTimeProperty{
//   					end: jsii.String("end"),
//   					start: jsii.String("start"),
//   				},
//   				startTime: jsii.String("startTime"),
//   				timeZone: jsii.String("timeZone"),
//   			},
//   			sizePercent: jsii.Number(123),
//   			templateConfiguration: &templateConfigurationProperty{
//   				emailTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				pushTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				smsTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				voiceTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   			},
//   			treatmentDescription: jsii.String("treatmentDescription"),
//   			treatmentName: jsii.String("treatmentName"),
//   		},
//   	},
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	holdoutPercent: jsii.Number(123),
//   	isPaused: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		session: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	messageConfiguration: &messageConfigurationProperty{
//   		admMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		apnsMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		baiduMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		customMessage: &campaignCustomMessageProperty{
//   			data: jsii.String("data"),
//   		},
//   		defaultMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		emailMessage: &campaignEmailMessageProperty{
//   			body: jsii.String("body"),
//   			fromAddress: jsii.String("fromAddress"),
//   			htmlBody: jsii.String("htmlBody"),
//   			title: jsii.String("title"),
//   		},
//   		gcmMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		inAppMessage: &campaignInAppMessageProperty{
//   			content: []interface{}{
//   				&inAppMessageContentProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					bodyConfig: &inAppMessageBodyConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						body: jsii.String("body"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					headerConfig: &inAppMessageHeaderConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						header: jsii.String("header"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					primaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   					secondaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			customConfig: customConfig,
//   			layout: jsii.String("layout"),
//   		},
//   		smsMessage: &campaignSmsMessageProperty{
//   			body: jsii.String("body"),
//   			entityId: jsii.String("entityId"),
//   			messageType: jsii.String("messageType"),
//   			originationNumber: jsii.String("originationNumber"),
//   			senderId: jsii.String("senderId"),
//   			templateId: jsii.String("templateId"),
//   		},
//   	},
//   	priority: jsii.Number(123),
//   	segmentVersion: jsii.Number(123),
//   	tags: tags,
//   	templateConfiguration: &templateConfigurationProperty{
//   		emailTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		pushTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		smsTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		voiceTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	treatmentDescription: jsii.String("treatmentDescription"),
//   	treatmentName: jsii.String("treatmentName"),
//   }
//
type CfnCampaignProps struct {
	// The unique identifier for the Amazon Pinpoint application that the campaign is associated with.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the campaign.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule settings for the campaign.
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The unique identifier for the segment to associate with the campaign.
	SegmentId *string `field:"required" json:"segmentId" yaml:"segmentId"`
	// An array of requests that defines additional treatments for the campaign, in addition to the default treatment for the campaign.
	AdditionalTreatments interface{} `field:"optional" json:"additionalTreatments" yaml:"additionalTreatments"`
	// Specifies the Lambda function to use as a code hook for a campaign.
	CampaignHook interface{} `field:"optional" json:"campaignHook" yaml:"campaignHook"`
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// A custom description of the campaign.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The allocated percentage of users (segment members) who shouldn't receive messages from the campaign.
	HoldoutPercent *float64 `field:"optional" json:"holdoutPercent" yaml:"holdoutPercent"`
	// Specifies whether to pause the campaign.
	//
	// A paused campaign doesn't run unless you resume it by changing this value to `false` . If you restart a campaign, the campaign restarts from the beginning and not at the point you paused it.
	IsPaused interface{} `field:"optional" json:"isPaused" yaml:"isPaused"`
	// The messaging limits for the campaign.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// The message configuration settings for the campaign.
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// An integer between 1 and 5, inclusive, that represents the priority of the in-app message campaign, where 1 is the highest priority and 5 is the lowest.
	//
	// If there are multiple messages scheduled to be displayed at the same time, the priority determines the order in which those messages are displayed.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The version of the segment to associate with the campaign.
	SegmentVersion *float64 `field:"optional" json:"segmentVersion" yaml:"segmentVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The message template to use for the treatment.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the default treatment for the campaign.
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name of the default treatment for the campaign, if the campaign has multiple treatments.
	//
	// A *treatment* is a variation of a campaign that's used for A/B testing.
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

// A CloudFormation `AWS::Pinpoint::EmailChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the email channel to send email to users. Before you can use Amazon Pinpoint to send email, you must enable the email channel for an Amazon Pinpoint application.
//
// The EmailChannel resource represents the status, identity, and other settings of the email channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailChannel := awscdk.Aws_pinpoint.NewCfnEmailChannel(this, jsii.String("MyCfnEmailChannel"), &cfnEmailChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	fromAddress: jsii.String("fromAddress"),
//   	identity: jsii.String("identity"),
//
//   	// the properties below are optional
//   	configurationSet: jsii.String("configurationSet"),
//   	enabled: jsii.Boolean(false),
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnEmailChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that you're specifying the email channel for.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The [Amazon SES configuration set](https://docs.aws.amazon.com/ses/latest/APIReference/API_ConfigurationSet.html) that you want to apply to messages that you send through the channel.
	ConfigurationSet() *string
	SetConfigurationSet(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the email channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
	// The verified email address that you want to send email from when you send email through the channel.
	FromAddress() *string
	SetFromAddress(val *string)
	// The Amazon Resource Name (ARN) of the identity, verified with Amazon Simple Email Service (Amazon SES), that you want to use when you send email through the channel.
	Identity() *string
	SetIdentity(val *string)
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
	// The ARN of the AWS Identity and Access Management (IAM) role that you want Amazon Pinpoint to use when it submits email-related event data for the channel.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnEmailChannel
type jsiiProxy_CfnEmailChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEmailChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) FromAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::EmailChannel`.
func NewCfnEmailChannel(scope constructs.Construct, id *string, props *CfnEmailChannelProps) CfnEmailChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnEmailChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::EmailChannel`.
func NewCfnEmailChannel_Override(c CfnEmailChannel, scope constructs.Construct, id *string, props *CfnEmailChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetConfigurationSet(val *string) {
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetFromAddress(val *string) {
	_jsii_.Set(
		j,
		"fromAddress",
		val,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetIdentity(val *string) {
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_CfnEmailChannel) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEmailChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEmailChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
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
func CfnEmailChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnEmailChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEmailChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEmailChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEmailChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEmailChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEmailChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEmailChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEmailChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEmailChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEmailChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEmailChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailChannelProps := &cfnEmailChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	fromAddress: jsii.String("fromAddress"),
//   	identity: jsii.String("identity"),
//
//   	// the properties below are optional
//   	configurationSet: jsii.String("configurationSet"),
//   	enabled: jsii.Boolean(false),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnEmailChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that you're specifying the email channel for.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The verified email address that you want to send email from when you send email through the channel.
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
	// The Amazon Resource Name (ARN) of the identity, verified with Amazon Simple Email Service (Amazon SES), that you want to use when you send email through the channel.
	Identity *string `field:"required" json:"identity" yaml:"identity"`
	// The [Amazon SES configuration set](https://docs.aws.amazon.com/ses/latest/APIReference/API_ConfigurationSet.html) that you want to apply to messages that you send through the channel.
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether to enable the email channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the AWS Identity and Access Management (IAM) role that you want Amazon Pinpoint to use when it submits email-related event data for the channel.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::Pinpoint::EmailTemplate`.
//
// Creates a message template that you can use in messages that are sent through the email channel. A *message template* is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnEmailTemplate := awscdk.Aws_pinpoint.NewCfnEmailTemplate(this, jsii.String("MyCfnEmailTemplate"), &cfnEmailTemplateProps{
//   	subject: jsii.String("subject"),
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	htmlPart: jsii.String("htmlPart"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   	textPart: jsii.String("textPart"),
//   })
//
type CfnEmailTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the message template.
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
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions() *string
	SetDefaultSubstitutions(val *string)
	// The message body, in HTML format, to use in email messages that are based on the message template.
	//
	// We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	HtmlPart() *string
	SetHtmlPart(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The subject line, or title, to use in email messages that are based on the message template.
	Subject() *string
	SetSubject(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// A custom description of the message template.
	TemplateDescription() *string
	SetTemplateDescription(val *string)
	// The name of the message template.
	TemplateName() *string
	SetTemplateName(val *string)
	// The message body, in plain text format, to use in email messages that are based on the message template.
	//
	// We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	TextPart() *string
	SetTextPart(val *string)
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

// The jsii proxy struct for CfnEmailTemplate
type jsiiProxy_CfnEmailTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEmailTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) HtmlPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) TemplateDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) TextPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::EmailTemplate`.
func NewCfnEmailTemplate(scope constructs.Construct, id *string, props *CfnEmailTemplateProps) CfnEmailTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnEmailTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::EmailTemplate`.
func NewCfnEmailTemplate_Override(c CfnEmailTemplate, scope constructs.Construct, id *string, props *CfnEmailTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetDefaultSubstitutions(val *string) {
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetHtmlPart(val *string) {
	_jsii_.Set(
		j,
		"htmlPart",
		val,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetSubject(val *string) {
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetTemplateDescription(val *string) {
	_jsii_.Set(
		j,
		"templateDescription",
		val,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetTemplateName(val *string) {
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_CfnEmailTemplate) SetTextPart(val *string) {
	_jsii_.Set(
		j,
		"textPart",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEmailTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEmailTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
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
func CfnEmailTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnEmailTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEmailTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEmailTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnEmailTemplateProps := &cfnEmailTemplateProps{
//   	subject: jsii.String("subject"),
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	htmlPart: jsii.String("htmlPart"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   	textPart: jsii.String("textPart"),
//   }
//
type CfnEmailTemplateProps struct {
	// The subject line, or title, to use in email messages that are based on the message template.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The name of the message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// The message body, in HTML format, to use in email messages that are based on the message template.
	//
	// We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
	// The message body, in plain text format, to use in email messages that are based on the message template.
	//
	// We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

// A CloudFormation `AWS::Pinpoint::EventStream`.
//
// Creates a new event stream for an application or updates the settings of an existing event stream for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventStream := awscdk.Aws_pinpoint.NewCfnEventStream(this, jsii.String("MyCfnEventStream"), &cfnEventStreamProps{
//   	applicationId: jsii.String("applicationId"),
//   	destinationStreamArn: jsii.String("destinationStreamArn"),
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnEventStream interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that you want to export data from.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the Amazon Kinesis data stream or Amazon Kinesis Data Firehose delivery stream that you want to publish event data to.
	//
	// For a Kinesis data stream, the ARN format is: `arn:aws:kinesis: region : account-id :stream/ stream_name`
	//
	// For a Kinesis Data Firehose delivery stream, the ARN format is: `arn:aws:firehose: region : account-id :deliverystream/ stream_name`.
	DestinationStreamArn() *string
	SetDestinationStreamArn(val *string)
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
	// The AWS Identity and Access Management (IAM) role that authorizes Amazon Pinpoint to publish event data to the stream in your AWS account.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnEventStream
type jsiiProxy_CfnEventStream struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventStream) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) DestinationStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventStream) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::EventStream`.
func NewCfnEventStream(scope constructs.Construct, id *string, props *CfnEventStreamProps) CfnEventStream {
	_init_.Initialize()

	j := jsiiProxy_CfnEventStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::EventStream`.
func NewCfnEventStream_Override(c CfnEventStream, scope constructs.Construct, id *string, props *CfnEventStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventStream) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnEventStream) SetDestinationStreamArn(val *string) {
	_jsii_.Set(
		j,
		"destinationStreamArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventStream) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventStream_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventStream_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
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
func CfnEventStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventStream_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnEventStream",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventStream) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventStream) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventStream) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventStream) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventStream) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventStream) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEventStream) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventStream) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventStream) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventStream) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventStream) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventStream) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventStream) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEventStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventStreamProps := &cfnEventStreamProps{
//   	applicationId: jsii.String("applicationId"),
//   	destinationStreamArn: jsii.String("destinationStreamArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnEventStreamProps struct {
	// The unique identifier for the Amazon Pinpoint application that you want to export data from.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis data stream or Amazon Kinesis Data Firehose delivery stream that you want to publish event data to.
	//
	// For a Kinesis data stream, the ARN format is: `arn:aws:kinesis: region : account-id :stream/ stream_name`
	//
	// For a Kinesis Data Firehose delivery stream, the ARN format is: `arn:aws:firehose: region : account-id :deliverystream/ stream_name`.
	DestinationStreamArn *string `field:"required" json:"destinationStreamArn" yaml:"destinationStreamArn"`
	// The AWS Identity and Access Management (IAM) role that authorizes Amazon Pinpoint to publish event data to the stream in your AWS account.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::Pinpoint::GCMChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. You can use the GCM channel to send push notification messages to the Firebase Cloud Messaging (FCM) service, which replaced the Google Cloud Messaging (GCM) service. Before you use Amazon Pinpoint to send notifications to FCM, you have to enable the GCM channel for an Amazon Pinpoint application.
//
// The GCMChannel resource represents the status and authentication settings of the GCM channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGCMChannel := awscdk.Aws_pinpoint.NewCfnGCMChannel(this, jsii.String("MyCfnGCMChannel"), &cfnGCMChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   })
//
type CfnGCMChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Web API key, also called the *server key* , that you received from Google to communicate with Google services.
	ApiKey() *string
	SetApiKey(val *string)
	// The unique identifier for the Amazon Pinpoint application that the GCM channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the GCM channel for the Amazon Pinpoint application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnGCMChannel
type jsiiProxy_CfnGCMChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGCMChannel) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGCMChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::GCMChannel`.
func NewCfnGCMChannel(scope constructs.Construct, id *string, props *CfnGCMChannelProps) CfnGCMChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnGCMChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::GCMChannel`.
func NewCfnGCMChannel_Override(c CfnGCMChannel, scope constructs.Construct, id *string, props *CfnGCMChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGCMChannel) SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_CfnGCMChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnGCMChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGCMChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGCMChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
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
func CfnGCMChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGCMChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnGCMChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGCMChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGCMChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGCMChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGCMChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGCMChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGCMChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGCMChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGCMChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGCMChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGCMChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGCMChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGCMChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGCMChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGCMChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGCMChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGCMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGCMChannelProps := &cfnGCMChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnGCMChannelProps struct {
	// The Web API key, also called the *server key* , that you received from Google to communicate with Google services.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that the GCM channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the GCM channel for the Amazon Pinpoint application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// A CloudFormation `AWS::Pinpoint::InAppTemplate`.
//
// Creates a message template that you can use to send in-app messages. A message template is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplate := awscdk.Aws_pinpoint.NewCfnInAppTemplate(this, jsii.String("MyCfnInAppTemplate"), &cfnInAppTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	content: []interface{}{
//   		&inAppMessageContentProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			bodyConfig: &bodyConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				body: jsii.String("body"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			headerConfig: &headerConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				header: jsii.String("header"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			primaryBtn: &buttonConfigProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   			secondaryBtn: &buttonConfigProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	customConfig: customConfig,
//   	layout: jsii.String("layout"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   })
//
type CfnInAppTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the message template.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An object that contains information about the content of an in-app message, including its title and body text, text colors, background colors, images, buttons, and behaviors.
	Content() interface{}
	SetContent(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig() interface{}
	SetCustomConfig(val interface{})
	// A string that determines the appearance of the in-app message. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	Layout() *string
	SetLayout(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// An optional description of the in-app template.
	TemplateDescription() *string
	SetTemplateDescription(val *string)
	// The name of the in-app message template.
	TemplateName() *string
	SetTemplateName(val *string)
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

// The jsii proxy struct for CfnInAppTemplate
type jsiiProxy_CfnInAppTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInAppTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CustomConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Layout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) TemplateDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplate(scope constructs.Construct, id *string, props *CfnInAppTemplateProps) CfnInAppTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnInAppTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplate_Override(c CfnInAppTemplate, scope constructs.Construct, id *string, props *CfnInAppTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInAppTemplate) SetContent(val interface{}) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate) SetCustomConfig(val interface{}) {
	_jsii_.Set(
		j,
		"customConfig",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate) SetLayout(val *string) {
	_jsii_.Set(
		j,
		"layout",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate) SetTemplateDescription(val *string) {
	_jsii_.Set(
		j,
		"templateDescription",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate) SetTemplateName(val *string) {
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInAppTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInAppTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
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
func CfnInAppTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInAppTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the configuration of the main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyConfigProperty := &bodyConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	body: jsii.String("body"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnInAppTemplate_BodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the behavior of buttons that appear in an in-app message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buttonConfigProperty := &buttonConfigProperty{
//   	android: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	defaultConfig: &defaultButtonConfigurationProperty{
//   		backgroundColor: jsii.String("backgroundColor"),
//   		borderRadius: jsii.Number(123),
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   		text: jsii.String("text"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	ios: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	web: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   }
//
type CfnInAppTemplate_ButtonConfigProperty struct {
	// Optional button configuration to use for in-app messages sent to Android devices.
	//
	// This button configuration overrides the default button configuration.
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// Specifies the default behavior of a button that appears in an in-app message.
	//
	// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// Optional button configuration to use for in-app messages sent to iOS devices.
	//
	// This button configuration overrides the default button configuration.
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// Optional button configuration to use for in-app messages sent to web applications.
	//
	// This button configuration overrides the default button configuration.
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

// Specifies the default behavior of a button that appears in an in-app message.
//
// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultButtonConfigurationProperty := &defaultButtonConfigurationProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	borderRadius: jsii.Number(123),
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   	text: jsii.String("text"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnInAppTemplate_DefaultButtonConfigurationProperty struct {
	// The background color of a button, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The border radius of a button.
	BorderRadius *float64 `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// The text that appears on a button in an in-app message.
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The color of the body text in a button, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the configuration and content of the header or title text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerConfigProperty := &headerConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	header: jsii.String("header"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnInAppTemplate_HeaderConfigProperty struct {
	// The text alignment of the title of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The title text of the in-app message.
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The color of the title text, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

// Specifies the configuration of an in-app message, including its header, body, buttons, colors, and images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageContentProperty := &inAppMessageContentProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	bodyConfig: &bodyConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		body: jsii.String("body"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	headerConfig: &headerConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		header: jsii.String("header"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	imageUrl: jsii.String("imageUrl"),
//   	primaryBtn: &buttonConfigProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   	secondaryBtn: &buttonConfigProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   }
//
type CfnInAppTemplate_InAppMessageContentProperty struct {
	// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// An object that contains configuration information about the header or title text of the in-app message.
	BodyConfig interface{} `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// An object that contains configuration information about the header or title text of the in-app message.
	HeaderConfig interface{} `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// The URL of the image that appears on an in-app message banner.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// An object that contains configuration information about the primary button in an in-app message.
	PrimaryBtn interface{} `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// An object that contains configuration information about the secondary button in an in-app message.
	SecondaryBtn interface{} `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}

// Specifies the configuration of a button with settings that are specific to a certain device type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideButtonConfigurationProperty := &overrideButtonConfigurationProperty{
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   }
//
type CfnInAppTemplate_OverrideButtonConfigurationProperty struct {
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	Link *string `field:"optional" json:"link" yaml:"link"`
}

// Properties for defining a `CfnInAppTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplateProps := &cfnInAppTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	content: []interface{}{
//   		&inAppMessageContentProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			bodyConfig: &bodyConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				body: jsii.String("body"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			headerConfig: &headerConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				header: jsii.String("header"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			primaryBtn: &buttonConfigProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   			secondaryBtn: &buttonConfigProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	customConfig: customConfig,
//   	layout: jsii.String("layout"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnInAppTemplateProps struct {
	// The name of the in-app message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// An object that contains information about the content of an in-app message, including its title and body text, text colors, background colors, images, buttons, and behaviors.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that determines the appearance of the in-app message. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An optional description of the in-app template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

// A CloudFormation `AWS::Pinpoint::PushTemplate`.
//
// Creates a message template that you can use in messages that are sent through a push notification channel. A *message template* is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnPushTemplate := awscdk.Aws_pinpoint.NewCfnPushTemplate(this, jsii.String("MyCfnPushTemplate"), &cfnPushTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	adm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	apns: &aPNSPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	baidu: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	default: &defaultPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	gcm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   })
//
type CfnPushTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The message template to use for the ADM (Amazon Device Messaging) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Adm() interface{}
	SetAdm(val interface{})
	// The message template to use for the APNs (Apple Push Notification service) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Apns() interface{}
	SetApns(val interface{})
	// The Amazon Resource Name (ARN) of the message template.
	AttrArn() *string
	// The message template to use for the Baidu (Baidu Cloud Push) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Baidu() interface{}
	SetBaidu(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default message template to use for push notification channels.
	Default() interface{}
	SetDefault(val interface{})
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions() *string
	SetDefaultSubstitutions(val *string)
	// The message template to use for the GCM channel, which is used to send notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Gcm() interface{}
	SetGcm(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// A custom description of the message template.
	TemplateDescription() *string
	SetTemplateDescription(val *string)
	// The name of the message template.
	TemplateName() *string
	SetTemplateName(val *string)
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

// The jsii proxy struct for CfnPushTemplate
type jsiiProxy_CfnPushTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPushTemplate) Adm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Apns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Baidu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baidu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Gcm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) TemplateDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::PushTemplate`.
func NewCfnPushTemplate(scope constructs.Construct, id *string, props *CfnPushTemplateProps) CfnPushTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnPushTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::PushTemplate`.
func NewCfnPushTemplate_Override(c CfnPushTemplate, scope constructs.Construct, id *string, props *CfnPushTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetAdm(val interface{}) {
	_jsii_.Set(
		j,
		"adm",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetApns(val interface{}) {
	_jsii_.Set(
		j,
		"apns",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetBaidu(val interface{}) {
	_jsii_.Set(
		j,
		"baidu",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetDefault(val interface{}) {
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetDefaultSubstitutions(val *string) {
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetGcm(val interface{}) {
	_jsii_.Set(
		j,
		"gcm",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetTemplateDescription(val *string) {
	_jsii_.Set(
		j,
		"templateDescription",
		val,
	)
}

func (j *jsiiProxy_CfnPushTemplate) SetTemplateName(val *string) {
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPushTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPushTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
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
func CfnPushTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPushTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnPushTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPushTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPushTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPushTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPushTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPushTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPushTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPushTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPushTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPushTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPushTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPushTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPushTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPushTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPushTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPushTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies channel-specific content and settings for a message template that can be used in push notifications that are sent through the APNs (Apple Push Notification service) channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPNSPushNotificationTemplateProperty := &aPNSPushNotificationTemplateProperty{
//   	action: jsii.String("action"),
//   	body: jsii.String("body"),
//   	mediaUrl: jsii.String("mediaUrl"),
//   	sound: jsii.String("sound"),
//   	title: jsii.String("title"),
//   	url: jsii.String("url"),
//   }
//
type CfnPushTemplate_APNSPushNotificationTemplateProperty struct {
	// The action to occur if a recipient taps a push notification that's based on the message template.
	//
	// Valid values are:
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This setting uses the deep-linking features of the iOS platform.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The message body to use in push notifications that are based on the message template.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of an image or video to display in push notifications that are based on the message template.
	MediaUrl *string `field:"optional" json:"mediaUrl" yaml:"mediaUrl"`
	// The key for the sound to play when the recipient receives a push notification that's based on the message template.
	//
	// The value for this key is the name of a sound file in your app's main bundle or the `Library/Sounds` folder in your app's data container. If the sound file can't be found or you specify `default` for the value, the system plays the default alert sound.
	Sound *string `field:"optional" json:"sound" yaml:"sound"`
	// The title to use in push notifications that are based on the message template.
	//
	// This title appears above the notification message on a recipient's device.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in the recipient's default mobile browser, if a recipient taps a push notification that's based on the message template and the value of the `Action` property is `URL` .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Specifies channel-specific content and settings for a message template that can be used in push notifications that are sent through the ADM (Amazon Device Messaging), Baidu (Baidu Cloud Push), or GCM (Firebase Cloud Messaging, formerly Google Cloud Messaging) channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   androidPushNotificationTemplateProperty := &androidPushNotificationTemplateProperty{
//   	action: jsii.String("action"),
//   	body: jsii.String("body"),
//   	imageIconUrl: jsii.String("imageIconUrl"),
//   	imageUrl: jsii.String("imageUrl"),
//   	smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   	sound: jsii.String("sound"),
//   	title: jsii.String("title"),
//   	url: jsii.String("url"),
//   }
//
type CfnPushTemplate_AndroidPushNotificationTemplateProperty struct {
	// The action to occur if a recipient taps a push notification that's based on the message template.
	//
	// Valid values are:
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This action uses the deep-linking features of the Android platform.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The message body to use in a push notification that's based on the message template.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The URL of the large icon image to display in the content view of a push notification that's based on the message template.
	ImageIconUrl *string `field:"optional" json:"imageIconUrl" yaml:"imageIconUrl"`
	// The URL of an image to display in a push notification that's based on the message template.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The URL of the small icon image to display in the status bar and the content view of a push notification that's based on the message template.
	SmallImageIconUrl *string `field:"optional" json:"smallImageIconUrl" yaml:"smallImageIconUrl"`
	// The sound to play when a recipient receives a push notification that's based on the message template.
	//
	// You can use the default stream or specify the file name of a sound resource that's bundled in your app. On an Android platform, the sound file must reside in `/res/raw/` .
	Sound *string `field:"optional" json:"sound" yaml:"sound"`
	// The title to use in a push notification that's based on the message template.
	//
	// This title appears above the notification message on a recipient's device.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in a recipient's default mobile browser, if a recipient taps a push notification that's based on the message template and the value of the `Action` property is `URL` .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Specifies the default settings and content for a message template that can be used in messages that are sent through a push notification channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultPushNotificationTemplateProperty := &defaultPushNotificationTemplateProperty{
//   	action: jsii.String("action"),
//   	body: jsii.String("body"),
//   	sound: jsii.String("sound"),
//   	title: jsii.String("title"),
//   	url: jsii.String("url"),
//   }
//
type CfnPushTemplate_DefaultPushNotificationTemplateProperty struct {
	// The action to occur if a recipient taps a push notification that's based on the message template.
	//
	// Valid values are:
	//
	// - `OPEN_APP` – Your app opens or it becomes the foreground app if it was sent to the background. This is the default action.
	// - `DEEP_LINK` – Your app opens and displays a designated user interface in the app. This setting uses the deep-linking features of the iOS and Android platforms.
	// - `URL` – The default mobile browser on the recipient's device opens and loads the web page at a URL that you specify.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The message body to use in push notifications that are based on the message template.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The sound to play when a recipient receives a push notification that's based on the message template.
	//
	// You can use the default stream or specify the file name of a sound resource that's bundled in your app. On an Android platform, the sound file must reside in `/res/raw/` .
	//
	// For an iOS platform, this value is the key for the name of a sound file in your app's main bundle or the `Library/Sounds` folder in your app's data container. If the sound file can't be found or you specify `default` for the value, the system plays the default alert sound.
	Sound *string `field:"optional" json:"sound" yaml:"sound"`
	// The title to use in push notifications that are based on the message template.
	//
	// This title appears above the notification message on a recipient's device.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL to open in a recipient's default mobile browser, if a recipient taps a push notification that's based on the message template and the value of the `Action` property is `URL` .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Properties for defining a `CfnPushTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnPushTemplateProps := &cfnPushTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	adm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	apns: &aPNSPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		mediaUrl: jsii.String("mediaUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	baidu: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	default: &defaultPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	gcm: &androidPushNotificationTemplateProperty{
//   		action: jsii.String("action"),
//   		body: jsii.String("body"),
//   		imageIconUrl: jsii.String("imageIconUrl"),
//   		imageUrl: jsii.String("imageUrl"),
//   		smallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		sound: jsii.String("sound"),
//   		title: jsii.String("title"),
//   		url: jsii.String("url"),
//   	},
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnPushTemplateProps struct {
	// The name of the message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The message template to use for the ADM (Amazon Device Messaging) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Adm interface{} `field:"optional" json:"adm" yaml:"adm"`
	// The message template to use for the APNs (Apple Push Notification service) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Apns interface{} `field:"optional" json:"apns" yaml:"apns"`
	// The message template to use for the Baidu (Baidu Cloud Push) channel.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Baidu interface{} `field:"optional" json:"baidu" yaml:"baidu"`
	// The default message template to use for push notification channels.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// The message template to use for the GCM channel, which is used to send notifications through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This message template overrides the default template for push notification channels ( `Default` ).
	Gcm interface{} `field:"optional" json:"gcm" yaml:"gcm"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

// A CloudFormation `AWS::Pinpoint::SMSChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. To send an SMS text message, you send the message through the SMS channel. Before you can use Amazon Pinpoint to send text messages, you have to enable the SMS channel for an Amazon Pinpoint application.
//
// The SMSChannel resource represents the status, sender ID, and other settings for the SMS channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSMSChannel := awscdk.Aws_pinpoint.NewCfnSMSChannel(this, jsii.String("MyCfnSMSChannel"), &cfnSMSChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	senderId: jsii.String("senderId"),
//   	shortCode: jsii.String("shortCode"),
//   })
//
type CfnSMSChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the SMS channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the SMS channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The identity that you want to display on recipients' devices when they receive messages from the SMS channel.
	//
	// > SenderIDs are only supported in certain countries and regions. For more information, see [Supported Countries and Regions](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) in the *Amazon Pinpoint User Guide* .
	SenderId() *string
	SetSenderId(val *string)
	// The registered short code that you want to use when you send messages through the SMS channel.
	//
	// > For information about obtaining a dedicated short code for sending SMS messages, see [Requesting Dedicated Short Codes for SMS Messaging with Amazon Pinpoint](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-short-code.html) in the *Amazon Pinpoint User Guide* .
	ShortCode() *string
	SetShortCode(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnSMSChannel
type jsiiProxy_CfnSMSChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSMSChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) SenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) ShortCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::SMSChannel`.
func NewCfnSMSChannel(scope constructs.Construct, id *string, props *CfnSMSChannelProps) CfnSMSChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnSMSChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::SMSChannel`.
func NewCfnSMSChannel_Override(c CfnSMSChannel, scope constructs.Construct, id *string, props *CfnSMSChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSMSChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnSMSChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnSMSChannel) SetSenderId(val *string) {
	_jsii_.Set(
		j,
		"senderId",
		val,
	)
}

func (j *jsiiProxy_CfnSMSChannel) SetShortCode(val *string) {
	_jsii_.Set(
		j,
		"shortCode",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSMSChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSMSChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
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
func CfnSMSChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSMSChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnSMSChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSMSChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSMSChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSMSChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSMSChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSMSChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSMSChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSMSChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSMSChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMSChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMSChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSMSChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSMSChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMSChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMSChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMSChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSMSChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSMSChannelProps := &cfnSMSChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	senderId: jsii.String("senderId"),
//   	shortCode: jsii.String("shortCode"),
//   }
//
type CfnSMSChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the SMS channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the SMS channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The identity that you want to display on recipients' devices when they receive messages from the SMS channel.
	//
	// > SenderIDs are only supported in certain countries and regions. For more information, see [Supported Countries and Regions](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) in the *Amazon Pinpoint User Guide* .
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The registered short code that you want to use when you send messages through the SMS channel.
	//
	// > For information about obtaining a dedicated short code for sending SMS messages, see [Requesting Dedicated Short Codes for SMS Messaging with Amazon Pinpoint](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-short-code.html) in the *Amazon Pinpoint User Guide* .
	ShortCode *string `field:"optional" json:"shortCode" yaml:"shortCode"`
}

// A CloudFormation `AWS::Pinpoint::Segment`.
//
// Updates the configuration, dimension, and other settings for an existing segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var tags interface{}
//   var userAttributes interface{}
//
//   cfnSegment := awscdk.Aws_pinpoint.NewCfnSegment(this, jsii.String("MyCfnSegment"), &cfnSegmentProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensions: &segmentDimensionsProperty{
//   		attributes: attributes,
//   		behavior: &behaviorProperty{
//   			recency: &recencyProperty{
//   				duration: jsii.String("duration"),
//   				recencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		demographic: &demographicProperty{
//   			appVersion: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			channel: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			deviceType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			make: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			model: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			platform: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		location: &locationProperty{
//   			country: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			gpsPoint: &gPSPointProperty{
//   				coordinates: &coordinatesProperty{
//   					latitude: jsii.Number(123),
//   					longitude: jsii.Number(123),
//   				},
//   				rangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		metrics: metrics,
//   		userAttributes: userAttributes,
//   	},
//   	segmentGroups: &segmentGroupsProperty{
//   		groups: []interface{}{
//   			&groupsProperty{
//   				dimensions: []interface{}{
//   					&segmentDimensionsProperty{
//   						attributes: attributes,
//   						behavior: &behaviorProperty{
//   							recency: &recencyProperty{
//   								duration: jsii.String("duration"),
//   								recencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						demographic: &demographicProperty{
//   							appVersion: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							channel: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							deviceType: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							make: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							model: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							platform: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						location: &locationProperty{
//   							country: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							gpsPoint: &gPSPointProperty{
//   								coordinates: &coordinatesProperty{
//   									latitude: jsii.Number(123),
//   									longitude: jsii.Number(123),
//   								},
//   								rangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						metrics: metrics,
//   						userAttributes: userAttributes,
//   					},
//   				},
//   				sourceSegments: []interface{}{
//   					&sourceSegmentsProperty{
//   						id: jsii.String("id"),
//
//   						// the properties below are optional
//   						version: jsii.Number(123),
//   					},
//   				},
//   				sourceType: jsii.String("sourceType"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		include: jsii.String("include"),
//   	},
//   	tags: tags,
//   })
//
type CfnSegment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the segment is associated with.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the segment.
	AttrArn() *string
	// The unique identifier for the segment.
	AttrSegmentId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The criteria that define the dimensions for the segment.
	Dimensions() interface{}
	SetDimensions(val interface{})
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
	// The name of the segment.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The segment group to use and the dimensions to apply to the group's base segments in order to build the segment.
	//
	// A segment group can consist of zero or more base segments. Your request can include only one segment group.
	SegmentGroups() interface{}
	SetSegmentGroups(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnSegment
type jsiiProxy_CfnSegment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSegment) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) AttrSegmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSegmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Dimensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) SegmentGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::Segment`.
func NewCfnSegment(scope constructs.Construct, id *string, props *CfnSegmentProps) CfnSegment {
	_init_.Initialize()

	j := jsiiProxy_CfnSegment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::Segment`.
func NewCfnSegment_Override(c CfnSegment, scope constructs.Construct, id *string, props *CfnSegmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSegment) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetDimensions(val interface{}) {
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetSegmentGroups(val interface{}) {
	_jsii_.Set(
		j,
		"segmentGroups",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSegment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSegment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
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
func CfnSegment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSegment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSegment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSegment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSegment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSegment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSegment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSegment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSegment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSegment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSegment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSegment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies attribute-based criteria for including or excluding endpoints from a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDimensionProperty := &attributeDimensionProperty{
//   	attributeType: jsii.String("attributeType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnSegment_AttributeDimensionProperty struct {
	// The type of segment dimension to use. Valid values are:.
	//
	// - `INCLUSIVE` – endpoints that have attributes matching the values are included in the segment.
	// - `EXCLUSIVE` – endpoints that have attributes matching the values are excluded from the segment.
	// - `CONTAINS` – endpoints that have attributes' substrings match the values are included in the segment.
	// - `BEFORE` – endpoints with attributes read as ISO_INSTANT datetimes before the value are included in the segment.
	// - `AFTER` – endpoints with attributes read as ISO_INSTANT datetimes after the value are included in the segment.
	// - `BETWEEN` – endpoints with attributes read as ISO_INSTANT datetimes between the values are included in the segment.
	// - `ON` – endpoints with attributes read as ISO_INSTANT dates on the value are included in the segment. Time is ignored in this comparison.
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `AttributeType` property, endpoints are included or excluded from the segment if their attribute values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Specifies behavior-based criteria for the segment, such as how recently users have used your app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   behaviorProperty := &behaviorProperty{
//   	recency: &recencyProperty{
//   		duration: jsii.String("duration"),
//   		recencyType: jsii.String("recencyType"),
//   	},
//   }
//
type CfnSegment_BehaviorProperty struct {
	// Specifies how recently segment members were active.
	Recency interface{} `field:"optional" json:"recency" yaml:"recency"`
}

// Specifies the GPS coordinates of a location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coordinatesProperty := &coordinatesProperty{
//   	latitude: jsii.Number(123),
//   	longitude: jsii.Number(123),
//   }
//
type CfnSegment_CoordinatesProperty struct {
	// The latitude coordinate of the location.
	Latitude *float64 `field:"required" json:"latitude" yaml:"latitude"`
	// The longitude coordinate of the location.
	Longitude *float64 `field:"required" json:"longitude" yaml:"longitude"`
}

// Specifies demographic-based criteria, such as device platform, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demographicProperty := &demographicProperty{
//   	appVersion: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	channel: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	deviceType: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	make: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	model: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	platform: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
type CfnSegment_DemographicProperty struct {
	// The app version criteria for the segment.
	AppVersion interface{} `field:"optional" json:"appVersion" yaml:"appVersion"`
	// The channel criteria for the segment.
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// The device type criteria for the segment.
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// The device make criteria for the segment.
	Make interface{} `field:"optional" json:"make" yaml:"make"`
	// The device model criteria for the segment.
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// The device platform criteria for the segment.
	Platform interface{} `field:"optional" json:"platform" yaml:"platform"`
}

// Specifies the GPS coordinates of the endpoint location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gPSPointProperty := &gPSPointProperty{
//   	coordinates: &coordinatesProperty{
//   		latitude: jsii.Number(123),
//   		longitude: jsii.Number(123),
//   	},
//   	rangeInKilometers: jsii.Number(123),
//   }
//
type CfnSegment_GPSPointProperty struct {
	// The GPS coordinates to measure distance from.
	Coordinates interface{} `field:"required" json:"coordinates" yaml:"coordinates"`
	// The range, in kilometers, from the GPS coordinates.
	RangeInKilometers *float64 `field:"required" json:"rangeInKilometers" yaml:"rangeInKilometers"`
}

// An array that defines the set of segment criteria to evaluate when handling segment groups for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var userAttributes interface{}
//
//   groupsProperty := &groupsProperty{
//   	dimensions: []interface{}{
//   		&segmentDimensionsProperty{
//   			attributes: attributes,
//   			behavior: &behaviorProperty{
//   				recency: &recencyProperty{
//   					duration: jsii.String("duration"),
//   					recencyType: jsii.String("recencyType"),
//   				},
//   			},
//   			demographic: &demographicProperty{
//   				appVersion: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				channel: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				deviceType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				make: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				model: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				platform: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			location: &locationProperty{
//   				country: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				gpsPoint: &gPSPointProperty{
//   					coordinates: &coordinatesProperty{
//   						latitude: jsii.Number(123),
//   						longitude: jsii.Number(123),
//   					},
//   					rangeInKilometers: jsii.Number(123),
//   				},
//   			},
//   			metrics: metrics,
//   			userAttributes: userAttributes,
//   		},
//   	},
//   	sourceSegments: []interface{}{
//   		&sourceSegmentsProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			version: jsii.Number(123),
//   		},
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	type: jsii.String("type"),
//   }
//
type CfnSegment_GroupsProperty struct {
	// An array that defines the dimensions to include or exclude from the segment.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The base segment to build the segment on.
	//
	// A base segment, also called a *source segment* , defines the initial population of endpoints for a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base segment by using the dimensions that you specify.
	//
	// You can specify more than one dimensional segment or only one imported segment. If you specify an imported segment, the segment size estimate that displays on the Amazon Pinpoint console indicates the size of the imported segment without any filters applied to it.
	SourceSegments interface{} `field:"optional" json:"sourceSegments" yaml:"sourceSegments"`
	// Specifies how to handle multiple base segments for the segment.
	//
	// For example, if you specify three base segments for the segment, whether the resulting segment is based on all, any, or none of the base segments.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Specifies how to handle multiple dimensions for the segment.
	//
	// For example, if you specify three dimensions for the segment, whether the resulting segment includes endpoints that match all, any, or none of the dimensions.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Specifies location-based criteria, such as region or GPS coordinates, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	country: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	gpsPoint: &gPSPointProperty{
//   		coordinates: &coordinatesProperty{
//   			latitude: jsii.Number(123),
//   			longitude: jsii.Number(123),
//   		},
//   		rangeInKilometers: jsii.Number(123),
//   	},
//   }
//
type CfnSegment_LocationProperty struct {
	// The country or region code, in ISO 3166-1 alpha-2 format, for the segment.
	Country interface{} `field:"optional" json:"country" yaml:"country"`
	// The GPS point dimension for the segment.
	GpsPoint interface{} `field:"optional" json:"gpsPoint" yaml:"gpsPoint"`
}

// Specifies how recently segment members were active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recencyProperty := &recencyProperty{
//   	duration: jsii.String("duration"),
//   	recencyType: jsii.String("recencyType"),
//   }
//
type CfnSegment_RecencyProperty struct {
	// The duration to use when determining which users have been active or inactive with your app.
	//
	// Possible values: `HR_24` | `DAY_7` | `DAY_14` | `DAY_30` .
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// The type of recency dimension to use for the segment.
	//
	// Valid values are: `ACTIVE` and `INACTIVE` . If the value is `ACTIVE` , the segment includes users who have used your app within the specified duration are included in the segment. If the value is `INACTIVE` , the segment includes users who haven't used your app within the specified duration are included in the segment.
	RecencyType *string `field:"required" json:"recencyType" yaml:"recencyType"`
}

// Specifies the dimension settings for a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var userAttributes interface{}
//
//   segmentDimensionsProperty := &segmentDimensionsProperty{
//   	attributes: attributes,
//   	behavior: &behaviorProperty{
//   		recency: &recencyProperty{
//   			duration: jsii.String("duration"),
//   			recencyType: jsii.String("recencyType"),
//   		},
//   	},
//   	demographic: &demographicProperty{
//   		appVersion: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		channel: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		deviceType: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		make: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		model: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		platform: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	location: &locationProperty{
//   		country: &setDimensionProperty{
//   			dimensionType: jsii.String("dimensionType"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		gpsPoint: &gPSPointProperty{
//   			coordinates: &coordinatesProperty{
//   				latitude: jsii.Number(123),
//   				longitude: jsii.Number(123),
//   			},
//   			rangeInKilometers: jsii.Number(123),
//   		},
//   	},
//   	metrics: metrics,
//   	userAttributes: userAttributes,
//   }
//
type CfnSegment_SegmentDimensionsProperty struct {
	// One or more custom attributes to use as criteria for the segment.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The behavior-based criteria, such as how recently users have used your app, for the segment.
	Behavior interface{} `field:"optional" json:"behavior" yaml:"behavior"`
	// The demographic-based criteria, such as device platform, for the segment.
	Demographic interface{} `field:"optional" json:"demographic" yaml:"demographic"`
	// The location-based criteria, such as region or GPS coordinates, for the segment.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// One or more custom metrics to use as criteria for the segment.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// One or more custom user attributes to use as criteria for the segment.
	UserAttributes interface{} `field:"optional" json:"userAttributes" yaml:"userAttributes"`
}

// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var userAttributes interface{}
//
//   segmentGroupsProperty := &segmentGroupsProperty{
//   	groups: []interface{}{
//   		&groupsProperty{
//   			dimensions: []interface{}{
//   				&segmentDimensionsProperty{
//   					attributes: attributes,
//   					behavior: &behaviorProperty{
//   						recency: &recencyProperty{
//   							duration: jsii.String("duration"),
//   							recencyType: jsii.String("recencyType"),
//   						},
//   					},
//   					demographic: &demographicProperty{
//   						appVersion: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						channel: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						deviceType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						make: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						model: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						platform: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   					location: &locationProperty{
//   						country: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						gpsPoint: &gPSPointProperty{
//   							coordinates: &coordinatesProperty{
//   								latitude: jsii.Number(123),
//   								longitude: jsii.Number(123),
//   							},
//   							rangeInKilometers: jsii.Number(123),
//   						},
//   					},
//   					metrics: metrics,
//   					userAttributes: userAttributes,
//   				},
//   			},
//   			sourceSegments: []interface{}{
//   				&sourceSegmentsProperty{
//   					id: jsii.String("id"),
//
//   					// the properties below are optional
//   					version: jsii.Number(123),
//   				},
//   			},
//   			sourceType: jsii.String("sourceType"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	include: jsii.String("include"),
//   }
//
type CfnSegment_SegmentGroupsProperty struct {
	// Specifies the set of segment criteria to evaluate when handling segment groups for the segment.
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Specifies how to handle multiple segment groups for the segment.
	//
	// For example, if the segment includes three segment groups, whether the resulting segment includes endpoints that match all, any, or none of the segment groups.
	Include *string `field:"optional" json:"include" yaml:"include"`
}

// Specifies the dimension type and values for a segment dimension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setDimensionProperty := &setDimensionProperty{
//   	dimensionType: jsii.String("dimensionType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnSegment_SetDimensionProperty struct {
	// The type of segment dimension to use.
	//
	// Valid values are: `INCLUSIVE` , endpoints that match the criteria are included in the segment; and, `EXCLUSIVE` , endpoints that match the criteria are excluded from the segment.
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `DimensionType` property, endpoints are included or excluded from the segment if their values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Specifies the base segment to build the segment on.
//
// A base segment, also called a *source segment* , defines the initial population of endpoints for a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base segment by using the dimensions that you specify.
//
// You can specify more than one dimensional segment or only one imported segment. If you specify an imported segment, the segment size estimate that displays on the Amazon Pinpoint console indicates the size of the imported segment without any filters applied to it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceSegmentsProperty := &sourceSegmentsProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnSegment_SourceSegmentsProperty struct {
	// The unique identifier for the source segment.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The version number of the source segment.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// Properties for defining a `CfnSegment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var tags interface{}
//   var userAttributes interface{}
//
//   cfnSegmentProps := &cfnSegmentProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensions: &segmentDimensionsProperty{
//   		attributes: attributes,
//   		behavior: &behaviorProperty{
//   			recency: &recencyProperty{
//   				duration: jsii.String("duration"),
//   				recencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		demographic: &demographicProperty{
//   			appVersion: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			channel: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			deviceType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			make: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			model: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			platform: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		location: &locationProperty{
//   			country: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			gpsPoint: &gPSPointProperty{
//   				coordinates: &coordinatesProperty{
//   					latitude: jsii.Number(123),
//   					longitude: jsii.Number(123),
//   				},
//   				rangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		metrics: metrics,
//   		userAttributes: userAttributes,
//   	},
//   	segmentGroups: &segmentGroupsProperty{
//   		groups: []interface{}{
//   			&groupsProperty{
//   				dimensions: []interface{}{
//   					&segmentDimensionsProperty{
//   						attributes: attributes,
//   						behavior: &behaviorProperty{
//   							recency: &recencyProperty{
//   								duration: jsii.String("duration"),
//   								recencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						demographic: &demographicProperty{
//   							appVersion: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							channel: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							deviceType: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							make: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							model: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							platform: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						location: &locationProperty{
//   							country: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							gpsPoint: &gPSPointProperty{
//   								coordinates: &coordinatesProperty{
//   									latitude: jsii.Number(123),
//   									longitude: jsii.Number(123),
//   								},
//   								rangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						metrics: metrics,
//   						userAttributes: userAttributes,
//   					},
//   				},
//   				sourceSegments: []interface{}{
//   					&sourceSegmentsProperty{
//   						id: jsii.String("id"),
//
//   						// the properties below are optional
//   						version: jsii.Number(123),
//   					},
//   				},
//   				sourceType: jsii.String("sourceType"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		include: jsii.String("include"),
//   	},
//   	tags: tags,
//   }
//
type CfnSegmentProps struct {
	// The unique identifier for the Amazon Pinpoint application that the segment is associated with.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the segment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The criteria that define the dimensions for the segment.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The segment group to use and the dimensions to apply to the group's base segments in order to build the segment.
	//
	// A segment group can consist of zero or more base segments. Your request can include only one segment group.
	SegmentGroups interface{} `field:"optional" json:"segmentGroups" yaml:"segmentGroups"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Pinpoint::SmsTemplate`.
//
// Creates a message template that you can use in messages that are sent through the SMS channel. A *message template* is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSmsTemplate := awscdk.Aws_pinpoint.NewCfnSmsTemplate(this, jsii.String("MyCfnSmsTemplate"), &cfnSmsTemplateProps{
//   	body: jsii.String("body"),
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   })
//
type CfnSmsTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the message template.
	AttrArn() *string
	// The message body to use in text messages that are based on the message template.
	Body() *string
	SetBody(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions() *string
	SetDefaultSubstitutions(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// A custom description of the message template.
	TemplateDescription() *string
	SetTemplateDescription(val *string)
	// The name of the message template.
	TemplateName() *string
	SetTemplateName(val *string)
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

// The jsii proxy struct for CfnSmsTemplate
type jsiiProxy_CfnSmsTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSmsTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) TemplateDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSmsTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::SmsTemplate`.
func NewCfnSmsTemplate(scope constructs.Construct, id *string, props *CfnSmsTemplateProps) CfnSmsTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnSmsTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::SmsTemplate`.
func NewCfnSmsTemplate_Override(c CfnSmsTemplate, scope constructs.Construct, id *string, props *CfnSmsTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSmsTemplate) SetBody(val *string) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_CfnSmsTemplate) SetDefaultSubstitutions(val *string) {
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CfnSmsTemplate) SetTemplateDescription(val *string) {
	_jsii_.Set(
		j,
		"templateDescription",
		val,
	)
}

func (j *jsiiProxy_CfnSmsTemplate) SetTemplateName(val *string) {
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSmsTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSmsTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
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
func CfnSmsTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSmsTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnSmsTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSmsTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSmsTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSmsTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSmsTemplateProps := &cfnSmsTemplateProps{
//   	body: jsii.String("body"),
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	defaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnSmsTemplateProps struct {
	// The message body to use in text messages that are based on the message template.
	Body *string `field:"required" json:"body" yaml:"body"`
	// The name of the message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// A JSON object that specifies the default values to use for message variables in the message template.
	//
	// This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A custom description of the message template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

// A CloudFormation `AWS::Pinpoint::VoiceChannel`.
//
// A *channel* is a type of platform that you can deliver messages to. To send a voice message, you send the message through the voice channel. Before you can use Amazon Pinpoint to send voice messages, you have to enable the voice channel for an Amazon Pinpoint application.
//
// The VoiceChannel resource represents the status and other information about the voice channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVoiceChannel := awscdk.Aws_pinpoint.NewCfnVoiceChannel(this, jsii.String("MyCfnVoiceChannel"), &cfnVoiceChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   })
//
type CfnVoiceChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the voice channel applies to.
	ApplicationId() *string
	SetApplicationId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to enable the voice channel for the application.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnVoiceChannel
type jsiiProxy_CfnVoiceChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVoiceChannel) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVoiceChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::VoiceChannel`.
func NewCfnVoiceChannel(scope constructs.Construct, id *string, props *CfnVoiceChannelProps) CfnVoiceChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnVoiceChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::VoiceChannel`.
func NewCfnVoiceChannel_Override(c CfnVoiceChannel, scope constructs.Construct, id *string, props *CfnVoiceChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVoiceChannel) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnVoiceChannel) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVoiceChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnVoiceChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
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
func CfnVoiceChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVoiceChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnVoiceChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVoiceChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVoiceChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnVoiceChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVoiceChannelProps := &cfnVoiceChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnVoiceChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the voice channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the voice channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

