package awsconnect


// Information about the Amazon Simple Storage Service (Amazon S3) storage type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &S3ConfigProperty{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//
//   	// the properties below are optional
//   	EncryptionConfig: &EncryptionConfigProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   }
//
type CfnInstanceStorageConfig_S3ConfigProperty struct {
	// The S3 bucket name.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The S3 bucket prefix.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The Amazon S3 encryption configuration.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

