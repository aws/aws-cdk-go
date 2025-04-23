package awss3tables


// Properties for defining a `CfnTableBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableBucketProps := &CfnTableBucketProps{
//   	TableBucketName: jsii.String("tableBucketName"),
//
//   	// the properties below are optional
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   	UnreferencedFileRemoval: &UnreferencedFileRemovalProperty{
//   		NoncurrentDays: jsii.Number(123),
//   		Status: jsii.String("status"),
//   		UnreferencedDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucket.html
//
type CfnTableBucketProps struct {
	// The name for the table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucket.html#cfn-s3tables-tablebucket-tablebucketname
	//
	TableBucketName *string `field:"required" json:"tableBucketName" yaml:"tableBucketName"`
	// Configuration specifying how data should be encrypted.
	//
	// This structure defines the encryption algorithm and optional KMS key to be used for server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucket.html#cfn-s3tables-tablebucket-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The unreferenced file removal settings for your table bucket.
	//
	// Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. For more information, see the [*Amazon S3 User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-table-buckets-maintenance.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablebucket.html#cfn-s3tables-tablebucket-unreferencedfileremoval
	//
	UnreferencedFileRemoval interface{} `field:"optional" json:"unreferencedFileRemoval" yaml:"unreferencedFileRemoval"`
}

