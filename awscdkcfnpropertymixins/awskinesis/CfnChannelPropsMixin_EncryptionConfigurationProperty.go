package awskinesis


// Server-side encryption configuration for data at rest in the destination.
//
// Data delivered to S3 / S3 Tables is encrypted with the same key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   	KeyId: jsii.String("keyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-encryptionconfiguration.html
//
type CfnChannelPropsMixin_EncryptionConfigurationProperty struct {
	// The encryption type.
	//
	// KMS is the only supported value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-encryptionconfiguration.html#cfn-kinesis-channel-encryptionconfiguration-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The customer-managed AWS KMS key.
	//
	// Accepts a key GUID, key ARN, alias ARN, or alias name prefixed by 'alias/'. The Kinesis Data Streams managed alias 'aws/kinesis' is not accepted - the key must be customer-owned so it can also be used by readers of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-encryptionconfiguration.html#cfn-kinesis-channel-encryptionconfiguration-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

