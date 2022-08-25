package awsappflow


// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourcePropertiesProperty := &s3SourcePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//
//   	// the properties below are optional
//   	s3InputFormatConfig: &s3InputFormatConfigProperty{
//   		s3InputFileType: jsii.String("s3InputFileType"),
//   	},
//   }
//
type CfnFlow_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// When you use Amazon S3 as the source, the configuration format that you provide the flow input data.
	S3InputFormatConfig interface{} `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

