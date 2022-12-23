package awsconnect


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
	// `CfnInstanceStorageConfig.S3ConfigProperty.BucketName`.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// `CfnInstanceStorageConfig.S3ConfigProperty.BucketPrefix`.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// `CfnInstanceStorageConfig.S3ConfigProperty.EncryptionConfig`.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

