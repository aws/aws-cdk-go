package awskinesisanalyticsv2


// The base location of the Amazon Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentBaseLocationProperty := &s3ContentBaseLocationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   }
//
type CfnApplication_S3ContentBaseLocationProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The base path for the S3 bucket.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
}

