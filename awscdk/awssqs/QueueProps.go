package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a new Queue.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   dlQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"), &QueueProps{
//   	QueueName: jsii.String("MySubscription_DLQ"),
//   	RetentionPeriod: awscdk.Duration_Days(jsii.Number(14)),
//   })
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &SubscriptionProps{
//   	Endpoint: jsii.String("endpoint"),
//   	Protocol: sns.SubscriptionProtocol_LAMBDA,
//   	Topic: Topic,
//   	DeadLetterQueue: dlQueue,
//   })
//
type QueueProps struct {
	// Specifies whether to enable content-based deduplication.
	//
	// During the deduplication interval (5 minutes), Amazon SQS treats
	// messages that are sent with identical content (excluding attributes) as
	// duplicates and delivers only one copy of the message.
	//
	// If you don't enable content-based deduplication and you want to deduplicate
	// messages, provide an explicit deduplication ID in your SendMessage() call.
	//
	// (Only applies to FIFO queues.)
	// Default: false.
	//
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// The length of time that Amazon SQS reuses a data key before calling KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24
	// hours). The default is 300 (5 minutes).
	// Default: Duration.minutes(5)
	//
	DataKeyReuse awscdk.Duration `field:"optional" json:"dataKeyReuse" yaml:"dataKeyReuse"`
	// Send messages to this queue if they were unsuccessfully dequeued a number of times.
	// Default: no dead-letter queue.
	//
	DeadLetterQueue *DeadLetterQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level.
	//
	// (Only applies to FIFO queues.)
	// Default: DeduplicationScope.QUEUE
	//
	DeduplicationScope DeduplicationScope `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// The time in seconds that the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of 0 to 900 (15 minutes). The default
	// value is 0.
	// Default: 0.
	//
	DeliveryDelay awscdk.Duration `field:"optional" json:"deliveryDelay" yaml:"deliveryDelay"`
	// Whether the contents of the queue are encrypted, and by what type of key.
	//
	// Be aware that encryption is not available in all regions, please see the docs
	// for current availability details.
	// Default: SQS_MANAGED (SSE-SQS) for newly created queues.
	//
	Encryption QueueEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for queue encryption.
	//
	// Individual messages will be encrypted using data keys. The data keys in
	// turn will be encrypted using this key, and reused for a maximum of
	// `dataKeyReuseSecs` seconds.
	//
	// If the 'encryptionMasterKey' property is set, 'encryption' type will be
	// implicitly set to "KMS".
	// Default: If encryption is set to KMS and not specified, a key will be created.
	//
	EncryptionMasterKey awskms.IKey `field:"optional" json:"encryptionMasterKey" yaml:"encryptionMasterKey"`
	// Enforce encryption of data in transit.
	// See: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-security-best-practices.html#enforce-encryption-data-in-transit
	//
	// Default: false.
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this a first-in-first-out (FIFO) queue.
	// Default: false, unless queueName ends in '.fifo' or 'contentBasedDeduplication' is true.
	//
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// (Only applies to FIFO queues.)
	// Default: FifoThroughputLimit.PER_QUEUE
	//
	FifoThroughputLimit FifoThroughputLimit `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from 1024 bytes (1 KiB) to 262144 bytes
	// (256 KiB). The default value is 262144 (256 KiB).
	// Default: 256KiB.
	//
	MaxMessageSizeBytes *float64 `field:"optional" json:"maxMessageSizeBytes" yaml:"maxMessageSizeBytes"`
	// A name for the queue.
	//
	// If specified and this is a FIFO queue, must end in the string '.fifo'.
	// Default: CloudFormation-generated name.
	//
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Default wait time for ReceiveMessage calls.
	//
	// Does not wait if set to 0, otherwise waits this amount of seconds
	// by default for messages to arrive.
	//
	// For more information, see Amazon SQS Long Poll.
	// Default: 0.
	//
	ReceiveMessageWaitTime awscdk.Duration `field:"optional" json:"receiveMessageWaitTime" yaml:"receiveMessageWaitTime"`
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues.
	// Default: - All source queues can designate this queue as their dead-letter queue.
	//
	RedriveAllowPolicy *RedriveAllowPolicy `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// Policy to apply when the queue is removed from the stack.
	//
	// Even though queues are technically stateful, their contents are transient and it
	// is common to add and remove Queues while rearchitecting your application. The
	// default is therefore `DESTROY`. Change it to `RETAIN` if the messages are so
	// valuable that accidentally losing them would be unacceptable.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from 60 seconds (1 minute) to 1209600
	// seconds (14 days). The default value is 345600 seconds (4 days).
	// Default: Duration.days(4)
	//
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message
	// and delete it from the queue before it becomes visible again for dequeueing
	// by another processor.
	//
	// Values must be from 0 to 43200 seconds (12 hours). If you don't specify
	// a value, AWS CloudFormation uses the default value of 30 seconds.
	// Default: Duration.seconds(30)
	//
	VisibilityTimeout awscdk.Duration `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

