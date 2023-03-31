package awstimestream


// Details on S3 location for error reports that result from running a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &s3ConfigurationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	encryptionOption: jsii.String("encryptionOption"),
//   	objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   }
//
type CfnScheduledQuery_S3ConfigurationProperty struct {
	// Name of the S3 bucket under which error reports will be created.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Encryption at rest options for the error reports.
	//
	// If no encryption option is specified, Timestream will choose SSE_S3 as default.
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Prefix for the error report key.
	//
	// Timestream by default adds the following prefix to the error report path.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

