package awsamazonmq


// Encryption options for the broker.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionOptionsProperty := &encryptionOptionsProperty{
//   	useAwsOwnedKey: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnBroker_EncryptionOptionsProperty struct {
	// Enables the use of an AWS owned CMK using AWS KMS (KMS).
	//
	// Set to `true` by default, if no value is provided, for example, for RabbitMQ brokers.
	UseAwsOwnedKey interface{} `field:"required" json:"useAwsOwnedKey" yaml:"useAwsOwnedKey"`
	// The customer master key (CMK) to use for the A AWS KMS (KMS).
	//
	// This key is used to encrypt your data at rest. If not provided, Amazon MQ will use a default CMK to encrypt your data.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

