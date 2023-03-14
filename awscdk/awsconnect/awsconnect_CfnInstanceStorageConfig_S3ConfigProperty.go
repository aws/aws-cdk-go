package awsconnect


// Information about the Amazon Simple Storage Service (Amazon S3) storage type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &s3ConfigProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//
//   	// the properties below are optional
//   	encryptionConfig: &encryptionConfigProperty{
//   		encryptionType: jsii.String("encryptionType"),
//   		keyId: jsii.String("keyId"),
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

