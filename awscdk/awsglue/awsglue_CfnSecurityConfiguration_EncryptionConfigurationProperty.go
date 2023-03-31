package awsglue


// Specifies an encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	cloudWatchEncryption: &cloudWatchEncryptionProperty{
//   		cloudWatchEncryptionMode: jsii.String("cloudWatchEncryptionMode"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryptionProperty{
//   		jobBookmarksEncryptionMode: jsii.String("jobBookmarksEncryptionMode"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	s3Encryptions: []interface{}{
//   		&s3EncryptionProperty{
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   			s3EncryptionMode: jsii.String("s3EncryptionMode"),
//   		},
//   	},
//   }
//
type CfnSecurityConfiguration_EncryptionConfigurationProperty struct {
	// The encryption configuration for Amazon CloudWatch.
	CloudWatchEncryption interface{} `field:"optional" json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for job bookmarks.
	JobBookmarksEncryption interface{} `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encyption configuration for Amazon Simple Storage Service (Amazon S3) data.
	S3Encryptions interface{} `field:"optional" json:"s3Encryptions" yaml:"s3Encryptions"`
}

