package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a Kinesis Stream.
//
// Example:
//   key := kms.NewKey(this, jsii.String("MyKey"))
//
//   kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &streamProps{
//   	encryption: kinesis.streamEncryption_KMS,
//   	encryptionKey: key,
//   })
//
type StreamProps struct {
	// The kind of server-side encryption to apply to this stream.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	Encryption StreamEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for stream encryption.
	//
	// The 'encryption' property must be set to "Kms".
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The number of shards for the stream.
	//
	// Can only be provided if streamMode is Provisioned.
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// The capacity mode of this stream.
	StreamMode StreamMode `field:"optional" json:"streamMode" yaml:"streamMode"`
	// Enforces a particular physical stream name.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

