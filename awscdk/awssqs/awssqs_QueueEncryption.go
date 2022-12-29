package awssqs


// What kind of encryption to apply to this queue.
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
//   // Use SQS managed server side encryption (SSE-SQS)
//   // Use SQS managed server side encryption (SSE-SQS)
//   sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
//   	encryption: sqs.*queueEncryption_SQS_MANAGED,
//   })
//
//   // Unencrypted queue
//   // Unencrypted queue
//   sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
//   	encryption: sqs.*queueEncryption_UNENCRYPTED,
//   })
//
type QueueEncryption string

const (
	// Messages in the queue are not encrypted.
	QueueEncryption_UNENCRYPTED QueueEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a KMS key managed by SQS.
	QueueEncryption_KMS_MANAGED QueueEncryption = "KMS_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	QueueEncryption_KMS QueueEncryption = "KMS"
	// Server-side encryption key managed by SQS (SSE-SQS).
	//
	// To learn more about SSE-SQS on Amazon SQS, please visit the
	// [Amazon SQS documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html).
	QueueEncryption_SQS_MANAGED QueueEncryption = "SQS_MANAGED"
)

