package awsglue


// Encryption mode for S3.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	SecurityConfigurationName: jsii.String("name"),
//   	CloudWatchEncryption: &CloudWatchEncryption{
//   		Mode: glue.CloudWatchEncryptionMode_KMS,
//   	},
//   	JobBookmarksEncryption: &JobBookmarksEncryption{
//   		Mode: glue.JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	S3Encryption: &S3Encryption{
//   		Mode: glue.S3EncryptionMode_KMS,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_S3Encryption.html#Glue-Type-S3Encryption-S3EncryptionMode
//
// Experimental.
type S3EncryptionMode string

const (
	// Server side encryption (SSE) with an Amazon S3-managed key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
	//
	// Experimental.
	S3EncryptionMode_S3_MANAGED S3EncryptionMode = "S3_MANAGED"
	// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
	//
	// Experimental.
	S3EncryptionMode_KMS S3EncryptionMode = "KMS"
)

