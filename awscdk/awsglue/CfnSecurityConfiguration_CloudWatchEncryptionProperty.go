package awsglue


// Specifies how Amazon CloudWatch data should be encrypted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchEncryptionProperty := &CloudWatchEncryptionProperty{
//   	CloudWatchEncryptionMode: jsii.String("cloudWatchEncryptionMode"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnSecurityConfiguration_CloudWatchEncryptionProperty struct {
	// The encryption mode to use for CloudWatch data.
	CloudWatchEncryptionMode *string `field:"optional" json:"cloudWatchEncryptionMode" yaml:"cloudWatchEncryptionMode"`
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

