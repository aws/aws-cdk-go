package awstimestream


// The configuration that specifies an S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &s3ConfigurationProperty{
//   	bucketName: jsii.String("bucketName"),
//   	encryptionOption: jsii.String("encryptionOption"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   }
//
type CfnTable_S3ConfigurationProperty struct {
	// The bucket name of the customer S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The encryption option for the customer S3 location.
	//
	// Options are S3 server-side encryption with an S3 managed key or AWS managed key.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// The AWS KMS key ID for the customer S3 location when encrypting with an AWS managed key.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The object key preview for the customer S3 location.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

