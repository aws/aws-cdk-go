package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for creating a new Queue.
//
// Example:
//   // Use managed key
//   // Use managed key
//   sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
//   	encryption: sqs.queueEncryption_KMS_MANAGED,
//   })
//
//   // Use custom key
//   myKey := kms.NewKey(this, jsii.String("Key"))
//
//   sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
//   	encryption: sqs.*queueEncryption_KMS,
//   	encryptionMasterKey: myKey,
//   })
//
// Experimental.
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
	// Experimental.
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// The length of time that Amazon SQS reuses a data key before calling KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24
	// hours). The default is 300 (5 minutes).
	// Experimental.
	DataKeyReuse awscdk.Duration `field:"optional" json:"dataKeyReuse" yaml:"dataKeyReuse"`
	// Send messages to this queue if they were unsuccessfully dequeued a number of times.
	// Experimental.
	DeadLetterQueue *DeadLetterQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level.
	//
	// (Only applies to FIFO queues.)
	// Experimental.
	DeduplicationScope DeduplicationScope `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// The time in seconds that the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of 0 to 900 (15 minutes). The default
	// value is 0.
	// Experimental.
	DeliveryDelay awscdk.Duration `field:"optional" json:"deliveryDelay" yaml:"deliveryDelay"`
	// Whether the contents of the queue are encrypted, and by what type of key.
	//
	// Be aware that encryption is not available in all regions, please see the docs
	// for current availability details.
	// Experimental.
	Encryption QueueEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS master key to use for queue encryption.
	//
	// Individual messages will be encrypted using data keys. The data keys in
	// turn will be encrypted using this key, and reused for a maximum of
	// `dataKeyReuseSecs` seconds.
	//
	// If the 'encryptionMasterKey' property is set, 'encryption' type will be
	// implicitly set to "KMS".
	// Experimental.
	EncryptionMasterKey awskms.IKey `field:"optional" json:"encryptionMasterKey" yaml:"encryptionMasterKey"`
	// Whether this a first-in-first-out (FIFO) queue.
	// Experimental.
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// (Only applies to FIFO queues.)
	// Experimental.
	FifoThroughputLimit FifoThroughputLimit `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from 1024 bytes (1 KiB) to 262144 bytes
	// (256 KiB). The default value is 262144 (256 KiB).
	// Experimental.
	MaxMessageSizeBytes *float64 `field:"optional" json:"maxMessageSizeBytes" yaml:"maxMessageSizeBytes"`
	// A name for the queue.
	//
	// If specified and this is a FIFO queue, must end in the string '.fifo'.
	// Experimental.
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Default wait time for ReceiveMessage calls.
	//
	// Does not wait if set to 0, otherwise waits this amount of seconds
	// by default for messages to arrive.
	//
	// For more information, see Amazon SQS Long Poll.
	// Experimental.
	ReceiveMessageWaitTime awscdk.Duration `field:"optional" json:"receiveMessageWaitTime" yaml:"receiveMessageWaitTime"`
	// Policy to apply when the queue is removed from the stack.
	//
	// Even though queues are technically stateful, their contents are transient and it
	// is common to add and remove Queues while rearchitecting your application. The
	// default is therefore `DESTROY`. Change it to `RETAIN` if the messages are so
	// valuable that accidentally losing them would be unacceptable.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from 60 seconds (1 minute) to 1209600
	// seconds (14 days). The default value is 345600 seconds (4 days).
	// Experimental.
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message
	// and delete it from the queue before it becomes visible again for dequeueing
	// by another processor.
	//
	// Values must be from 0 to 43200 seconds (12 hours). If you don't specify
	// a value, AWS CloudFormation uses the default value of 30 seconds.
	// Experimental.
	VisibilityTimeout awscdk.Duration `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

