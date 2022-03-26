package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssns/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Between condition for a numeric attribute.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
// Experimental.
type BetweenCondition struct {
	// The start value.
	// Experimental.
	Start *float64 `json:"start" yaml:"start"`
	// The stop value.
	// Experimental.
	Stop *float64 `json:"stop" yaml:"stop"`
}

// A CloudFormation `AWS::SNS::Subscription`.
//
// The `AWS::SNS::Subscription` resource subscribes an endpoint to an Amazon SNS topic. For a subscription to be created, the owner of the endpoint must confirm the subscription.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//
//   var deliveryPolicy interface{}
//   var filterPolicy interface{}
//   var redrivePolicy interface{}
//   cfnSubscription := sns.NewCfnSubscription(this, jsii.String("MyCfnSubscription"), &cfnSubscriptionProps{
//   	protocol: jsii.String("protocol"),
//   	topicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	deliveryPolicy: deliveryPolicy,
//   	endpoint: jsii.String("endpoint"),
//   	filterPolicy: filterPolicy,
//   	rawMessageDelivery: jsii.Boolean(false),
//   	redrivePolicy: redrivePolicy,
//   	region: jsii.String("region"),
//   	subscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   })
//
type CfnSubscription interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The delivery policy JSON assigned to the subscription.
	//
	// Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message delivery retries](https://docs.aws.amazon.com/sns/latest/dg/sns-message-delivery-retries.html) in the *Amazon SNS Developer Guide* .
	DeliveryPolicy() interface{}
	SetDeliveryPolicy(val interface{})
	// The subscription's endpoint.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Endpoint() *string
	SetEndpoint(val *string)
	// The filter policy JSON assigned to the subscription.
	//
	// Enables the subscriber to filter out unwanted messages. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message filtering](https://docs.aws.amazon.com/sns/latest/dg/sns-message-filtering.html) in the *Amazon SNS Developer Guide* .
	FilterPolicy() interface{}
	SetFilterPolicy(val interface{})
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
	Node() awscdk.ConstructNode
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Protocol() *string
	SetProtocol(val *string)
	// When set to `true` , enables raw message delivery.
	//
	// Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* .
	RawMessageDelivery() interface{}
	SetRawMessageDelivery(val interface{})
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	//
	// Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.
	//
	// For more information about the redrive policy and dead-letter queues, see [Amazon SQS dead-letter queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html) in the *Amazon SQS Developer Guide* .
	RedrivePolicy() interface{}
	SetRedrivePolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// For cross-region subscriptions, the region in which the topic resides.
	//
	// If no region is specified, AWS CloudFormation uses the region of the caller as the default.
	//
	// If you perform an update operation that only updates the `Region` property of a `AWS::SNS::Subscription` resource, that operation will fail unless you are either:
	//
	// - Updating the `Region` from `NULL` to the caller region.
	// - Updating the `Region` from the caller region to `NULL` .
	Region() *string
	SetRegion(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// This property applies only to Amazon Kinesis Data Firehose delivery stream subscriptions.
	//
	// Specify the ARN of the IAM role that has the following:
	//
	// - Permission to write to the Amazon Kinesis Data Firehose delivery stream
	// - Amazon SNS listed as a trusted entity
	//
	// Specifying a valid ARN for this attribute is required for Kinesis Data Firehose delivery stream subscriptions. For more information, see [Fanout to Amazon Kinesis Data Firehose delivery streams](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html) in the *Amazon SNS Developer Guide.*
	SubscriptionRoleArn() *string
	SetSubscriptionRoleArn(val *string)
	// The ARN of the topic to subscribe to.
	TopicArn() *string
	SetTopicArn(val *string)
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
	AddDependsOn(target awscdk.CfnResource)
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
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
	Inspect(inspector awscdk.TreeInspector)
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
	Synthesize(session awscdk.ISynthesisSession)
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

// The jsii proxy struct for CfnSubscription
type jsiiProxy_CfnSubscription struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSubscription) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) DeliveryPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) FilterPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) RawMessageDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawMessageDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) RedrivePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) SubscriptionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscription) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SNS::Subscription`.
func NewCfnSubscription(scope awscdk.Construct, id *string, props *CfnSubscriptionProps) CfnSubscription {
	_init_.Initialize()

	j := jsiiProxy_CfnSubscription{}

	_jsii_.Create(
		"monocdk.aws_sns.CfnSubscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SNS::Subscription`.
