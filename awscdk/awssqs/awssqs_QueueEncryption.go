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
// Experimental.
type QueueEncryption string

const (
	// Messages in the queue are not encrypted.
	// Experimental.
	QueueEncryption_UNENCRYPTED QueueEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a master key managed by SQS.
	// Experimental.
	QueueEncryption_KMS_MANAGED QueueEncryption = "KMS_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	// Experimental.
	QueueEncryption_KMS QueueEncryption = "KMS"
)

