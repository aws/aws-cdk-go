package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// AddHeaderAction configuration.
//
// TODO: EXAMPLE
//
type AddHeaderActionConfig struct {
	// The name of the header that you want to add to the incoming message.
	HeaderName *string `json:"headerName" yaml:"headerName"`
	// The content that you want to include in the header.
	HeaderValue *string `json:"headerValue" yaml:"headerValue"`
}

// An allow list receipt filter.
//
// TODO: EXAMPLE
//
type AllowListReceiptFilter interface {
	constructs.Construct
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for AllowListReceiptFilter
type jsiiProxy_AllowListReceiptFilter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AllowListReceiptFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAllowListReceiptFilter(scope constructs.Construct, id *string, props *AllowListReceiptFilterProps) AllowListReceiptFilter {
	_init_.Initialize()

	j := jsiiProxy_AllowListReceiptFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAllowListReceiptFilter_Override(a AllowListReceiptFilter, scope constructs.Construct, id *string, props *AllowListReceiptFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AllowListReceiptFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.AllowListReceiptFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AllowListReceiptFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for am AllowListReceiptFilter.
//
// TODO: EXAMPLE
//
type AllowListReceiptFilterProps struct {
	// A list of ip addresses or ranges to allow list.
	Ips *[]*string `json:"ips" yaml:"ips"`
}

// BoundAction configuration.
//
// TODO: EXAMPLE
//
type BounceActionConfig struct {
	// Human-readable text to include in the bounce message.
	Message *string `json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address that the bounce message is sent from.
	Sender *string `json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by RFC 5321.
	SmtpReplyCode *string `json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// A CloudFormation `AWS::SES::ConfigurationSet`.
//
// The name of the configuration set.
//
// Configuration sets let you create groups of rules that you can apply to the emails you send using Amazon SES. For more information about using configuration sets, see [Using Amazon SES Configuration Sets](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/using-configuration-sets.html) in the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/) .
//
// TODO: EXAMPLE
//
type CfnConfigurationSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
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

// The jsii proxy struct for CfnConfigurationSet
type jsiiProxy_CfnConfigurationSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ConfigurationSet`.
func NewCfnConfigurationSet(scope constructs.Construct, id *string, props *CfnConfigurationSetProps) CfnConfigurationSet {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ConfigurationSet`.
func NewCfnConfigurationSet_Override(c CfnConfigurationSet, scope constructs.Construct, id *string, props *CfnConfigurationSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetName(val *string) {
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
func CfnConfigurationSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConfigurationSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
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
func CfnConfigurationSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnConfigurationSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnConfigurationSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConfigurationSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConfigurationSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfigurationSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnConfigurationSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnConfigurationSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfigurationSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnConfigurationSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnConfigurationSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnConfigurationSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnConfigurationSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnConfigurationSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnConfigurationSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::SES::ConfigurationSetEventDestination`.
//
// Specifies a configuration set event destination. An event destination is an AWS service that Amazon SES publishes email sending events to. When you specify an event destination, you provide one, and only one, destination. You can send event data to Amazon CloudWatch or Amazon Kinesis Data Firehose.
//
// > You can't specify Amazon SNS event destinations in CloudFormation templates.
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestination interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigurationSetName() *string
	SetConfigurationSetName(val *string)
	CreationStack() *[]*string
	EventDestination() interface{}
	SetEventDestination(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
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

// The jsii proxy struct for CfnConfigurationSetEventDestination
type jsiiProxy_CfnConfigurationSetEventDestination struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) EventDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestination(scope constructs.Construct, id *string, props *CfnConfigurationSetEventDestinationProps) CfnConfigurationSetEventDestination {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationSetEventDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestination_Override(c CfnConfigurationSetEventDestination, scope constructs.Construct, id *string, props *CfnConfigurationSetEventDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) SetConfigurationSetName(val *string) {
	_jsii_.Set(
		j,
		"configurationSetName",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) SetEventDestination(val interface{}) {
	_jsii_.Set(
		j,
		"eventDestination",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConfigurationSetEventDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConfigurationSetEventDestination_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
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
func CfnConfigurationSetEventDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSetEventDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnConfigurationSetEventDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnConfigurationSetEventDestination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnConfigurationSetEventDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information associated with an Amazon CloudWatch event destination to which email sending events are published.
//
// Event destinations, such as Amazon CloudWatch, are associated with configuration sets, which enable you to publish email sending events. For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html) .
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
	DimensionConfigurations interface{} `json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

// Contains the dimension configuration to use when you publish email sending events to Amazon CloudWatch.
//
// For information about publishing email sending events to Amazon CloudWatch, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html) .
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestination_DimensionConfigurationProperty struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.
	//
	// The default value must meet the following requirements:
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), dashes (-), at signs (@), or periods (.).
	// - Contain 256 characters or fewer.
	DefaultDimensionValue *string `json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	//
	// The name must meet the following requirements:
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 256 characters or fewer.
	DimensionName *string `json:"dimensionName" yaml:"dimensionName"`
	// The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch.
	//
	// To use the message tags that you specify using an `X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` / `SendRawEmail` API, specify `messageTag` . To use your own email headers, specify `emailHeader` . To put a custom tag on any link included in your email, specify `linkTag` .
	DimensionValueSource *string `json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

// Contains information about an event destination.
//
// > When you create or update an event destination, you must provide one, and only one, destination. The destination can be Amazon CloudWatch, Amazon Kinesis Firehose or Amazon Simple Notification Service (Amazon SNS).
//
// Event destinations are associated with configuration sets, which enable you to publish email sending events to Amazon CloudWatch, Amazon Kinesis Firehose, or Amazon Simple Notification Service (Amazon SNS). For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html) .
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestination_EventDestinationProperty struct {
	// The type of email sending events to publish to the event destination.
	//
	// - `send` - The call was successful and Amazon SES is attempting to deliver the email.
	// - `reject` - Amazon SES determined that the email contained a virus and rejected it.
	// - `bounce` - The recipient's mail server permanently rejected the email. This corresponds to a hard bounce.
	// - `complaint` - The recipient marked the email as spam.
	// - `delivery` - Amazon SES successfully delivered the email to the recipient's mail server.
	// - `open` - The recipient received the email and opened it in their email client.
	// - `click` - The recipient clicked one or more links in the email.
	// - `renderingFailure` - Amazon SES did not send the email because of a template rendering issue.
	MatchingEventTypes *[]*string `json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
	CloudWatchDestination interface{} `json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set.
	//
	// Set to `true` to enable publishing to this destination; set to `false` to prevent publishing to this destination. The default value is `false` .
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
	KinesisFirehoseDestination interface{} `json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// The name of the event destination. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Contain 64 characters or fewer.
	Name *string `json:"name" yaml:"name"`
}

// Contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
//
// Event destinations, such as Amazon Kinesis Firehose, are associated with configuration sets, which enable you to publish email sending events. For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html) .
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestination_KinesisFirehoseDestinationProperty struct {
	// The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.
	DeliveryStreamArn *string `json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.
	IamRoleArn *string `json:"iamRoleArn" yaml:"iamRoleArn"`
}

// Properties for defining a `CfnConfigurationSetEventDestination`.
//
// TODO: EXAMPLE
//
type CfnConfigurationSetEventDestinationProps struct {
	// The name of the configuration set that contains the event destination.
	ConfigurationSetName *string `json:"configurationSetName" yaml:"configurationSetName"`
	// The event destination object.
	EventDestination interface{} `json:"eventDestination" yaml:"eventDestination"`
}

// Properties for defining a `CfnConfigurationSet`.
//
// TODO: EXAMPLE
//
type CfnConfigurationSetProps struct {
	// `AWS::SES::ConfigurationSet.Name`.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::SES::ContactList`.
//
// A list that contains contacts that have subscribed to a particular topic or topics.
//
// TODO: EXAMPLE
//
type CfnContactList interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContactListName() *string
	SetContactListName(val *string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Topics() interface{}
	SetTopics(val interface{})
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

// The jsii proxy struct for CfnContactList
type jsiiProxy_CfnContactList struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnContactList) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) ContactListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) Topics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactList) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ContactList`.
func NewCfnContactList(scope constructs.Construct, id *string, props *CfnContactListProps) CfnContactList {
	_init_.Initialize()

	j := jsiiProxy_CfnContactList{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnContactList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ContactList`.
func NewCfnContactList_Override(c CfnContactList, scope constructs.Construct, id *string, props *CfnContactListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnContactList",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContactList) SetContactListName(val *string) {
	_jsii_.Set(
		j,
		"contactListName",
		val,
	)
}

func (j *jsiiProxy_CfnContactList) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnContactList) SetTopics(val interface{}) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnContactList_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnContactList",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnContactList_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnContactList",
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
func CfnContactList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnContactList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContactList_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnContactList",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnContactList) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnContactList) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnContactList) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnContactList) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnContactList) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnContactList) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnContactList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnContactList) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnContactList) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnContactList) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnContactList) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContactList) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnContactList) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnContactList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContactList) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An interest group, theme, or label within a list.
//
// Lists can have multiple topics.
//
// TODO: EXAMPLE
//
type CfnContactList_TopicProperty struct {
	// The default subscription status to be applied to a contact if the contact has not noted their preference for subscribing to a topic.
	DefaultSubscriptionStatus *string `json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// The name of the topic the contact will see.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// The name of the topic.
	TopicName *string `json:"topicName" yaml:"topicName"`
	// A description of what the topic is about, which the contact will see.
	Description *string `json:"description" yaml:"description"`
}

// Properties for defining a `CfnContactList`.
//
// TODO: EXAMPLE
//
type CfnContactListProps struct {
	// The name of the contact list.
	ContactListName *string `json:"contactListName" yaml:"contactListName"`
	// A description of what the contact list is about.
	Description *string `json:"description" yaml:"description"`
	// The tags associated with a contact list.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// An interest group, theme, or label within a list.
	//
	// A contact list can have multiple topics.
	Topics interface{} `json:"topics" yaml:"topics"`
}

// A CloudFormation `AWS::SES::ReceiptFilter`.
//
// Specify a new IP address filter. You use IP address filters when you receive email with Amazon SES.
//
// TODO: EXAMPLE
//
type CfnReceiptFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Filter() interface{}
	SetFilter(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
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

// The jsii proxy struct for CfnReceiptFilter
type jsiiProxy_CfnReceiptFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReceiptFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) Filter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ReceiptFilter`.
func NewCfnReceiptFilter(scope constructs.Construct, id *string, props *CfnReceiptFilterProps) CfnReceiptFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnReceiptFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ReceiptFilter`.
func NewCfnReceiptFilter_Override(c CfnReceiptFilter, scope constructs.Construct, id *string, props *CfnReceiptFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReceiptFilter) SetFilter(val interface{}) {
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReceiptFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReceiptFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
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
func CfnReceiptFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReceiptFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnReceiptFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReceiptFilter) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReceiptFilter) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReceiptFilter) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReceiptFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReceiptFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReceiptFilter) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReceiptFilter) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReceiptFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReceiptFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReceiptFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReceiptFilter) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReceiptFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReceiptFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an IP address filter.
//
// TODO: EXAMPLE
//
type CfnReceiptFilter_FilterProperty struct {
	// A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.
	IpFilter interface{} `json:"ipFilter" yaml:"ipFilter"`
	// The name of the IP address filter. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	Name *string `json:"name" yaml:"name"`
}

// A receipt IP address filter enables you to specify whether to accept or reject mail originating from an IP address or range of IP addresses.
//
// For information about setting up IP address filters, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-ip-filters.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptFilter_IpFilterProperty struct {
	// A single IP address or a range of IP addresses to block or allow, specified in Classless Inter-Domain Routing (CIDR) notation.
	//
	// An example of a single email address is 10.0.0.1. An example of a range of IP addresses is 10.0.0.1/24. For more information about CIDR notation, see [RFC 2317](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2317) .
	Cidr *string `json:"cidr" yaml:"cidr"`
	// Indicates whether to block or allow incoming mail from the specified IP addresses.
	Policy *string `json:"policy" yaml:"policy"`
}

// Properties for defining a `CfnReceiptFilter`.
//
// TODO: EXAMPLE
//
type CfnReceiptFilterProps struct {
	// A data structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.
	Filter interface{} `json:"filter" yaml:"filter"`
}

// A CloudFormation `AWS::SES::ReceiptRule`.
//
// Specifies a receipt rule.
//
// TODO: EXAMPLE
//
type CfnReceiptRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	After() *string
	SetAfter(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Rule() interface{}
	SetRule(val interface{})
	RuleSetName() *string
	SetRuleSetName(val *string)
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

// The jsii proxy struct for CfnReceiptRule
type jsiiProxy_CfnReceiptRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReceiptRule) After() *string {
	var returns *string
	_jsii_.Get(
		j,
		"after",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) Rule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ReceiptRule`.
func NewCfnReceiptRule(scope constructs.Construct, id *string, props *CfnReceiptRuleProps) CfnReceiptRule {
	_init_.Initialize()

	j := jsiiProxy_CfnReceiptRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ReceiptRule`.
func NewCfnReceiptRule_Override(c CfnReceiptRule, scope constructs.Construct, id *string, props *CfnReceiptRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReceiptRule) SetAfter(val *string) {
	_jsii_.Set(
		j,
		"after",
		val,
	)
}

func (j *jsiiProxy_CfnReceiptRule) SetRule(val interface{}) {
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_CfnReceiptRule) SetRuleSetName(val *string) {
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReceiptRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReceiptRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
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
func CfnReceiptRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReceiptRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnReceiptRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReceiptRule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReceiptRule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReceiptRule) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReceiptRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReceiptRule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReceiptRule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReceiptRule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReceiptRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReceiptRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReceiptRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReceiptRule) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReceiptRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An action that Amazon SES can take when it receives an email on behalf of one or more email addresses or domains that you own.
//
// An instance of this data type can represent only one action.
//
// For information about setting up receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_ActionProperty struct {
	// Adds a header to the received email.
	AddHeaderAction interface{} `json:"addHeaderAction" yaml:"addHeaderAction"`
	// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
	BounceAction interface{} `json:"bounceAction" yaml:"bounceAction"`
	// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
	LambdaAction interface{} `json:"lambdaAction" yaml:"lambdaAction"`
	// Saves the received message to an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes a notification to Amazon SNS.
	S3Action interface{} `json:"s3Action" yaml:"s3Action"`
	// Publishes the email content within a notification to Amazon SNS.
	SnsAction interface{} `json:"snsAction" yaml:"snsAction"`
	// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
	StopAction interface{} `json:"stopAction" yaml:"stopAction"`
	// Calls Amazon WorkMail and, optionally, publishes a notification to Amazon Amazon SNS.
	WorkmailAction interface{} `json:"workmailAction" yaml:"workmailAction"`
}

// When included in a receipt rule, this action adds a header to the received email.
//
// For information about adding a header using a receipt rule, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-add-header.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_AddHeaderActionProperty struct {
	// The name of the header to add to the incoming message.
	//
	// The name must contain at least one character, and can contain up to 50 characters. It consists of alphanumeric (a–z, A–Z, 0–9) characters and dashes.
	HeaderName *string `json:"headerName" yaml:"headerName"`
	// The content to include in the header.
	//
	// This value can contain up to 2048 characters. It can't contain newline ( `\n` ) or carriage return ( `\r` ) characters.
	HeaderValue *string `json:"headerValue" yaml:"headerValue"`
}

// When included in a receipt rule, this action rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// For information about sending a bounce message in response to a received email, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-bounce.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_BounceActionProperty struct {
	// Human-readable text to include in the bounce message.
	Message *string `json:"message" yaml:"message"`
	// The email address of the sender of the bounced email.
	//
	// This is the address from which the bounce message is sent.
	Sender *string `json:"sender" yaml:"sender"`
	// The SMTP reply code, as defined by [RFC 5321](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc5321) .
	SmtpReplyCode *string `json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by [RFC 3463](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3463) .
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is taken.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// When included in a receipt rule, this action calls an AWS Lambda function and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// To enable Amazon SES to call your AWS Lambda function or to publish to an Amazon SNS topic of another account, Amazon SES must have permission to access those resources. For information about giving permissions, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html) .
//
// For information about using AWS Lambda actions in receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-lambda.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_LambdaActionProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	//
	// An example of an AWS Lambda function ARN is `arn:aws:lambda:us-west-2:account-id:function:MyFunction` . For more information about AWS Lambda, see the [AWS Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html) .
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
	// The invocation type of the AWS Lambda function.
	//
	// An invocation type of `RequestResponse` means that the execution of the function immediately results in a response, and a value of `Event` means that the function is invoked asynchronously. The default value is `Event` . For information about AWS Lambda invocation types, see the [AWS Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html) .
	//
	// > There is a 30-second timeout on `RequestResponse` invocations. You should use `Event` invocation in most cases. Use `RequestResponse` only to make a mail flow decision, such as whether to stop the receipt rule or the receipt rule set.
	InvocationType *string `json:"invocationType" yaml:"invocationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the Lambda action is executed.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// Receipt rules enable you to specify which actions Amazon SES should take when it receives mail on behalf of one or more email addresses or domains that you own.
//
// Each receipt rule defines a set of email addresses or domains that it applies to. If the email addresses or domains match at least one recipient address of the message, Amazon SES executes all of the receipt rule's actions on the message.
//
// For information about setting up receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_RuleProperty struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions interface{} `json:"actions" yaml:"actions"`
	// If `true` , the receipt rule is active.
	//
	// The default value is `false` .
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The name of the receipt rule. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	Name *string `json:"name" yaml:"name"`
	// The recipient domains and email addresses that the receipt rule applies to.
	//
	// If this field is not specified, this rule matches all recipients on all verified domains.
	Recipients *[]*string `json:"recipients" yaml:"recipients"`
	// If `true` , then messages that this receipt rule applies to are scanned for spam and viruses.
	//
	// The default value is `false` .
	ScanEnabled interface{} `json:"scanEnabled" yaml:"scanEnabled"`
	// Specifies whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	//
	// If this parameter is set to `Require` , Amazon SES bounces emails that are not received over TLS. The default is `Optional` .
	TlsPolicy *string `json:"tlsPolicy" yaml:"tlsPolicy"`
}

// When included in a receipt rule, this action saves the received message to an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// To enable Amazon SES to write emails to your Amazon S3 bucket, use an AWS KMS key to encrypt your emails, or publish to an Amazon SNS topic of another account, Amazon SES must have permission to access those resources. For information about granting permissions, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html) .
//
// > When you save your emails to an Amazon S3 bucket, the maximum email size (including headers) is 30 MB. Emails larger than that bounces.
//
// For information about specifying Amazon S3 actions in receipt rules, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-s3.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_S3ActionProperty struct {
	// The name of the Amazon S3 bucket for incoming email.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The customer master key that Amazon SES should use to encrypt your emails before saving them to the Amazon S3 bucket.
	//
	// You can use the default master key or a custom master key that you created in AWS KMS as follows:
	//
	// - To use the default master key, provide an ARN in the form of `arn:aws:kms:REGION:ACCOUNT-ID-WITHOUT-HYPHENS:alias/aws/ses` . For example, if your AWS account ID is 123456789012 and you want to use the default master key in the US West (Oregon) Region, the ARN of the default master key would be `arn:aws:kms:us-west-2:123456789012:alias/aws/ses` . If you use the default master key, you don't need to perform any extra steps to give Amazon SES permission to use the key.
	// - To use a custom master key that you created in AWS KMS, provide the ARN of the master key and ensure that you add a statement to your key's policy to give Amazon SES permission to use it. For more information about giving permissions, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html) .
	//
	// For more information about key policies, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html) . If you do not specify a master key, Amazon SES does not encrypt your emails.
	//
	// > Your mail is encrypted by Amazon SES using the Amazon S3 encryption client before the mail is submitted to Amazon S3 for storage. It is not encrypted using Amazon S3 server-side encryption. This means that you must use the Amazon S3 encryption client to decrypt the email after retrieving it from Amazon S3, as the service has no access to use your AWS KMS keys for decryption. This encryption client is currently available with the [AWS SDK for Java](https://docs.aws.amazon.com/sdk-for-java/) and [AWS SDK for Ruby](https://docs.aws.amazon.com/sdk-for-ruby/) only. For more information about client-side encryption using AWS KMS master keys, see the [Amazon S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html) .
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The key prefix of the Amazon S3 bucket.
	//
	// The key prefix is similar to a directory name that enables you to store similar data under the same directory in a bucket.
	ObjectKeyPrefix *string `json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// The ARN of the Amazon SNS topic to notify when the message is saved to the Amazon S3 bucket.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// When included in a receipt rule, this action publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// This action includes a complete copy of the email content in the Amazon SNS notifications. Amazon SNS notifications for all other actions simply provide information about the email. They do not include the email content itself.
//
// If you own the Amazon SNS topic, you don't need to do anything to give Amazon SES permission to publish emails to it. However, if you don't own the Amazon SNS topic, you need to attach a policy to the topic to give Amazon SES permissions to access it. For information about giving permissions, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html) .
//
// > You can only publish emails that are 150 KB or less (including the header) to Amazon SNS. Larger emails bounce. If you anticipate emails larger than 150 KB, use the S3 action instead.
//
// For information about using a receipt rule to publish an Amazon SNS notification, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-sns.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_SNSActionProperty struct {
	// The encoding to use for the email within the Amazon SNS notification.
	//
	// UTF-8 is easier to use, but may not preserve all special characters when a message was encoded with a different encoding format. Base64 preserves all special characters. The default value is UTF-8.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// When included in a receipt rule, this action terminates the evaluation of the receipt rule set and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// For information about setting a stop action in a receipt rule, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-stop.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_StopActionProperty struct {
	// The scope of the StopAction.
	//
	// The only acceptable value is `RuleSet` .
	Scope *string `json:"scope" yaml:"scope"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the stop action is taken.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) Amazon SNS operation.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// When included in a receipt rule, this action calls Amazon WorkMail and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// It usually isn't necessary to set this up manually, because Amazon WorkMail adds the rule automatically during its setup procedure.
//
// For information using a receipt rule to call Amazon WorkMail, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-workmail.html) .
//
// TODO: EXAMPLE
//
type CfnReceiptRule_WorkmailActionProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon WorkMail organization. Amazon WorkMail ARNs use the following format:.
	//
	// `arn:aws:workmail:<region>:<awsAccountId>:organization/<workmailOrganizationId>`
	//
	// You can find the ID of your organization by using the [ListOrganizations](https://docs.aws.amazon.com/workmail/latest/APIReference/API_ListOrganizations.html) operation in Amazon WorkMail. Amazon WorkMail organization IDs begin with " `m-` ", followed by a string of alphanumeric characters.
	//
	// For information about Amazon WorkMail organizations, see the [Amazon WorkMail Administrator Guide](https://docs.aws.amazon.com/workmail/latest/adminguide/organizations_overview.html) .
	OrganizationArn *string `json:"organizationArn" yaml:"organizationArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the WorkMail action is called.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) operation in Amazon SNS.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// Properties for defining a `CfnReceiptRule`.
//
// TODO: EXAMPLE
//
type CfnReceiptRuleProps struct {
	// A data structure that contains the specified rule's name, actions, recipients, domains, enabled status, scan status, and TLS policy.
	Rule interface{} `json:"rule" yaml:"rule"`
	// The name of the rule set where the receipt rule is added.
	RuleSetName *string `json:"ruleSetName" yaml:"ruleSetName"`
	// The name of an existing rule after which the new rule is placed.
	//
	// If this parameter is null, the new rule is inserted at the beginning of the rule list.
	After *string `json:"after" yaml:"after"`
}

// A CloudFormation `AWS::SES::ReceiptRuleSet`.
//
// Creates an empty receipt rule set.
//
// For information about setting up receipt rule sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html) .
//
// You can execute this operation no more than once per second.
//
// TODO: EXAMPLE
//
type CfnReceiptRuleSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RuleSetName() *string
	SetRuleSetName(val *string)
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

// The jsii proxy struct for CfnReceiptRuleSet
type jsiiProxy_CfnReceiptRuleSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReceiptRuleSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRuleSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::ReceiptRuleSet`.
func NewCfnReceiptRuleSet(scope constructs.Construct, id *string, props *CfnReceiptRuleSetProps) CfnReceiptRuleSet {
	_init_.Initialize()

	j := jsiiProxy_CfnReceiptRuleSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::ReceiptRuleSet`.
func NewCfnReceiptRuleSet_Override(c CfnReceiptRuleSet, scope constructs.Construct, id *string, props *CfnReceiptRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReceiptRuleSet) SetRuleSetName(val *string) {
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReceiptRuleSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReceiptRuleSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
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
func CfnReceiptRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReceiptRuleSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnReceiptRuleSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReceiptRuleSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReceiptRuleSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReceiptRuleSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptRuleSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReceiptRuleSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReceiptRuleSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReceiptRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReceiptRuleSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReceiptRuleSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReceiptRuleSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReceiptRuleSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReceiptRuleSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReceiptRuleSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReceiptRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReceiptRuleSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnReceiptRuleSet`.
//
// TODO: EXAMPLE
//
type CfnReceiptRuleSetProps struct {
	// The name of the receipt rule set to reorder.
	RuleSetName *string `json:"ruleSetName" yaml:"ruleSetName"`
}

// A CloudFormation `AWS::SES::Template`.
//
// Specifies an email template. Email templates enable you to send personalized email to one or more destinations in a single API operation.
//
// TODO: EXAMPLE
//
type CfnTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Template() interface{}
	SetTemplate(val interface{})
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

// The jsii proxy struct for CfnTemplate
type jsiiProxy_CfnTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Template() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::Template`.
func NewCfnTemplate(scope constructs.Construct, id *string, props *CfnTemplateProps) CfnTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::Template`.
func NewCfnTemplate_Override(c CfnTemplate, scope constructs.Construct, id *string, props *CfnTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTemplate) SetTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"template",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnTemplate",
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
func CfnTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTemplate) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTemplate) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTemplate) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTemplate) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTemplate) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTemplate) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The content of the email, composed of a subject line, an HTML part, and a text-only part.
//
// TODO: EXAMPLE
//
type CfnTemplate_TemplateProperty struct {
	// The HTML body of the email.
	HtmlPart *string `json:"htmlPart" yaml:"htmlPart"`
	// The subject line of the email.
	SubjectPart *string `json:"subjectPart" yaml:"subjectPart"`
	// The name of the template.
	TemplateName *string `json:"templateName" yaml:"templateName"`
	// The email body that is visible to recipients whose email clients do not display HTML content.
	TextPart *string `json:"textPart" yaml:"textPart"`
}

// Properties for defining a `CfnTemplate`.
//
// TODO: EXAMPLE
//
type CfnTemplateProps struct {
	// The content of the email, composed of a subject line, an HTML part, and a text-only part.
	Template interface{} `json:"template" yaml:"template"`
}

// A rule added at the top of the rule set to drop spam/virus.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-lambda-example-functions.html
//
type DropSpamReceiptRule interface {
	constructs.Construct
	Node() constructs.Node
	Rule() ReceiptRule
	ToString() *string
}

// The jsii proxy struct for DropSpamReceiptRule
type jsiiProxy_DropSpamReceiptRule struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DropSpamReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropSpamReceiptRule) Rule() ReceiptRule {
	var returns ReceiptRule
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}


func NewDropSpamReceiptRule(scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) DropSpamReceiptRule {
	_init_.Initialize()

	j := jsiiProxy_DropSpamReceiptRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDropSpamReceiptRule_Override(d DropSpamReceiptRule, scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DropSpamReceiptRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DropSpamReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type DropSpamReceiptRuleProps struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions *[]IReceiptRuleAction `json:"actions" yaml:"actions"`
	// An existing rule after which the new rule will be placed.
	After IReceiptRule `json:"after" yaml:"after"`
	// Whether the rule is active.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the rule.
	ReceiptRuleName *string `json:"receiptRuleName" yaml:"receiptRuleName"`
	// The recipient domains and email addresses that the receipt rule applies to.
	Recipients *[]*string `json:"recipients" yaml:"recipients"`
	// Whether to scan for spam and viruses.
	ScanEnabled *bool `json:"scanEnabled" yaml:"scanEnabled"`
	// Whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	TlsPolicy TlsPolicy `json:"tlsPolicy" yaml:"tlsPolicy"`
	// The name of the rule set that the receipt rule will be added to.
	RuleSet IReceiptRuleSet `json:"ruleSet" yaml:"ruleSet"`
}

// A receipt rule.
type IReceiptRule interface {
	awscdk.IResource
	// The name of the receipt rule.
	ReceiptRuleName() *string
}

// The jsii proxy for IReceiptRule
type jsiiProxy_IReceiptRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IReceiptRule) ReceiptRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleName",
		&returns,
	)
	return returns
}

// An abstract action for a receipt rule.
type IReceiptRuleAction interface {
	// Returns the receipt rule action specification.
	Bind(receiptRule IReceiptRule) *ReceiptRuleActionConfig
}

// The jsii proxy for IReceiptRuleAction
type jsiiProxy_IReceiptRuleAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IReceiptRuleAction) Bind(receiptRule IReceiptRule) *ReceiptRuleActionConfig {
	var returns *ReceiptRuleActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{receiptRule},
		&returns,
	)

	return returns
}

// A receipt rule set.
type IReceiptRuleSet interface {
	awscdk.IResource
	// Adds a new receipt rule in this rule set.
	//
	// The new rule is added after
	// the last added rule unless `after` is specified.
	AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule
	// The receipt rule set name.
	ReceiptRuleSetName() *string
}

// The jsii proxy for IReceiptRuleSet
type jsiiProxy_IReceiptRuleSet struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IReceiptRuleSet) AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule {
	var returns ReceiptRule

	_jsii_.Invoke(
		i,
		"addRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReceiptRuleSet) ReceiptRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleSetName",
		&returns,
	)
	return returns
}

// LambdaAction configuration.
//
// TODO: EXAMPLE
//
type LambdaActionConfig struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
	// The invocation type of the AWS Lambda function.
	InvocationType *string `json:"invocationType" yaml:"invocationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the Lambda action is executed.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// A receipt filter.
//
// When instantiated without props, it creates a
// block all receipt filter.
//
// TODO: EXAMPLE
//
type ReceiptFilter interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ReceiptFilter
type jsiiProxy_ReceiptFilter struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ReceiptFilter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptFilter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewReceiptFilter(scope constructs.Construct, id *string, props *ReceiptFilterProps) ReceiptFilter {
	_init_.Initialize()

	j := jsiiProxy_ReceiptFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewReceiptFilter_Override(r ReceiptFilter, scope constructs.Construct, id *string, props *ReceiptFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptFilter",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ReceiptFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ReceiptFilter_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptFilter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_ReceiptFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReceiptFilter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_ReceiptFilter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_ReceiptFilter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ReceiptFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The policy for the receipt filter.
type ReceiptFilterPolicy string

const (
	ReceiptFilterPolicy_ALLOW ReceiptFilterPolicy = "ALLOW"
	ReceiptFilterPolicy_BLOCK ReceiptFilterPolicy = "BLOCK"
)

// Construction properties for a ReceiptFilter.
//
// TODO: EXAMPLE
//
type ReceiptFilterProps struct {
	// The ip address or range to filter.
	Ip *string `json:"ip" yaml:"ip"`
	// The policy for the filter.
	Policy ReceiptFilterPolicy `json:"policy" yaml:"policy"`
	// The name for the receipt filter.
	ReceiptFilterName *string `json:"receiptFilterName" yaml:"receiptFilterName"`
}

// A new receipt rule.
//
// TODO: EXAMPLE
//
type ReceiptRule interface {
	awscdk.Resource
	IReceiptRule
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ReceiptRuleName() *string
	Stack() awscdk.Stack
	AddAction(action IReceiptRuleAction)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ReceiptRule
type jsiiProxy_ReceiptRule struct {
	internal.Type__awscdkResource
	jsiiProxy_IReceiptRule
}

func (j *jsiiProxy_ReceiptRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRule) ReceiptRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewReceiptRule(scope constructs.Construct, id *string, props *ReceiptRuleProps) ReceiptRule {
	_init_.Initialize()

	j := jsiiProxy_ReceiptRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewReceiptRule_Override(r ReceiptRule, scope constructs.Construct, id *string, props *ReceiptRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRule",
		[]interface{}{scope, id, props},
		r,
	)
}

func ReceiptRule_FromReceiptRuleName(scope constructs.Construct, id *string, receiptRuleName *string) IReceiptRule {
	_init_.Initialize()

	var returns IReceiptRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRule",
		"fromReceiptRuleName",
		[]interface{}{scope, id, receiptRuleName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ReceiptRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ReceiptRule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an action to this receipt rule.
func (r *jsiiProxy_ReceiptRule) AddAction(action IReceiptRuleAction) {
	_jsii_.InvokeVoid(
		r,
		"addAction",
		[]interface{}{action},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_ReceiptRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReceiptRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_ReceiptRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_ReceiptRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a receipt rule action.
//
// TODO: EXAMPLE
//
type ReceiptRuleActionConfig struct {
	// Adds a header to the received email.
	AddHeaderAction *AddHeaderActionConfig `json:"addHeaderAction" yaml:"addHeaderAction"`
	// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon SNS.
	BounceAction *BounceActionConfig `json:"bounceAction" yaml:"bounceAction"`
	// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
	LambdaAction *LambdaActionConfig `json:"lambdaAction" yaml:"lambdaAction"`
	// Saves the received message to an Amazon S3 bucket and, optionally, publishes a notification to Amazon SNS.
	S3Action *S3ActionConfig `json:"s3Action" yaml:"s3Action"`
	// Publishes the email content within a notification to Amazon SNS.
	SnsAction *SNSActionConfig `json:"snsAction" yaml:"snsAction"`
	// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
	StopAction *StopActionConfig `json:"stopAction" yaml:"stopAction"`
	// Calls Amazon WorkMail and, optionally, publishes a notification to Amazon SNS.
	WorkmailAction *WorkmailActionConfig `json:"workmailAction" yaml:"workmailAction"`
}

// Options to add a receipt rule to a receipt rule set.
//
// TODO: EXAMPLE
//
type ReceiptRuleOptions struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions *[]IReceiptRuleAction `json:"actions" yaml:"actions"`
	// An existing rule after which the new rule will be placed.
	After IReceiptRule `json:"after" yaml:"after"`
	// Whether the rule is active.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the rule.
	ReceiptRuleName *string `json:"receiptRuleName" yaml:"receiptRuleName"`
	// The recipient domains and email addresses that the receipt rule applies to.
	Recipients *[]*string `json:"recipients" yaml:"recipients"`
	// Whether to scan for spam and viruses.
	ScanEnabled *bool `json:"scanEnabled" yaml:"scanEnabled"`
	// Whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	TlsPolicy TlsPolicy `json:"tlsPolicy" yaml:"tlsPolicy"`
}

// Construction properties for a ReceiptRule.
//
// TODO: EXAMPLE
//
type ReceiptRuleProps struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions *[]IReceiptRuleAction `json:"actions" yaml:"actions"`
	// An existing rule after which the new rule will be placed.
	After IReceiptRule `json:"after" yaml:"after"`
	// Whether the rule is active.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the rule.
	ReceiptRuleName *string `json:"receiptRuleName" yaml:"receiptRuleName"`
	// The recipient domains and email addresses that the receipt rule applies to.
	Recipients *[]*string `json:"recipients" yaml:"recipients"`
	// Whether to scan for spam and viruses.
	ScanEnabled *bool `json:"scanEnabled" yaml:"scanEnabled"`
	// Whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	TlsPolicy TlsPolicy `json:"tlsPolicy" yaml:"tlsPolicy"`
	// The name of the rule set that the receipt rule will be added to.
	RuleSet IReceiptRuleSet `json:"ruleSet" yaml:"ruleSet"`
}

// A new receipt rule set.
//
// TODO: EXAMPLE
//
type ReceiptRuleSet interface {
	awscdk.Resource
	IReceiptRuleSet
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ReceiptRuleSetName() *string
	Stack() awscdk.Stack
	AddDropSpamRule()
	AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ReceiptRuleSet
type jsiiProxy_ReceiptRuleSet struct {
	internal.Type__awscdkResource
	jsiiProxy_IReceiptRuleSet
}

func (j *jsiiProxy_ReceiptRuleSet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) ReceiptRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewReceiptRuleSet(scope constructs.Construct, id *string, props *ReceiptRuleSetProps) ReceiptRuleSet {
	_init_.Initialize()

	j := jsiiProxy_ReceiptRuleSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewReceiptRuleSet_Override(r ReceiptRuleSet, scope constructs.Construct, id *string, props *ReceiptRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an exported receipt rule set.
func ReceiptRuleSet_FromReceiptRuleSetName(scope constructs.Construct, id *string, receiptRuleSetName *string) IReceiptRuleSet {
	_init_.Initialize()

	var returns IReceiptRuleSet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"fromReceiptRuleSetName",
		[]interface{}{scope, id, receiptRuleSetName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ReceiptRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ReceiptRuleSet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a drop spam rule.
func (r *jsiiProxy_ReceiptRuleSet) AddDropSpamRule() {
	_jsii_.InvokeVoid(
		r,
		"addDropSpamRule",
		nil, // no parameters
	)
}

// Adds a new receipt rule in this rule set.
//
// The new rule is added after
// the last added rule unless `after` is specified.
func (r *jsiiProxy_ReceiptRuleSet) AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule {
	var returns ReceiptRule

	_jsii_.Invoke(
		r,
		"addRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_ReceiptRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReceiptRuleSet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_ReceiptRuleSet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_ReceiptRuleSet) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ReceiptRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a ReceiptRuleSet.
//
// TODO: EXAMPLE
//
type ReceiptRuleSetProps struct {
	// Whether to add a first rule to stop processing messages that have at least one spam indicator.
	DropSpam *bool `json:"dropSpam" yaml:"dropSpam"`
	// The name for the receipt rule set.
	ReceiptRuleSetName *string `json:"receiptRuleSetName" yaml:"receiptRuleSetName"`
	// The list of rules to add to this rule set.
	//
	// Rules are added in the same
	// order as they appear in the list.
	Rules *[]*ReceiptRuleOptions `json:"rules" yaml:"rules"`
}

// S3Action configuration.
//
// TODO: EXAMPLE
//
type S3ActionConfig struct {
	// The name of the Amazon S3 bucket that you want to send incoming mail to.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The customer master key that Amazon SES should use to encrypt your emails before saving them to the Amazon S3 bucket.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The key prefix of the Amazon S3 bucket.
	ObjectKeyPrefix *string `json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// The ARN of the Amazon SNS topic to notify when the message is saved to the Amazon S3 bucket.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// SNSAction configuration.
//
// TODO: EXAMPLE
//
type SNSActionConfig struct {
	// The encoding to use for the email within the Amazon SNS notification.
	Encoding *string `json:"encoding" yaml:"encoding"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// StopAction configuration.
//
// TODO: EXAMPLE
//
type StopActionConfig struct {
	// The scope of the StopAction.
	//
	// The only acceptable value is RuleSet.
	Scope *string `json:"scope" yaml:"scope"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the stop action is taken.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// The type of TLS policy for a receipt rule.
type TlsPolicy string

const (
	TlsPolicy_OPTIONAL TlsPolicy = "OPTIONAL"
	TlsPolicy_REQUIRE TlsPolicy = "REQUIRE"
)

// WorkmailAction configuration.
//
// TODO: EXAMPLE
//
type WorkmailActionConfig struct {
	// The Amazon Resource Name (ARN) of the Amazon WorkMail organization.
	OrganizationArn *string `json:"organizationArn" yaml:"organizationArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the WorkMail action is called.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

