// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Constructions properties of {@link SecurityConfiguration}.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type SecurityConfigurationProps struct {
	// The encryption configuration for Amazon CloudWatch Logs.
	// Experimental.
	CloudWatchEncryption *CloudWatchEncryption `field:"optional" json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for Glue Job Bookmarks.
	// Experimental.
	JobBookmarksEncryption *JobBookmarksEncryption `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.
	// Experimental.
	S3Encryption *S3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName *string `field:"optional" json:"securityConfigurationName" yaml:"securityConfigurationName"`
}

