package awsglue


// Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EncryptionProperty := &s3EncryptionProperty{
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	s3EncryptionMode: jsii.String("s3EncryptionMode"),
//   }
//
type CfnSecurityConfiguration_S3EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The encryption mode to use for Amazon S3 data.
	S3EncryptionMode *string `field:"optional" json:"s3EncryptionMode" yaml:"s3EncryptionMode"`
}

