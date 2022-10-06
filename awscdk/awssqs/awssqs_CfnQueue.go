package awssqs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SQS::Queue`.
//
// The `AWS::SQS::Queue` resource creates an Amazon SQS standard or FIFO queue.
//
// Keep the following caveats in mind:
//
// - If you don't specify the `FifoQueue` property, Amazon SQS creates a standard queue.
//
// > You can't change the queue type after you create it and you can't convert an existing standard queue into a FIFO queue. You must either create a new FIFO queue for your application or delete your existing standard queue and recreate it as a FIFO queue. For more information, see [Moving from a standard queue to a FIFO queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-moving.html) in the *Amazon SQS Developer Guide* .
// - If you don't provide a value for a property, the queue is created with the default value for the property.
// - If you delete a queue, you must wait at least 60 seconds before creating a queue with the same name.
// - To successfully create a new queue, you must provide a queue name that adheres to the [limits related to queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html) and is unique within the scope of your queues.
//
// For more information about creating FIFO (first-in-first-out) queues, see [Creating an Amazon SQS queue ( AWS CloudFormation )](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/screate-queue-cloudformation.html) in the *Amazon SQS Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var redriveAllowPolicy interface{}
//   var redrivePolicy interface{}
//
//   cfnQueue := awscdk.Aws_sqs.NewCfnQueue(this, jsii.String("MyCfnQueue"), &cfnQueueProps{
//   	contentBasedDeduplication: jsii.Boolean(false),
//   	deduplicationScope: jsii.String("deduplicationScope"),
//   	delaySeconds: jsii.Number(123),
//   	fifoQueue: jsii.Boolean(false),
//   	fifoThroughputLimit: jsii.String("fifoThroughputLimit"),
//   	kmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	maximumMessageSize: jsii.Number(123),
//   	messageRetentionPeriod: jsii.Number(123),
//   	queueName: jsii.String("queueName"),
//   	receiveMessageWaitTimeSeconds: jsii.Number(123),
//   	redriveAllowPolicy: redriveAllowPolicy,
//   	redrivePolicy: redrivePolicy,
//   	sqsManagedSseEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visibilityTimeout: jsii.Number(123),
//   })
//
type CfnQueue interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the queue.
	//
	// For example: `arn:aws:sqs:us-east-2:123456789012:mystack-myqueue-15PG5C2FC1CW8` .
	AttrArn() *string
	// Returns the queue name.
	//
	// For example: `mystack-myqueue-1VF9BKQH5BJVI` .
	AttrQueueName() *string
	AttrQueueUrl() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication.
	//
	// During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the `ContentBasedDeduplication` attribute for the `[CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)` action in the *Amazon SQS API Reference* .
	ContentBasedDeduplication() interface{}
	SetContentBasedDeduplication(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level.
	//
	// Valid values are `messageGroup` and `queue` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `messageGroup` *and* set the `FifoThroughputLimit` attribute to `perMessageGroupId` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	DeduplicationScope() *string
	SetDeduplicationScope(val *string)
	// The time in seconds for which the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of `0` to `900` (15 minutes). The default value is `0` .
	DelaySeconds() *float64
	SetDelaySeconds(val *float64)
	// If set to true, creates a FIFO queue.
	//
	// If you don't specify this property, Amazon SQS creates a standard queue. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Amazon SQS Developer Guide* .
	FifoQueue() interface{}
	SetFifoQueue(val interface{})
	// For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// Valid values are `perQueue` and `perMessageGroupId` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `perMessageGroupId` *and* set the `DeduplicationScope` attribute to `messageGroup` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	FifoThroughputLimit() *string
	SetFifoThroughputLimit(val *string)
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	//
	// > A shorter time period provides better security, but results in more calls to AWS KMS , which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Amazon SQS Developer Guide* .
	KmsDataKeyReusePeriodSeconds() *float64
	SetKmsDataKeyReusePeriodSeconds(val *float64)
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK.
	//
	// To use the AWS managed CMK for Amazon SQS , specify the (default) alias `alias/aws/sqs` . For more information, see the following:
	//
	// - [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Amazon SQS Developer Guide*
	// - [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *Amazon SQS API Reference*
	// - The Customer Master Keys section of the [AWS Key Management Service Best Practices](https://docs.aws.amazon.com/https://d0.awsstatic.com/whitepapers/aws-kms-best-practices.pdf) whitepaper
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
	LogicalId() *string
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from `1,024` bytes (1 KiB) to `262,144` bytes (256 KiB). The default value is `262,144` (256 KiB).
	MaximumMessageSize() *float64
	SetMaximumMessageSize(val *float64)
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from `60` seconds (1 minute) to `1,209,600` seconds (14 days). The default value is `345,600` seconds (4 days).
	MessageRetentionPeriod() *float64
	SetMessageRetentionPeriod(val *float64)
	// The tree node.
	Node() constructs.Node
	// A name for the queue.
	//
	// To create a FIFO queue, the name of your FIFO queue must end with the `.fifo` suffix. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Amazon SQS Developer Guide* .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *AWS CloudFormation User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	QueueName() *string
	SetQueueName(val *string)
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available.
	//
	// You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Amazon SQS Developer Guide* .
	ReceiveMessageWaitTimeSeconds() *float64
	SetReceiveMessageWaitTimeSeconds(val *float64)
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	//
	// The parameters are as follows:
	//
	// - `redrivePermission` : The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:
	//
	// - `allowAll` : (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.
	// - `denyAll` : No source queues can specify this queue as the dead-letter queue.
	// - `byQueue` : Only queues specified by the `sourceQueueArns` parameter can specify this queue as the dead-letter queue.
	// - `sourceQueueArns` : The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the `redrivePermission` parameter is set to `byQueue` . You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the `redrivePermission` parameter to `allowAll` .
	RedriveAllowPolicy() interface{}
	SetRedriveAllowPolicy(val interface{})
	// The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object.
	//
	// The parameters are as follows:
	//
	// - `deadLetterTargetArn` : The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon SQS moves messages after the value of `maxReceiveCount` is exceeded.
	// - `maxReceiveCount` : The number of times a message is delivered to the source queue before being moved to the dead-letter queue. When the `ReceiveCount` for a message exceeds the `maxReceiveCount` for a queue, Amazon SQS moves the message to the dead-letter-queue.
	//
	// > The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.
	//
	// *JSON*
	//
	// `{ "deadLetterTargetArn" : *String* , "maxReceiveCount" : *Integer* }`
	//
	// *YAML*
	//
	// `deadLetterTargetArn : *String*`
	//
	// `maxReceiveCount : *Integer*`.
	RedrivePolicy() interface{}
	SetRedrivePolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::SQS::Queue.SqsManagedSseEnabled`.
	SqsManagedSseEnabled() interface{}
	SetSqsManagedSseEnabled(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags that you attach to this queue.
	//
	// For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
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
	// The length of time during which a message will be unavailable after a message is delivered from the queue.
	//
	// This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.
	//
	// Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	//
	// For more information about Amazon SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Amazon SQS Developer Guide* .
	VisibilityTimeout() *float64
	SetVisibilityTimeout(val *float64)
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

// The jsii proxy struct for CfnQueue
type jsiiProxy_CfnQueue struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnQueue) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) AttrQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrQueueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) AttrQueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrQueueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) DeduplicationScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) DelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) FifoQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) FifoThroughputLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) KmsDataKeyReusePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) MaximumMessageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumMessageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) MessageRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) ReceiveMessageWaitTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveMessageWaitTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) RedriveAllowPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redriveAllowPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) RedrivePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) SqsManagedSseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsManagedSseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueue) VisibilityTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeout",
		&returns,
	)
	return returns
}


