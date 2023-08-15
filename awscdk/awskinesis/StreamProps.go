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
//   kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
//   	Encryption: kinesis.StreamEncryption_KMS,
//   	EncryptionKey: key,
//   })
//
type StreamProps struct {
	// The kind of server-side encryption to apply to this stream.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	// Default: - StreamEncryption.KMS if encrypted Streams are supported in the region
	// or StreamEncryption.UNENCRYPTED otherwise.
	// StreamEncryption.KMS if an encryption key is supplied through the encryptionKey property
	//
	Encryption StreamEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for stream encryption.
	//
	// The 'encryption' property must be set to "Kms".
	// Default: - Kinesis Data Streams master key ('/alias/aws/kinesis').
	// If encryption is set to StreamEncryption.KMS and this property is undefined, a new KMS key
	// will be created and associated with this stream.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	// Default: Duration.hours(24)
	//
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The number of shards for the stream.
	//
	// Can only be provided if streamMode is Provisioned.
	// Default: 1.
	//
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// The capacity mode of this stream.
	// Default: StreamMode.PROVISIONED
	//
	StreamMode StreamMode `field:"optional" json:"streamMode" yaml:"streamMode"`
	// Enforces a particular physical stream name.
	// Default: <generated>.
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

