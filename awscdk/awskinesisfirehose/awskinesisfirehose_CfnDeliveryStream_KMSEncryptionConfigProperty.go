package awskinesisfirehose


// The `KMSEncryptionConfig` property type specifies the AWS Key Management Service ( AWS KMS) encryption key that Amazon Simple Storage Service (Amazon S3) uses to encrypt data delivered by the Amazon Kinesis Data Firehose (Kinesis Data Firehose) stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kMSEncryptionConfigProperty := &kMSEncryptionConfigProperty{
//   	awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   }
//
type CfnDeliveryStream_KMSEncryptionConfigProperty struct {
	// The Amazon Resource Name (ARN) of the AWS KMS encryption key that Amazon S3 uses to encrypt data delivered by the Kinesis Data Firehose stream.
	//
	// The key must belong to the same region as the destination S3 bucket.
	AwskmsKeyArn *string `field:"required" json:"awskmsKeyArn" yaml:"awskmsKeyArn"`
}