// Create a new `AWS::SQS::Queue`.
func NewCfnQueue(scope constructs.Construct, id *string, props *CfnQueueProps) CfnQueue {
	_init_.Initialize()

	if err := validateNewCfnQueueParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQueue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SQS::Queue`.
func NewCfnQueue_Override(c CfnQueue, scope constructs.Construct, id *string, props *CfnQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnQueue)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetDeduplicationScope(val *string) {
	_jsii_.Set(
		j,
		"deduplicationScope",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetDelaySeconds(val *float64) {
	_jsii_.Set(
		j,
		"delaySeconds",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetFifoQueue(val interface{}) {
	if err := j.validateSetFifoQueueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoQueue",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetFifoThroughputLimit(val *string) {
	_jsii_.Set(
		j,
		"fifoThroughputLimit",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetKmsDataKeyReusePeriodSeconds(val *float64) {
	_jsii_.Set(
		j,
		"kmsDataKeyReusePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetKmsMasterKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetMaximumMessageSize(val *float64) {
	_jsii_.Set(
		j,
		"maximumMessageSize",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetMessageRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"messageRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetQueueName(val *string) {
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetReceiveMessageWaitTimeSeconds(val *float64) {
	_jsii_.Set(
		j,
		"receiveMessageWaitTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetRedriveAllowPolicy(val interface{}) {
	if err := j.validateSetRedriveAllowPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redriveAllowPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetRedrivePolicy(val interface{}) {
	if err := j.validateSetRedrivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetSqsManagedSseEnabled(val interface{}) {
	if err := j.validateSetSqsManagedSseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsManagedSseEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnQueue)SetVisibilityTimeout(val *float64) {
	_jsii_.Set(
		j,
		"visibilityTimeout",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnQueue_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueue_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnQueue_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnQueue_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.CfnQueue",
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
func CfnQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueue_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueue) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnQueue) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnQueue) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnQueue) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnQueue) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnQueue) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnQueue) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueue) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnQueue) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnQueue) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnQueue) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnQueue) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueue) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

