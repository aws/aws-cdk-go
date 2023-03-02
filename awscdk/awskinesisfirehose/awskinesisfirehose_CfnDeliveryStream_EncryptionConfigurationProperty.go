package awskinesisfirehose


// The `EncryptionConfiguration` property type specifies the encryption settings that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses when delivering data to Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   		awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   	},
//   	noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   }
//
type CfnDeliveryStream_EncryptionConfigurationProperty struct {
	// The AWS Key Management Service ( AWS KMS) encryption key that Amazon S3 uses to encrypt your data.
	KmsEncryptionConfig interface{} `field:"optional" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
	// Disables encryption.
	//
	// For valid values, see the `NoEncryptionConfig` content for the [EncryptionConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_EncryptionConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	NoEncryptionConfig *string `field:"optional" json:"noEncryptionConfig" yaml:"noEncryptionConfig"`
}

