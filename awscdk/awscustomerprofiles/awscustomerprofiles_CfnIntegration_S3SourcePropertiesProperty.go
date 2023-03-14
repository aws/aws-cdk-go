package awscustomerprofiles


// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourcePropertiesProperty := &s3SourcePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnIntegration_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

