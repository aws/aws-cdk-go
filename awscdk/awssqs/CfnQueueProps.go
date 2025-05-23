package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQueue`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var redriveAllowPolicy interface{}
//   var redrivePolicy interface{}
//
//   cfnQueueProps := &CfnQueueProps{
//   	ContentBasedDeduplication: jsii.Boolean(false),
//   	DeduplicationScope: jsii.String("deduplicationScope"),
//   	DelaySeconds: jsii.Number(123),
//   	FifoQueue: jsii.Boolean(false),
//   	FifoThroughputLimit: jsii.String("fifoThroughputLimit"),
//   	KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	MaximumMessageSize: jsii.Number(123),
//   	MessageRetentionPeriod: jsii.Number(123),
//   	QueueName: jsii.String("queueName"),
//   	ReceiveMessageWaitTimeSeconds: jsii.Number(123),
//   	RedriveAllowPolicy: redriveAllowPolicy,
//   	RedrivePolicy: redrivePolicy,
//   	SqsManagedSseEnabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisibilityTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html
//
type CfnQueueProps struct {
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication.
	//
	// During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the `ContentBasedDeduplication` attribute for the `[CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)` action in the *Amazon SQS API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-contentbaseddeduplication
	//
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level.
	//
	// Valid values are `messageGroup` and `queue` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `messageGroup` *and* set the `FifoThroughputLimit` attribute to `perMessageGroupId` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-deduplicationscope
	//
	DeduplicationScope *string `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// The time in seconds for which the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of `0` to `900` (15 minutes). The default value is `0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-delayseconds
	//
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// If set to true, creates a FIFO queue.
	//
	// If you don't specify this property, Amazon SQS creates a standard queue. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-fifoqueue
	//
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// Valid values are `perQueue` and `perMessageGroupId` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `perMessageGroupId` *and* set the `DeduplicationScope` attribute to `messageGroup` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-fifothroughputlimit
	//
	FifoThroughputLimit *string `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	//
	// > A shorter time period provides better security, but results in more calls to AWS KMS , which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-kmsdatakeyreuseperiodseconds
	//
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS Key Management Service (KMS) for Amazon SQS , or a custom KMS.
	//
	// To use the AWS managed KMS for Amazon SQS , specify a (default) alias ARN, alias name (for example `alias/aws/sqs` ), key ARN, or key ID. For more information, see the following:
	//
	// - [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Amazon SQS Developer Guide*
	// - [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *Amazon SQS API Reference*
	// - [Request Parameters](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the *AWS Key Management Service API Reference*
	// - The Key Management Service (KMS) section of the [Security best practices for AWS Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/best-practices.html) in the *AWS Key Management Service Developer Guide*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-kmsmasterkeyid
	//
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from `1,024` bytes (1 KiB) to `262,144` bytes (256 KiB). The default value is `262,144` (256 KiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-maximummessagesize
	//
	MaximumMessageSize *float64 `field:"optional" json:"maximumMessageSize" yaml:"maximumMessageSize"`
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from `60` seconds (1 minute) to `1,209,600` seconds (14 days). The default value is `345,600` seconds (4 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-messageretentionperiod
	//
	MessageRetentionPeriod *float64 `field:"optional" json:"messageRetentionPeriod" yaml:"messageRetentionPeriod"`
	// A name for the queue.
	//
	// To create a FIFO queue, the name of your FIFO queue must end with the `.fifo` suffix. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Amazon SQS Developer Guide* .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *AWS CloudFormation User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-queuename
	//
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available.
	//
	// You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-receivemessagewaittimeseconds
	//
	ReceiveMessageWaitTimeSeconds *float64 `field:"optional" json:"receiveMessageWaitTimeSeconds" yaml:"receiveMessageWaitTimeSeconds"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-redriveallowpolicy
	//
	RedriveAllowPolicy interface{} `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object.
	//
	// The parameters are as follows:
	//
	// - `deadLetterTargetArn` : The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon SQS moves messages after the value of `maxReceiveCount` is exceeded.
	// - `maxReceiveCount` : The number of times a message is received by a consumer of the source queue before being moved to the dead-letter queue. When the `ReceiveCount` for a message exceeds the `maxReceiveCount` for a queue, Amazon SQS moves the message to the dead-letter-queue.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-redrivepolicy
	//
	RedrivePolicy interface{} `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Enables server-side queue encryption using SQS owned encryption keys.
	//
	// Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html) ). When `SqsManagedSseEnabled` is not defined, `SSE-SQS` encryption is enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-sqsmanagedsseenabled
	//
	SqsManagedSseEnabled interface{} `field:"optional" json:"sqsManagedSseEnabled" yaml:"sqsManagedSseEnabled"`
	// The tags that you attach to this queue.
	//
	// For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The length of time during which a message will be unavailable after a message is delivered from the queue.
	//
	// This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.
	//
	// Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	//
	// For more information about Amazon SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html#cfn-sqs-queue-visibilitytimeout
	//
	VisibilityTimeout *float64 `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

