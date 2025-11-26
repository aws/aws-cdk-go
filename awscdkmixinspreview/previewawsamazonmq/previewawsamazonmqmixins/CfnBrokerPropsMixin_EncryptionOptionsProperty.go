package previewawsamazonmqmixins


// Encryption options for the broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionOptionsProperty := &EncryptionOptionsProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	UseAwsOwnedKey: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-encryptionoptions.html
//
type CfnBrokerPropsMixin_EncryptionOptionsProperty struct {
	// The customer master key (CMK) to use for the A AWS  (KMS).
	//
	// This key is used to encrypt your data at rest. If not provided, Amazon MQ will use a default CMK to encrypt your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-encryptionoptions.html#cfn-amazonmq-broker-encryptionoptions-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Enables the use of an AWS owned CMK using AWS  (KMS).
	//
	// Set to `true` by default, if no value is provided, for example, for RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-encryptionoptions.html#cfn-amazonmq-broker-encryptionoptions-useawsownedkey
	//
	UseAwsOwnedKey interface{} `field:"optional" json:"useAwsOwnedKey" yaml:"useAwsOwnedKey"`
}

