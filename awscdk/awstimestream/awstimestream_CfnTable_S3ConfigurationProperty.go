package awstimestream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	EncryptionOption: jsii.String("encryptionOption"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   }
//
type CfnTable_S3ConfigurationProperty struct {
	// `CfnTable.S3ConfigurationProperty.BucketName`.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// `CfnTable.S3ConfigurationProperty.EncryptionOption`.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// `CfnTable.S3ConfigurationProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `CfnTable.S3ConfigurationProperty.ObjectKeyPrefix`.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