func NewCfnSubscription_Override(c CfnSubscription, scope awscdk.Construct, id *string, props *CfnSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.CfnSubscription",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSubscription) SetDeliveryPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetFilterPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"filterPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetRawMessageDelivery(val interface{}) {
	_jsii_.Set(
		j,
		"rawMessageDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetRedrivePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetSubscriptionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"subscriptionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnSubscription) SetTopicArn(val *string) {
	_jsii_.Set(
		j,
		"topicArn",
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
func CfnSubscription_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnSubscription",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSubscription_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnSubscription",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscription_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sns.CfnSubscription",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscription) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSubscription) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubscription) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSubscription) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSubscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSubscription) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSubscription) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSubscription) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSubscription) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSubscription) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSubscription) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSubscription`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//
//   var deliveryPolicy interface{}
//   var filterPolicy interface{}
//   var redrivePolicy interface{}
//   cfnSubscriptionProps := &cfnSubscriptionProps{
//   	protocol: jsii.String("protocol"),
//   	topicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	deliveryPolicy: deliveryPolicy,
//   	endpoint: jsii.String("endpoint"),
//   	filterPolicy: filterPolicy,
//   	rawMessageDelivery: jsii.Boolean(false),
//   	redrivePolicy: redrivePolicy,
//   	region: jsii.String("region"),
//   	subscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
type CfnSubscriptionProps struct {
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The ARN of the topic to subscribe to.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
	// The delivery policy JSON assigned to the subscription.
	//
	// Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message delivery retries](https://docs.aws.amazon.com/sns/latest/dg/sns-message-delivery-retries.html) in the *Amazon SNS Developer Guide* .
	DeliveryPolicy interface{} `json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The subscription's endpoint.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The filter policy JSON assigned to the subscription.
	//
	// Enables the subscriber to filter out unwanted messages. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message filtering](https://docs.aws.amazon.com/sns/latest/dg/sns-message-filtering.html) in the *Amazon SNS Developer Guide* .
	FilterPolicy interface{} `json:"filterPolicy" yaml:"filterPolicy"`
	// When set to `true` , enables raw message delivery.
	//
	// Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* .
	RawMessageDelivery interface{} `json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	//
	// Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.
	//
	// For more information about the redrive policy and dead-letter queues, see [Amazon SQS dead-letter queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html) in the *Amazon SQS Developer Guide* .
	RedrivePolicy interface{} `json:"redrivePolicy" yaml:"redrivePolicy"`
	// For cross-region subscriptions, the region in which the topic resides.
	//
	// If no region is specified, AWS CloudFormation uses the region of the caller as the default.
	//
	// If you perform an update operation that only updates the `Region` property of a `AWS::SNS::Subscription` resource, that operation will fail unless you are either:
	//
	// - Updating the `Region` from `NULL` to the caller region.
	// - Updating the `Region` from the caller region to `NULL` .
	Region *string `json:"region" yaml:"region"`
	// This property applies only to Amazon Kinesis Data Firehose delivery stream subscriptions.
	//
	// Specify the ARN of the IAM role that has the following:
	//
	// - Permission to write to the Amazon Kinesis Data Firehose delivery stream
	// - Amazon SNS listed as a trusted entity
	//
	// Specifying a valid ARN for this attribute is required for Kinesis Data Firehose delivery stream subscriptions. For more information, see [Fanout to Amazon Kinesis Data Firehose delivery streams](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html) in the *Amazon SNS Developer Guide.*
	SubscriptionRoleArn *string `json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

// A CloudFormation `AWS::SNS::Topic`.
//
// The `AWS::SNS::Topic` resource creates a topic to which notifications can be published.
//
// > One account can create a maximum of 100,000 standard topics and 1,000 FIFO topics. For more information, see [Amazon SNS endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/sns.html) in the *AWS General Reference* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//   cfnTopic := sns.NewCfnTopic(this, jsii.String("MyCfnTopic"), &cfnTopicProps{
//   	contentBasedDeduplication: jsii.Boolean(false),
//   	displayName: jsii.String("displayName"),
//   	fifoTopic: jsii.Boolean(false),
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	subscription: []interface{}{
//   		&subscriptionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	topicName: jsii.String("topicName"),
//   })
//
type CfnTopic interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the name of an Amazon SNS topic.
	AttrTopicName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Enables content-based deduplication for FIFO topics.
	//
	// - By default, `ContentBasedDeduplication` is set to `false` . If you create a FIFO topic and this attribute is `false` , you must specify a value for the `MessageDeduplicationId` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	// - When you set `ContentBasedDeduplication` to `true` , Amazon SNS uses a SHA-256 hash to generate the `MessageDeduplicationId` using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the `MessageDeduplicationId` parameter for the `Publish` action.
	ContentBasedDeduplication() interface{}
	SetContentBasedDeduplication(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName() *string
	SetDisplayName(val *string)
	// Set to true to create a FIFO topic.
	FifoTopic() interface{}
	SetFifoTopic(val interface{})
	// The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK.
	//
	// For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms) . For more examples, see `[KeyId](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)` in the *AWS Key Management Service API Reference* .
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html) .
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
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
	Node() awscdk.ConstructNode
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
	Stack() awscdk.Stack
	// The Amazon SNS subscriptions (endpoints) for this topic.
	Subscription() interface{}
	SetSubscription(val interface{})
	// The list of tags to add to a new topic.
	//
	// > To be able to tag a topic on creation, you must have the `sns:CreateTopic` and `sns:TagResource` permissions.
	Tags() awscdk.TagManager
	// The name of the topic you want to create.
	//
	// Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with `.fifo` .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TopicName() *string
	SetTopicName(val *string)
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
	AddDependsOn(target awscdk.CfnResource)
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
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
	Inspect(inspector awscdk.TreeInspector)
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
	Synthesize(session awscdk.ISynthesisSession)
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

// The jsii proxy struct for CfnTopic
type jsiiProxy_CfnTopic struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopic) AttrTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Subscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SNS::Topic`.
func NewCfnTopic(scope awscdk.Construct, id *string, props *CfnTopicProps) CfnTopic {
	_init_.Initialize()

	j := jsiiProxy_CfnTopic{}

	_jsii_.Create(
		"monocdk.aws_sns.CfnTopic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SNS::Topic`.
func NewCfnTopic_Override(c CfnTopic, scope awscdk.Construct, id *string, props *CfnTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.CfnTopic",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopic) SetContentBasedDeduplication(val interface{}) {
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_CfnTopic) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnTopic) SetFifoTopic(val interface{}) {
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_CfnTopic) SetKmsMasterKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnTopic) SetSubscription(val interface{}) {
	_jsii_.Set(
		j,
		"subscription",
		val,
	)
}

func (j *jsiiProxy_CfnTopic) SetTopicName(val *string) {
	_jsii_.Set(
		j,
		"topicName",
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
func CfnTopic_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopic",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTopic_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopic",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopic_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sns.CfnTopic",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopic) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopic) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopic) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopic) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopic) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopic) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopic) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopic) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopic) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopic) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopic) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopic) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `Subscription` is an embedded property that describes the subscription endpoints of an Amazon SNS topic.
//
// > For full control over subscription behavior (for example, delivery policy, filtering, raw message delivery, and cross-region subscriptions), use the [AWS::SNS::Subscription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html) resource.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//   subscriptionProperty := &subscriptionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnTopic_SubscriptionProperty struct {
	// The endpoint that receives notifications from the Amazon SNS topic.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Protocol *string `json:"protocol" yaml:"protocol"`
}

// A CloudFormation `AWS::SNS::TopicPolicy`.
//
// The `AWS::SNS::TopicPolicy` resource associates Amazon SNS topics with a policy. For an example snippet, see [Declaring an Amazon SNS policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sns-policy) in the *AWS CloudFormation User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//
//   var policyDocument interface{}
//   cfnTopicPolicy := sns.NewCfnTopicPolicy(this, jsii.String("MyCfnTopicPolicy"), &cfnTopicPolicyProps{
//   	policyDocument: policyDocument,
//   	topics: []*string{
//   		jsii.String("topics"),
//   	},
//   })
//
type CfnTopicPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// A policy document that contains permissions to add to the specified SNS topics.
	PolicyDocument() interface{}
	SetPolicyDocument(val interface{})
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
	Stack() awscdk.Stack
	// The Amazon Resource Names (ARN) of the topics to which you want to add the policy.
	//
	// You can use the `[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an `[AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html)` resource.
	Topics() *[]*string
	SetTopics(val *[]*string)
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
	AddDependsOn(target awscdk.CfnResource)
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
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
	Inspect(inspector awscdk.TreeInspector)
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
	Synthesize(session awscdk.ISynthesisSession)
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

// The jsii proxy struct for CfnTopicPolicy
type jsiiProxy_CfnTopicPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopicPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) PolicyDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SNS::TopicPolicy`.
func NewCfnTopicPolicy(scope awscdk.Construct, id *string, props *CfnTopicPolicyProps) CfnTopicPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnTopicPolicy{}

	_jsii_.Create(
		"monocdk.aws_sns.CfnTopicPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SNS::TopicPolicy`.
func NewCfnTopicPolicy_Override(c CfnTopicPolicy, scope awscdk.Construct, id *string, props *CfnTopicPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.CfnTopicPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopicPolicy) SetPolicyDocument(val interface{}) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnTopicPolicy) SetTopics(val *[]*string) {
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
// Experimental.
func CfnTopicPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopicPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTopicPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopicPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTopicPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.CfnTopicPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sns.CfnTopicPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopicPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopicPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopicPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnTopicPolicy`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//
//   var policyDocument interface{}
//   cfnTopicPolicyProps := &cfnTopicPolicyProps{
//   	policyDocument: policyDocument,
//   	topics: []*string{
//   		jsii.String("topics"),
//   	},
//   }
//
type CfnTopicPolicyProps struct {
	// A policy document that contains permissions to add to the specified SNS topics.
	PolicyDocument interface{} `json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Names (ARN) of the topics to which you want to add the policy.
	//
	// You can use the `[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an `[AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html)` resource.
	Topics *[]*string `json:"topics" yaml:"topics"`
}

// Properties for defining a `CfnTopic`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"
//   cfnTopicProps := &cfnTopicProps{
//   	contentBasedDeduplication: jsii.Boolean(false),
//   	displayName: jsii.String("displayName"),
//   	fifoTopic: jsii.Boolean(false),
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	subscription: []interface{}{
//   		&subscriptionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	topicName: jsii.String("topicName"),
//   }
//
type CfnTopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	//
	// - By default, `ContentBasedDeduplication` is set to `false` . If you create a FIFO topic and this attribute is `false` , you must specify a value for the `MessageDeduplicationId` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	// - When you set `ContentBasedDeduplication` to `true` , Amazon SNS uses a SHA-256 hash to generate the `MessageDeduplicationId` using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the `MessageDeduplicationId` parameter for the `Publish` action.
	ContentBasedDeduplication interface{} `json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	FifoTopic interface{} `json:"fifoTopic" yaml:"fifoTopic"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK.
	//
	// For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms) . For more examples, see `[KeyId](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)` in the *AWS Key Management Service API Reference* .
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html) .
	KmsMasterKeyId *string `json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The Amazon SNS subscriptions (endpoints) for this topic.
	Subscription interface{} `json:"subscription" yaml:"subscription"`
	// The list of tags to add to a new topic.
	//
	// > To be able to tag a topic on creation, you must have the `sns:CreateTopic` and `sns:TagResource` permissions.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The name of the topic you want to create.
	//
	// Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with `.fifo` .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TopicName *string `json:"topicName" yaml:"topicName"`
}

// Represents an SNS topic.
// Experimental.
type ITopic interface {
	awscodestarnotifications.INotificationRuleTarget
	awscdk.IResource
	// Subscribe some endpoint to this topic.
	// Experimental.
	AddSubscription(subscription ITopicSubscription)
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant topic publishing permissions to the given identity.
	// Experimental.
	GrantPublish(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	// Experimental.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	// Experimental.
	Fifo() *bool
	// The ARN of the topic.
	// Experimental.
	TopicArn() *string
	// The name of the topic.
	// Experimental.
	TopicName() *string
}

// The jsii proxy for ITopic
type jsiiProxy_ITopic struct {
	internal.Type__awscodestarnotificationsINotificationRuleTarget
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ITopic) AddSubscription(subscription ITopicSubscription) {
	_jsii_.InvokeVoid(
		i,
		"addSubscription",
		[]interface{}{subscription},
	)
}

func (i *jsiiProxy_ITopic) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) GrantPublish(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublish",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfMessagesPublished",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsDelivered",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOutInvalidAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOutNoMessageAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPublishSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSMSMonthToDateSpentUSD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSMSSuccessRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ITopic) BindAsNotificationRuleTarget(scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITopic) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Topic subscription.
// Experimental.
type ITopicSubscription interface {
	// Returns a configuration used to subscribe to an SNS topic.
	// Experimental.
	Bind(topic ITopic) *TopicSubscriptionConfig
}

// The jsii proxy for ITopicSubscription
type jsiiProxy_ITopicSubscription struct {
	_ byte // padding
}

func (i *jsiiProxy_ITopicSubscription) Bind(topic ITopic) *TopicSubscriptionConfig {
	var returns *TopicSubscriptionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

// Conditions that can be applied to numeric attributes.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
// Experimental.
type NumericConditions struct {
	// Match one or more values.
	// Experimental.
	Allowlist *[]*float64 `json:"allowlist" yaml:"allowlist"`
	// Match values that are between the specified values.
	// Experimental.
	Between *BetweenCondition `json:"between" yaml:"between"`
	// Match values that are strictly between the specified values.
	// Experimental.
	BetweenStrict *BetweenCondition `json:"betweenStrict" yaml:"betweenStrict"`
	// Match values that are greater than the specified value.
	// Experimental.
	GreaterThan *float64 `json:"greaterThan" yaml:"greaterThan"`
	// Match values that are greater than or equal to the specified value.
	// Experimental.
	GreaterThanOrEqualTo *float64 `json:"greaterThanOrEqualTo" yaml:"greaterThanOrEqualTo"`
	// Match values that are less than the specified value.
	// Experimental.
	LessThan *float64 `json:"lessThan" yaml:"lessThan"`
	// Match values that are less than or equal to the specified value.
	// Experimental.
	LessThanOrEqualTo *float64 `json:"lessThanOrEqualTo" yaml:"lessThanOrEqualTo"`
	// Match one or more values.
	// Deprecated: use `allowlist`.
	Whitelist *[]*float64 `json:"whitelist" yaml:"whitelist"`
}

// Conditions that can be applied to string attributes.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
// Experimental.
type StringConditions struct {
	// Match one or more values.
	// Experimental.
	Allowlist *[]*string `json:"allowlist" yaml:"allowlist"`
	// Match any value that doesn't include any of the specified values.
	// Deprecated: use `denylist`.
	Blacklist *[]*string `json:"blacklist" yaml:"blacklist"`
	// Match any value that doesn't include any of the specified values.
	// Experimental.
	Denylist *[]*string `json:"denylist" yaml:"denylist"`
	// Matches values that begins with the specified prefixes.
	// Experimental.
	MatchPrefixes *[]*string `json:"matchPrefixes" yaml:"matchPrefixes"`
	// Match one or more values.
	// Deprecated: use `allowlist`.
	Whitelist *[]*string `json:"whitelist" yaml:"whitelist"`
}

// A new subscription.
//
// Prefer to use the `ITopic.addSubscription()` methods to create instances of
// this class.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DeliveryStream awscdk.DeliveryStream
//   var stream deliveryStream
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
//   })
//
// Experimental.
type Subscription interface {
	awscdk.Resource
	// The DLQ associated with this subscription if present.
	// Experimental.
	DeadLetterQueue() awssqs.IQueue
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for Subscription
type jsiiProxy_Subscription struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Subscription) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subscription) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subscription) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subscription) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subscription) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSubscription(scope constructs.Construct, id *string, props *SubscriptionProps) Subscription {
	_init_.Initialize()

	j := jsiiProxy_Subscription{}

	_jsii_.Create(
		"monocdk.aws_sns.Subscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSubscription_Override(s Subscription, scope constructs.Construct, id *string, props *SubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.Subscription",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Subscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.Subscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Subscription_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.Subscription",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Subscription) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Subscription) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Subscription) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Subscription) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Subscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subscription) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A subscription filter for an attribute.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   // Lambda should receive only message matching the following conditions on attributes:
//   // color: 'red' or 'orange' or begins with 'bl'
//   // size: anything but 'small' or 'medium'
//   // price: between 100 and 200 or greater than 300
//   // store: attribute must be present
//   myTopic.addSubscription(subscriptions.NewLambdaSubscription(fn, &lambdaSubscriptionProps{
//   	filterPolicy: map[string]subscriptionFilter{
//   		"color": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"allowlist": []*string{
//   				jsii.String("red"),
//   				jsii.String("orange"),
//   			},
//   			"matchPrefixes": []*string{
//   				jsii.String("bl"),
//   			},
//   		}),
//   		"size": sns.*subscriptionFilter.stringFilter(&StringConditions{
//   			"denylist": []*string{
//   				jsii.String("small"),
//   				jsii.String("medium"),
//   			},
//   		}),
//   		"price": sns.*subscriptionFilter.numericFilter(&NumericConditions{
//   			"between": &BetweenCondition{
//   				"start": jsii.Number(100),
//   				"stop": jsii.Number(200),
//   			},
//   			"greaterThan": jsii.Number(300),
//   		}),
//   		"store": sns.*subscriptionFilter.existsFilter(),
//   	},
//   }))
//
// Experimental.
type SubscriptionFilter interface {
	// conditions that specify the message attributes that should be included, excluded, matched, etc.
	// Experimental.
	Conditions() *[]interface{}
}

// The jsii proxy struct for SubscriptionFilter
type jsiiProxy_SubscriptionFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_SubscriptionFilter) Conditions() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}


// Experimental.
func NewSubscriptionFilter(conditions *[]interface{}) SubscriptionFilter {
	_init_.Initialize()

	j := jsiiProxy_SubscriptionFilter{}

	_jsii_.Create(
		"monocdk.aws_sns.SubscriptionFilter",
		[]interface{}{conditions},
		&j,
	)

	return &j
}

// Experimental.
func NewSubscriptionFilter_Override(s SubscriptionFilter, conditions *[]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.SubscriptionFilter",
		[]interface{}{conditions},
		s,
	)
}

// Returns a subscription filter for attribute key matching.
// Experimental.
func SubscriptionFilter_ExistsFilter() SubscriptionFilter {
	_init_.Initialize()

	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.SubscriptionFilter",
		"existsFilter",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a subscription filter for a numeric attribute.
// Experimental.
func SubscriptionFilter_NumericFilter(numericConditions *NumericConditions) SubscriptionFilter {
	_init_.Initialize()

	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.SubscriptionFilter",
		"numericFilter",
		[]interface{}{numericConditions},
		&returns,
	)

	return returns
}

// Returns a subscription filter for a string attribute.
// Experimental.
func SubscriptionFilter_StringFilter(stringConditions *StringConditions) SubscriptionFilter {
	_init_.Initialize()

	var returns SubscriptionFilter

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.SubscriptionFilter",
		"stringFilter",
		[]interface{}{stringConditions},
		&returns,
	)

	return returns
}

// Options for creating a new subscription.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import sqs "github.com/aws/aws-cdk-go/awscdk/aws_sqs"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//   subscriptionOptions := &subscriptionOptions{
//   	endpoint: jsii.String("endpoint"),
//   	protocol: sns.subscriptionProtocol_HTTP,
//
//   	// the properties below are optional
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	rawMessageDelivery: jsii.Boolean(false),
//   	region: jsii.String("region"),
//   	subscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
// Experimental.
type SubscriptionOptions struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	// Experimental.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	// Experimental.
	Protocol SubscriptionProtocol `json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]SubscriptionFilter `json:"filterPolicy" yaml:"filterPolicy"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	// Experimental.
	RawMessageDelivery *bool `json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	// Experimental.
	SubscriptionRoleArn *string `json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

// Properties for creating a new subscription.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DeliveryStream awscdk.DeliveryStream
//   var stream deliveryStream
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
//   })
//
// Experimental.
type SubscriptionProps struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	// Experimental.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	// Experimental.
	Protocol SubscriptionProtocol `json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]SubscriptionFilter `json:"filterPolicy" yaml:"filterPolicy"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	// Experimental.
	RawMessageDelivery *bool `json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	// Experimental.
	SubscriptionRoleArn *string `json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
	// The topic to subscribe to.
	// Experimental.
	Topic ITopic `json:"topic" yaml:"topic"`
}

// The type of subscription, controlling the type of the endpoint parameter.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DeliveryStream awscdk.DeliveryStream
//   var stream deliveryStream
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
//   })
//
// Experimental.
type SubscriptionProtocol string

const (
	// JSON-encoded message is POSTED to an HTTP url.
	// Experimental.
	SubscriptionProtocol_HTTP SubscriptionProtocol = "HTTP"
	// JSON-encoded message is POSTed to an HTTPS url.
	// Experimental.
	SubscriptionProtocol_HTTPS SubscriptionProtocol = "HTTPS"
	// Notifications are sent via email.
	// Experimental.
	SubscriptionProtocol_EMAIL SubscriptionProtocol = "EMAIL"
	// Notifications are JSON-encoded and sent via mail.
	// Experimental.
	SubscriptionProtocol_EMAIL_JSON SubscriptionProtocol = "EMAIL_JSON"
	// Notification is delivered by SMS.
	// Experimental.
	SubscriptionProtocol_SMS SubscriptionProtocol = "SMS"
	// Notifications are enqueued into an SQS queue.
	// Experimental.
	SubscriptionProtocol_SQS SubscriptionProtocol = "SQS"
	// JSON-encoded notifications are sent to a mobile app endpoint.
	// Experimental.
	SubscriptionProtocol_APPLICATION SubscriptionProtocol = "APPLICATION"
	// Notifications trigger a Lambda function.
	// Experimental.
	SubscriptionProtocol_LAMBDA SubscriptionProtocol = "LAMBDA"
	// Notifications put records into a firehose delivery stream.
	// Experimental.
	SubscriptionProtocol_FIREHOSE SubscriptionProtocol = "FIREHOSE"
)

// A new SNS topic.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type Topic interface {
	TopicBase
	// Controls automatic creation of policy objects.
	//
	// Set by subclasses.
	// Experimental.
	AutoCreatePolicy() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	// Experimental.
	Fifo() *bool
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The ARN of the topic.
	// Experimental.
	TopicArn() *string
	// The name of the topic.
	// Experimental.
	TopicName() *string
	// Subscribe some endpoint to this topic.
	// Experimental.
	AddSubscription(subscription ITopicSubscription)
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Represents a notification target That allows SNS topic to associate with this rule target.
	// Experimental.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant topic publishing permissions to the given identity.
	// Experimental.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	// Experimental.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for Topic
type jsiiProxy_Topic struct {
	jsiiProxy_TopicBase
}

func (j *jsiiProxy_Topic) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopic(scope constructs.Construct, id *string, props *TopicProps) Topic {
	_init_.Initialize()

	j := jsiiProxy_Topic{}

	_jsii_.Create(
		"monocdk.aws_sns.Topic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopic_Override(t Topic, scope constructs.Construct, id *string, props *TopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.Topic",
		[]interface{}{scope, id, props},
		t,
	)
}

// Import an existing SNS topic provided an ARN.
// Experimental.
func Topic_FromTopicArn(scope constructs.Construct, id *string, topicArn *string) ITopic {
	_init_.Initialize()

	var returns ITopic

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.Topic",
		"fromTopicArn",
		[]interface{}{scope, id, topicArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Topic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.Topic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Topic_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.Topic",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) AddSubscription(subscription ITopicSubscription) {
	_jsii_.InvokeVoid(
		t,
		"addSubscription",
		[]interface{}{subscription},
	)
}

func (t *jsiiProxy_Topic) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		t,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_Topic) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		t,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfMessagesPublished",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsDelivered",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutInvalidAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutNoMessageAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricPublishSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSMonthToDateSpentUSD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSSuccessRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Topic) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Topic) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Topic) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Topic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Either a new or imported Topic.
// Experimental.
type TopicBase interface {
	awscdk.Resource
	ITopic
	// Controls automatic creation of policy objects.
	//
	// Set by subclasses.
	// Experimental.
	AutoCreatePolicy() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	// Experimental.
	Fifo() *bool
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The ARN of the topic.
	// Experimental.
	TopicArn() *string
	// The name of the topic.
	// Experimental.
	TopicName() *string
	// Subscribe some endpoint to this topic.
	// Experimental.
	AddSubscription(subscription ITopicSubscription)
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Represents a notification target That allows SNS topic to associate with this rule target.
	// Experimental.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant topic publishing permissions to the given identity.
	// Experimental.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	// Experimental.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for TopicBase
type jsiiProxy_TopicBase struct {
	internal.Type__awscdkResource
	jsiiProxy_ITopic
}

func (j *jsiiProxy_TopicBase) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicBase_Override(t TopicBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.TopicBase",
		[]interface{}{scope, id, props},
		t,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func TopicBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TopicBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) AddSubscription(subscription ITopicSubscription) {
	_jsii_.InvokeVoid(
		t,
		"addSubscription",
		[]interface{}{subscription},
	)
}

func (t *jsiiProxy_TopicBase) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		t,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TopicBase) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		t,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfMessagesPublished",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsDelivered",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutInvalidAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutNoMessageAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricPublishSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSMonthToDateSpentUSD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSSuccessRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The policy for an SNS Topic.
//
// Policies define the operations that are allowed on this resource.
//
// You almost never need to define this construct directly.
//
// All AWS resources that support resource policies have a method called
// `addToResourcePolicy()`, which will automatically create a new resource
// policy if one doesn't exist yet, otherwise it will add to the existing
// policy.
//
// Prefer to use `addToResourcePolicy()` instead.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &topicPolicyProps{
//   	topics: []iTopic{
//   		topic,
//   	},
//   })
//
//   topicPolicy.document.addStatements(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("sns:Subscribe"),
//   	},
//   	principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	resources: []*string{
//   		topic.topicArn,
//   	},
//   }))
//
// Experimental.
type TopicPolicy interface {
	awscdk.Resource
	// The IAM policy document for this policy.
	// Experimental.
	Document() awsiam.PolicyDocument
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for TopicPolicy
type jsiiProxy_TopicPolicy struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_TopicPolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicPolicy(scope constructs.Construct, id *string, props *TopicPolicyProps) TopicPolicy {
	_init_.Initialize()

	j := jsiiProxy_TopicPolicy{}

	_jsii_.Create(
		"monocdk.aws_sns.TopicPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicPolicy_Override(t TopicPolicy, scope constructs.Construct, id *string, props *TopicPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.TopicPolicy",
		[]interface{}{scope, id, props},
		t,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func TopicPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TopicPolicy_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TopicPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to associate SNS topics with a policy.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("TopicPolicy"), &topicPolicyProps{
//   	topics: []iTopic{
//   		topic,
//   	},
//   })
//
//   topicPolicy.document.addStatements(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("sns:Subscribe"),
//   	},
//   	principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	resources: []*string{
//   		topic.topicArn,
//   	},
//   }))
//
// Experimental.
type TopicPolicyProps struct {
	// The set of topics this policy applies to.
	// Experimental.
	Topics *[]ITopic `json:"topics" yaml:"topics"`
	// IAM policy document to apply to topic(s).
	// Experimental.
	PolicyDocument awsiam.PolicyDocument `json:"policyDocument" yaml:"policyDocument"`
}

// Properties for a new SNS topic.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"), &topicProps{
//   	displayName: jsii.String("Customer subscription topic"),
//   })
//
// Experimental.
type TopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	// Experimental.
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// A developer-defined string that can be used to identify this SNS topic.
	// Experimental.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	// Experimental.
	Fifo *bool `json:"fifo" yaml:"fifo"`
	// A KMS Key, either managed by this CDK app, or imported.
	// Experimental.
	MasterKey awskms.IKey `json:"masterKey" yaml:"masterKey"`
	// A name for the topic.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique
	// physical ID and uses that ID for the topic name. For more information,
	// see Name Type.
	// Experimental.
	TopicName *string `json:"topicName" yaml:"topicName"`
}

// Subscription configuration.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import sqs "github.com/aws/aws-cdk-go/awscdk/aws_sqs"
//
//   var construct construct
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//   topicSubscriptionConfig := &topicSubscriptionConfig{
//   	endpoint: jsii.String("endpoint"),
//   	protocol: sns.subscriptionProtocol_HTTP,
//   	subscriberId: jsii.String("subscriberId"),
//
//   	// the properties below are optional
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	rawMessageDelivery: jsii.Boolean(false),
//   	region: jsii.String("region"),
//   	subscriberScope: construct,
//   	subscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
// Experimental.
type TopicSubscriptionConfig struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	// Experimental.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	// Experimental.
	Protocol SubscriptionProtocol `json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]SubscriptionFilter `json:"filterPolicy" yaml:"filterPolicy"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	// Experimental.
	RawMessageDelivery *bool `json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	// Experimental.
	SubscriptionRoleArn *string `json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
	// The id of the SNS subscription resource created under `scope`.
	//
	// In most
	// cases, it is recommended to use the `uniqueId` of the topic you are
	// subscribing to.
	// Experimental.
	SubscriberId *string `json:"subscriberId" yaml:"subscriberId"`
	// The scope in which to create the SNS subscription resource.
	//
	// Normally you'd
	// want the subscription to be created on the consuming stack because the
	// topic is usually referenced by the consumer's resource policy (e.g. SQS
	// queue policy). Otherwise, it will cause a cyclic reference.
	//
	// If this is undefined, the subscription will be created on the topic's stack.
	// Experimental.
	SubscriberScope awscdk.Construct `json:"subscriberScope" yaml:"subscriberScope"`
}

