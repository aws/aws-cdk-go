package awskinesisfirehose


// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryStreamEncryptionConfigurationInputProperty := &deliveryStreamEncryptionConfigurationInputProperty{
//   	keyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty struct {
	// Indicates the type of customer master key (CMK) to use for encryption.
	//
	// The default setting is `AWS_OWNED_CMK` . For more information about CMKs, see [Customer Master Keys (CMKs)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys) .
	//
	// You can use a CMK of type CUSTOMER_MANAGED_CMK to encrypt up to 500 delivery streams.
	//
	// > To encrypt your delivery stream, use symmetric CMKs. Kinesis Data Firehose doesn't support asymmetric CMKs. For information about symmetric and asymmetric CMKs, see [About Symmetric and Asymmetric CMKs](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html) in the AWS Key Management Service developer guide.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// If you set `KeyType` to `CUSTOMER_MANAGED_CMK` , you must specify the Amazon Resource Name (ARN) of the CMK.
	//
	// If you set `KeyType` to `AWS _OWNED_CMK` , Kinesis Data Firehose uses a service-account CMK.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

