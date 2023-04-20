package awstimestream


// Configuration required for error reporting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorReportConfigurationProperty := &ErrorReportConfigurationProperty{
//   	S3Configuration: &S3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		EncryptionOption: jsii.String("encryptionOption"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   }
//
type CfnScheduledQuery_ErrorReportConfigurationProperty struct {
	// The S3 configuration for the error reports.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
}

